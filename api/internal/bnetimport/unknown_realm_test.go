// api/internal/bnetimport/unknown_realm_test.go
package bnetimport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/character"
)

// unknownRealm is a realm id the committed Forever realm table does not hold.
const unknownRealm = 6999

// unknownRealmAccount makes the fixture's Forever account one character on
// realm, whose profile is not public: nothing but the realm id names its
// ruleset.
func unknownRealmAccount(f *blizzardFixture, first, last string, id, realm int64) {
	f.rawJSON(foreverAccountPath, []byte(fmt.Sprintf(`{"wow_accounts":[{"id":1,"characters":[
		{"name":%q,"surname":%q,"id":%d,"playable_class":{"name":"Rogue"},
		 "playable_race":{"name":"Gnome"},"gender":{"type":"FEMALE"},"faction":{"type":"ALLIANCE"},"level":3,
		 "protected_character":{"href":"https://us.api.blizzard.com/profile/user/wow/protected-character/%d-%d?namespace=profile-forever-us"}}]}]}`,
		first, last, id, realm, id)))
	slug := character.Slug(first + "-" + last)
	for _, sub := range []string{"", "/equipment", "/character-media"} {
		f.json(http.MethodGet, "/profile/wow/character/"+slug+sub+"?namespace=profile-forever-us", http.StatusNotFound, nil)
	}
}

// unknownRealmService is a Forever service whose realm index answers 404.
func unknownRealmService(t *testing.T) (*Service, *blizzardFixture, *bytes.Buffer) {
	t.Helper()
	f := newBlizzardFixture(t)
	f.json(http.MethodGet, "/data/wow/realm/index?namespace=dynamic-forever-us", http.StatusNotFound, nil)
	svc := foreverService(t, f)
	var logs bytes.Buffer
	svc.Log = slog.New(slog.NewTextHandler(&logs, nil))
	return svc, f, &logs
}

// importOne signs a fresh account in with one character on realm and returns
// the account and the character's key.
func importOne(t *testing.T, svc *Service, f *blizzardFixture, first, last string, id, realm int64) (int64, string) {
	t.Helper()
	unknownRealmAccount(f, first, last, id, realm)
	uid := seedUser(t, svc.Pool)
	summary, err := svc.ImportAccount(context.Background(), uid, "user-oauth-token")
	if err != nil {
		t.Fatal(err)
	}
	if summary.Characters != 1 || summary.Skipped != 0 {
		t.Fatalf("summary = %+v, want the one character imported", summary)
	}
	return uid, keyOf(t, svc, uid)
}

func keyOf(t *testing.T, svc *Service, uid int64) string {
	t.Helper()
	var key string
	if err := svc.Pool.QueryRow(context.Background(), `select key from characters where user_id = $1`, uid).Scan(&key); err != nil {
		t.Fatal(err)
	}
	return key
}

// With the realm index gone, the realm id unknown to the table and the profile
// not yet public, nothing names the ruleset: the character is imported under
// the unknown ruleset (never a guess), with its realm id, and the log line
// names the realm id and that the player is asked.
func TestImportAccountKeysAnUnknownRealmCharacterAsUnknown(t *testing.T) {
	svc, f, logs := unknownRealmService(t)
	uid, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm)

	if key != "us/unknown/newly-rolled" {
		t.Fatalf("key = %q, want us/unknown/newly-rolled", key)
	}
	var ruleset string
	var realmID int64
	if err := svc.Pool.QueryRow(context.Background(),
		`select ruleset, bnet_realm_id from characters where user_id = $1`, uid).Scan(&ruleset, &realmID); err != nil {
		t.Fatal(err)
	}
	if ruleset != character.RulesetUnknown || realmID != unknownRealm {
		t.Fatalf("ruleset %q realm %d, want unknown on %d", ruleset, realmID, unknownRealm)
	}
	for _, want := range []string{"op=needs_ruleset", "realm_id=6999", "realm_known=false", "fallback=ask"} {
		if !strings.Contains(logs.String(), want) {
			t.Errorf("needs_ruleset line lacks %s:\n%s", want, logs.String())
		}
	}
}

