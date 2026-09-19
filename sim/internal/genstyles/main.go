// Command genstyles writes sim/request/styles.json, the fight-style
// table the web lane's own style table is tested against.
//
//	go run ./internal/genstyles
//
// TestStylesJSONIsCommitted fails when the committed file and this
// output disagree, so retuning a preset in Go and forgetting the page
// is a failing test rather than two products disagreeing quietly.
package main

import (
	"log"
	"os"

	"github.com/jhunthrop/foreversixty/sim/request"
)

func main() {
	out := "request/styles.json"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	b, err := request.StylesJSON()
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile(out, b, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", out)
}
