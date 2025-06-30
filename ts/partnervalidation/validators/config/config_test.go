package config

import (
	"testing"
)

func TestValidatorConfigINSHA_TEST2(t *testing.T) {
	// Test the default configuration for INSHA_TEST2
	config := GetDefaultConfigForINSHA_TEST2()

	// Test on-demand checks support
	if !config.IsOnDemandChecksSupported("INSHA_TEST2") {
		t.Errorf("Expected on-demand checks to be supported for INSHA_TEST2")
	}

	// Test validity checks support
	if !config.IsValidityChecksSupported("INSHA_TEST2") {
		t.Errorf("Expected validity checks to be supported for INSHA_TEST2")
	}

	// Test that validity updates are not disabled
	if config.IsValidityUpdatesDisabled("INSHA_TEST2") {
		t.Errorf("Expected validity updates to be enabled for INSHA_TEST2")
	}

	// Test that unsupported token types return false
	if config.IsOnDemandChecksSupported("UNSUPPORTED_TOKEN") {
		t.Errorf("Expected on-demand checks to be unsupported for UNSUPPORTED_TOKEN")
	}

	if config.IsValidityChecksSupported("UNSUPPORTED_TOKEN") {
		t.Errorf("Expected validity checks to be unsupported for UNSUPPORTED_TOKEN")
	}
}

func TestValidatorConfigMethods(t *testing.T) {
	config := NewValidatorConfig()

	// Test WithOnDemandChecksSupportedForTokenTypes
	config.WithOnDemandChecksSupportedForTokenTypes("INSHA_TEST2")
	supportedTypes := config.GetOnDemandChecksSupportedTokenTypes()
	if len(supportedTypes) != 1 || supportedTypes[0] != "INSHA_TEST2" {
		t.Errorf("Expected INSHA_TEST2 to be in on-demand supported types, got %v", supportedTypes)
	}

	// Test WithValidityChecksSupportedForTokenTypes
	config.WithValidityChecksSupportedForTokenTypes("INSHA_TEST2")
	validityTypes := config.GetValidityChecksSupportedTokenTypes()
	if len(validityTypes) != 1 || validityTypes[0] != "INSHA_TEST2" {
		t.Errorf("Expected INSHA_TEST2 to be in validity supported types, got %v", validityTypes)
	}

	// Test DisableValidityUpdatesForTokenTypes
	config.DisableValidityUpdatesForTokenTypes("SOME_OTHER_TOKEN")
	disabledTypes := config.GetDisabledValidityUpdatesTokenTypes()
	if len(disabledTypes) != 1 || disabledTypes[0] != "SOME_OTHER_TOKEN" {
		t.Errorf("Expected SOME_OTHER_TOKEN to be in disabled validity types, got %v", disabledTypes)
	}

	// Test that INSHA_TEST2 is not disabled
	if config.IsValidityUpdatesDisabled("INSHA_TEST2") {
		t.Errorf("Expected INSHA_TEST2 validity updates to not be disabled")
	}
}

func TestMultipleTokenTypes(t *testing.T) {
	config := NewValidatorConfig()

	// Add multiple token types
	config.WithOnDemandChecksSupportedForTokenTypes("INSHA_TEST2", "ANOTHER_TOKEN", "THIRD_TOKEN")
	config.WithValidityChecksSupportedForTokenTypes("INSHA_TEST2", "ANOTHER_TOKEN")

	// Test all are supported for on-demand checks
	for _, tokenType := range []string{"INSHA_TEST2", "ANOTHER_TOKEN", "THIRD_TOKEN"} {
		if !config.IsOnDemandChecksSupported(tokenType) {
			t.Errorf("Expected %s to support on-demand checks", tokenType)
		}
	}

	// Test only first two are supported for validity checks
	for _, tokenType := range []string{"INSHA_TEST2", "ANOTHER_TOKEN"} {
		if !config.IsValidityChecksSupported(tokenType) {
			t.Errorf("Expected %s to support validity checks", tokenType)
		}
	}

	// Test third token is not supported for validity checks
	if config.IsValidityChecksSupported("THIRD_TOKEN") {
		t.Errorf("Expected THIRD_TOKEN to not support validity checks")
	}
}
