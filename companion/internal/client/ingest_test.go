package client_test

import (
	"errors"
	"net/http"
	"slices"
	"testing"
	"time"

	"github.com/klauspost/compress/zstd"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fakeapi"
	"github.com/jhunthrop/foreversixty/logs/engine/fight"
	"github.com/jhunthrop/foreversixty/logs/engine/session"
	"github.com/jhunthrop/foreversixty/logs/engine/summary"
)

func dial(t *testing.T, srv *fakeapi.Server) *client.Client {
	t.Helper()
	c, err := client.New(client.Options{
		BaseURL: srv.URL,
		Token:   func() string { return fakeapi.Token },
		Retry:   client.Retry{MaxAttempts: 3, Base: time.Millisecond, Max: time.Millisecond},
		Frac:    func() float64 { return 1 },
	})
	if err != nil {
		t.Fatal(err)
	}
	return c
}

// pack compresses a chunk the way the pipeline does.
func pack(t *testing.T, plain []byte) []byte {
	t.Helper()
	enc, err := zstd.NewWriter(nil, zstd.WithEncoderLevel(zstd.SpeedDefault))
	if err != nil {
		t.Fatal(err)
	}
	defer enc.Close()
	return enc.EncodeAll(plain, nil)
}

func sampleFight() (fight.Fight, summary.Summary) {
	f := fight.Fight{
		Index: 2, Kind: fight.Encounter, Name: "Warden Kelthas", EncounterID: 9001,
		Difficulty: 14, Size: 20, Kill: true,
		Start: time.Unix(1000, 0).UTC(), End: time.Unix(1240, 0).UTC(),
		StartOffset: 4096, EndOffset: 90112,
	}
	s := summary.Summary{
		EngineVersion: session.Version, FightIndex: 2, DurationMS: 240000,
		Roster: []summary.RosterRow{
			{GUID: "Player-4184-000000A1", Name: "Morrowlyn", Class: "Paladin", Spec: "Holy",
				Role: "healer", ItemLevel: 183, ActiveMS: 200000, Deaths: 1,
				DamageTaken: 40123, DPS: 120.5, HPS: 980.25},
			{GUID: "Player-4184-000000A2", Name: "Brannic", Class: "Warrior", Spec: "Fury",
				Role: "dps", ItemLevel: 176, ActiveMS: 230000, Deaths: 0,
				DamageTaken: 91002, DPS: 1420.75, HPS: 0},
		},
	}
	return f, s
}

func TestMetricsRowsArePivotedFromTheRoster(t *testing.T) {
	f, s := sampleFight()
	rows := client.MetricsRowsOf(f, s)
	if len(rows) != 2 {
		t.Fatalf("got %d rows, want 2", len(rows))
	}
	first := rows[0]
	if first.PlayerGUID != "Player-4184-000000A1" || first.Name != "Morrowlyn" ||
		first.Class != "Paladin" || first.Spec != "Holy" || first.Role != "healer" ||
		first.ItemLevel != 183 || first.MetricDPS != 120.5 || first.MetricHPS != 980.25 ||
		first.DamageTaken != 40123 || first.ActiveMS != 200000 || first.Deaths != 1 {
		t.Errorf("row = %+v", first)
	}
	if first.EncounterID != 9001 || first.Difficulty != 14 || first.Size != 20 ||
		first.DurationMS != 240000 || !first.Kill {
		t.Errorf("fight fields = %+v", first)
	}
}

func TestAFightBundleRoundTripsThroughTheMultipartRoute(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)

	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "public", Zone: "Molten Core"})
	if err != nil {
		t.Fatal(err)
	}
	f, s := sampleFight()
	b := client.FightBundle{
		Summary: s,
		Events:  []byte("PAR1fake"),
		Metrics: client.MetricsRowsOf(f, s),
		RawRange: client.RawRange{StartOffset: f.StartOffset, EndOffset: f.EndOffset,
			SHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	}
	ct, body, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	stored, err := c.PutFight(t.Context(), rep.ID, 2, ct, body)
	if err != nil {
		t.Fatal(err)
	}
	if stored != client.Created {
		t.Errorf("first PUT = %v, want Created", stored)
	}
	again, err := c.PutFight(t.Context(), rep.ID, 2, ct, body)
	if err != nil {
		t.Fatal(err)
	}
	if again != client.Duplicate {
		t.Errorf("re-sent PUT = %v, want Duplicate", again)
	}

	got := srv.Reports()[rep.ID].Fights[2]
	if string(got.Events) != "PAR1fake" {
		t.Errorf("events = %q", got.Events)
	}
	if len(got.Metrics) != 2 || got.Metrics[1].Name != "Brannic" {
		t.Errorf("metrics = %+v", got.Metrics)
	}
	if got.RawRange.StartOffset != 4096 || got.RawRange.EndOffset != 90112 {
		t.Errorf("raw range = %+v", got.RawRange)
	}
	want := []string{"report " + rep.ID, "fight 2"}
	if got := srv.Order(); !slices.Equal(got, want) {
		t.Errorf("order = %v, want %v", got, want)
	}
}

