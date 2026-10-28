package auth

import "testing"

func TestCharacterLevelsAnswersOnlyCharactersWithALevel(t *testing.T) {
	store := &Store{Pool: testPool(t)}
	for _, row := range []struct {
		key, name string
		level     any
	}{
		{"us/normal/lvl-sixty", "Lvlsixty", 60},
		{"us/normal/lvl-none", "Lvlnone", nil},
	} {
		if _, err := store.Pool.Exec(t.Context(),
			`insert into characters (key, region, ruleset, name, level) values ($1, 'us', 'normal', $2, $3)
			 on conflict (key) do update set level = excluded.level`, row.key, row.name, row.level); err != nil {
			t.Fatal(err)
		}
	}
	got, err := store.CharacterLevels(t.Context(),
		[]string{"us/normal/lvl-sixty", "us/normal/lvl-none", "us/normal/never-seen"})
	if err != nil {
		t.Fatal(err)
	}
	if got["us/normal/lvl-sixty"] != 60 || len(got) != 1 {
		t.Fatalf("levels = %v, want only the level-60 character", got)
	}
	empty, err := store.CharacterLevels(t.Context(), nil)
	if err != nil || len(empty) != 0 {
		t.Fatalf("no keys: %v %v", empty, err)
	}
}
