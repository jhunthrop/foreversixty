// companion/internal/app/connect.go
// Sign in with Battle.net, through the site (companion-connect-contract.md).
// The companion mints a connect, opens the site in the player's browser
// and polls until the site decides. It never sees a password and never
// receives a redirect. One connect runs at a time; the poll runs in the
// background, so the player can leave the Device page and the Status
// headline keeps saying what is happening.
package app

import (
	"context"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// The states of a sign-in as the window sees them.
const (
	ConnectIdle    = "idle"
	ConnectWaiting = "waiting"
	// ConnectBlocked is waiting with no browser: the system would not
	// open one, so the window shows the link to copy. Polling goes on.
	ConnectBlocked = "blocked"
	ConnectFailed  = "failed"
)

// The reasons a sign-in failed.
const (
	ReasonDenied  = "denied"
	ReasonExpired = "expired"
	ReasonOffline = "offline"
	ReasonServer  = "server"
	ReasonSave    = "save"
)

// TroubleAfter is how many polls in a row may fail before the window
// is told the site is hard to reach. Until then the retry is silent.
const TroubleAfter = 3

// ConnectView is the sign-in as the snapshot carries it.
type ConnectView struct {
	State    string     `json:"state"`
	UserCode string     `json:"user_code,omitempty"`
	URL      string     `json:"url,omitempty"`
	Expires  *time.Time `json:"expires_at,omitempty"`
	// Trouble is true after TroubleAfter failed polls in a row.
	Trouble bool   `json:"trouble"`
	Reason  string `json:"reason,omitempty"`
}

// connectState is the sign-in's state; mu guards all of it. gen counts
// connects, so a poll that was cancelled can tell and stop writing.
type connectState struct {
	view     ConnectView
	id       string // the connect id: the secret, never shown
	minting  bool
	gen      int
	failures int
	cancel   context.CancelFunc
}

// connectView is a copy of the sign-in for the snapshot.
func (a *App) connectView() ConnectView {
	a.mu.Lock()
	defer a.mu.Unlock()
	v := a.conn.view
	if v.State == "" {
		v.State = ConnectIdle
	}
	return v
}

func (a *App) connectActive() bool {
	s := a.conn.view.State
	return a.conn.minting || s == ConnectWaiting || s == ConnectBlocked
}

// StartConnect begins a sign-in: it mints a connect, opens the
// browser and starts polling. A second call while one is running does
// nothing, so a double click cannot mint two; a paired device refuses.
// Errors are Faults; a connect that could not start is not an error but
// the failed state in the snapshot, because that is where the page
// reads it.
func (a *App) StartConnect(ctx context.Context) error {
	if a.deviceToken() != "" {
		return fault(CodeConnected, nil)
	}
	a.mu.Lock()
	if a.connectActive() {
		a.mu.Unlock()
		return nil
	}
	a.conn.minting = true
	a.conn.view = ConnectView{State: ConnectIdle}
	a.mu.Unlock()

	con, err := a.client.Connect(ctx, client.ConnectRequest{Name: deviceName(), Platform: platform()})

	a.mu.Lock()
	a.conn.minting = false
	if err != nil {
		a.log.Warn("could not start a sign-in", "component", "app", "err", err.Error())
		a.conn.view = ConnectView{State: ConnectFailed, Reason: reasonOf(err)}
		a.mu.Unlock()
		return nil
	}
	expires := a.now().Add(time.Duration(con.ExpiresIn) * time.Second)
	if con.ExpiresIn <= 0 {
		expires = a.now().Add(10 * time.Minute)
	}
	a.conn.gen++
	gen := a.conn.gen
	a.conn.id, a.conn.failures = con.ConnectID, 0
	a.conn.view = ConnectView{State: ConnectWaiting, UserCode: con.UserCode, URL: con.URL, Expires: &expires}
	pollCtx, cancel := context.WithCancel(a.bg)
	a.conn.cancel = cancel
	a.mu.Unlock()

	if err := a.o.Open(con.URL); err != nil {
		a.log.Warn("could not open the browser for the sign-in", "component", "app", "err", err.Error())
		a.mu.Lock()
		if a.conn.gen == gen {
			a.conn.view.State = ConnectBlocked
		}
		a.mu.Unlock()
	}
	interval := time.Duration(con.Interval) * time.Second
	if a.o.ConnectInterval > 0 {
		interval = a.o.ConnectInterval
	}
	a.bgWG.Add(1)
	go func() {
		defer a.bgWG.Done()
		a.pollConnect(pollCtx, gen, con.ConnectID, expires, interval)
	}()
	return nil
}

// reasonOf is the failure reason for an error from the API.
func reasonOf(err error) string {
	if netCode(err, CodeServer) == CodeOffline {
		return ReasonOffline
	}
	return ReasonServer
}

// CancelConnect stops the sign-in and returns to idle.
func (a *App) CancelConnect() { a.cancelConnect() }

func (a *App) cancelConnect() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.endConnect(ConnectView{State: ConnectIdle})
}

