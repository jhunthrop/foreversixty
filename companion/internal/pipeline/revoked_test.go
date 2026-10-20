package pipeline_test

import (
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
	"github.com/jhunthrop/foreversixty/companion/internal/fixture"
)

// A device the site no longer knows is told apart from an outage: the
// drain's error says Unauthorized, and the bundle that met it stays
// queued, so the night uploads after the next sign-in.
func TestAFightUploadTheSiteAnswers401KeepsTheBundleQueued(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone + fixture.Encounter(0) + fixture.Heartbeat(0))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.srv.RevokeToken()
	r.write(fixture.Encounter(1) + fixture.Heartbeat(1))
	err := r.tick()
	if err == nil || !client.Unauthorized(err) {
		t.Fatalf("the drain answered %v; want an unauthorized error", err)
	}
	n, qerr := r.q.Len()
	if qerr != nil {
		t.Fatal(qerr)
	}
	if n == 0 {
		t.Fatal("the bundle that met the 401 left the queue")
	}
	// Signing in again resumes it.
	r.srv.RestoreToken()
	if err := r.pipe.Drain(t.Context()); err != nil {
		t.Fatalf("drain after the sign-in: %v", err)
	}
	if n, _ := r.q.Len(); n != 0 {
		t.Fatalf("%d items still queued after the sign-in", n)
	}
}

func TestALiveSnapshotTheSiteAnswers401IsReported(t *testing.T) {
	r := newRig(t, 1<<20)
	r.write(fixture.Header + fixture.Zone)
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	// An encounter with no end: the fight stays open.
	open := strings.SplitAfter(fixture.Encounter(0), "\n")
	r.write(strings.Join(open[:len(open)-2], ""))
	if err := r.tick(); err != nil {
		t.Fatal(err)
	}
	r.srv.RevokeToken()
	r.clock = r.clock.Add(10 * time.Second)
	if err := r.pipe.Live(t.Context(), r.clock); err == nil || !client.Unauthorized(err) {
		t.Fatalf("Live answered %v; want an unauthorized error", err)
	}
}
