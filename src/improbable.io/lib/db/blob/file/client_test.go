// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package file

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContentWithCommonPrefix(t *testing.T) {
	dir := "tmp"
	_, err := NewBucket(dir)
	assert.EqualError(t, err, "code = NotFound desc = the directory tmp does not exist")

	os.Mkdir(dir, 0755)
	defer os.RemoveAll(dir)

	b, err := NewBucket(dir)
	assert.NoError(t, err)

	objs, err := b.GetContentWithCommonPrefix(context.Background(), "abc", 10)
	assert.NoError(t, err)
	assert.Empty(t, objs)

	f, _ := os.Create(dir + "/abcd")
	f.Close()

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc", 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abcd"}, objs)

	os.Mkdir(dir+"/abce", 0755)

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc", 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abcd", "abce/"}, objs)

	f, _ = os.Create(dir + "/abce/123")
	f.Close()

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc", 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abcd", "abce/", "abce/123"}, objs)

	os.Mkdir(dir+"/abc", 0755)

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc", 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abc/", "abcd", "abce/", "abce/123"}, objs)

	os.Mkdir(dir+"/abc/123", 0755)

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc/", 10)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abc/", "abc/123/"}, objs)

	objs, err = b.GetContentWithCommonPrefix(context.Background(), "abc/", 1)
	assert.NoError(t, err)
	assert.Equal(t, []string{"abc/"}, objs)
}
