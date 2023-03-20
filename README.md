# Go API client for couchdb_go

*Note*
This is not a definitive implementation of the CouchDB API, it's missing a lot of the endpoints for server/database managment and everything for attachments all COPY operations and probably a few other parts.

It also targets golang, as such the '#/definitions/Document' is intentionally left empty to generate a go `interface`, which you can then cast to a `map[string]interface{}`.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import couchdb_go "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), couchdb_go.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), couchdb_go.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), couchdb_go.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), couchdb_go.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://virtserver.swaggerhub.com/RossMerr/CouchDB/4.0.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DatabaseApi* | [**BulkDocs**](docs/DatabaseApi.md#bulkdocs) | **Post** /{db}/_bulk_docs | The bulk document API allows you to create and update multiple documents at the same time within a single request.
*DatabaseApi* | [**BulkGet**](docs/DatabaseApi.md#bulkget) | **Post** /{db}/_bulk_get | This method can be called to query several documents in bulk.
*DatabaseApi* | [**DbSecurityGet**](docs/DatabaseApi.md#dbsecurityget) | **Get** /{db}/_security | Returns the current security object from the specified database.
*DatabaseApi* | [**Delete**](docs/DatabaseApi.md#delete) | **Delete** /{db} | Deletes the specified database, and all the documents and attachments contained within it.
*DatabaseApi* | [**DesignDocAllGet**](docs/DatabaseApi.md#designdocallget) | **Get** /{db}/_design_docs | Returns a JSON structure of all of the design documents in a given database.
*DatabaseApi* | [**DesignDocAllPost**](docs/DatabaseApi.md#designdocallpost) | **Post** /{db}/_design_docs | POST _design_docs functionality supports identical parameters and behavior as specified in the GET /{db}/_design_docs
*DatabaseApi* | [**DocGetAll**](docs/DatabaseApi.md#docgetall) | **Get** /{db}/_all_docs | Executes the built-in _all_docs view
*DatabaseApi* | [**DocPostAll**](docs/DatabaseApi.md#docpostall) | **Post** /{db}/_all_docs | Executes the built-in _all_docs view
*DatabaseApi* | [**Exists**](docs/DatabaseApi.md#exists) | **Head** /{db} | Returns the HTTP Headers containing a minimal amount of information about the specified database.
*DatabaseApi* | [**Get**](docs/DatabaseApi.md#get) | **Get** /{db} | Gets information about the specified database.
*DatabaseApi* | [**Put**](docs/DatabaseApi.md#put) | **Put** /{db} | Creates a new database.
*DatabaseApi* | [**SbSecurityPut**](docs/DatabaseApi.md#sbsecurityput) | **Put** /{db}/_security | Sets the security object for the given database.
*DesignDocumentsApi* | [**DesignDocDelete**](docs/DesignDocumentsApi.md#designdocdelete) | **Delete** /{db}/_design/{ddoc} | Deletes the specified document from the database. You must supply the current (latest) revision, either by using the rev parameter to specify the revision.
*DesignDocumentsApi* | [**DesignDocExists**](docs/DesignDocumentsApi.md#designdocexists) | **Head** /{db}/_design/{ddoc} | Returns the HTTP Headers containing a minimal amount of information about the specified design document.
*DesignDocumentsApi* | [**DesignDocGet**](docs/DesignDocumentsApi.md#designdocget) | **Get** /{db}/_design/{ddoc} | Returns the contents of the design document specified with the name of the design document and from the specified database from the URL.
*DesignDocumentsApi* | [**DesignDocInfo**](docs/DesignDocumentsApi.md#designdocinfo) | **Head** /{db}/_design/{ddoc}/_info | Obtains information about the specified design document, including the index, index size and current status of the design document and associated index information.
*DesignDocumentsApi* | [**DesignDocPut**](docs/DesignDocumentsApi.md#designdocput) | **Put** /{db}/_design/{ddoc} | The PUT method creates a new named design document, or creates a new revision of the existing design document.
*DesignDocumentsApi* | [**DesignDocSearch**](docs/DesignDocumentsApi.md#designdocsearch) | **Get** /{db}/_design/{ddoc}/_search/{index} | Executes a search request against the named index in the specified design document.
*DesignDocumentsApi* | [**DesignDocSearchInfo**](docs/DesignDocumentsApi.md#designdocsearchinfo) | **Get** /{db}/_design/{ddoc}/_search_info/{index} | Executes a search request against the named index in the specified design document.
*DesignDocumentsApi* | [**DesignDocView**](docs/DesignDocumentsApi.md#designdocview) | **Get** /{db}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
*DesignDocumentsApi* | [**DesignDocViewPost**](docs/DesignDocumentsApi.md#designdocviewpost) | **Post** /{db}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
*DocumentApi* | [**DocDelete**](docs/DocumentApi.md#docdelete) | **Delete** /{db}/{docid} | Marks the specified document as deleted by adding a field _deleted with the value true.
*DocumentApi* | [**DocGet**](docs/DocumentApi.md#docget) | **Get** /{db}/{docid} | Returns document by the specified docid from the specified db. Unless you request a specific revision, the latest revision of the document will always be returned.
*DocumentApi* | [**DocInfo**](docs/DocumentApi.md#docinfo) | **Head** /{db}/{docid} | Returns the HTTP Headers containing a minimal amount of information about the specified document.
*DocumentApi* | [**DocPut**](docs/DocumentApi.md#docput) | **Put** /{db}/{docid} | The PUT method creates a new named document, or creates a new revision of the existing document. Unlike the POST /{db}, you must specify the document ID in the request URL.
*DocumentApi* | [**Post**](docs/DocumentApi.md#post) | **Post** /{db} | Creates a new document in the specified database, using the supplied JSON document structure.
*IndexApi* | [**DbFindGet**](docs/IndexApi.md#dbfindget) | **Post** /{db}/_find | Finds the document.
*IndexApi* | [**DbIndexGet**](docs/IndexApi.md#dbindexget) | **Get** /{db}/_index | Returns the current indexes object from the specified database.
*IndexApi* | [**DbPartitionFindGet**](docs/IndexApi.md#dbpartitionfindget) | **Post** /{db}/_partition/{partition}/_find | Finds the document.
*IndexApi* | [**IndexDelete**](docs/IndexApi.md#indexdelete) | **Delete** /{db}/_index/{designdoc}/json/{name} | 
*IndexApi* | [**SbIndexPost**](docs/IndexApi.md#sbindexpost) | **Post** /{db}/_index | Sets the index for the given database.
*PartitionApi* | [**PartitionDesignDocSearch**](docs/PartitionApi.md#partitiondesigndocsearch) | **Get** /{db}/_partition/{partition}/_design/{ddoc}/_search/{index} | Executes a search request against the named index in the specified design document.
*PartitionApi* | [**PartitionDesignDocView**](docs/PartitionApi.md#partitiondesigndocview) | **Get** /{db}/_partition/{partition}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
*PartitionApi* | [**PartitionDesignDocViewPost**](docs/PartitionApi.md#partitiondesigndocviewpost) | **Post** /{db}/_partition/{partition}/_design/{ddoc}/_view/{view} | Executes the specified view function from the specified design document.
*PartitionApi* | [**PartitionDocGetAll**](docs/PartitionApi.md#partitiondocgetall) | **Get** /{db}/_partition/{partition}/_all_docs | Executes the built-in _all_docs view
*PartitionApi* | [**PartitionInfo**](docs/PartitionApi.md#partitioninfo) | **Get** /{db}/_partition/{partition} | This endpoint returns information describing the provided partition. It includes document and deleted document counts along with external and active data sizes.
*ServerApi* | [**ActiveTasks**](docs/ServerApi.md#activetasks) | **Get** /_active_tasks | List of running tasks, including the task type, name, status and process ID.
*ServerApi* | [**AllDBs**](docs/ServerApi.md#alldbs) | **Get** /_all_dbs | Returns a list of all the databases in the CouchDB instance.
*ServerApi* | [**ClusterSetupGet**](docs/ServerApi.md#clustersetupget) | **Get** /_cluster_setup | Returns the status of the node or cluster, per the cluster setup wizard.
*ServerApi* | [**ClusterSetupPost**](docs/ServerApi.md#clustersetuppost) | **Post** /_cluster_setup | Configure a node as a single (standalone) node, as part of a cluster, or finalise a cluster.
*ServerApi* | [**DBsInfo**](docs/ServerApi.md#dbsinfo) | **Post** /_dbs_info | Returns information of a list of the specified databases in the CouchDB instance.
*ServerApi* | [**Membership**](docs/ServerApi.md#membership) | **Get** /_membership | Displays the nodes that are part of the cluster as cluster_nodes.
*ServerApi* | [**MetaInformation**](docs/ServerApi.md#metainformation) | **Get** / | Accessing the root of a CouchDB instance returns meta information about the instance.
*ServerApi* | [**Replication**](docs/ServerApi.md#replication) | **Post** /_replicate | Request, configure, or stop, a replication operation.
*ServerApi* | [**SearchAnalyze**](docs/ServerApi.md#searchanalyze) | **Post** /_search_analyze | Tests the results of Lucene analyzer tokenization on sample text.
*ServerApi* | [**Up**](docs/ServerApi.md#up) | **Get** /_up | Confirms that the server is up, running, and ready to respond to requests.


## Documentation For Models

 - [Admins](docs/Admins.md)
 - [All](docs/All.md)
 - [BasicDoc](docs/BasicDoc.md)
 - [Body](docs/Body.md)
 - [Body1](docs/Body1.md)
 - [Body2](docs/Body2.md)
 - [Body3](docs/Body3.md)
 - [Body4](docs/Body4.md)
 - [Body5](docs/Body5.md)
 - [BulkResponse](docs/BulkResponse.md)
 - [Cluster](docs/Cluster.md)
 - [Database](docs/Database.md)
 - [DatabaseCluster](docs/DatabaseCluster.md)
 - [DatabaseProps](docs/DatabaseProps.md)
 - [DatabaseSizes](docs/DatabaseSizes.md)
 - [DesignDoc](docs/DesignDoc.md)
 - [DocumentOK](docs/DocumentOK.md)
 - [ErrorBulkGetResponse](docs/ErrorBulkGetResponse.md)
 - [ErrorResponse](docs/ErrorResponse.md)
 - [Index](docs/Index.md)
 - [IndexDefinitions](docs/IndexDefinitions.md)
 - [IndexDefinitionsDef](docs/IndexDefinitionsDef.md)
 - [IndexDefinitionsDefField](docs/IndexDefinitionsDefField.md)
 - [IndexResponse](docs/IndexResponse.md)
 - [Indexes](docs/Indexes.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse2006ExecutionStats](docs/InlineResponse2006ExecutionStats.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [Keys](docs/Keys.md)
 - [Members](docs/Members.md)
 - [OK](docs/OK.md)
 - [Pagination](docs/Pagination.md)
 - [Partition](docs/Partition.md)
 - [PartitionSizes](docs/PartitionSizes.md)
 - [Query](docs/Query.md)
 - [Replicate](docs/Replicate.md)
 - [ReplicateCreateTargetParams](docs/ReplicateCreateTargetParams.md)
 - [Replication](docs/Replication.md)
 - [ReplicationHistory](docs/ReplicationHistory.md)
 - [Request](docs/Request.md)
 - [RequestHeaders](docs/RequestHeaders.md)
 - [Results](docs/Results.md)
 - [ResultsDocs](docs/ResultsDocs.md)
 - [ResultsResults](docs/ResultsResults.md)
 - [Row](docs/Row.md)
 - [SearchIndex](docs/SearchIndex.md)
 - [Server](docs/Server.md)
 - [ServerVendor](docs/ServerVendor.md)
 - [Task](docs/Task.md)
 - [ViewIndex](docs/ViewIndex.md)
 - [ViewIndexSizes](docs/ViewIndexSizes.md)


## Documentation For Authorization



### BasicAuth

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


### JwtAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



