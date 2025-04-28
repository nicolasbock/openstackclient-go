package lib

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
	expectedAuthProtocol := "http"
	expectedAuthType := "passwordx"
	expectedAuthURL := "https://1.2.3.4:5000/v3"
	os.Setenv("OS_AUTH_PROTOCOL", expectedAuthProtocol)
	os.Setenv("OS_AUTH_TYPE", expectedAuthType)
	os.Setenv("OS_AUTH_URL", expectedAuthURL)
	options := parseEnvironment()
	if options.OS_AUTH_TYPE != expectedAuthType {
		t.Errorf("expected %s, got %s", expectedAuthType, options.OS_AUTH_TYPE)
	}
	if options.OS_AUTH_PROTOCOL != expectedAuthProtocol {
		t.Errorf("expected %s, got %s", expectedAuthProtocol, options.OS_AUTH_PROTOCOL)
	}
	if options.OS_AUTH_URL != expectedAuthURL {
		t.Errorf("expected %s, got %s", expectedAuthURL, options.OS_AUTH_URL)
	}
}
