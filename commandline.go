package main

import (
	"encoding/json"
	"flag"
	"os"
)

type Options struct {
	OS_AUTH_PROTOCOL        string `json:"os_auth_protocol"`
	OS_AUTH_TYPE            string `json:"os_auth_type"`
	OS_AUTH_URL             string `json:"os_auth_url,omitempty"`
	os_identity_api_version string
	os_password             string
	os_project_domain_name  string
	os_project_name         string
	os_region_name          string
	os_username             string
	os_user_domain_name     string
}

func NewOptions() Options {
	options := Options{}
	options.OS_AUTH_PROTOCOL = "http"
	options.OS_AUTH_TYPE = "password"
	return options
}

func (o Options) String() string {
	result, _ := json.Marshal(o)
	return string(result)
}

func getEnvWithDefault(key string, defaultValue string) string {
	value, isSet := os.LookupEnv(key)
	if isSet {
		return value
	} else {
		return defaultValue
	}
}

func getEnv(key string, value *string) {
	envValue, isSet := os.LookupEnv(key)
	if isSet {
		*value = envValue
	}
}

func parseEnvironment() Options {
	var options Options = NewOptions()
	getEnv("OS_AUTH_PROTOCOL", &options.OS_AUTH_PROTOCOL)
	getEnv("OS_AUTH_TYPE", &options.OS_AUTH_TYPE)
	getEnv("OS_AUTH_URL", &options.OS_AUTH_URL)
	return options
}

func parseCommandline() Options {
	options := parseEnvironment()
	flag.StringVar(&options.OS_AUTH_PROTOCOL, "os-auth-protocol", "http", "protocol")
	flag.Parse()
	return options
}
