package main

import (
	"os"
	"testing"
)

func TestNewOptions(t *testing.T) {
	options := NewOptions()
	if options.OS_AUTH_PROTOCOL != "http" {
		t.Errorf("os_auth_protocol != http (%s)", options.OS_AUTH_PROTOCOL)
	}
	if options.OS_AUTH_TYPE != "password" {
		t.Errorf("options.os_auth_type != password (%s)", options.OS_AUTH_TYPE)
	}
}

func TestParseEnvironment(t *testing.T) {
	expected := "https://1.2.3.4:5000/v3"
	os.Setenv("OS_AUTH_URL", expected)
	options := parseEnvironment()
	if options.OS_AUTH_URL != expected {
		t.Errorf("exptect %s, got %s", expected, options.OS_AUTH_URL)
	}
}
