package main

import (
	"fmt"
	"psci-3200-final-project/ts/partnervalidation/validators/config"
)

func main() {
	// Create the default configuration with INSHA_TEST enabled
	cfg := config.DefaultConfig()
	
	fmt.Println("=== INSHA_TEST Token Configuration ===")
	fmt.Printf("On-Demand Checks Supported: %v\n", cfg.GetOnDemandSupportedTokenTypes())
	fmt.Printf("Validity Checks Supported: %v\n", cfg.GetValidityChecksSupportedTokens())
	fmt.Printf("Disabled Validity Tokens: %v\n", cfg.GetDisabledValidityTokens())
	
	// Verify INSHA_TEST is properly configured
	onDemandTokens := cfg.GetOnDemandSupportedTokenTypes()
	validityTokens := cfg.GetValidityChecksSupportedTokens()
	disabledTokens := cfg.GetDisabledValidityTokens()
	
	hasOnDemand := false
	hasValidity := false
	isDisabled := false
	
	for _, token := range onDemandTokens {
		if token == "INSHA_TEST" {
			hasOnDemand = true
			break
		}
	}
	
	for _, token := range validityTokens {
		if token == "INSHA_TEST" {
			hasValidity = true
			break
		}
	}
	
	for _, token := range disabledTokens {
		if token == "INSHA_TEST" {
			isDisabled = true
			break
		}
	}
	
	fmt.Println("\n=== Verification Results ===")
	fmt.Printf("✓ INSHA_TEST has on-demand checks: %t\n", hasOnDemand)
	fmt.Printf("✓ INSHA_TEST has validity checks: %t\n", hasValidity)
	fmt.Printf("✓ INSHA_TEST is NOT disabled: %t\n", !isDisabled)
	
	if hasOnDemand && hasValidity && !isDisabled {
		fmt.Println("\n🎉 SUCCESS: INSHA_TEST is properly configured!")
	} else {
		fmt.Println("\n❌ ERROR: INSHA_TEST configuration is incomplete!")
	}
}