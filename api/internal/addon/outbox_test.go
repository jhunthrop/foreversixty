package addon

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/sim/enginever"
)

// recordingJobs is a jobs.Runner that remembers every dispatch.
type recordingJobs struct {
	mu    sync.Mutex
	calls [][]string
}

func (r *recordingJobs) Run(_ context.Context, args ...string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.calls = append(r.calls, args)
	return nil
}

func (r *recordingJobs) count() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.calls)
}

// entitled answers every server-sims question with allowed.
type entitled bool

func (e entitled) Can(context.Context, int64, entitlements.Feature) (bool, entitlements.Reason, error) {
	return bool(e), "", nil
}

type outboxHarness struct {
	*harness
	jobs *recordingJobs
}

// newOutboxHarness is newHarness with the sim request path wired to a real
// sims.Store on the same database and a job runner that only records.
func newOutboxHarness(t *testing.T, allowed bool) *outboxHarness {
	t.Helper()
	return newOutboxHarnessOn(t, allowed, enginever.Version)
}

// newOutboxHarnessOn is newOutboxHarness with the sims service claiming engineVersion: any
// other than enginever.Version makes every request one the engine refuses to run.
func newOutboxHarnessOn(t *testing.T, allowed bool, engineVersion string) *outboxHarness {
	t.Helper()
	h := newHarness(t)
	h.server.Close()
	jobs := &recordingJobs{}
	simSvc := &sims.Service{
		Store: &sims.Store{Pool: h.pool}, Jobs: jobs, Accounts: entitled(allowed),
		EngineVersion: engineVersion,
	}
	mux := http.NewServeMux()
	Mount(mux, &Service{Store: h.store, Builds: h.builds, Data: h.trees, Accounts: &auth.Store{Pool: h.pool},
		SimRequests: simSvc, Log: slog.New(slog.NewTextHandler(io.Discard, nil))}, 0)
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return &outboxHarness{harness: h, jobs: jobs}
}

type exportsReply struct {
	Stored int `json:"stored"`
	Outbox *struct {
		Status         string `json:"status"`
		Export         string `json:"export"`
		SimRequest     string `json:"sim_request"`
		RosterSnapshot string `json:"roster_snapshot"`
		DroppedMembers int    `json:"dropped_members"`
	} `json:"outbox"`
}

func (h *outboxHarness) post(body string) exportsReply {
	h.t.Helper()
	res := h.do(http.MethodPost, "/v1/addon/exports", body)
	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		res.Body.Close()
		h.t.Fatalf("POST /v1/addon/exports = %d: %s", res.StatusCode, b)
	}
	var out exportsReply
	h.data(res, &out)
	return out
}

// obnoxious is the account's own copy of the contract fixture's character,
// synced the way the companion already sends every character today.
const obnoxious = `{"name":"Obnoxious","ruleset":"normal","region":"us","export":"FS1:1.60.1.70291:warrior:human:0/0/0:|guild=OLYMPUS%20XXVII:3"}`

const obnoxiousKey = "us/normal/obnoxious"

// withOutbox is a sync body carrying obnoxious and the given outbox.
func withOutbox(outbox string) string {
	return `{"characters":[` + obnoxious + `],"outbox":` + outbox + `}`
}

// simRequests is the character's recorded requests that started a run; a
// refused one is recorded too (state needs_premium, duplicate or invalid, for
// the inbox ack) and is read by refusedRequests.
func (h *outboxHarness) simRequests() []struct{ requestedAt, simID string } {
	h.t.Helper()
	rows, err := h.pool.Query(context.Background(),
		`select requested_at, coalesce(sim_id, '') from addon_sim_requests
		 where character_key = $1 and state = 'queued' order by created_at`, obnoxiousKey)
	if err != nil {
		h.t.Fatal(err)
	}
	defer rows.Close()
	var out []struct{ requestedAt, simID string }
	for rows.Next() {
		var r struct{ requestedAt, simID string }
		if err := rows.Scan(&r.requestedAt, &r.simID); err != nil {
			h.t.Fatal(err)
		}
		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		h.t.Fatal(err)
	}
	return out
}

