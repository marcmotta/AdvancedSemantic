// cmd/advancedsemantic/main.go
package main

import (
	"flag"
	"log"
	"os"

	"advancedsemantic/internal/advancedsemantic"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := advancedsemantic.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
