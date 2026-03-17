package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	source := strings.NewReader("hello world")

	upper := &UpperReader{inner: source}

	counter := &WordCountReader{inner: upper}

	pipeline := &ROT13Reader{inner: counter}

	result, err := io.ReadAll(pipeline)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("input:      %q\n", "hello world")
	fmt.Printf("after upper + rot13: %q\n", string(result))
	fmt.Printf("word count: %d\n", counter.Count)
}
