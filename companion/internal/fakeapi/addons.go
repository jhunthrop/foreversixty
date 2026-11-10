package fakeapi

import (
	"encoding/json"
	"net/http"

	"github.com/jhunthrop/foreversixty/companion/internal/client"
)

// AddonReports are the bodies PUT /v1/devices/self/addons accepted, in order.
func (s *Server) AddonReports() [][]client.InstalledAddon {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([][]client.InstalledAddon(nil), s.addonReports...)
}

// putAddons is PUT /v1/devices/self/addons: the device's AddOns scan, stored as sent.
func (s *Server) putAddons(w http.ResponseWriter, r *http.Request) {
	if s.shouldFail("PUT /v1/devices/self/addons") {
		fail(w, http.StatusInternalServerError, "could not save the addons just now")
		return
	}
	var in struct {
		Addons []client.InstalledAddon `json:"addons"`
	}
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil || in.Addons == nil {
		fail(w, http.StatusBadRequest, "body must be JSON with an addons list")
		return
	}
	s.mu.Lock()
	s.addonReports = append(s.addonReports, in.Addons)
	s.mu.Unlock()
	w.WriteHeader(http.StatusNoContent)
}
