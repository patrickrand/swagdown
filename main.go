package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/patrickrand/swagdown/markdown"
)

func main() {
	if len(os.Args) != 2 {
		exit(errors.New("swagdown requires exactly one argument, a path to a local swagger.json file"))
	}

	filename := os.Args[1]
	if filename == "test" {
		filename = "./swagger/petstore/petstore-expanded.json"
	}

	f, err := os.Open(filename)
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
