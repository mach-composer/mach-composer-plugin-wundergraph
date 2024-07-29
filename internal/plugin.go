package internal

import (
	"fmt"

	"github.com/mach-composer/mach-composer-plugin-helpers/helpers"
	"github.com/mach-composer/mach-composer-plugin-sdk/plugin"
	"github.com/mach-composer/mach-composer-plugin-sdk/schema"
	"github.com/mitchellh/mapstructure"
)

type WundergraphPlugin struct {
	provider     string
	environment  string
	globalConfig *WundergraphGlobalConfig
	siteConfigs  map[string]*WundergraphSiteConfig
	enabled      bool
}

func NewWundergraphPlugin() schema.MachComposerPlugin {
	state := &WundergraphPlugin{
		provider:    "0.1.0",
		siteConfigs: map[string]*WundergraphSiteConfig{},
	}
	return plugin.NewPlugin(&schema.PluginSchema{
		Identifier:          "wundergraph",
		Configure:           state.Configure,
		IsEnabled:           func() bool { return state.enabled },
		GetValidationSchema: state.GetValidationSchema,

		SetGlobalConfig: state.SetGlobalConfig,
		SetSiteConfig:   state.SetSiteConfig,

		// Renders
		RenderTerraformProviders: state.RenderTerraformProviders,
		RenderTerraformResources: state.RenderTerraformResources,
		RenderTerraformComponent: state.RenderTerraformComponent,
	})
}

func (p *WundergraphPlugin) Configure(environment string, provider string) error {
	p.environment = environment
	if provider != "" {
		p.provider = provider
	}
	return nil
}

func (p *WundergraphPlugin) GetValidationSchema() (*schema.ValidationSchema, error) {
	result := getSchema()
	return result, nil
}

func (p *WundergraphPlugin) SetGlobalConfig(data map[string]any) error {
	cfg := WundergraphGlobalConfig{}

	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.globalConfig = &cfg
	p.enabled = true

	return nil
}

func (p *WundergraphPlugin) SetSiteConfig(site string, data map[string]any) error {
	cfg := WundergraphSiteConfig{}
	if err := mapstructure.Decode(data, &cfg); err != nil {
		return err
	}
	p.siteConfigs[site] = &cfg
	p.enabled = true
	return nil
}

func (p *WundergraphPlugin) getSiteConfig(site string) *WundergraphSiteConfig {
	result := &WundergraphSiteConfig{}
	if p.globalConfig != nil {
		result.ApiKey = p.globalConfig.ApiKey
		result.ApiUrl = p.globalConfig.ApiUrl
	}

	cfg, ok := p.siteConfigs[site]
	if ok {
		if cfg.ApiKey != "" {
			result.ApiKey = cfg.ApiKey
			result.ApiUrl = cfg.ApiUrl
		}
	}

	return result
}

func (p *WundergraphPlugin) RenderTerraformStateBackend(_ string) (string, error) {
	return "", nil
}

func (p *WundergraphPlugin) RenderTerraformProviders(site string) (string, error) {
	cfg := p.getSiteConfig(site)

	if cfg == nil {
		return "", nil
	}

	result := fmt.Sprintf(`
		wundergraph = {
			source = "labd/wundergraph"
			version = "%s"
		}
	`, helpers.VersionConstraint(p.provider))

	return result, nil
}

func (p *WundergraphPlugin) RenderTerraformResources(site string) (string, error) {
	cfg := p.getSiteConfig(site)

	if cfg == nil {
		return "", nil
	}

	template := `
		provider "wundergraph" {
			{{ renderProperty "api_key" .ApiKey }}
			{{ renderOptionalProperty "api_url" .ApiUrl }}
		}
	`
	return helpers.RenderGoTemplate(template, cfg)
}

func (p *WundergraphPlugin) RenderTerraformComponent(site string, _ string) (*schema.ComponentSchema, error) {
	cfg := p.getSiteConfig(site)
	if cfg == nil {
		return nil, nil
	}

	template := `
		wundergraph = {
			{{ renderProperty "api_key" .ApiKey }}
			{{ renderOptionalProperty "api_url" .ApiUrl }}
		}
	`

	vars, err := helpers.RenderGoTemplate(template, cfg)
	if err != nil {
		return nil, err
	}

	result := &schema.ComponentSchema{
		Variables: vars,
	}

	return result, nil
}
