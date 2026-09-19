// Command genids writes sim/request/IDS.md, the buff and consumable id
// vocabulary the settings bar sends, from the engine's own protobuf
// descriptors.
//
//	go run ./internal/genids
//
// TestIDsMarkdownIsCommitted fails when the committed file and this
// output disagree, so a new flask in the engine shows up as a failing
// test rather than as a settings bar that cannot be built.
package main

import (
	"log"
	"os"

	"github.com/jhunthrop/foreversixty/sim/request"
)

func main() {
	out := "request/IDS.md"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}
	if err := os.WriteFile(out, []byte(request.IDsMarkdown()), 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote %s", out)
}
