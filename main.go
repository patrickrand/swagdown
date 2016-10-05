package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/patrickrand/swagdown/markdown"
)

func main() {
	if len(os.Args) != 2 {
		exit(errors.New("swagdown requires exactly one argument, a path to a local swagger.json file"))
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		exit(err)
	}

	switch filepath.Ext(f.Name()) {
	case ".json":
		if err := markdown.RenderFromJSON(os.Stdout, f); err != nil {
			exit(err)
		}
	case ".yml", ".yaml":
		if err := markdown.RenderFromYAML(os.Stdout, f); err != nil {
			exit(err)
		}
	default:
		exit(errors.New("swagdown only recognizes JSON or YAML Swagger files"))
	}

}

func exit(err error) {
	fmt.Println("error: ", err)
	os.Exit(1)
}
