package sims

import "net/http"

// This handler is written in Task 7, which deletes its stub from this
// file; the file goes with it.
func (s *Service) simInput(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) }
