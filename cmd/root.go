/*
Copyright © 2023 Nicolas Bock <nicolasbock@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "openstack-go",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags
// appropriately. This is called by main.main(). It only needs to happen once to
// the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings. Cobra
	// supports persistent flags, which, if defined here, will be global for
	// your application.

	// Cobra also supports local flags, which will only run when this action is
	// called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
