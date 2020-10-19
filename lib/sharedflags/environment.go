// Copyright (c) Improbable Worlds Ltd, All Rights Reserved

package sharedflags

import (
	//"fmt"
	"os"
	//"strings"

	//"improbable.io/lib/cloud/cloudaccess"
)

const (
	// SecureEnvironmentFlagName when enabled will force all the external connections from the Spatial CLI to the outside world to be tunnelled through Aquaman proxy.
	// For more details, check https://improbableio.atlassian.net/wiki/spaces/GBU/pages/244318309/Aquaman+secure+customer+proxy.
	SecureEnvironmentFlagName = "secure_environment"
	DomainNameWest            = "improbable.io"
	DomainNameChina           = "spatialoschina.com"
)

var (
	//DO NOT REMOVE defaultEnvironment var. This is set during builds
	//defaultEnvironment MUST be a constant (see `-X importpath.name=value` in https://golang.org/cmd/link/)
	defaultEnvironment = "testing"

	//Environment indicates environment that we are in.
	Environment = Set.String(
		"environment",
		getEnvironment(),
		"The environment this server is running in (eg: production, staging)")

	//ClusterName indicates cluster that we are in.
	ClusterName = Set.String(
		"cluster_name",
		"cp-gce-europe-west1-testing",
		"This indicates cluster that we are in.")

	_ = Set.String(
		"orchestration_environment",
		"kubernetes",
		"DEPRECATED: Defines which orchestration is backing this cluster")

	//DevComponents indicates components that should be expected to be running locally.
	DevComponents = Set.StringSlice(
		"dev_components",
		[]string{},
		"The components to override to local instances when running. eg: thor,gcontroller",
	)

	_ = Set.String("cloud_provider", "", "DEPRECATED")

	// Domain indicates the top level domain name for public APIs (e.g. improbable.io or spatialoschina.com).
	//
	// Deprecated: The domain variable can now be derived automatically from the environment variable. Use GetDomain()
	// instead.
	domain = Set.String(
		"domain",
		"",
		"DEPRECATED: The top level domain name for public facing APIs. eg: "+DomainNameWest,
	)
)

//func GetDomain() string {
//	if *domain != "" {
//		fmt.Printf("WARNING: Use of the Domain flag is deprecated, if unset the domain will be derived from value of the '--environment' flag\n")
//		return *domain
//	}
//
//	environment, contains := cloudaccess.EnvByName[*Environment]
//	if !contains {
//		fmt.Printf("Unrecognized environment value: %s\n", *Environment)
//
//		// There's no real way to recover from this. We could use the default environment but that is almost certainly
//		// not what the user wanted. We could return an empty string or an error, but the unrecognized environment is
//		// just going to cause a crash later, so why not just stop now right after we printed the reason why.
//		os.Exit(1)
//	}
//	return environment.RootDomain
//}

func init() {
	envCluster := os.ExpandEnv("$CLUSTER_NAME")
	if envCluster != "" {
		if err := Set.Set("cluster_name", envCluster); err != nil {
			panic(err)
		}
	}
}

//func IsChinaEnvironment() bool {
//	return GetDomain() == DomainNameChina || strings.HasPrefix(*Environment, "cn-")
//}

func getEnvironment() string {
	ciEnv := os.Getenv("CI_ENVIRONMENT")
	if ciEnv != "" {
		return ciEnv
	}
	return defaultEnvironment
}
