// companion/internal/fakeapi/github.go
// The part of GitHub's REST API the companion's updaters read: the release
// list of a repository and the downloads behind its assets. It lives in the
// fake beside the ingest contract so one test server stands in for every
// host the companion talks to, and so a capture rig can point the updaters
// at it with their API base option.
package fakeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// GitHubAsset is one file on a fake release.
type GitHubAsset struct {
	Name string
	Body []byte
}

// GitHubRelease is one fake release.
type GitHubRelease struct {
	Tag        string
	Prerelease bool
	Assets     []GitHubAsset
}

// SetReleases replaces the release list, newest first, as GitHub orders it.
func (s *Server) SetReleases(rels ...GitHubRelease) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.releases = append([]GitHubRelease(nil), rels...)
}

// ReleaseListReads is how many times the release list has been read.
func (s *Server) ReleaseListReads() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.releaseReads
}

// isGitHub reports a path the GitHub routes answer, which carry no device
// token.
func isGitHub(path string) bool {
	return strings.HasPrefix(path, "/repos/") || strings.HasPrefix(path, "/releasedl/")
}

func (s *Server) listReleases(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.releaseReads++
	rels := append([]GitHubRelease(nil), s.releases...)
	s.mu.Unlock()
	per, _ := strconv.Atoi(r.URL.Query().Get("per_page"))
	if per <= 0 {
		per = 30
	}
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	if page <= 0 {
		page = 1
	}
	type asset struct {
		Name string `json:"name"`
		URL  string `json:"browser_download_url"`
	}
	type row struct {
		Tag        string  `json:"tag_name"`
		Draft      bool    `json:"draft"`
		Prerelease bool    `json:"prerelease"`
		Assets     []asset `json:"assets"`
	}
	out := []row{}
	for i, rel := range rels {
		if i < (page-1)*per || i >= page*per {
			continue
		}
		rr := row{Tag: rel.Tag, Prerelease: rel.Prerelease, Assets: []asset{}}
		for j, a := range rel.Assets {
			rr.Assets = append(rr.Assets, asset{Name: a.Name, URL: fmt.Sprintf("%s/releasedl/%d/%d", s.URL, i, j)})
		}
		out = append(out, rr)
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(out)
}

func (s *Server) downloadAsset(w http.ResponseWriter, r *http.Request) {
	i, errI := strconv.Atoi(r.PathValue("release"))
	j, errJ := strconv.Atoi(r.PathValue("asset"))
	s.mu.Lock()
	defer s.mu.Unlock()
	if errI != nil || errJ != nil || i < 0 || i >= len(s.releases) || j < 0 || j >= len(s.releases[i].Assets) {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	_, _ = w.Write(s.releases[i].Assets[j].Body)
}
