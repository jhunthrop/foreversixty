// companion/internal/app/http.go
// The local HTTP interface the webview talks to. Everything hangs off
// the session token in the path; the JSON below is the whole contract
// between ui/app.js and the Go side.
package app

import (
	"encoding/json"
	"io"
	"io/fs"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/browser"
	"github.com/jhunthrop/foreversixty/companion/ui"
)

// MaxRequestBody bounds a request from the local page. Nothing it
// sends is larger than a settings object.
const MaxRequestBody = 64 << 10

// Handler is the local UI server's routes.
func (a *App) Handler() http.Handler {
	assets, err := fs.Sub(ui.Assets, ".")
	if err != nil {
		panic(err) // the embedded filesystem is a build-time fact
	}
	inner := http.NewServeMux()
	inner.Handle("GET /", http.FileServerFS(assets))
	inner.HandleFunc("GET /api/status", a.handleStatus)
	inner.HandleFunc("POST /api/pair", a.handlePair)
	inner.HandleFunc("POST /api/unpair", a.handleUnpair)
	inner.HandleFunc("POST /api/settings", a.handleSettings)
	inner.HandleFunc("POST /api/connect", a.handleConnect)
	inner.HandleFunc("POST /api/connect/cancel", a.handleConnectCancel)
	inner.HandleFunc("POST /api/connect/reopen", a.handleConnectReopen)
	inner.HandleFunc("POST /api/open", a.handleOpen)
	inner.HandleFunc("POST /api/folders/pick", a.handleFolderPick)
	inner.HandleFunc("POST /api/folders/remove", a.handleFolderRemove)

	outer := http.NewServeMux()
	outer.Handle("/"+a.token+"/", http.StripPrefix("/"+a.token, inner))
	outer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "not found", http.StatusNotFound)
	})
	return outer
}

// writeJSON answers with the same envelope shape the API uses, so
// ui/app.js has one way to read a response.
func writeJSON(w http.ResponseWriter, status int, data any, message string) {
	writeBody(w, status, data, message, "")
}

func writeBody(w http.ResponseWriter, status int, data any, message, code string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	body := map[string]any{"ok": message == "", "data": data, "error": nil}
	if message != "" {
		e := map[string]string{"message": message}
		if code != "" {
			e["code"] = code
		}
		body["error"] = e
	}
	// The status line is already on the wire, so a failure here can
	// only mean the local page went away mid-response: there is no
	// second answer to send and nothing for the player to do.
	_ = json.NewEncoder(w).Encode(body)
}

// writeFault answers a failure with the sentence for its code, never
// with the error's own text: that goes to the log, and the window shows
// this to the player as it is.
func (a *App) writeFault(w http.ResponseWriter, status int, err error, fallback string) {
	code := codeOf(err, fallback)
	if code == CodeInternal {
		a.log.Error("a request failed", "component", "ui", "err", err.Error())
	}
	writeBody(w, status, nil, Message(code), code)
}

func (a *App) decode(w http.ResponseWriter, r *http.Request, out any) bool {
	b, err := io.ReadAll(io.LimitReader(r.Body, MaxRequestBody))
	if err == nil {
		err = json.Unmarshal(b, out)
	}
	if err != nil {
		a.writeFault(w, http.StatusBadRequest, fault(CodeBadRequest, err), CodeBadRequest)
		return false
	}
	return true
}

// statusOf is the HTTP status for a failure code.
func statusOf(err error) int {
	switch codeOf(err, CodeInternal) {
	case CodeInternal, CodeSave, CodeSettings:
		return http.StatusInternalServerError
	case CodeConnected, CodeNoConnect, CodeKnownFolder, CodeNotAdded:
		return http.StatusConflict
	case CodeOffline, CodeServer:
		return http.StatusBadGateway
	case CodeNoPicker, CodeNoBrowser:
		return http.StatusServiceUnavailable
	}
	return http.StatusBadRequest
}

func (a *App) answer(w http.ResponseWriter, err error) {
	if err != nil {
		a.writeFault(w, statusOf(err), err, CodeInternal)
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handleStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handlePair(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Code string `json:"code"`
	}
	if !a.decode(w, r, &in) {
		return
	}
	a.answer(w, a.Pair(r.Context(), in.Code))
}

func (a *App) handleUnpair(w http.ResponseWriter, r *http.Request) {
	a.answer(w, a.Unpair())
}

func (a *App) handleSettings(w http.ResponseWriter, r *http.Request) {
	var in Settings
	if !a.decode(w, r, &in) {
		return
	}
	a.answer(w, a.SaveSettings(in))
}

func (a *App) handleConnect(w http.ResponseWriter, r *http.Request) {
	a.answer(w, a.StartConnect(r.Context()))
}

func (a *App) handleConnectCancel(w http.ResponseWriter, r *http.Request) {
	a.CancelConnect()
	a.answer(w, nil)
}

func (a *App) handleConnectReopen(w http.ResponseWriter, r *http.Request) {
	a.answer(w, a.ReopenConnect())
}

// handleOpen opens a link in the system browser: the one way the
// window leaves itself. Only https links go (and loopback http, for a
// development site); anything else is refused before it can reach the
// operating system.
func (a *App) handleOpen(w http.ResponseWriter, r *http.Request) {
	var in struct {
		URL string `json:"url"`
	}
	if !a.decode(w, r, &in) {
		return
	}
	if err := browser.Check(in.URL); err != nil {
		a.answer(w, fault(CodeBadURL, err))
		return
	}
	if err := a.o.Open(in.URL); err != nil {
		a.log.Warn("could not open a link", "component", "ui", "err", err.Error())
		a.answer(w, fault(CodeNoBrowser, err))
		return
	}
	a.answer(w, nil)
}

func (a *App) handleFolderPick(w http.ResponseWriter, r *http.Request) {
	a.answer(w, a.PickFolder())
}

func (a *App) handleFolderRemove(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Path string `json:"path"`
	}
	if !a.decode(w, r, &in) {
		return
	}
	a.answer(w, a.RemoveFolder(in.Path))
}