func TestTheContractOutboxIsStoredAndItsSimRequestQueuesOneRun(t *testing.T) {
	h := newOutboxHarness(t, true)
	out := h.post(withOutbox(string(contractFixture(t, "outbox-v1.json"))))
	if out.Outbox == nil || out.Outbox.Status != "read" || out.Outbox.Export != "stored" || out.Outbox.SimRequest != "queued" {
		t.Fatalf("outbox = %+v, want read/stored/queued", out.Outbox)
	}
	if out.Stored != 1 {
		t.Fatalf("stored = %d, want 1: the list and the outbox name the same character", out.Stored)
	}

	var export string
	if err := h.pool.QueryRow(t.Context(),
		`select export from addon_exports where character_key = $1`, obnoxiousKey).Scan(&export); err != nil {
		t.Fatal(err)
	}
	if export != "FS1:1.60.1.70291:warrior:human:371000000/0/0:chest=6627" {
		t.Fatalf("stored export = %q, want the outbox's code", export)
	}

	reqs := h.simRequests()
	if len(reqs) != 1 || reqs[0].requestedAt != "2026-11-04 14:10" || reqs[0].simID == "" {
		t.Fatalf("sim requests = %+v, want one keyed by the request time with a sim id", reqs)
	}
	var (
		state, spec string
		owner       int64
		body        []byte
	)
	if err := h.pool.QueryRow(t.Context(),
		`select state, spec, user_id, result from sims where id = $1`, reqs[0].simID).
		Scan(&state, &spec, &owner, &body); err != nil {
		t.Fatal(err)
	}
	if state != sims.StateQueued || spec != "warrior-arms" || owner != h.owner {
		t.Fatalf("sims row = %s/%s/%d, want queued/warrior-arms/%d", state, spec, owner, h.owner)
	}
	var stored struct {
		Request struct {
			Source struct {
				Kind, Ref  string
				CapturedAt string `json:"captured_at"`
			} `json:"source"`
			Character struct {
				Name, Race, Class, Talents string
				Level                      int
				Gear                       []struct {
					Slot   string `json:"slot"`
					ItemID int    `json:"item_id"`
				} `json:"gear"`
			} `json:"character"`
		} `json:"request"`
	}
	if err := json.Unmarshal(body, &stored); err != nil {
		t.Fatal(err)
	}
	r := stored.Request
	if r.Source.Kind != "addon" || r.Source.Ref != obnoxiousKey || r.Source.CapturedAt != "2026-11-04 14:10" {
		t.Errorf("source = %+v, want the addon, the character and the request time", r.Source)
	}
	c := r.Character
	if c.Class != "warrior" || c.Race != "human" || c.Level != 20 || c.Talents != "371--" ||
		len(c.Gear) != 1 || c.Gear[0].Slot != "chest" || c.Gear[0].ItemID != 6627 {
		t.Errorf("character = %+v, want the request's own gear code", c)
	}
	if h.jobs.count() != 1 || h.jobs.calls[0][0] != sims.SimRunJobCommand || h.jobs.calls[0][1] != reqs[0].simID {
		t.Fatalf("job calls = %v, want one sim-run of %s", h.jobs.calls, reqs[0].simID)
	}
}