// The answer moves the character and every row keyed by it, marks the ruleset
// as the player's, records one piece of evidence for the realm, and a later
// sign-in keeps the answered key rather than going back to unknown.
func TestAnswerRealmMovesTheCharacterAndItsRows(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	ctx := context.Background()
	uid, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm)
	if _, err := svc.Pool.Exec(ctx, `update users set main_character_key = $1 where id = $2`, key, uid); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.Pool.Exec(ctx,
		`insert into addon_exports (character_key, user_id, region, ruleset, name, export, source, captured_at, updated_at)
		 values ($1, $2, 'us', 'unknown', 'Newly Rolled', 'x', 'addon', now(), now())
		 on conflict (character_key) do update set export = excluded.export`, key, uid); err != nil {
		t.Fatal(err)
	}

	newKey, err := svc.AnswerRealm(ctx, uid, key, character.RulesetPvP)
	if err != nil {
		t.Fatal(err)
	}
	if newKey != "us/pvp/newly-rolled" {
		t.Fatalf("new key = %q", newKey)
	}
	var answered bool
	var main, exportRuleset string
	if err := svc.Pool.QueryRow(ctx,
		`select c.ruleset_answered, u.main_character_key, e.ruleset from characters c
		   join users u on u.id = c.user_id join addon_exports e on e.character_key = c.key
		  where c.key = $1`, newKey).Scan(&answered, &main, &exportRuleset); err != nil {
		t.Fatal(err)
	}
	if !answered || main != newKey || exportRuleset != "pvp" {
		t.Fatalf("answered %v main %q export ruleset %q, want the rows moved with it", answered, main, exportRuleset)
	}
	var evidence int
	if err := svc.Pool.QueryRow(ctx,
		`select count(*) from realm_ruleset_answers where region = 'us' and realm_id = $1 and user_id = $2 and ruleset = 'pvp'`,
		unknownRealm, uid).Scan(&evidence); err != nil {
		t.Fatal(err)
	}
	if evidence != 1 {
		t.Fatalf("evidence rows = %d, want 1", evidence)
	}

	if _, err := svc.ImportAccount(ctx, uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}
	if got := keyOf(t, svc, uid); got != newKey {
		t.Fatalf("after a later sign-in key = %q, want the answered %q kept", got, newKey)
	}
}

func TestAnswerRealmRefusesWhatIsNotTheirsToAnswer(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	ctx := context.Background()
	uid, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm)
	other := seedUser(t, svc.Pool)

	if _, err := svc.AnswerRealm(ctx, uid, key, "pve"); !errors.Is(err, ErrInvalidAnswer) {
		t.Errorf("unknown ruleset err = %v, want ErrInvalidAnswer", err)
	}
	if _, err := svc.AnswerRealm(ctx, other, key, character.RulesetPvP); !errors.Is(err, ErrNotAwaitingAnswer) {
		t.Errorf("another account's character err = %v, want ErrNotAwaitingAnswer", err)
	}
	if _, err := svc.AnswerRealm(ctx, uid, key, character.RulesetRP); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.AnswerRealm(ctx, uid, "us/rp/newly-rolled", character.RulesetPvP); !errors.Is(err, ErrNotAwaitingAnswer) {
		t.Errorf("an answered character err = %v, want ErrNotAwaitingAnswer", err)
	}
	if _, err := svc.Pool.Exec(ctx,
		`insert into characters (key, region, ruleset, name, user_id) values ('us/pvp/second-roll', 'us', 'pvp', 'Second Roll', $1)`,
		other); err != nil {
		t.Fatal(err)
	}
	_, secondKey := importOne(t, svc, f, "Second", "Roll", 9200, unknownRealm)
	secondUID := ownerOf(t, svc, secondKey)
	if _, err := svc.AnswerRealm(ctx, secondUID, secondKey, character.RulesetPvP); !errors.Is(err, ErrAnswerKeyTaken) {
		t.Errorf("a key another account holds err = %v, want ErrAnswerKeyTaken", err)
	}
}

