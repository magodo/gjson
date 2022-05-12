package main

import (
	"fmt"
	"io"
	"os"

	"github.com/tidwall/gjson"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" {
		fmt.Fprintln(os.Stderr, "usage: gjson <path>")
		os.Exit(0)
	}

	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading from stdin: %v\n", err)
		os.Exit(1)
	}

	result := gjson.GetBytes(b, os.Args[1])
	fmt.Println(result.String())
}
