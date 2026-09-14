package addon

import (
	"strings"
	"testing"
)

func TestDeeplyNestedTablesAreRefusedRatherThanOverflowingTheStack(t *testing.T) {
	deep := "X = " + strings.Repeat("{", MaxTableDepth+50) +
		strings.Repeat("}", MaxTableDepth+50)
	_, err := ParseLua(deep)
	if err == nil {
		t.Fatal("a file nested past the limit was accepted")
	}
	if !strings.HasPrefix(err.Error(), "lua line ") {
		t.Errorf("the error has no line prefix: %v", err)
	}
	if !strings.Contains(err.Error(), "nested") {
		t.Errorf("the error does not say why: %v", err)
	}

	// The limit is well clear of anything the game writes.
	fine := "X = " + strings.Repeat("{", MaxTableDepth-1) + strings.Repeat("}", MaxTableDepth-1)
	if _, err := ParseLua(fine); err != nil {
		t.Fatalf("a file inside the limit was refused: %v", err)
	}
}
