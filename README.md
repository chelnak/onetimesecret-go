# Go Api Client for onetimesecret.com

![onetimesecret](https://github.com/chelnak/onetimesecret-go/actions/workflows/main.yml/badge.svg)

This is a Go implementation of the [onetimesecret.com api](https://onetimesecret.com/docs/api) api. The module aims to provide a developer friendly interface for all endpoints exposed by the api.

## Installation

```go
go get github.com/chelnak/onetimesecret-go/onetimesecret
```

## Usage

```go
package main

import (
    ots "github.com/chelnak/onetimesecret-go/onetimesecret"
    "context"
)

func main() {

    // Build a new client
    client := ots.NewClient(
        WithUsername("otsuser@domain.com"),
        WithApiKey("xxxxxxxx"),
    )

    // Send a request with context
    ctx := context.Background()
    response := clent.GetStatus(ctx)

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
        WithUsername("otsuser@domain.com"),
        WithApiKey("xxxxxxxx"),
        WithHttpClient(customHttpClient)
    )

    // Send a request with context
    ctx := context.Background()
    response := clent.GetStatus(ctx)

    fmt.Println(response.Status)
}
```

> Note: By default, a default http.Client instance is used if one is not passed when creating a client instance.

## Testing

```go
go test ./... -v
```
