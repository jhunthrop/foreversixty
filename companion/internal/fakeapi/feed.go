// companion/internal/fakeapi/feed.go
// The release feed the companion's updaters read: one manifest per family at
// /dl/<family>/latest.json and the downloads it names. It lives in the fake
// beside the ingest contract so one test server stands in for every host the
// companion talks to, and so a capture rig can point the updaters at it with
// their feed base option (<server URL>/dl).
package fakeapi

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
)

// FeedAsset is one file on a fake release.
type FeedAsset struct {
	Name string
	Body []byte
}

// FeedRelease is one fake release. Its family is the part of the tag before
// "-v" (addon-v2.0.16 is the addon family's).
type FeedRelease struct {
	Tag    string
	Assets []FeedAsset
}

func (r FeedRelease) family() string {
	f, _, _ := strings.Cut(r.Tag, "-v")
	return f
}

// SetReleases replaces what the feed holds, newest first: each family's
// manifest names the first release of that family. A family with none
// answers 404, as the real feed does before its first release.
func (s *Server) SetReleases(rels ...FeedRelease) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.releases = append([]FeedRelease(nil), rels...)
}

// ManifestReads is how many times a manifest has been read.
func (s *Server) ManifestReads() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.manifestReads
}

// isFeed reports a path the feed routes answer, which carry no device token.
func isFeed(path string) bool {
	return strings.HasPrefix(path, "/dl/")
}

// newestOf is the newest release of family; the caller holds s.mu.
func (s *Server) newestOf(family string) (FeedRelease, bool) {
	for _, rel := range s.releases {
		if rel.family() == family {
			return rel, true
		}
	}
	return FeedRelease{}, false
}

func (s *Server) feedManifest(w http.ResponseWriter, r *http.Request) {
	family := r.PathValue("family")
	s.mu.Lock()
	s.manifestReads++
	rel, ok := s.newestOf(family)
	s.mu.Unlock()
	if !ok {
		http.NotFound(w, r)
		return
	}
	type asset struct {
		Name   string `json:"name"`
		URL    string `json:"browser_download_url"`
		Size   int    `json:"size"`
		SHA256 string `json:"sha256"`
	}
	assets := []asset{}
	for _, a := range rel.Assets {
		sum := sha256.Sum256(a.Body)
		assets = append(assets, asset{Name: a.Name, URL: s.URL + "/dl/" + family + "/" + rel.Tag + "/" + a.Name,
			Size: len(a.Body), SHA256: hex.EncodeToString(sum[:])})
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]any{
		"schema": 1, "family": family, "tag_name": rel.Tag,
		"published_at": "2026-11-06T10:00:00Z", "assets": assets,
	})
}

func (s *Server) feedDownload(w http.ResponseWriter, r *http.Request) {
	family, tag, name := r.PathValue("family"), r.PathValue("tag"), r.PathValue("name")
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, rel := range s.releases {
		if rel.family() != family || rel.Tag != tag {
			continue
		}
		for _, a := range rel.Assets {
			if a.Name == name {
				w.Header().Set("Content-Type", "application/octet-stream")
				_, _ = w.Write(a.Body)
				return
			}
		}
	}
	http.NotFound(w, r)
}
