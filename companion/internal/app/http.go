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
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")
	w.WriteHeader(status)
	body := map[string]any{"ok": message == "", "data": data, "error": nil}
	if message != "" {
		body["error"] = map[string]string{"message": message}
	}
	json.NewEncoder(w).Encode(body)
}

func decode(w http.ResponseWriter, r *http.Request, out any) bool {
	b, err := io.ReadAll(io.LimitReader(r.Body, MaxRequestBody))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, nil, "could not read the request")
		return false
	}
	if err := json.Unmarshal(b, out); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, "could not read the request: "+err.Error())
		return false
	}
	return true
}

func (a *App) handleStatus(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handlePair(w http.ResponseWriter, r *http.Request) {
	var in struct {
		Code string `json:"code"`
	}
	if !decode(w, r, &in) {
		return
	}
	if err := a.Pair(r.Context(), in.Code); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handleUnpair(w http.ResponseWriter, r *http.Request) {
	if err := a.Unpair(); err != nil {
		writeJSON(w, http.StatusInternalServerError, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}

func (a *App) handleSettings(w http.ResponseWriter, r *http.Request) {
	var in Settings
	if !decode(w, r, &in) {
		return
	}
	if err := a.SaveSettings(in); err != nil {
		writeJSON(w, http.StatusBadRequest, nil, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, a.Snapshot(), "")
}
