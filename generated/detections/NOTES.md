## Differences between bundled.yaml and OpenAPI spec in Kibana repo

- Extra tags are removed from API route definitions. Code gen duplicates code when it sees multiple tags, causing compilation problems.
  - "API" postfix on tag names is removed so generated code doesn't duplicate the "API" string in names
- Many routes are removed
- spaceId, kbn_xsrf, and elastic_api_version added to request parameters for each API route
  - /s/{spaceId}/ added as API path prefix for each API
- `apiKeyAuth` added to top-level `security` config
- Changed RuleSource from `oneOf` to `anyOf` because otherwise the generated validator in go rejects inputs - inputs to the schema match both schemas, but `oneOf` technically should reject inputs that match multiple schemas

## Running changes locally

- Create .terraformrc file in home directory with

```
provider_installation {
  filesystem_mirror {
    path    = "/home/user/.terraform.d/plugins"
  }
  direct {
    exclude = ["terraform.local/*/*"]
  }
}
```

Then after running `make install` per the repo readme terraform will use the local provider. Maybe have to bump the version?
