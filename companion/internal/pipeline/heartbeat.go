// companion/internal/pipeline/heartbeat.go
// The heartbeat: while a report is open and its log is being followed,
// the companion tells the site once a minute that it is alive, so the
// report page can say when it has gone quiet instead of guessing from
// the last fight.
package pipeline

import (
	"context"
	"errors"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// HeartbeatEvery is the API's own limit: one heartbeat a minute per
// device. Sending faster only earns a 429.
const HeartbeatEvery = time.Minute

// heartbeat is the goroutine of one open report.
type heartbeat struct {
	cancel context.CancelFunc
}

// startHeartbeat begins the heartbeat for the report that just opened
// or resumed, and ends the previous report's if it somehow still runs.
// It runs under p.mu. The goroutine ends with the report (complete), with
// the pipeline (Close), or on a refused token.
func (p *Pipeline) startHeartbeat() {
	p.stopHeartbeat()
	if p.cur == nil || p.hbCtx.Err() != nil {
		return
	}
	ctx, cancel := context.WithCancel(p.hbCtx)
	p.hb = &heartbeat{cancel: cancel}
	p.hbRefused = nil
	p.hbWG.Add(1)
	go p.beat(ctx, p.cur.Key)
}

// stopHeartbeat cancels the running heartbeat, if any, without waiting
// for it: the goroutine takes p.mu, which the caller holds. It runs
// under p.mu.
func (p *Pipeline) stopHeartbeat() {
	if p.hb != nil {
		p.hb.cancel()
		p.hb = nil
	}
}

// ResumeHeartbeat starts the heartbeat again for the open report after
// a refused token has been replaced by a new sign-in. It does nothing
// when no report is open or the heartbeat is already running.
func (p *Pipeline) ResumeHeartbeat() {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil || p.cur.Closed || p.hb != nil {
		return
	}
	p.startHeartbeat()
}

// HeartbeatRefused returns, once, the error with which the site refused
// the device token on a heartbeat. The heartbeat has stopped by then;
// the caller hands the error to the signed-out path, which is the one
// place that decides what a dead token means.
func (p *Pipeline) HeartbeatRefused() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	err := p.hbRefused
	p.hbRefused = nil
	return err
}

// beat posts the heartbeat for the report with this key every
// HeartbeatEvery until ctx ends. A failure is logged once per report at
// WARN and retried on the next tick, never in a loop; a 429 waits for
// the Retry-After the site named when that is longer than a tick; a
// refused token ends the goroutine and is left to the signed-out path.
func (p *Pipeline) beat(ctx context.Context, key string) {
	defer p.hbWG.Done()
	every := p.o.HeartbeatEvery
	wait := every
	warned := false
	for {
		t := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			t.Stop()
			return
		case <-t.C:
		}
		wait = every
		id := p.openReportID(key)
		if id == "" {
			continue // the site has not been told of this report yet
		}
		err := p.o.Client.Heartbeat(ctx, id)
		switch {
		case err == nil:
			p.o.Log.Debug("heartbeat sent", "component", "pipeline", "report", key, "report_id", id)
		case ctx.Err() != nil:
			return
		case client.Unauthorized(err):
			p.mu.Lock()
			if p.hb != nil && ctx.Err() == nil {
				p.hbRefused, p.hb = err, nil
			}
			p.mu.Unlock()
			return
		default:
			var ae *client.Error
			if errors.As(err, &ae) && ae.RetryAfter > wait {
				wait = ae.RetryAfter
			}
			if !warned {
				warned = true
				p.o.Log.Warn("a heartbeat did not land; trying again on the next tick",
					"component", "pipeline", "report", key, "report_id", id, "err", err.Error())
			}
		}
	}
}

// openReportID is the server id of the open report with this key, or
// the empty string when that report has closed or has none yet.
func (p *Pipeline) openReportID(key string) string {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.cur == nil || p.cur.Key != key {
		return ""
	}
	return p.cur.ReportID
}