func TestASecondSimRequestWithinTheHourIsDeduplicated(t *testing.T) {
	h := newOutboxHarness(t, true)
	fixture := string(contractFixture(t, "outbox-v1.json"))
	h.post(withOutbox(fixture))

	// The companion posts the same outbox again on the next file change.
	again := h.post(withOutbox(fixture))
	if again.Outbox.SimRequest != "duplicate" {
		t.Fatalf("the identical request = %q, want duplicate", again.Outbox.SimRequest)
	}
	// A new request for the same character inside the hour is one too many.
	later := h.post(withOutbox(strings.Replace(fixture, "2026-11-04 14:10", "2026-11-04 14:40", 1)))
	if later.Outbox.SimRequest != "duplicate" {
		t.Fatalf("a second request inside the hour = %q, want duplicate", later.Outbox.SimRequest)
	}
	if n := len(h.simRequests()); n != 1 {
		t.Fatalf("sim requests = %d, want 1", n)
	}
	// The second request is recorded as a duplicate, for the inbox ack; the
	// re-posted identical one is the first request again and adds nothing.
	if got := h.refusedRequests(); len(got) != 1 || got[0] != [2]string{"2026-11-04 14:40", "duplicate"} {
		t.Fatalf("refused requests = %v, want the 14:40 one as a duplicate", got)
	}
	var runs int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from sims where user_id = $1`, h.owner).Scan(&runs); err != nil {
		t.Fatal(err)
	}
	if runs != 1 || h.jobs.count() != 1 {
		t.Fatalf("sims rows = %d, job calls = %d, want 1 and 1", runs, h.jobs.count())
	}

	// An hour on, the next request runs; the old identical one still never does.
	if _, err := h.pool.Exec(t.Context(),
		`update addon_sim_requests set created_at = now() - interval '61 minutes'`); err != nil {
		t.Fatal(err)
	}
	if got := h.post(withOutbox(fixture)).Outbox.SimRequest; got != "duplicate" {
		t.Fatalf("the identical request an hour on = %q, want duplicate", got)
	}
	if got := h.post(withOutbox(strings.Replace(fixture, "2026-11-04 14:10", "2026-11-04 15:30", 1))).Outbox.SimRequest; got != "queued" {
		t.Fatalf("a new request an hour on = %q, want queued", got)
	}
	if n := len(h.simRequests()); n != 2 || h.jobs.count() != 2 {
		t.Fatalf("sim requests = %d, job calls = %d, want 2 and 2", n, h.jobs.count())
	}
}

func TestASimRequestWithoutServerSimsQueuesNothing(t *testing.T) {
	h := newOutboxHarness(t, false)
	out := h.post(withOutbox(string(contractFixture(t, "outbox-v1.json"))))
	if out.Outbox.SimRequest != "premium_required" || out.Outbox.Export != "stored" {
		t.Fatalf("outbox = %+v, want the export stored and the sim refused", out.Outbox)
	}
	if n := len(h.simRequests()); n != 0 || h.jobs.count() != 0 {
		t.Fatalf("sim requests = %d, job calls = %d, want none", n, h.jobs.count())
	}
	// The refusal is recorded, so the inbox ack can say why nothing happened.
	if got := h.refusedRequests(); len(got) != 1 || got[0] != [2]string{"2026-11-04 14:10", "needs_premium"} {
		t.Fatalf("refused requests = %v, want the request as needs_premium", got)
	}
}

// refusedRequests is the character's recorded requests that started no run,
// as (requested_at, state), oldest first.
func (h *outboxHarness) refusedRequests() [][2]string {
	h.t.Helper()
	rows, err := h.pool.Query(context.Background(),
		`select requested_at, state from addon_sim_requests
		 where character_key = $1 and state <> 'queued' and sim_id is null order by created_at`, obnoxiousKey)
	if err != nil {
		h.t.Fatal(err)
	}
	defer rows.Close()
	var out [][2]string
	for rows.Next() {
		var r [2]string
		if err := rows.Scan(&r[0], &r[1]); err != nil {
			h.t.Fatal(err)
		}
		out = append(out, r)
	}
	if err := rows.Err(); err != nil {
		h.t.Fatal(err)
	}
	return out
}

const rosterOutbox = `{"format":1,
 "export":{"character":"Obnoxious-Yell","code":"FS1:1.60.1.70291:warrior:human:0/0/0:|guild=OLYMPUS%20XXVII:3","saved_at":"2026-11-04 14:02"},
 "roster_snapshot":{"guild":"%s","members":[
   {"name":"Obnoxious-Yell","level":20,"spec":"warrior-arms","gear_gap":14,"build":"Arms leveling","seen_at":"2026-11-04 14:01"},
   {"name":"Healbot-Yell","level":18,"spec":"priest-holy","gear_gap":3,"build":"Holy leveling","seen_at":"2026-11-04 13:55"}]}}`

func TestTheRosterSnapshotIsStoredForTheCharactersOwnGuild(t *testing.T) {
	h := newOutboxHarness(t, true)
	out := h.post(withOutbox(strings.Replace(rosterOutbox, "%s", "olympus xxvii", 1)))
	if out.Outbox.RosterSnapshot != "stored" || out.Outbox.SimRequest != "" {
		t.Fatalf("outbox = %+v, want the snapshot stored and no sim request", out.Outbox)
	}
	var (
		guild, reporter string
		members         []map[string]any
	)
	if err := h.pool.QueryRow(t.Context(),
		`select g.name, s.character_key, s.members from guild_roster_snapshots s
		 join guilds g on g.id = s.guild_id`).Scan(&guild, &reporter, &members); err != nil {
		t.Fatal(err)
	}
	if guild != "OLYMPUS XXVII" || reporter != obnoxiousKey || len(members) != 2 ||
		members[1]["name"] != "Healbot-Yell" || members[1]["gear_gap"] != float64(3) {
		t.Fatalf("snapshot = %s by %s: %v", guild, reporter, members)
	}

	// A newer snapshot replaces the guild's row rather than adding one.
	h.post(withOutbox(strings.Replace(rosterOutbox, "%s", "OLYMPUS XXVII", 1)))
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from guild_roster_snapshots`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 1 {
		t.Fatalf("snapshots = %d, want 1 per guild", n)
	}
}

