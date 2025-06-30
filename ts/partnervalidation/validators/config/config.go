package config

// ValidatorConfig represents the configuration for partner validation
type ValidatorConfig struct {
	onDemandChecksSupportedTokenTypes []string
	validityChecksSupportedTokenTypes []string
	disabledValidityUpdatesTokenTypes []string
}

// NewValidatorConfig creates a new instance of ValidatorConfig
func NewValidatorConfig() *ValidatorConfig {
	return &ValidatorConfig{
		onDemandChecksSupportedTokenTypes: make([]string, 0),
		validityChecksSupportedTokenTypes: make([]string, 0),
		disabledValidityUpdatesTokenTypes: make([]string, 0),
	}
}

// WithOnDemandChecksSupportedForTokenTypes adds token types that support on-demand checks
func (c *ValidatorConfig) WithOnDemandChecksSupportedForTokenTypes(tokenTypes ...string) *ValidatorConfig {
	c.onDemandChecksSupportedTokenTypes = append(c.onDemandChecksSupportedTokenTypes, tokenTypes...)
	return c
}

// WithValidityChecksSupportedForTokenTypes adds token types that support validity checks
func (c *ValidatorConfig) WithValidityChecksSupportedForTokenTypes(tokenTypes ...string) *ValidatorConfig {
	c.validityChecksSupportedTokenTypes = append(c.validityChecksSupportedTokenTypes, tokenTypes...)
	return c
}

// DisableValidityUpdatesForTokenTypes disables validity updates for specific token types
func (c *ValidatorConfig) DisableValidityUpdatesForTokenTypes(tokenTypes ...string) *ValidatorConfig {
	c.disabledValidityUpdatesTokenTypes = append(c.disabledValidityUpdatesTokenTypes, tokenTypes...)
	return c
}

// GetOnDemandChecksSupportedTokenTypes returns the list of token types that support on-demand checks
func (c *ValidatorConfig) GetOnDemandChecksSupportedTokenTypes() []string {
	return c.onDemandChecksSupportedTokenTypes
}

// GetValidityChecksSupportedTokenTypes returns the list of token types that support validity checks
func (c *ValidatorConfig) GetValidityChecksSupportedTokenTypes() []string {
	return c.validityChecksSupportedTokenTypes
}

// GetDisabledValidityUpdatesTokenTypes returns the list of token types with disabled validity updates
func (c *ValidatorConfig) GetDisabledValidityUpdatesTokenTypes() []string {
	return c.disabledValidityUpdatesTokenTypes
}

// IsOnDemandChecksSupported checks if on-demand checks are supported for a token type
func (c *ValidatorConfig) IsOnDemandChecksSupported(tokenType string) bool {
	for _, supportedType := range c.onDemandChecksSupportedTokenTypes {
		if supportedType == tokenType {
			return true
		}
	}
	return false
}

// IsValidityChecksSupported checks if validity checks are supported for a token type
func (c *ValidatorConfig) IsValidityChecksSupported(tokenType string) bool {
	for _, supportedType := range c.validityChecksSupportedTokenTypes {
		if supportedType == tokenType {
			return true
		}
	}
	return false
}

// IsValidityUpdatesDisabled checks if validity updates are disabled for a token type
func (c *ValidatorConfig) IsValidityUpdatesDisabled(tokenType string) bool {
	for _, disabledType := range c.disabledValidityUpdatesTokenTypes {
		if disabledType == tokenType {
			return true
		}
	}
	return false
}

// GetDefaultConfigForINSHA_TEST2 returns a configuration with INSHA_TEST2 token type enabled
func GetDefaultConfigForINSHA_TEST2() *ValidatorConfig {
	return NewValidatorConfig().
		WithOnDemandChecksSupportedForTokenTypes("INSHA_TEST2").
		WithValidityChecksSupportedForTokenTypes("INSHA_TEST2")
}
