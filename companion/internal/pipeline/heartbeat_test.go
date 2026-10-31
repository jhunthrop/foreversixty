package pipeline_test

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
)

// beatRig is a rig whose pipeline beats every 20 ms and which has a
// report open and known to the site, the state L.4 calls "waiting for
// the first fight".
func beatRig(t *testing.T) *rig {
	t.Helper()
	r := newRig(t, 1<<20)
	r.beatEvery = 20 * time.Millisecond
	r.open()
	// A fresh watcher must see the empty folder before the log appears.
	if err := r.pipe.Tick(r.clock); err != nil {
		t.Fatal(err)
	}
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	return r
}

// until polls cond for up to two seconds.
func until(t *testing.T, what string, cond func() bool) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func TestTheHeartbeatPostsAtTheIntervalAndStopsWhenTheReportCloses(t *testing.T) {
	r := beatRig(t)
	until(t, "three heartbeats", func() bool { return r.srv.HeartbeatRequests() >= 3 })
	r.settle() // the log goes quiet and the report closes
	if st := r.pipe.Status(); st.Logging {
		t.Fatal("the report did not close")
	}
	time.Sleep(60 * time.Millisecond) // a beat already in flight lands
	n := r.srv.HeartbeatRequests()
	time.Sleep(120 * time.Millisecond)
	if got := r.srv.HeartbeatRequests(); got != n {
		t.Fatalf("heartbeats went from %d to %d after the report closed", n, got)
	}
	if r.srv.Heartbeats() == 0 {
		t.Fatal("the open report received no heartbeat")
	}
}

func TestAHeartbeatThatFailsIsLoggedOnceAndRetriedOnTheNextTick(t *testing.T) {
	r := beatRig(t)
	r.srv.FailNext("POST /v1/reports/{id}/heartbeat", 3)
	// A beat already sent before FailNext can land first, so wait for the
	// three refusals themselves, then for a beat accepted after them.
	refused := func() int { return r.srv.HeartbeatRequests() - r.srv.Heartbeats() }
	until(t, "three refused heartbeats", func() bool { return refused() >= 3 })
	accepted := r.srv.Heartbeats()
	until(t, "a heartbeat after three failures", func() bool { return r.srv.Heartbeats() > accepted })
	r.pipe.Close() // waits for the goroutine, so the log is safe to read
	if n := strings.Count(r.logs.String(), "a heartbeat did not land"); n != 1 {
		t.Fatalf("the failure was logged %d times; want once per report\n%s", n, r.logs.String())
	}
}

func TestA429OnTheHeartbeatIsHonouredToItsRetryAfter(t *testing.T) {
	r := beatRig(t)
	r.srv.LimitHeartbeats(1, 1)
	until(t, "the refused heartbeat", func() bool { return r.srv.HeartbeatRequests() >= 1 })
	// The tick is 20 ms and the site said one second: nothing else is
	// sent in between.
	time.Sleep(400 * time.Millisecond)
	if n := r.srv.HeartbeatRequests(); n != 1 {
		t.Fatalf("%d heartbeats inside the Retry-After; want 1", n)
	}
	until(t, "the heartbeat after the wait", func() bool { return r.srv.Heartbeats() > 0 })
}

func TestAHeartbeatTheSiteAnswers401StopsAndIsLeftToTheSignedOutPath(t *testing.T) {
	r := beatRig(t)
	until(t, "a first heartbeat", func() bool { return r.srv.HeartbeatRequests() >= 1 })
	r.srv.RevokeToken()
	var refused error
	until(t, "the refusal", func() bool {
		refused = r.pipe.HeartbeatRefused()
		return refused != nil
	})
	if !client.Unauthorized(refused) {
		t.Fatalf("the refusal is %v; want an unauthorized error", refused)
	}
	if again := r.pipe.HeartbeatRefused(); again != nil {
		t.Fatalf("the refusal was handed out twice: %v", again)
	}
	time.Sleep(60 * time.Millisecond)
	n := r.srv.HeartbeatRequests()
	time.Sleep(120 * time.Millisecond)
	if got := r.srv.HeartbeatRequests(); got != n {
		t.Fatalf("heartbeats went from %d to %d after the 401", n, got)
	}
	// A new sign-in starts it again for the report still open.
	r.srv.RestoreToken()
	r.pipe.ResumeHeartbeat()
	until(t, "a heartbeat after the sign-in", func() bool { return r.srv.HeartbeatRequests() > n })
}

func TestCloseStopsTheHeartbeat(t *testing.T) {
	r := beatRig(t)
	until(t, "a heartbeat", func() bool { return r.srv.HeartbeatRequests() >= 1 })
	r.pipe.Close()
	n := r.srv.HeartbeatRequests()
	time.Sleep(100 * time.Millisecond)
	if got := r.srv.HeartbeatRequests(); got != n {
		t.Fatalf("heartbeats went from %d to %d after Close", n, got)
	}
}
