# Go Api Client for onetimesecret.com

![onetimesecret](https://github.com/chelnak/onetimesecret-go/actions/workflows/main.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/chelnak/onetimesecret-go)](https://goreportcard.com/report/github.com/chelnak/onetimesecret-go)

This is a Go implementation of the [onetimesecret.com api](https://onetimesecret.com/docs/api) api. The module aims to provide a developer friendly interface for all endpoints exposed by the api.

## Installation

```go
go get github.com/chelnak/onetimesecret-go
```

## Usage

```go
package main

import (
 "context"
 "fmt"

 ots "github.com/chelnak/onetimesecret-go/onetimesecret"
)

func main() {

 // Build a new client
 client := ots.NewClient(
  ots.WithUsername("otsuser@domain.com"),
  ots.WithAPIKey("xxxxxxxx"),
 )

 // Send a request with context
 ctx := context.Background()
 response, err := client.GetStatus(ctx)
 if err != nil {
  panic(err)
 }

 fmt.Println(response.Status)

}
```

### Using a custom http.Client instance

```go
package main

import (
    ots "github.com/chelnak/onetimesecret-go/onetimesecret"
    "context"
)

func main() {

    // Create a custom http client instance
    customHttpClient := &http.Client{ ... }

    // Build a new ots client and use WithHttpClient option
    client := ots.NewClient(
        ots.WithUsername("otsuser@domain.com"),
        ots.WithApiKey("xxxxxxxx"),
        ots.WithHttpClient(customHttpClient)
    )

    // Send a request with context
    ctx := context.Background()
    response, err := client.GetStatus(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println(response.Status)
}
```

> Note: By default, a default http.Client instance is used if one is not passed when creating a client instance.

More documentation can be found [here](https://pkg.go.dev/github.com/chelnak/onetimesecret-go).

## Testing

```go
go test ./... -v
```

### Running example tests

```go
export OTS_USERNAME=otsuser@domain.com
export OTS_APIKEY=xxxxx

go test ./... -timeout 30s -run ^Example_shareJourney$
go test ./... -timeout 30s -run ^Example_generateJourney$
```

## Releasing

Releases are tag based. From the main branch do the following:

```go
git tag -a vX.X.X -m "Release vX.X.X some reason"
git push --follow-tags
```
