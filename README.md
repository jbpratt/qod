[![GoDoc](https://godoc.org/github.com/jbpratt78/qod?status.svg)](https://godoc.org/github.com/jbpratt78/qod)
# Quote of the Day API wrapper

### usage

```go
package main

import (
	"fmt"

	"github.com/jbpratt78/qod"
)

func main() {
	client := qod.Client{}

	res, err := client.GetQuoteOfTheDay()
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
```
