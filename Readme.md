# forismatic.go

## Summary

Forismatic Public API implementation in Go. Not complete. Works for me. No warrenty expressed, given, or implied.

### Why?

I like happy quotes.

## Usage

```go
package main

import (
	"fmt"

	"github.com/dinedal/forismatic"
)

func main() {
	quote, err := forismatic.GetQuote()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(quote)
}
```

## License
MIT
