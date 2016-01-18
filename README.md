# OFX [![Build Status](https://drone.io/github.com/luizbranco/ofx/status.png)](https://drone.io/github.com/luizbranco/ofx/latest) [![GoDoc](https://godoc.org/github.com/luizbranco/ofx?status.svg)](https://godoc.org/github.com/luizbranco/ofx)

Open Financial Exchange (OFX) response reader in go. Compatible with versions 1.02 and
1.03.

## Response Services

- [ ] Sign On
- [x] Banking
- [x] Credit Card
- [ ] Payments
- [ ] Investments

## Example

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/luizbranco/ofx"
)

func main() {
	log.SetFlags(0)
	for _, a := range os.Args[1:] {
		f, err := os.Open(a)
		if err != nil {
			log.Fatalf("error reading file %s, %s", a, err)
		}
		t, err := ofx.Parse(f)
		if err != nil {
			log.Fatalf("error parsing file %s, %s", a, err)
		}
		fmt.Printf("%s", t)
	}
}
```

```shell
go run main.go example.ofx
```
