package main

import (
	"encoding/json"
	"fmt"
	"os"

	apex "github.com/apex/go-apex"
)

// stdout is reserved for the node shim so all messages to the console
// need to go through stderr
func printline(message string) {
	fmt.Fprintln(os.Stderr, message)
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		printline("Hello world")

		return "Success", nil
	})
}
