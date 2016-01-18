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
