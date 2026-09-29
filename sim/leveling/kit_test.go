package leveling

import "testing"

func TestKitConsumesIsRoguePoisonsFromTwenty(t *testing.T) {
	if got := KitConsumes("rogue", 19); got != nil {
		t.Fatalf("rogue at 19 = %v, want nothing (the poison quest is level 20)", got)
	}
	if got := KitConsumes("rogue", 20); len(got) != 2 {
		t.Fatalf("rogue at 20 = %v, want Instant Poison on both weapons", got)
	}
	if got := KitConsumes("warrior", 60); got != nil {
		t.Fatalf("warrior = %v, want nothing", got)
	}
}