// endConnect stops polling and sets the final view. Callers hold a.mu.
func (a *App) endConnect(v ConnectView) {
	a.conn.gen++
	if a.conn.cancel != nil {
		a.conn.cancel()
		a.conn.cancel = nil
	}
	a.conn.id, a.conn.failures = "", 0
	a.conn.view = v
}

// ReopenConnect opens the same link again. It never mints a new
// connect: the code on the screen must stay the code the site shows.
func (a *App) ReopenConnect() error {
	a.mu.Lock()
	if s := a.conn.view.State; s != ConnectWaiting && s != ConnectBlocked {
		a.mu.Unlock()
		return fault(CodeNoConnect, nil)
	}
	gen, link := a.conn.gen, a.conn.view.URL
	a.mu.Unlock()

	err := a.o.Open(link)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.conn.gen != gen {
		return nil
	}
	if err != nil {
		a.log.Warn("could not open the browser for the sign-in", "component", "app", "err", err.Error())
		a.conn.view.State = ConnectBlocked
		return nil
	}
	a.conn.view.State = ConnectWaiting
	return nil
}

// pollConnect polls until the site decides, the connect expires, or
// the sign-in is cancelled (the context ends, or gen moves on).
func (a *App) pollConnect(ctx context.Context, gen int, id string, expires time.Time, every time.Duration) {
	t := time.NewTicker(every)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
		}
		st, err := a.client.ConnectStatus(ctx, id)
		if ctx.Err() != nil {
			return
		}
		if err != nil {
			// The network is allowed to fail: retry silently, and say
			// so only once it has failed several times in a row.
			a.mu.Lock()
			if a.conn.gen == gen {
				a.conn.failures++
				a.conn.view.Trouble = a.conn.failures >= TroubleAfter
			}
			a.mu.Unlock()
			a.log.Debug("a sign-in poll failed", "component", "app", "err", err.Error())
			continue
		}
		a.mu.Lock()
		if a.conn.gen != gen {
			a.mu.Unlock()
			return
		}
		a.conn.failures, a.conn.view.Trouble = 0, false
		a.mu.Unlock()

		switch st.Status {
		case client.ConnectApproved:
			a.finishConnect(gen, st)
			return
		case client.ConnectDenied:
			a.failConnect(gen, ReasonDenied)
			return
		case client.ConnectExpired:
			a.failConnect(gen, ReasonExpired)
			return
		case client.ConnectUsed:
			// Collected already, by an earlier answer this poll never
			// saw: the token is gone and cannot be fetched again.
			a.log.Warn("the sign-in was already collected", "component", "app")
			a.failConnect(gen, ReasonServer)
			return
		}
		if a.now().After(expires) {
			a.failConnect(gen, ReasonExpired)
			return
		}
	}
}

func (a *App) failConnect(gen int, reason string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.conn.gen != gen {
		return
	}
	a.endConnect(ConnectView{State: ConnectFailed, Reason: reason})
}

// finishConnect stores what the approved answer carried, once.
func (a *App) finishConnect(gen int, st client.ConnectState) {
	a.mu.Lock()
	current := a.conn.gen == gen
	a.mu.Unlock()
	if !current {
		return
	}
	err := a.saveDevice(st.Token, st.AccountLabel)
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.conn.gen != gen {
		return
	}
	if err != nil {
		a.endConnect(ConnectView{State: ConnectFailed, Reason: ReasonSave})
		return
	}
	a.endConnect(ConnectView{State: ConnectIdle})
	a.log.Info("this device is connected", "component", "app", "device_id", st.DeviceID)
}
