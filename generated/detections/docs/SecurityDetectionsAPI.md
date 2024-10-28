# \SecurityDetectionsAPI

All URIs are relative to *http://localhost:5601*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkPatchRules**](SecurityDetectionsAPI.md#BulkPatchRules) | **Patch** /api/detection_engine/rules/_bulk_action | Patch multiple detection rules
[**BulkUpdateRules**](SecurityDetectionsAPI.md#BulkUpdateRules) | **Put** /api/detection_engine/rules/_bulk_action | Update multiple detection rules
[**CreateRule**](SecurityDetectionsAPI.md#CreateRule) | **Post** /s/{spaceId}/api/detection_engine/rules | Create a detection rule
[**DeleteRule**](SecurityDetectionsAPI.md#DeleteRule) | **Delete** /s/{spaceId}/api/detection_engine/rules | Delete a detection rule
[**PatchRule**](SecurityDetectionsAPI.md#PatchRule) | **Patch** /s/{spaceId}/api/detection_engine/rules | Patch a detection rule
[**ReadRule**](SecurityDetectionsAPI.md#ReadRule) | **Get** /s/{spaceId}/api/detection_engine/rules | Retrieve a detection rule
[**ReadTags**](SecurityDetectionsAPI.md#ReadTags) | **Get** /api/detection_engine/rules/_bulk_action | List all detection rule tags
[**UpdateRule**](SecurityDetectionsAPI.md#UpdateRule) | **Put** /s/{spaceId}/api/detection_engine/rules | Update a detection rule



## BulkPatchRules

> []BulkCrudRulesResponseInner BulkPatchRules(ctx).KbnXsrf(kbnXsrf).RulePatchProps(rulePatchProps).ElasticApiVersion(elasticApiVersion).Execute()

Patch multiple detection rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    rulePatchProps := []openapiclient.RulePatchProps{*openapiclient.NewRulePatchProps()} // []RulePatchProps | A JSON array of rules, where each rule contains the required fields.
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.BulkPatchRules(context.Background()).KbnXsrf(kbnXsrf).RulePatchProps(rulePatchProps).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.BulkPatchRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkPatchRules`: []BulkCrudRulesResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.BulkPatchRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkPatchRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **rulePatchProps** | [**[]RulePatchProps**](RulePatchProps.md) | A JSON array of rules, where each rule contains the required fields. | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**[]BulkCrudRulesResponseInner**](BulkCrudRulesResponseInner.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkUpdateRules

> []BulkCrudRulesResponseInner BulkUpdateRules(ctx).KbnXsrf(kbnXsrf).RuleUpdateProps(ruleUpdateProps).ElasticApiVersion(elasticApiVersion).Execute()

Update multiple detection rules



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    ruleUpdateProps := []openapiclient.RuleUpdateProps{*openapiclient.NewRuleUpdateProps("Description_example", "Name_example", int32(123), openapiclient.Severity("low"), openapiclient.EsqlQueryLanguage("esql"), "Query_example", "Type_example", "SavedId_example", *openapiclient.NewThreshold(openapiclient.ThresholdField{ArrayOfString: new([]string)}, int32(123)), []string{"ThreatIndex_example"}, []openapiclient.ThreatMappingInner{*openapiclient.NewThreatMappingInner([]openapiclient.ThreatMappingInnerEntriesInner{*openapiclient.NewThreatMappingInnerEntriesInner("Field_example", "Type_example", "Value_example")})}, "ThreatQuery_example", int32(123), openapiclient.MachineLearningJobId{ArrayOfString: new([]string)}, "HistoryWindowStart_example", []string{"NewTermsFields_example"})} // []RuleUpdateProps | A JSON array where each element includes the `id` or `rule_id` field of the rule you want to update and the fields you want to modify.
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.BulkUpdateRules(context.Background()).KbnXsrf(kbnXsrf).RuleUpdateProps(ruleUpdateProps).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.BulkUpdateRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkUpdateRules`: []BulkCrudRulesResponseInner
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.BulkUpdateRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkUpdateRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **ruleUpdateProps** | [**[]RuleUpdateProps**](RuleUpdateProps.md) | A JSON array where each element includes the &#x60;id&#x60; or &#x60;rule_id&#x60; field of the rule you want to update and the fields you want to modify. | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**[]BulkCrudRulesResponseInner**](BulkCrudRulesResponseInner.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRule

> RuleResponse CreateRule(ctx, spaceId).KbnXsrf(kbnXsrf).RuleCreateProps(ruleCreateProps).ElasticApiVersion(elasticApiVersion).Execute()

Create a detection rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    spaceId := "spaceId_example" // string | 
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    ruleCreateProps := *openapiclient.NewRuleCreateProps("Description_example", "Name_example", int32(123), openapiclient.Severity("low"), openapiclient.EsqlQueryLanguage("esql"), "Query_example", "Type_example", "SavedId_example", *openapiclient.NewThreshold(openapiclient.ThresholdField{ArrayOfString: new([]string)}, int32(123)), []string{"ThreatIndex_example"}, []openapiclient.ThreatMappingInner{*openapiclient.NewThreatMappingInner([]openapiclient.ThreatMappingInnerEntriesInner{*openapiclient.NewThreatMappingInnerEntriesInner("Field_example", "Type_example", "Value_example")})}, "ThreatQuery_example", int32(123), openapiclient.MachineLearningJobId{ArrayOfString: new([]string)}, "HistoryWindowStart_example", []string{"NewTermsFields_example"}) // RuleCreateProps | 
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.CreateRule(context.Background(), spaceId).KbnXsrf(kbnXsrf).RuleCreateProps(ruleCreateProps).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.CreateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **ruleCreateProps** | [**RuleCreateProps**](RuleCreateProps.md) |  | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> RuleResponse DeleteRule(ctx, spaceId).KbnXsrf(kbnXsrf).Id(id).RuleId(ruleId).ElasticApiVersion(elasticApiVersion).Execute()

Delete a detection rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    spaceId := "spaceId_example" // string | 
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The rule's `id` value. (optional)
    ruleId := "ruleId_example" // string | The rule's `rule_id` value. (optional)
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.DeleteRule(context.Background(), spaceId).KbnXsrf(kbnXsrf).Id(id).RuleId(ruleId).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.DeleteRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **id** | **string** | The rule&#39;s &#x60;id&#x60; value. | 
 **ruleId** | **string** | The rule&#39;s &#x60;rule_id&#x60; value. | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRule

> RuleResponse PatchRule(ctx, spaceId).KbnXsrf(kbnXsrf).RulePatchProps(rulePatchProps).ElasticApiVersion(elasticApiVersion).Execute()

Patch a detection rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    spaceId := "spaceId_example" // string | 
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    rulePatchProps := *openapiclient.NewRulePatchProps() // RulePatchProps | 
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.PatchRule(context.Background(), spaceId).KbnXsrf(kbnXsrf).RulePatchProps(rulePatchProps).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.PatchRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PatchRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.PatchRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **rulePatchProps** | [**RulePatchProps**](RulePatchProps.md) |  | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadRule

> RuleResponse ReadRule(ctx, spaceId).KbnXsrf(kbnXsrf).Id(id).RuleId(ruleId).ElasticApiVersion(elasticApiVersion).Execute()

Retrieve a detection rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    spaceId := "spaceId_example" // string | 
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The rule's `id` value. (optional)
    ruleId := "ruleId_example" // string | The rule's `rule_id` value. (optional)
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.ReadRule(context.Background(), spaceId).KbnXsrf(kbnXsrf).Id(id).RuleId(ruleId).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.ReadRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.ReadRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **id** | **string** | The rule&#39;s &#x60;id&#x60; value. | 
 **ruleId** | **string** | The rule&#39;s &#x60;rule_id&#x60; value. | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadTags

> []string ReadTags(ctx).KbnXsrf(kbnXsrf).ElasticApiVersion(elasticApiVersion).Execute()

List all detection rule tags



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.ReadTags(context.Background()).KbnXsrf(kbnXsrf).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.ReadTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadTags`: []string
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.ReadTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

**[]string**

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRule

> RuleResponse UpdateRule(ctx, spaceId).KbnXsrf(kbnXsrf).RuleUpdateProps(ruleUpdateProps).ElasticApiVersion(elasticApiVersion).Execute()

Update a detection rule



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/elastic/terraform-provider-elasticstack/detections"
)

func main() {
    spaceId := "spaceId_example" // string | 
    kbnXsrf := "kbnXsrf_example" // string | Cross-site request forgery protection
    ruleUpdateProps := *openapiclient.NewRuleUpdateProps("Description_example", "Name_example", int32(123), openapiclient.Severity("low"), openapiclient.EsqlQueryLanguage("esql"), "Query_example", "Type_example", "SavedId_example", *openapiclient.NewThreshold(openapiclient.ThresholdField{ArrayOfString: new([]string)}, int32(123)), []string{"ThreatIndex_example"}, []openapiclient.ThreatMappingInner{*openapiclient.NewThreatMappingInner([]openapiclient.ThreatMappingInnerEntriesInner{*openapiclient.NewThreatMappingInnerEntriesInner("Field_example", "Type_example", "Value_example")})}, "ThreatQuery_example", int32(123), openapiclient.MachineLearningJobId{ArrayOfString: new([]string)}, "HistoryWindowStart_example", []string{"NewTermsFields_example"}) // RuleUpdateProps | 
    elasticApiVersion := "elasticApiVersion_example" // string | Api version (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SecurityDetectionsAPI.UpdateRule(context.Background(), spaceId).KbnXsrf(kbnXsrf).RuleUpdateProps(ruleUpdateProps).ElasticApiVersion(elasticApiVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityDetectionsAPI.UpdateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRule`: RuleResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityDetectionsAPI.UpdateRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**spaceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kbnXsrf** | **string** | Cross-site request forgery protection | 
 **ruleUpdateProps** | [**RuleUpdateProps**](RuleUpdateProps.md) |  | 
 **elasticApiVersion** | **string** | Api version | 

### Return type

[**RuleResponse**](RuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth), [apiKeyAuth](../README.md#apiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

