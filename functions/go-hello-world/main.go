package main

import (
	"encoding/json"
	"fmt"
	"os"

	apex "github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		fmt.Fprintln(os.Stderr, "Hello world")

		return "Success", nil
	})
}
