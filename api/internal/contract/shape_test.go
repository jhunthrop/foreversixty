package contract

import (
	"strings"
	"testing"
)

func mustDecode(t *testing.T, s string) any {
	t.Helper()
	v, err := decode([]byte(s))
	if err != nil {
		t.Fatal(err)
	}
	return v
}

func TestCompare(t *testing.T) {
	cases := []struct {
		name, want, got string
		opts            Options
		diff            string
	}{
		{"same shape, different values", `{"a":1,"b":"x"}`, `{"a":9,"b":"y"}`, Options{}, ""},
		{"extra response field is allowed", `{"a":1}`, `{"a":1,"b":2}`, Options{}, ""},
		{"missing field", `{"a":1,"b":2}`, `{"a":1}`, Options{}, "b: is missing"},
		{"optional field may be missing", `{"a":1,"b":2}`, `{"a":1}`, Options{Optional: []string{"b"}}, ""},
		{"wrong kind", `{"a":1}`, `{"a":"1"}`, Options{}, "a: is a string, the fixture has a number"},
		{"null in the response matches", `{"a":1}`, `{"a":null}`, Options{}, ""},
		{"null in the fixture matches", `{"a":null}`, `{"a":{"x":1}}`, Options{}, ""},
		{"array elements are each checked", `{"r":[{"id":1}]}`, `{"r":[{"id":1},{"name":"x"}]}`, Options{}, "r[].id: is missing"},
		{"empty fixture array matches anything", `{"r":[]}`, `{"r":[{"id":1}]}`, Options{}, ""},
		{"empty response array matches", `{"r":[{"id":1}]}`, `{"r":[]}`, Options{}, ""},
		{"open object compares value kinds", `{"m":{"head":1}}`, `{"m":{"chest":2,"legs":"x"}}`, Options{Open: []string{"m"}}, "m.legs: is a string"},
		{"array against object", `{"r":[]}`, `{"r":{}}`, Options{}, "r: is a object, the fixture has a array"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			diffs := Compare(mustDecode(t, tc.want), mustDecode(t, tc.got), tc.opts)
			joined := strings.Join(diffs, "\n")
			if tc.diff == "" && joined != "" {
				t.Fatalf("unexpected diffs: %s", joined)
			}
			if tc.diff != "" && !strings.Contains(joined, tc.diff) {
				t.Fatalf("diffs = %q, want it to contain %q", joined, tc.diff)
			}
		})
	}
}
