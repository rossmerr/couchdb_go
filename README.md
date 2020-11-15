# couchdb go client

CouchDB go client auto generated from a OpenAPI v2 schema.

## Sample

Import couchdb_go

``` go
import (
    apiclient "github.com/rossmerr/couchdb_go/client"
    httptransport "github.com/go-openapi/runtime/client"
)
```

Create a Client

``` go
transport := httptransport.New("localhost:5984", "", nil)
transport.DefaultAuthentication = httptransport.BasicAuth("username", "password")
client := apiclient.New(transport, strfmt.Default)
```    

Get a docuemnt

``` go
params := document.NewDocGetParams()
params.WithDb("db").WithDocid("docid")

// GET /{db}/{docid}
response, err := client.Document.DocGet(params)
if err != nil {
    return nil, fmt.Errorf("couchdb not found: %w", err)
}
```


## Using swagger to generate a client

`swagger generate client -f RossMerr-CouchDB-3.1.1-resolved.json -A couchdb_go`


## Contribute to the Schema

[SwaggerHub](https://app.swaggerhub.com/apis/RossMerr/CouchDB/3.1.1)

[API Docs](https://app.swaggerhub.com/apis-docs/RossMerr/CouchDB/3.1.1)
