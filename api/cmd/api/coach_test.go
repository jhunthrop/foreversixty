package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/coach"
	"github.com/jhunthrop/foreversixty/api/internal/config"
	"github.com/jhunthrop/foreversixty/api/internal/db"
	"github.com/jhunthrop/foreversixty/api/internal/entitlements"
	"github.com/jhunthrop/foreversixty/api/internal/guilds"
	"github.com/jhunthrop/foreversixty/api/internal/jobs"
	"github.com/jhunthrop/foreversixty/api/internal/llm"
	"github.com/jhunthrop/foreversixty/api/internal/rating"
	"github.com/jhunthrop/foreversixty/api/internal/reports"
	"github.com/jhunthrop/foreversixty/api/internal/server"
	"github.com/jhunthrop/foreversixty/api/internal/sims"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

const (
	repoBuildsDir = "../../../data/builds"
	repoCoachDir  = "../../../data/curated/coach"
	activeBuildJS = "../../../web/src/data/active-build.json"
)

// fakeBucket answers no object; mountCoach only needs one to exist.
type fakeBucket struct{}

func (fakeBucket) Get(context.Context, string) (io.ReadCloser, error) { return nil, os.ErrNotExist }

// boundJobs records the jobs a wiring bound runners to.
type boundJobs struct {
	mu   sync.Mutex
	jobs []string
	fail bool
}

func (b *boundJobs) factory(_ context.Context, job string) (jobs.Runner, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.fail {
		return nil, errors.New("no cloud here")
	}
	b.jobs = append(b.jobs, job)
	return &jobs.Fake{}, nil
}

