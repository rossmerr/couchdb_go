# \DocumentApi

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DocDelete**](DocumentApi.md#DocDelete) | **Delete** /{db}/{docid} | Marks the specified document as deleted by adding a field _deleted with the value true.
[**DocGet**](DocumentApi.md#DocGet) | **Get** /{db}/{docid} | Returns document by the specified docid from the specified db. Unless you request a specific revision, the latest revision of the document will always be returned.
[**DocInfo**](DocumentApi.md#DocInfo) | **Head** /{db}/{docid} | Returns the HTTP Headers containing a minimal amount of information about the specified document.
[**DocPut**](DocumentApi.md#DocPut) | **Put** /{db}/{docid} | The PUT method creates a new named document, or creates a new revision of the existing document. Unlike the POST /{db}, you must specify the document ID in the request URL.
[**Post**](DocumentApi.md#Post) | **Post** /{db} | Creates a new document in the specified database, using the supplied JSON document structure.



## DocDelete

> DocumentOK DocDelete(ctx, db, docid).IfMatch(ifMatch).Rev(rev).Batch(batch).Execute()

Marks the specified document as deleted by adding a field _deleted with the value true.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    docid := "docid_example" // string | DDocument ID
    ifMatch := "ifMatch_example" // string | Document’s revision. Alternative to rev query parameter (optional)
    rev := "rev_example" // string | Actual document’s revision (optional)
    batch := "batch_example" // string | Stores document in batch mode Possible values: ok. Optional  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DocDelete(context.Background(), db, docid).IfMatch(ifMatch).Rev(rev).Batch(batch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DocDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocDelete`: DocumentOK
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DocDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**docid** | **string** | DDocument ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Document’s revision. Alternative to rev query parameter | 
 **rev** | **string** | Actual document’s revision | 
 **batch** | **string** | Stores document in batch mode Possible values: ok. Optional  | 

### Return type

[**DocumentOK**](DocumentOK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocGet

> map[string]interface{} DocGet(ctx, db, docid).IfNoneMatch(ifNoneMatch).Attachments(attachments).AttEncodingInfo(attEncodingInfo).AttsSince(attsSince).Conflicts(conflicts).DeletedConflicts(deletedConflicts).Latest(latest).LocalSeq(localSeq).Meta(meta).OpenRevs(openRevs).Rev(rev).Revs(revs).RevsInfo(revsInfo).Execute()

Returns document by the specified docid from the specified db. Unless you request a specific revision, the latest revision of the document will always be returned.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    docid := "docid_example" // string | DDocument ID
    ifNoneMatch := "ifNoneMatch_example" // string | Double quoted document’s revision token
    attachments := true // bool | Includes attachments bodies in response. Default is false (optional)
    attEncodingInfo := true // bool | Includes encoding information in attachment stubs if the particular attachment is compressed. Default is false. (optional)
    attsSince := []string{"Inner_example"} // []string | Includes attachments only since specified revisions. Doesn’t includes attachments for specified revisions. Optional (optional)
    conflicts := true // bool | Includes information about conflicts in document. Default is false (optional)
    deletedConflicts := true // bool | Includes information about deleted conflicted revisions. Default is false (optional)
    latest := true // bool | Forces retrieving latest “leaf” revision, no matter what rev was requested. Default is false (optional)
    localSeq := true // bool | Includes last update sequence for the document. Default is false (optional)
    meta := true // bool | Acts same as specifying all conflicts, deleted_conflicts and revs_info query parameters. Default is false (optional)
    openRevs := []string{"Inner_example"} // []string | Retrieves documents of specified leaf revisions. Additionally, it accepts value as all to return all leaf revisions. Optional (optional)
    rev := "rev_example" // string | Retrieves document of specified revision. Optional (optional)
    revs := true // bool | Includes list of all known document revisions. Default is false (optional)
    revsInfo := true // bool | Includes detailed information for all known document revisions. Default is false (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DocGet(context.Background(), db, docid).IfNoneMatch(ifNoneMatch).Attachments(attachments).AttEncodingInfo(attEncodingInfo).AttsSince(attsSince).Conflicts(conflicts).DeletedConflicts(deletedConflicts).Latest(latest).LocalSeq(localSeq).Meta(meta).OpenRevs(openRevs).Rev(rev).Revs(revs).RevsInfo(revsInfo).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DocGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocGet`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DocGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**docid** | **string** | DDocument ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifNoneMatch** | **string** | Double quoted document’s revision token | 
 **attachments** | **bool** | Includes attachments bodies in response. Default is false | 
 **attEncodingInfo** | **bool** | Includes encoding information in attachment stubs if the particular attachment is compressed. Default is false. | 
 **attsSince** | **[]string** | Includes attachments only since specified revisions. Doesn’t includes attachments for specified revisions. Optional | 
 **conflicts** | **bool** | Includes information about conflicts in document. Default is false | 
 **deletedConflicts** | **bool** | Includes information about deleted conflicted revisions. Default is false | 
 **latest** | **bool** | Forces retrieving latest “leaf” revision, no matter what rev was requested. Default is false | 
 **localSeq** | **bool** | Includes last update sequence for the document. Default is false | 
 **meta** | **bool** | Acts same as specifying all conflicts, deleted_conflicts and revs_info query parameters. Default is false | 
 **openRevs** | **[]string** | Retrieves documents of specified leaf revisions. Additionally, it accepts value as all to return all leaf revisions. Optional | 
 **rev** | **string** | Retrieves document of specified revision. Optional | 
 **revs** | **bool** | Includes list of all known document revisions. Default is false | 
 **revsInfo** | **bool** | Includes detailed information for all known document revisions. Default is false | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, multipart/related, multipart/mixed, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocInfo

> DocInfo(ctx, db, docid).Execute()

Returns the HTTP Headers containing a minimal amount of information about the specified document.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    docid := "docid_example" // string | DDocument ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DocumentApi.DocInfo(context.Background(), db, docid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DocInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**docid** | **string** | DDocument ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DocPut

> DocumentOK DocPut(ctx, db, docid).Body(body).IfMatch(ifMatch).Rev(rev).Batch(batch).NewEdits(newEdits).Execute()

The PUT method creates a new named document, or creates a new revision of the existing document. Unlike the POST /{db}, you must specify the document ID in the request URL.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    docid := "docid_example" // string | DDocument ID
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    ifMatch := "ifMatch_example" // string | Document’s revision. Alternative to rev query parameter or document key. Optional (optional)
    rev := "rev_example" // string | Document’s revision if updating an existing document. Alternative to If-Match header or document key. Optional (optional)
    batch := "batch_example" // string | Stores document in batch mode. Possible values: ok. Optional  (optional)
    newEdits := true // bool | Prevents insertion of a conflicting document. Possible values: true (default) and false. If false,  a well-formed _rev must be included in the document. new_edits=false is used by the replicator  to insert documents into the target database even if that leads to the creation of conflicts. Optional  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.DocPut(context.Background(), db, docid).Body(body).IfMatch(ifMatch).Rev(rev).Batch(batch).NewEdits(newEdits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.DocPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DocPut`: DocumentOK
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.DocPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 
**docid** | **string** | DDocument ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDocPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 
 **ifMatch** | **string** | Document’s revision. Alternative to rev query parameter or document key. Optional | 
 **rev** | **string** | Document’s revision if updating an existing document. Alternative to If-Match header or document key. Optional | 
 **batch** | **string** | Stores document in batch mode. Possible values: ok. Optional  | 
 **newEdits** | **bool** | Prevents insertion of a conflicting document. Possible values: true (default) and false. If false,  a well-formed _rev must be included in the document. new_edits&#x3D;false is used by the replicator  to insert documents into the target database even if that leads to the creation of conflicts. Optional  | 

### Return type

[**DocumentOK**](DocumentOK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Post

> DocumentOK Post(ctx, db).Body(body).Batch(batch).Execute()

Creates a new document in the specified database, using the supplied JSON document structure.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
    db := "db_example" // string | Database name
    body := map[string]interface{}{ ... } // map[string]interface{} | 
    batch := "batch_example" // string | Stores document in batch mode Possible values: ok. Optional  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DocumentApi.Post(context.Background(), db).Body(body).Batch(batch).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DocumentApi.Post``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Post`: DocumentOK
    fmt.Fprintf(os.Stdout, "Response from `DocumentApi.Post`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**db** | **string** | Database name | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **batch** | **string** | Stores document in batch mode Possible values: ok. Optional  | 

### Return type

[**DocumentOK**](DocumentOK.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, text/plain
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