func ownerOf(t *testing.T, svc *Service, key string) int64 {
	t.Helper()
	var uid int64
	if err := svc.Pool.QueryRow(context.Background(), `select user_id from characters where key = $1`, key).Scan(&uid); err != nil {
		t.Fatal(err)
	}
	return uid
}

// Two players agreeing on a realm, with no one disagreeing, resolve it: the
// realm's other unknown characters move at once and the next import keys the
// realm by it. A third player's different answer is logged and stops the
// realm resolving for anyone after.
func TestRealmAnswersResolveTheRealmAtQuorumAndNotOnDisagreement(t *testing.T) {
	svc, f, logs := unknownRealmService(t)
	ctx := context.Background()
	first, firstKey := importOne(t, svc, f, "First", "Player", 9101, unknownRealm)
	second, secondKey := importOne(t, svc, f, "Second", "Player", 9102, unknownRealm)
	_, waitingKey := importOne(t, svc, f, "Waiting", "Player", 9103, unknownRealm)

	if _, err := svc.AnswerRealm(ctx, first, firstKey, character.RulesetPvP); err != nil {
		t.Fatal(err)
	}
	if got := keyOfKey(t, svc, waitingKey); got != waitingKey {
		t.Fatalf("one answer moved another player's character to %q", got)
	}
	if _, err := svc.AnswerRealm(ctx, second, secondKey, character.RulesetPvP); err != nil {
		t.Fatal(err)
	}
	var waitingRuleset string
	var waitingAnswered bool
	if err := svc.Pool.QueryRow(ctx,
		`select ruleset, ruleset_answered from characters where name = 'Waiting Player'`).Scan(&waitingRuleset, &waitingAnswered); err != nil {
		t.Fatal(err)
	}
	if waitingRuleset != "pvp" || waitingAnswered {
		t.Fatalf("waiting player's ruleset %q answered %v, want pvp from the realm's agreed answers", waitingRuleset, waitingAnswered)
	}
	_, freshKey := importOne(t, svc, f, "Fresh", "Player", 9104, unknownRealm)
	if freshKey != "us/pvp/fresh-player" {
		t.Fatalf("a sign-in after the quorum keyed %q, want us/pvp/fresh-player", freshKey)
	}

	dissent, dissentKey := importOne(t, svc, f, "Dissent", "Player", 9105, 7001)
	if _, err := svc.Pool.Exec(ctx, `update characters set bnet_realm_id = $1 where key = $2`, unknownRealm, dissentKey); err != nil {
		t.Fatal(err)
	}
	if _, err := svc.AnswerRealm(ctx, dissent, dissentKey, character.RulesetNormal); err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(logs.String(), "op=realm_answers_disagree") || !strings.Contains(logs.String(), "answers=normal:1,pvp:2") {
		t.Fatalf("no disagreement line:\n%s", logs.String())
	}
	_, laterKey := importOne(t, svc, f, "Later", "Player", 9106, unknownRealm)
	if laterKey != "us/unknown/later-player" {
		t.Fatalf("a sign-in after a disagreement keyed %q, want unknown", laterKey)
	}
}

// keyOfKey is key's current key: itself while its row still holds it, else
// the key its Battle.net character moved to.
func keyOfKey(t *testing.T, svc *Service, key string) string {
	t.Helper()
	var n int
	if err := svc.Pool.QueryRow(context.Background(), `select count(*) from characters where key = $1`, key).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n == 1 {
		return key
	}
	return "moved"
}

