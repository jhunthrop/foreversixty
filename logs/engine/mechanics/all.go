package mechanics

import (
	"fmt"
	"io/fs"
)

// All is every embedded encounter table, in file order.
func All() ([]Table, error) {
	entries, err := fs.ReadDir(tables, "tables")
	if err != nil {
		return nil, fmt.Errorf("mechanics: list tables: %w", err)
	}
	out := make([]Table, 0, len(entries))
	for _, e := range entries {
		data, err := tables.ReadFile("tables/" + e.Name())
		if err != nil {
			return nil, fmt.Errorf("mechanics: read %s: %w", e.Name(), err)
		}
		t, err := Parse(data)
		if err != nil {
			return nil, fmt.Errorf("mechanics: parse %s: %w", e.Name(), err)
		}
		out = append(out, t)
	}
	return out, nil
}
