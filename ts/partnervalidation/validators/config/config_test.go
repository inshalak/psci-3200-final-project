package config

import (
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	// Test that INSHA_TEST is in on-demand supported tokens
	expectedOnDemand := []string{"INSHA_TEST"}
	if !reflect.DeepEqual(config.GetOnDemandSupportedTokenTypes(), expectedOnDemand) {
		t.Errorf("Expected on-demand supported tokens %v, got %v", expectedOnDemand, config.GetOnDemandSupportedTokenTypes())
	}

	// Test that INSHA_TEST is in validity checks supported tokens
	expectedValidity := []string{"INSHA_TEST"}
	if !reflect.DeepEqual(config.GetValidityChecksSupportedTokens(), expectedValidity) {
		t.Errorf("Expected validity checks supported tokens %v, got %v", expectedValidity, config.GetValidityChecksSupportedTokens())
	}

	// Test that INSHA_TEST is NOT in disabled validity tokens
	expectedDisabled := []string{}
	if !reflect.DeepEqual(config.GetDisabledValidityTokens(), expectedDisabled) {
		t.Errorf("Expected disabled validity tokens %v, got %v", expectedDisabled, config.GetDisabledValidityTokens())
	}
}

func TestConfigMethods(t *testing.T) {
	config := NewConfig()

	// Test WithOnDemandChecksSupportedForTokenTypes
	config.WithOnDemandChecksSupportedForTokenTypes("INSHA_TEST")
	expected := []string{"INSHA_TEST"}
	if !reflect.DeepEqual(config.GetOnDemandSupportedTokenTypes(), expected) {
		t.Errorf("Expected %v, got %v", expected, config.GetOnDemandSupportedTokenTypes())
	}

	// Test WithValidityChecksSupportedForTokenTypes
	config.WithValidityChecksSupportedForTokenTypes("INSHA_TEST")
	if !reflect.DeepEqual(config.GetValidityChecksSupportedTokens(), expected) {
		t.Errorf("Expected %v, got %v", expected, config.GetValidityChecksSupportedTokens())
	}

	// Test that DisableValidityUpdatesForTokenTypes is available but not used for INSHA_TEST
	emptyExpected := []string{}
	if !reflect.DeepEqual(config.GetDisabledValidityTokens(), emptyExpected) {
		t.Errorf("Expected empty disabled tokens, got %v", config.GetDisabledValidityTokens())
	}
}
