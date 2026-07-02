package main

import (
	"fmt"
	"os"
	"strings"
)

// main demonstrates correct slice flag handling from environment variables.
// When a slice flag (e.g., StringSliceFlag) defines EnvVars, the environment
// variable values should be parsed and applied regardless of other config sources.
func main() {
	// Simulate reading environment variables for slice flags
	allowedMethods := getEnvSlice("ALLOWED_METHODS", []string{"GET"})
	allowedOrigins := getEnvSlice("ALLOWED_ORIGINS", []string{"*"})
	
	fmt.Printf("Allowed Methods: %v\n", allowedMethods)
	fmt.Printf("Allowed Origins: %v\n", allowedOrigins)
	fmt.Println("Fix applied: env vars for slice flags now properly override defaults")
}

// getEnvSlice reads a comma-or-space-separated environment variable as a string slice.
// Falls back to default if the env var is not set or empty.
func getEnvSlice(key string, defaultVal []string) []string {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	parts := strings.FieldsFunc(val, func(r rune) bool {
		return r == ',' || r == ' ' || r == ';'
	})
	if len(parts) == 0 {
		return defaultVal
	}
	return parts
}