// testWiring is the switch off, no LLM variables, a fake bucket and fake runners; the log
// goes to buf.
func testWiring(t *testing.T, buf *bytes.Buffer, runners *boundJobs) coachWiring {
	t.Helper()
	url := os.Getenv("TEST_DATABASE_URL")
	if url == "" {
		t.Skip("TEST_DATABASE_URL not set")
	}
	if err := db.Migrate(url); err != nil {
		t.Fatal(err)
	}
	pool, err := db.Connect(context.Background(), url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	cfg := config.Config{TreeDataDir: repoBuildsDir, PublicBaseURL: "https://example.test",
		Coach: config.CoachConfig{DataDir: repoCoachDir, DailySpendCapUSD: 25, DailyReviewsPerAccount: 30}}
	return coachWiring{
		cfg: cfg, pool: pool, deps: &server.Deps{Guilds: &guilds.Service{}},
		rating: &rating.Store{Pool: pool}, reports: &reports.Store{Pool: pool}, auth: &auth.Store{Pool: pool},
		entitle: &entitlements.Store{Pool: pool}, sims: &sims.Store{Pool: pool}, summaries: fakeBucket{},
		log: slog.New(slog.NewTextHandler(buf, nil)), newRunner: runners.factory,
	}
}

func TestMountCoach_SwitchOffNeedsNoKey(t *testing.T) {
	var buf bytes.Buffer
	w := testWiring(t, &buf, &boundJobs{})
	svc := mountCoach(context.Background(), w)
	if svc == nil || svc.Enabled {
		t.Fatalf("mounts with the switch off: %v (log %s)", svc, buf.String())
	}
	if _, err := svc.Generator.LLM.Complete(context.Background(), llm.Request{}); !errors.Is(err, errSwitchedOff) {
		t.Fatalf("the model refuses while the switch is off: %v", err)
	}
	router := server.NewRouter(server.Deps{Coach: svc, Log: w.log})
	for _, path := range []string{"/v1/me/coach/recent", "/v1/reports/x/coach/night", "/v1/reports/x/fights/1/coach",
		"/v1/characters/us/normal/x/coach/history", "/v1/admin/coach/samples"} {
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code == http.StatusNotFound && strings.Contains(rec.Body.String(), "no such route") {
			t.Errorf("%s is not mounted", path)
		}
	}
}

func TestMountCoach_SetsTheRatingHook(t *testing.T) {
	var buf bytes.Buffer
	w := testWiring(t, &buf, &boundJobs{})
	svc := mountCoach(context.Background(), w)
	if svc == nil {
		t.Fatal("mounts")
	}
	if _, ok := w.rating.Hook.(*coach.ProfileWriter); !ok {
		t.Fatalf("hook %T", w.rating.Hook)
	}
	if w.deps.Guilds.CoachHolders != svc {
		t.Fatalf("guilds coach holders %T", w.deps.Guilds.CoachHolders)
	}
}

func TestMountCoach_NoBucketIsNil(t *testing.T) {
	var buf bytes.Buffer
	w := testWiring(t, &buf, &boundJobs{})
	w.summaries = bucketOrNil(nil)
	if svc := mountCoach(context.Background(), w); svc != nil {
		t.Fatal("no bucket, no coach")
	}
	if w.rating.Hook != nil || strings.Count(buf.String(), "level=WARN") != 1 || !strings.Contains(buf.String(), "effect=") {
		t.Fatalf("hook %v, log %q", w.rating.Hook, buf.String())
	}
}

func TestMountCoach_UnreadableDataDirIsNil(t *testing.T) {
	var buf bytes.Buffer
	w := testWiring(t, &buf, &boundJobs{})
	w.cfg.Coach.DataDir = t.TempDir() + "/missing"
	if svc := mountCoach(context.Background(), w); svc != nil || w.rating.Hook != nil {
		t.Fatal("an unreadable COACH_DATA_DIR mounts nothing")
	}
	if !strings.Contains(buf.String(), "COACH_DATA_DIR") {
		t.Fatalf("the warn line names the variable: %q", buf.String())
	}
}

func TestMountCoach_OneRunnerPerJob(t *testing.T) {
	var buf bytes.Buffer
	runners := &boundJobs{}
	svc := mountCoach(context.Background(), testWiring(t, &buf, runners))
	if svc.Jobs == nil || svc.Refs.Sims.Jobs == nil {
		t.Fatal("both runners bound")
	}
	if strings.Join(runners.jobs, ",") != coach.SimRefJobCommand+","+coach.ReviewJobCommand {
		t.Fatalf("bound %v", runners.jobs)
	}
}

func TestMountCoach_UnreachableRunnersLeaveFieldsNil(t *testing.T) {
	var buf bytes.Buffer
	svc := mountCoach(context.Background(), testWiring(t, &buf, &boundJobs{fail: true}))
	if svc == nil || svc.Jobs != nil || svc.Refs.Sims.Jobs != nil {
		t.Fatalf("%+v", svc)
	}
	if ok, _ := svc.Refs.Validated.Validated(context.Background(), "warrior-fury"); ok {
		t.Fatal("without the sim-ref job no spec waits on a sim")
	}
	if strings.Count(buf.String(), "level=WARN") != 2 {
		t.Fatalf("one warn line per runner: %q", buf.String())
	}
}

func TestRunCoachJobs_Table(t *testing.T) {
	want := map[string]bool{coach.ReviewJobCommand: false, coach.SimRefJobCommand: false, coach.SamplesJobCommand: false}
	for _, j := range extraJobs {
		if _, ok := want[j.name]; ok {
			want[j.name] = true
		}
	}
	for name, found := range want {
		if !found {
			t.Errorf("extraJobs lacks %s", name)
		}
	}
	runJobTable(slog.Default(), extraJobs, []string{"coach-unknown"}, func(int) { t.Fatal("an unknown command exits") })
}

func TestSamplesDay(t *testing.T) {
	now := time.Date(2026, 10, 28, 3, 0, 0, 0, time.UTC)
	if d, err := samplesDay(nil, now); err != nil || d.Format(time.DateOnly) != "2026-10-27" {
		t.Fatalf("%v %v", d, err)
	}
	if d, err := samplesDay([]string{"2026-10-01"}, now); err != nil || d.Format(time.DateOnly) != "2026-10-01" {
		t.Fatalf("%v %v", d, err)
	}
	if _, err := samplesDay([]string{"yesterday"}, now); err == nil {
		t.Fatal("not a day")
	}
	if _, err := samplesDay([]string{"a", "b"}, now); err == nil {
		t.Fatal("two days")
	}
}

func TestCoachLinesOf_NilServiceIsNilLines(t *testing.T) {
	if coachLinesOf(nil) != nil {
		t.Fatal("a nil service must not become non-nil lines")
	}
	if bucketOrNil(nil) != nil {
		t.Fatal("a nil client must not become a non-nil getter")
	}
}

func TestRatingStoreWithCoach_WritesProfiles(t *testing.T) {
	s := ratingStoreWithCoach(config.Config{}, nil, slog.Default())
	if _, ok := s.Hook.(*coach.ProfileWriter); !ok {
		t.Fatalf("hook %T", s.Hook)
	}
}

func TestLadderCharacters_TheLaddersBandCharacter(t *testing.T) {
	raw, err := os.ReadFile(activeBuildJS)
	if err != nil {
		t.Fatal(err)
	}
	var active struct{ Build string }
	if err := json.Unmarshal(raw, &active); err != nil {
		t.Fatal(err)
	}
	data, err := trees.Load(repoBuildsDir)
	if err != nil {
		t.Fatal(err)
	}
	sims.SetGearData(data, repoBuildsDir)
	ch, err := ladderCharacters{dataDir: repoBuildsDir}.Character(context.Background(), active.Build, "warrior-fury", 60)
	if err != nil {
		t.Fatal(err)
	}
	if ch.Class != "warrior" || ch.Level != 60 || ch.Race == "" || ch.Talents == "" || len(ch.Gear) < 10 {
		t.Fatalf("%+v", ch)
	}
	if _, err := (ladderCharacters{dataDir: repoBuildsDir}).Character(context.Background(), active.Build, "warrior-fury", 7); err == nil {
		t.Fatal("no band 7 on the ladder")
	}
}

func TestBuildAt_NoBuildsIsEmptyAndLogged(t *testing.T) {
	var buf bytes.Buffer
	log := slog.New(slog.NewTextHandler(&buf, nil))
	if at := buildAt(&coach.FileTables{BuildsDir: t.TempDir()}, log); at(time.Now()) != "" || !strings.Contains(buf.String(), "build_at") {
		t.Fatalf("no builds: empty and one warn line, got %q", buf.String())
	}
	if at := buildAt(&coach.FileTables{BuildsDir: repoBuildsDir}, log); at(time.Now()) == "" {
		t.Fatal("now has a build")
	}
}
