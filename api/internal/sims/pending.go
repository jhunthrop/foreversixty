package sims

import "net/http"

// These three handlers are written in Tasks 5, 6 and 7. Each of those
// tasks deletes its stub from this file; the file goes with the last
// of them.
func (s *Service) run(w http.ResponseWriter, r *http.Request)      { http.NotFound(w, r) }
func (s *Service) specs(w http.ResponseWriter, r *http.Request)    { http.NotFound(w, r) }
func (s *Service) simInput(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) }
