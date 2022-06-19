# Coin Market Cap Client

A Golang client for the Coin Market Cap API

### Install

```shell
go get github.com/d3code/cmc-client-go
```

### Import
```go
import (
    "github.com/d3code/cmc-client-go/cmc"
)
```

## Client

To use the client, create a CMC Client

```go
package main

import (
	"github.com/d3code/cmc-client-go/cmc"
)

func main() {
	client := cmc.Client("<your-api-key>")
}
```