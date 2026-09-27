package bnetimport

import (
	"context"
	"testing"
)

func TestStubImportsNothingAndSaysSo(t *testing.T) {
	t.Parallel()
	stub := &Stub{Regions: []string{"us", "eu"}}
	summary, err := stub.ImportAccount(context.Background(), 1, "token")
	if err != nil {
		t.Fatalf("stub returned an error: %v", err)
	}
	if summary.Characters != 0 || summary.Guilds != 0 {
		t.Fatalf("stub imported something: %+v", summary)
	}
	if summary.Unavailable != 2 || len(summary.Regions) != 2 {
		t.Fatalf("stub should report every region unavailable: %+v", summary)
	}
}
