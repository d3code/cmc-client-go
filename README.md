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

### Testing

To use the CMC testnet, call `Test(true)` on the client. This will replace the Client ID and the Base URL for the requests to use the sandbox API.

You do not need to provide your Client ID, though this can be provided and calling `Test(false)` will have no effect. This has been done to simplify setting this via config.

```go
cmc.Client("").Test(true)
```

### Logging

There is no logging in this library, an `error` is returned if there are request / unmarshalling errors and can be handled accordingly.

To output the raw response's, call `PrintResponse(true)` on the client.

```go
cmc.Client("").PrintResponse(true)
```