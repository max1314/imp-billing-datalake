// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package tests

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
	"testing"
	"time"

	"improbable.io/lib/db/blob"
	"improbable.io/lib/errors"
	"improbable.io/lib/extclients/memcachedext"
	mocks_blob "improbable.io/mocks/lib/db/blob"
	mocks_cache "improbable.io/mocks/lib/extclients/memcachedext"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/mock" // nolint
	"github.com/stretchr/testify/require"
)

func TestReadCacheDecorator_CacheMissWritesBytesToCache(t *testing.T) {
	bucket := newMockBucket()

	cacheClient := &mocks_cache.Client{}
	cacheClient.On("Get", mock.Anything).Return(nil, memcache.ErrNotStored)
	cacheClient.On("Set", mock.Anything).Return(nil)

	cachedReader := blob.ReadCacheDecorator(bucket, cacheClient, "mynamespace")

	r, err := cachedReader.NewReader("some/resource.md", nil)
	require.NoError(t, err, "must return reader")
	actual, err := ioutil.ReadAll(r)
	r.Close()
	require.NoError(t, err, "must return reader")
	require.Equal(t, []byte("hello"), actual, "must return contents")

	setCalls := getCacheSetCalls(cacheClient)
	actualCacheItem, actualCacheEntry, err := getCacheItemAndEntry(setCalls[0])
	require.NoError(t, err, "must read cache item")

	require.Equal(t, 1, len(setCalls), "must call client.Set() once")
	require.Equal(t, []byte("hello"), actualCacheEntry.Bytes, "must cache expected bytes")
	require.Equal(t, memcachedext.ExpirationSec(time.Hour*1), actualCacheItem.Expiration, "must set default expiration for a successful read")
	require.True(t, strings.HasPrefix(actualCacheItem.Key, "mynamespace/"), "cacheKey prefixed with supplied namespace")
}

func TestReadCacheDecorator_CacheMissWritesErrorToCache(t *testing.T) {
	bucket := newMockBucket()

	cacheClient := &mocks_cache.Client{}
	cacheClient.On("Get", mock.Anything).Return(nil, memcache.ErrNotStored)
	cacheClient.On("Set", mock.Anything).Return(nil)

	cachedReader := blob.ReadCacheDecorator(bucket, cacheClient, "mynamespace")

	_, err := cachedReader.NewReader("missing/resource.md", nil)
	require.Error(t, err, "must return reader")

	setCalls := getCacheSetCalls(cacheClient)
	actualCacheItem, actualCacheEntry, err := getCacheItemAndEntry(setCalls[0])
	require.NoError(t, err, "must read cache item")

	require.Equal(t, 1, len(setCalls), "must call client.Set() once")
	require.Equal(t, errors.NotFound, actualCacheEntry.ErrCode, "must cache BucketReader error code")
	require.Equal(t, "no bueno", actualCacheEntry.ErrMsg, "must cache BucketReader error message")
	require.Equal(t, memcachedext.ExpirationSec(time.Minute*5), actualCacheItem.Expiration, "must set default expiration for a successful read")
	require.True(t, strings.HasPrefix(actualCacheItem.Key, "mynamespace/"), "cacheKey prefixed with supplied namespace")
}

func TestReadCacheDecorator_CacheMissWritesWithCustomExpirations(t *testing.T) {
	bucket := newMockBucket()

	cacheClient := &mocks_cache.Client{}
	cacheClient.On("Get", mock.Anything).Return(nil, memcache.ErrNotStored)
	cacheClient.On("Set", mock.Anything).Return(nil)

	cachedReader := blob.ReadCacheDecorator(bucket, cacheClient, "mynamespace",
		blob.WithBytesExpiration(time.Minute*13),
		blob.WithErrorExpiration(time.Minute*22),
	)

	r, err := cachedReader.NewReader("some/resource.md", nil)
	if err == nil {
		ioutil.ReadAll(r)
		r.Close()
	}
	r, err = cachedReader.NewReader("missing/resource.md", nil)
	if err == nil {
		ioutil.ReadAll(r)
		r.Close()
	}

	setCalls := getCacheSetCalls(cacheClient)
	require.Equal(t, 2, len(setCalls), "must cache two resources")

	successEntry, _, err := getCacheItemAndEntry(setCalls[0])
	require.NoError(t, err, "must read cache item")
	require.Equal(t, memcachedext.ExpirationSec(time.Minute*13), successEntry.Expiration, "must set custom byte expiration time")

	errorEntry, _, err := getCacheItemAndEntry(setCalls[1])
	require.NoError(t, err, "must read cache item")
	require.Equal(t, memcachedext.ExpirationSec(time.Minute*22), errorEntry.Expiration, "must set custom error expiration time")
}

