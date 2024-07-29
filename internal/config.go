package internal

type WundergraphGlobalConfig struct {
	ApiKey string  `mapstructure:"api_key"`
	ApiUrl *string `mapstructure:"api_url"`
}

type WundergraphSiteConfig struct {
	ApiKey string  `mapstructure:"api_key"`
	ApiUrl *string `mapstructure:"api_url"`
}
