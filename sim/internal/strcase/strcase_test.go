package strcase

import "testing"

func TestSnakeMatchesTheSettingsBarVocabulary(t *testing.T) {
	cases := map[string]string{
		"ElixirOfTheMongoose": "elixir_of_the_mongoose",
		"Attack":              "attack",
		"attack":              "attack",
		"":                    "",
		"A":                   "a",
		"MP5":                 "m_p5",
		"RallyingCry":         "rallying_cry",
	}
	for in, want := range cases {
		if got := Snake(in); got != want {
			t.Errorf("Snake(%q) = %q, want %q", in, got, want)
		}
	}
}
