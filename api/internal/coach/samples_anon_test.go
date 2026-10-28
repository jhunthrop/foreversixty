package coach

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
)

// A moderator never sees an anonymised character's name or review in the samples, as
// AccessFor hides that character's row from moderators everywhere else.
func TestSamplesRoute_BlanksAnAnonymisedCharacter(t *testing.T) {
	pool := writerPool(t)
	ctx := context.Background()
	day := sampleToday.Add(48 * time.Hour)
	seedReviews(t, pool, day.Add(time.Hour), map[string]int{"anon": 1, "open": 1})
	anonKey := "us/normal/" + fmt.Sprintf("t37-%s-anon-0", day.Add(time.Hour).Format("0102"))
	user := capsUser(t, pool, fmt.Sprintf("t37-anon-%d@example.com", time.Now().UnixNano()))
	for _, q := range []string{
		`update users set anonymize = true where id = $2 and $1::text is not null`,
		`insert into characters (key, region, ruleset, name, user_id) values ($1, 'us', 'normal', 't37 anon', $2)
		 on conflict (key) do update set user_id = excluded.user_id`,
	} {
		if _, err := pool.Exec(ctx, q, anonKey, user); err != nil {
			t.Fatalf("%s: %v", q, err)
		}
	}
	t.Cleanup(func() { pool.Exec(context.Background(), `delete from characters where key = $1`, anonKey) })
	if _, err := SampleDay(ctx, pool, day, sampleRand(), slog.Default()); err != nil {
		t.Fatal(err)
	}
	r := httptest.NewRequest(http.MethodGet, "/v1/admin/coach/samples?day="+day.Format(sampleDayFmt), nil)
	w := httptest.NewRecorder()
	SamplesHandler(pool)(w, r.WithContext(auth.WithActor(r.Context(), auth.Actor{UserID: 1, Role: "moderator"})))
	var body struct{ Data struct{ Samples []sampleItem } }
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil || w.Code != http.StatusOK || len(body.Data.Samples) != 2 {
		t.Fatalf("%d %s %v", w.Code, w.Body, err)
	}
	for _, s := range body.Data.Samples {
		anon := s.PlayerKey == anonKey
		blank := s.CharacterName == "" && string(s.Review) == "null"
		if anon != blank {
			t.Fatalf("anonymised %v: name %q review %s", anon, s.CharacterName, s.Review)
		}
	}
}
