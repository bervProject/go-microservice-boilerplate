package utils

import (
	"strings"
	"testing"
)

func TestGetEnvAvailableEnv(t *testing.T) {
	t.Setenv("TESTING", "null")
	result := GetEnv("TESTING", "")
	if !strings.EqualFold(result, "null") {
		t.Fatalf("Get wrong result: %s", result)
	}
}

func TestGetEnvNotAvailableEnv(t *testing.T) {
	result := GetEnv("TESTING", "meow")
	if !strings.EqualFold(result, "meow") {
		t.Fatalf("Get wrong result: %s", result)
	}
}
