package main

import (
	"fmt"
	"os"

	"github.com/patrickrand/swagdown/markdown"
)

func main() {
	f, err := os.Open("./petstore/petstore-expanded.json")
	if err != nil {
		exit(err)
	}

	if err := markdown.Render(os.Stdout, f); err != nil {
		exit(err)
	}
}

func exit(err error) {
	fmt.Println("error: ", err)
	os.Exit(1)
}
