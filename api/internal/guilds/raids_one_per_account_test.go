package guilds

import "testing"

// One row per account (raids spec §3.7): the account's best place leads, and its other answered
// characters come back best place first for the row's Also line, never as rows of their own.
func TestOnePerAccountKeepsTheOtherCharactersForTheAlsoLine(t *testing.T) {
	signups := []Signup{
		{CharacterKey: "rootbound", Name: "Rootbound", Class: "druid", Spec: "Restoration", Role: RoleHealer, State: SignupBench, SignedBy: 1},
		{CharacterKey: "brew", Name: "Brew", Class: "mage", Spec: "Fire", Role: RoleDPS, State: SignupOut, SignedBy: 1},
		{CharacterKey: "quill", Name: "Quillfeather", Class: "priest", Spec: "Holy", Role: RoleHealer, State: SignupLate, SignedBy: 1},
		{CharacterKey: "bow", Name: "Bow Jackzon", Class: "hunter", Spec: "Beast Mastery", Role: RoleDPS, State: SignupIn, SignedBy: 2},
	}
	shown, also := onePerAccount(signups)
	if len(shown) != 2 || shown[0].CharacterKey != "quill" || shown[1].CharacterKey != "bow" {
		t.Fatalf("shown = %+v, want Quillfeather then Bow Jackzon", shown)
	}
	got := alsoSignups(also[1])
	if len(got) != 2 || got[0].Name != "Rootbound" || got[0].State != SignupBench || got[1].Name != "Brew" {
		t.Fatalf("also = %+v, want Rootbound (Bench) then Brew (Out)", got)
	}
	if got[0].SpecKey == nil || *got[0].SpecKey != "druid-restoration" {
		t.Fatalf("Rootbound's spec_key = %v, want druid-restoration", got[0].SpecKey)
	}
	if none := alsoSignups(also[2]); none == nil || len(none) != 0 {
		t.Fatalf("a one-character account's also = %#v, want an empty list", none)
	}
}
