# Wundergraph Plugin for MACH composer

This repository contains the Wundergraph plugin for Mach Composer. It requires MACH
composer >= 2.5.x

This plugin uses the [Wundergraph Terraform Provider](https://github.com/labd/terraform-provider-wundergraph)




## Usage

```yaml
mach_composer:
  version: 1
  plugins:
    wundergraph:
      source: mach-composer/wundergraph
      version: 0.1.0

global:
  # ...
  wundergraph:
    api_key: "my-apikey"

sites:
  - identifier: my-site
    # ...
    wundergraph:
      api_key: "my-apikey"
    components:
      - name: my-component
        # ...
```
