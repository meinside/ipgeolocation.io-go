# ipgeolocation.io-go

A Go library for [ipgeolocation](https://ipgeolocation.io/) APIs.

## Usage

```go
package main

import (
    "fmt"

    ipgeolocation "github.com/meinside/ipgeolocation.io-go"
)

func main() {
    apiKey := "xxxxxxxxxxxxxxxxxx"

    client := ipgeolocation.NewClient(apiKey)
    if result, err := client.GetGeolocation("8.8.8.8"); err == nil {
        fmt.Printf("country name = %s\n", result.CountryName)
    }
}
```

## Test

```bash
$ API_KEY=xxxxxxxxxxxxxxxxxx go test
```

## Todos

- [ ] Implement [Timezone API](https://ipgeolocation.io/documentation/timezone-api.html) functions.
- [ ] Implement [User-Agent API](https://ipgeolocation.io/documentation/user-agent-api.html) functions.
- [ ] Implement [Astronomy API](https://ipgeolocation.io/documentation/astronomy-api.html) functions.
- [ ] Implement functionalities available only on paid plans.

## License

MIT