func TestEncodingABundleIsDeterministic(t *testing.T) {
	f, s := sampleFight()
	b := client.FightBundle{Summary: s, Events: []byte("PAR1"), Metrics: client.MetricsRowsOf(f, s)}
	_, first, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	_, second, err := b.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatal("two encodings of the same bundle differ")
	}
}

func TestLiveRawAndCompleteReachTheServer(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)
	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "unlisted"})
	if err != nil {
		t.Fatal(err)
	}
	_, s := sampleFight()
	if err := c.PutLive(t.Context(), rep.ID, 2, client.Live{
		Summary: s, ElapsedMS: 12000, UpdatedAt: time.Unix(1200, 0).UTC()}); err != nil {
		t.Fatal(err)
	}
	plain := []byte("9/26 20:10:00.000  ZONE_CHANGE,2284,\"Sanguine Depths\",8\n")
	packed := pack(t, plain)
	digest := client.SHA256(plain)
	if stored, err := c.PutRaw(t.Context(), rep.ID, 4194304, digest, packed); err != nil || stored != client.Created {
		t.Fatalf("PutRaw = %v, %v", stored, err)
	}
	if stored, err := c.PutRaw(t.Context(), rep.ID, 4194304, digest, packed); err != nil || stored != client.Duplicate {
		t.Fatalf("re-sent PutRaw = %v, %v", stored, err)
	}
	if err := c.Complete(t.Context(), rep.ID, client.Complete{
		FinalOffset: 131072, EngineVersion: session.Version,
		Health: session.Health{Layout: "retail-v16", Lines: 4200}}); err != nil {
		t.Fatal(err)
	}
	stored := srv.Reports()[rep.ID]
	if stored.Live[2].ElapsedMS != 12000 {
		t.Errorf("live = %+v", stored.Live[2])
	}
	if string(stored.Raw[4194304]) != string(packed) {
		t.Errorf("raw = %q", stored.Raw[4194304])
	}
	if stored.Complete == nil || stored.Complete.FinalOffset != 131072 ||
		stored.Complete.Health.Lines != 4200 {
		t.Errorf("complete = %+v", stored.Complete)
	}
}

func TestARawOffsetThatHoldsDifferentBytesIsAConflict(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)
	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "public"})
	if err != nil {
		t.Fatal(err)
	}
	first, second := []byte("first chunk"), []byte("second chunk")
	if _, err := c.PutRaw(t.Context(), rep.ID, 0, client.SHA256(first), pack(t, first)); err != nil {
		t.Fatal(err)
	}
	_, err = c.PutRaw(t.Context(), rep.ID, 0, client.SHA256(second), pack(t, second))
	var ae *client.Error
	if !errors.As(err, &ae) || ae.Status != http.StatusConflict {
		t.Fatalf("err = %v, want a 409", err)
	}
}

func TestARawChunkWhoseHashIsOverTheCompressedFrameIsRefused(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c := dial(t, srv)
	rep, err := c.CreateReport(t.Context(), client.CreateReport{Visibility: "public"})
	if err != nil {
		t.Fatal(err)
	}
	plain := []byte("the decoded bytes are what the server hashes")
	packed := pack(t, plain)
	// Hashing the frame rather than its contents is the mistake the
	// contract's amendment exists to prevent.
	_, err = c.PutRaw(t.Context(), rep.ID, 0, client.SHA256(packed), packed)
	var ae *client.Error
	if !errors.As(err, &ae) || ae.Status != http.StatusBadRequest {
		t.Fatalf("err = %v, want a 400", err)
	}
}

func TestAnUnknownTokenIsReportedAsUnauthorized(t *testing.T) {
	srv := fakeapi.New()
	defer srv.Close()
	c, err := client.New(client.Options{BaseURL: srv.URL, Token: func() string { return "fsd_wrong" }})
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.CreateReport(t.Context(), client.CreateReport{Visibility: "public"})
	if !client.Unauthorized(err) {
		t.Fatalf("err = %v, want unauthorized", err)
	}
}
