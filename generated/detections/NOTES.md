## Differences between bundled.yaml and OpenAPI spec in Kibana repo

- Extra tags are removed from API route definitions. Code gen duplicates code when it sees multiple tags, causing compilation problems.
  - "API" postfix on tag names is removed so generated code doesn't duplicate the "API" string in names
- Many routes are removed

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