func TestReadCacheDecorator_ReadsBytesFromCacheNotUnderlyingBucket(t *testing.T) {
	bucket := newMockBucket()

	cacheEntry, err := memcachedext.Marshal(blob.CacheEntry{Bytes: []byte("read-from-cache")})
	require.NoError(t, err, "must marshal CacheEntry")
	cacheItem := &memcache.Item{Value: cacheEntry}

	cacheClient := &mocks_cache.Client{}
	cacheClient.On("Get", cacheKeyMatcher(t, "mynamespace", "must-never-request")).
		Return(cacheItem, nil)

	cachedReader := blob.ReadCacheDecorator(bucket, cacheClient, "mynamespace")
	r, err := cachedReader.NewReader("must-never-request", nil)
	require.NoError(t, err, "must not fail to read")
	actual, err := ioutil.ReadAll(r)
	r.Close()

	require.NoError(t, err, "must not fail to read")
	require.Equal(t, []byte("read-from-cache"), actual, "must pull value from cache")
}

func TestReadCacheDecorator_ReadsErrorFromCacheNotUnderlyingBucket(t *testing.T) {
	bucket := newMockBucket()

	cacheEntry, err := memcachedext.Marshal(blob.CacheEntry{ErrCode: errors.NotFound, ErrMsg: "err-came-from-cache"})
	require.NoError(t, err, "must marshal CacheEntry")
	cacheItem := &memcache.Item{Value: cacheEntry}

	cacheClient := &mocks_cache.Client{}
	cacheClient.On("Get", cacheKeyMatcher(t, "mynamespace", "must-never-request")).
		Return(cacheItem, nil)

	cachedReader := blob.ReadCacheDecorator(bucket, cacheClient, "mynamespace")
	_, err = cachedReader.NewReader("must-never-request", nil)

	require.Error(t, err, "must fail to read")
	require.Equal(t, errors.NotFound, errors.Code(err), "must return error code from cache")
	require.Equal(t, "err-came-from-cache", errors.Desc(err), "must return error message from cache")
}

func newMockBucket() *mocks_blob.Bucket {
	bucket := &mocks_blob.Bucket{}
	bucket.On("NewReader", "some/resource.md", mock.Anything).Return(objectReader("hello"), nil)
	bucket.On("NewReader", "missing/resource.md", mock.Anything).Return(nil, errors.New(nil, errors.NotFound, "no bueno"))
	return bucket
}

func getCacheSetCalls(cacheClient *mocks_cache.Client) []mock.Call {
	setCalls := []mock.Call{}
	for _, call := range cacheClient.Calls {
		if call.Method == "Set" {
			setCalls = append(setCalls, call)
		}
	}
	return setCalls
}

func getCacheItemAndEntry(call mock.Call) (*memcache.Item, blob.CacheEntry, error) {
	item, ok := call.Arguments[0].(*memcache.Item)
	if !ok {
		item = &memcache.Item{}
	}
	var entry blob.CacheEntry
	if err := memcachedext.Unmarshal(item, &entry); err != nil {
		return nil, blob.CacheEntry{}, err
	}
	return item, entry, nil
}

func objectReader(str string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBufferString(str))
}

func cacheKeyMatcher(t *testing.T, prefix string, p string) interface{} {
	expected, err := blob.MakeCacheKey(prefix, p)
	require.NoError(t, err, "must be able to compute cache key")
	return mock.MatchedBy(func(v interface{}) bool {
		if actual, ok := v.(string); ok {
			return actual == expected
		}
		return false
	})
}
