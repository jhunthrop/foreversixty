// Package strcase holds the one spelling of the protobuf-name-to-id
// conversion the module needs.
//
// Two lanes turn an engine enum value name into the lower-snake
// vocabulary the settings bar and the report speak: sim/request maps a
// buff or consumable id onto its proto field, and sim/adapter names a
// non-spell action. They have to agree - the same enum value reaching
// the browser under two spellings is two rows nothing can join - so the
// conversion is written once.
package strcase

import (
	"strings"
	"unicode"
)

// Snake turns an upper-camel protobuf name into lower snake case:
// ElixirOfTheMongoose becomes elixir_of_the_mongoose. A leading capital
// starts no underscore, and a name that is already lower case, or that
// carries digits, is returned unchanged apart from the separators.
func Snake(name string) string {
	var b strings.Builder
	b.Grow(len(name) + 4)
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i > 0 {
				b.WriteByte('_')
			}
			b.WriteRune(unicode.ToLower(r))
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}