func TestARosterSnapshotForAnotherGuildIsNotStored(t *testing.T) {
	h := newOutboxHarness(t, true)
	out := h.post(withOutbox(strings.Replace(rosterOutbox, "%s", "Somebody Else", 1)))
	if out.Outbox.RosterSnapshot != "not_in_guild" {
		t.Fatalf("roster_snapshot = %q, want not_in_guild", out.Outbox.RosterSnapshot)
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from guild_roster_snapshots`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("snapshots = %d, want none", n)
	}
}

func TestAnOutboxForACharacterTheAccountHasNeverSyncedStoresNothing(t *testing.T) {
	h := newOutboxHarness(t, true)
	out := h.post(`{"outbox":` + string(contractFixture(t, "outbox-v1.json")) + `}`)
	if out.Stored != 0 || out.Outbox.Export != "unknown_character" || out.Outbox.SimRequest != "unknown_character" ||
		out.Outbox.RosterSnapshot != "unknown_character" {
		t.Fatalf("reply = %+v, want nothing stored", out)
	}
	if n := len(h.simRequests()); n != 0 || h.jobs.count() != 0 {
		t.Fatalf("sim requests = %d, job calls = %d, want none", n, h.jobs.count())
	}
}

func TestAnOutboxOfAnUnknownFormatIsIgnored(t *testing.T) {
	h := newOutboxHarness(t, true)
	fixture := strings.Replace(string(contractFixture(t, "outbox-v1.json")), `"format": 1`, `"format": 2`, 1)
	out := h.post(withOutbox(fixture))
	if out.Stored != 1 || out.Outbox.Status != "ignored" || out.Outbox.Export != "" {
		t.Fatalf("reply = %+v, want the characters stored and the outbox ignored", out)
	}
	if n := len(h.simRequests()); n != 0 {
		t.Fatalf("sim requests = %d, want none", n)
	}
}

func TestAnOutboxThatIsNotOneIsRefused(t *testing.T) {
	h := newOutboxHarness(t, true)
	for name, body := range map[string]string{
		"neither list":        `{}`,
		"not an object":       withOutbox(`[1]`),
		"no export":           withOutbox(`{"format":1}`),
		"export not object":   withOutbox(`{"format":1,"export":"FS1:a"}`),
		"no export character": withOutbox(`{"format":1,"export":{"character":"","code":"FS1:a"}}`),
		"over-long code": withOutbox(`{"format":1,"export":{"character":"Obnoxious-Yell","code":"` +
			strings.Repeat("x", MaxExportLen+1) + `"}}`),
	} {
		res := h.do(http.MethodPost, "/v1/addon/exports", body)
		res.Body.Close()
		if res.StatusCode != http.StatusBadRequest {
			t.Errorf("%s = %d, want 400", name, res.StatusCode)
		}
	}
	if h.jobs.count() != 0 {
		t.Fatalf("job calls = %d, want none", h.jobs.count())
	}
}

// altbert is a second character of the account, synced in the same body.
const altbert = `{"name":"Altbert","ruleset":"normal","region":"us","export":"FS1:1.60.1.70291:priest:human:0/0/0:"}`

// Members come from other players' addons: one bad line must cost only that
// line, never the player's own sync.
func TestABadRosterMemberIsDroppedAndTheRestOfTheSyncStored(t *testing.T) {
	h := newOutboxHarness(t, true)
	members := `[
	  {"name":"Obnoxious-Yell","level":20,"spec":"warrior-arms","gear_gap":14,"build":"Arms leveling","seen_at":"2026-11-04 14:01"},
	  {"name":"Overlevel-Yell","level":99,"spec":"mage-fire","gear_gap":1,"build":"Fire","seen_at":"2026-11-04 14:01"},
	  {"name":"Wrongtype-Yell","level":"twenty","spec":"rogue-combat","gear_gap":1,"build":"Combat","seen_at":"2026-11-04 14:01"},
	  {"name":"Ahead-Yell","level":30,"spec":"priest-holy","gear_gap":-2,"build":"Holy","seen_at":"2026-11-04 13:50"}]`
	outbox := `{"format":1,
	 "export":{"character":"Obnoxious-Yell","code":"FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627|guild=OLYMPUS%20XXVII:3","saved_at":"2026-11-04 14:02"},
	 "roster_snapshot":{"guild":"OLYMPUS XXVII","members":` + members + `}}`
	out := h.post(`{"characters":[` + obnoxious + `,` + altbert + `],"outbox":` + outbox + `}`)
	if out.Stored != 2 || out.Outbox.Export != OutboxStored || out.Outbox.RosterSnapshot != OutboxStored ||
		out.Outbox.DroppedMembers != 2 {
		t.Fatalf("reply = %+v %+v, want 2 stored, the snapshot stored and 2 members dropped", out, out.Outbox)
	}

	var stored []RosterMember
	if err := h.pool.QueryRow(t.Context(), `select members from guild_roster_snapshots`).Scan(&stored); err != nil {
		t.Fatal(err)
	}
	if len(stored) != 2 || stored[0].Name != "Obnoxious-Yell" || stored[1].Name != "Ahead-Yell" || stored[1].GearGap != -2 {
		t.Fatalf("stored members = %+v, want the two good lines, a negative gear gap included", stored)
	}
	exports, err := h.store.Exports(t.Context(), h.owner)
	if err != nil {
		t.Fatal(err)
	}
	byName := map[string]string{}
	for _, e := range exports {
		byName[e.Name] = e.Export
	}
	if byName["Altbert"] != "FS1:1.60.1.70291:priest:human:0/0/0:" ||
		byName["Obnoxious"] != "FS1:1.60.1.70291:warrior:human:0/0/0:chest=6627|guild=OLYMPUS%20XXVII:3" {
		t.Fatalf("exports = %v, want the list's alt and the outbox's export", byName)
	}
}

func TestARosterWithNoUsableMemberIsAnsweredInvalid(t *testing.T) {
	h := newOutboxHarness(t, true)
	for name, roster := range map[string]string{
		"every member bad": `{"guild":"OLYMPUS XXVII","members":[{"name":"","level":0}]}`,
		"no guild name":    `{"guild":"","members":[]}`,
		"not an object":    `"OLYMPUS XXVII"`,
	} {
		outbox := strings.Replace(rosterOutbox, `"roster_snapshot":{"guild":"%s","members":[`,
			`"roster_snapshot":`+roster+`,"unused":{"members":[`, 1)
		out := h.post(withOutbox(outbox))
		if out.Outbox.Export != OutboxStored || out.Outbox.RosterSnapshot != OutboxInvalid {
			t.Errorf("%s: outbox = %+v, want the export stored and the snapshot invalid", name, out.Outbox)
		}
	}
	var n int
	if err := h.pool.QueryRow(t.Context(), `select count(*) from guild_roster_snapshots`).Scan(&n); err != nil {
		t.Fatal(err)
	}
	if n != 0 {
		t.Fatalf("snapshots = %d, want none", n)
	}
}

func TestABadSimRequestIsAnsweredInvalidNotRefused(t *testing.T) {
	h := newOutboxHarness(t, true)
	fixture := string(contractFixture(t, "outbox-v1.json"))
	// The sim request's spec, not the roster member's.
	unknownSpec := strings.Replace(fixture, "\"spec\": \"warrior-arms\",\n    \"requested_at\"",
		"\"spec\": \"warrior-bard\",\n    \"requested_at\"", 1)
	if unknownSpec == fixture {
		t.Fatal("the fixture's sim_request spec moved; update this replacement")
	}
	for name, outbox := range map[string]string{
		"bad request time": strings.Replace(fixture, "2026-11-04 14:10", "soon", 1),
		"unknown spec":     unknownSpec,
		"not an object":    `{"format":1,"export":{"character":"Obnoxious-Yell","code":"FS1:1.60.1.70291:warrior:human:0/0/0:"},"sim_request":5}`,
	} {
		out := h.post(withOutbox(outbox))
		if out.Outbox.Export != OutboxStored || out.Outbox.SimRequest != OutboxInvalid {
			t.Errorf("%s: outbox = %+v, want the export stored and the request invalid", name, out.Outbox)
		}
	}
	if n := len(h.simRequests()); n != 0 || h.jobs.count() != 0 {
		t.Fatalf("sim requests = %d, job calls = %d, want none", n, h.jobs.count())
	}
}

// The log names a two-word character First-Last-Realm-, and the outbox carries that unit; it
// must find the export of the same two-word name, not a half-name phantom.
func TestResolveOwnMatchesATwoWordUnit(t *testing.T) {
	exports := []Export{
		{Name: "Bow Jackzon", Ruleset: "pvp", Region: "us"},
		{Name: "Bow", Ruleset: "pvp", Region: "us"},
	}
	for _, tc := range []struct {
		unit, want string
		ok         bool
	}{
		{"Bow-Jackzon-ClassicBetaPvP-", "Bow Jackzon", true},
		{"Bow-ClassicBetaPvP-", "Bow", true},
		{"Nobody-Here-ClassicBetaPvP-", "", false},
	} {
		got, ok := resolveOwn(exports, tc.unit)
		if ok != tc.ok || got.Name != tc.want {
			t.Errorf("resolveOwn(%q) = %q, %v; want %q, %v", tc.unit, got.Name, ok, tc.want, tc.ok)
		}
		if tc.ok && !validUnit(tc.unit) {
			t.Errorf("validUnit(%q) = false", tc.unit)
		}
	}
}
