package config

// Config represents the validator configuration
type Config struct {
	onDemandSupportedTokenTypes   []string
	validityChecksSupportedTokens []string
	disabledValidityTokens        []string
}

// NewConfig creates a new configuration instance
func NewConfig() *Config {
	return &Config{
		onDemandSupportedTokenTypes:   make([]string, 0),
		validityChecksSupportedTokens: make([]string, 0),
		disabledValidityTokens:        make([]string, 0),
	}
}

// WithOnDemandChecksSupportedForTokenTypes adds token types that support on-demand checks
func (c *Config) WithOnDemandChecksSupportedForTokenTypes(tokenTypes ...string) *Config {
	c.onDemandSupportedTokenTypes = append(c.onDemandSupportedTokenTypes, tokenTypes...)
	return c
}

// WithValidityChecksSupportedForTokenTypes adds token types that support validity checks
func (c *Config) WithValidityChecksSupportedForTokenTypes(tokenTypes ...string) *Config {
	c.validityChecksSupportedTokens = append(c.validityChecksSupportedTokens, tokenTypes...)
	return c
}

// DisableValidityUpdatesForTokenTypes disables validity updates for specific token types
func (c *Config) DisableValidityUpdatesForTokenTypes(tokenTypes ...string) *Config {
	c.disabledValidityTokens = append(c.disabledValidityTokens, tokenTypes...)
	return c
}

// GetOnDemandSupportedTokenTypes returns the list of token types that support on-demand checks
func (c *Config) GetOnDemandSupportedTokenTypes() []string {
	return c.onDemandSupportedTokenTypes
}

// GetValidityChecksSupportedTokens returns the list of token types that support validity checks
func (c *Config) GetValidityChecksSupportedTokens() []string {
	return c.validityChecksSupportedTokens
}

// GetDisabledValidityTokens returns the list of token types with disabled validity updates
func (c *Config) GetDisabledValidityTokens() []string {
	return c.disabledValidityTokens
}

// DefaultConfig creates a default configuration with INSHA_TEST token enabled
func DefaultConfig() *Config {
	return NewConfig().
		WithOnDemandChecksSupportedForTokenTypes("INSHA_TEST").
		WithValidityChecksSupportedForTokenTypes("INSHA_TEST")
}
