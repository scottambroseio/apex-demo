package main

import (
	"encoding/json"
	"fmt"

	apex "github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		fmt.Println("Hello world")

		return "Success", nil
	})
}