// A character imported as unknown whose realm resolves later (here by the
// realm's agreed answers; a new realm table entry takes the same path) moves
// at the next sign-in with every row keyed by it, and the nightly refresh
// moves one that has not signed in since.
func TestALaterResolutionMovesAnUnknownCharactersRows(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	ctx := context.Background()
	uid, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm)
	if _, err := svc.Pool.Exec(ctx, `update users set main_character_key = $1 where id = $2`, key, uid); err != nil {
		t.Fatal(err)
	}
	agree := func(ruleset string) {
		for range realmAnswerQuorum {
			voter := seedUser(t, svc.Pool)
			if _, err := svc.Pool.Exec(ctx,
				`insert into realm_ruleset_answers (region, realm_id, user_id, ruleset) values ('us', $1, $2, $3)`,
				unknownRealm, voter, ruleset); err != nil {
				t.Fatal(err)
			}
		}
	}
	agree("rp")
	if _, err := svc.ImportAccount(ctx, uid, "user-oauth-token"); err != nil {
		t.Fatal(err)
	}
	var main string
	if err := svc.Pool.QueryRow(ctx, `select main_character_key from users where id = $1`, uid).Scan(&main); err != nil {
		t.Fatal(err)
	}
	if got := keyOf(t, svc, uid); got != "us/rp/newly-rolled" || main != got {
		t.Fatalf("key %q main %q, want both moved to us/rp/newly-rolled", got, main)
	}

	// The refresh: a second unknown character whose realm has since resolved.
	other, otherKey := importOne(t, svc, f, "Quiet", "Sleeper", 9300, 7002)
	if otherKey != "us/unknown/quiet-sleeper" {
		t.Fatalf("key = %q, want unknown before its realm resolves", otherKey)
	}
	if _, err := svc.Pool.Exec(ctx, `update characters set bnet_realm_id = $1 where key = $2`, unknownRealm, otherKey); err != nil {
		t.Fatal(err)
	}
	tx, err := svc.Pool.Begin(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = tx.Rollback(ctx) }()
	newKey, ruleset, err := svc.resolveUnknownInPlace(ctx, tx, otherKey, "us", character.RulesetUnknown, "Quiet Sleeper", unknownRealm, 0)
	if err != nil {
		t.Fatal(err)
	}
	if newKey != "us/rp/quiet-sleeper" || ruleset != "rp" {
		t.Fatalf("refresh resolved %q %q, want us/rp/quiet-sleeper", newKey, ruleset)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatal(err)
	}
	if got := keyOf(t, svc, other); got != newKey {
		t.Fatalf("row key %q, want %q", got, newKey)
	}
}

func TestRealmAnswerEndpoint(t *testing.T) {
	svc, f, _ := unknownRealmService(t)
	uid, key := importOne(t, svc, f, "Newly", "Rolled", 9100, unknownRealm)
	post := func(body string) *httptest.ResponseRecorder {
		r := httptest.NewRequest(http.MethodPost, RealmAnswerPath, strings.NewReader(body))
		r = r.WithContext(auth.WithActor(r.Context(), auth.Actor{UserID: uid}))
		w := httptest.NewRecorder()
		svc.answerRealm(w, r)
		return w
	}
	if w := post(`not json`); w.Code != http.StatusBadRequest {
		t.Errorf("bad body = %d, want 400", w.Code)
	}
	if w := post(`{"key":"` + key + `","ruleset":"pve"}`); w.Code != http.StatusBadRequest {
		t.Errorf("bad ruleset = %d, want 400", w.Code)
	}
	if w := post(`{"key":"us/unknown/nobody-here","ruleset":"pvp"}`); w.Code != http.StatusNotFound {
		t.Errorf("not theirs = %d, want 404", w.Code)
	}
	w := post(`{"key":"` + key + `","ruleset":"hardcore"}`)
	if w.Code != http.StatusOK {
		t.Fatalf("answer = %d %s", w.Code, w.Body.String())
	}
	var env struct {
		Data struct {
			Key     string `json:"key"`
			Ruleset string `json:"ruleset"`
		} `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &env); err != nil {
		t.Fatal(err)
	}
	if env.Data.Key != "us/hardcore/newly-rolled" || env.Data.Ruleset != "hardcore" {
		t.Fatalf("answer body = %s", w.Body.String())
	}
}
