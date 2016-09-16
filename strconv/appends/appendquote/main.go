package main

import (
	"fmt"
	"strconv"
)

func main() {
	b := []byte("quote:")

	// AppendQuote appends a double-quoted Go string literal
	// representing s, as generated by Quote, to dst and returns
	// the extended buffer.
	b = strconv.AppendQuote(b, `"I'm inside of quotes", "stuff"`)

	fmt.Println(string(b))
}