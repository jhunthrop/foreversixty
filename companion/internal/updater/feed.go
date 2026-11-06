// companion/internal/updater/feed.go
package updater

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// DefaultFeedBase is where the release manifests live: one per family at
// <base>/<family>/latest.json, written by the release workflows on our own
// domain so the updaters depend on nothing but the site.
const DefaultFeedBase = "https://foreversixty.gg/dl"

// FeedFamily is the companion's own family in the feed.
const FeedFamily = "companion"

// ManifestSchema is the manifest layout this build reads. A manifest of
// another schema is an error rather than a guess at what its fields mean.
const ManifestSchema = 1

// maxManifest bounds a manifest read; one names a handful of assets.
const maxManifest = 1 << 20

// Asset is one file a manifest names. The field names match the GitHub
// release object the workflows used to publish, so one shape covers both.
type Asset struct {
	Name string `json:"name"`
	URL  string `json:"browser_download_url"`
	// SHA256 is the hex digest of the file, checked when present. The
	// signature stays the authority; the digest only catches a truncated
	// or swapped download early and says so plainly.
	SHA256 string `json:"sha256"`
}

// Release is one family's manifest: the newest signed release of it. The
// workflows write a release into the feed only once it is signed, and the
// manifest only ever names the newest, so there is no list to page through
// and no draft or prerelease to step over.
type Release struct {
	Schema  int     `json:"schema"`
	Family  string  `json:"family"`
	TagName string  `json:"tag_name"`
	Assets  []Asset `json:"assets"`
}

// Feed reads manifests and the downloads they name. Every download must sit
// on the feed's own origin: a manifest that points elsewhere is refused
// before anything is fetched, so the feed cannot be used to pull bytes from
// a host we do not run.
type Feed struct {
	Base string
	HTTP *http.Client
}

// Latest reads family's manifest. A 404 means the family has no release in
// the feed yet, which is "no update", not a failure: the zero Release comes
// back with no error.
func (f Feed) Latest(ctx context.Context, family string) (Release, error) {
	base, err := f.base()
	if err != nil {
		return Release{}, err
	}
	manifest := base.JoinPath(family, "latest.json").String()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, manifest, nil)
	if err != nil {
		return Release{}, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := f.HTTP.Do(req)
	if err != nil {
		return Release{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return Release{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return Release{}, fmt.Errorf("the release feed answered %d for the %s manifest", resp.StatusCode, family)
	}
	var rel Release
	if err := json.NewDecoder(io.LimitReader(resp.Body, maxManifest)).Decode(&rel); err != nil {
		return Release{}, fmt.Errorf("the %s manifest is unreadable: %w", family, err)
	}
	if rel.Schema != ManifestSchema {
		return Release{}, fmt.Errorf("the %s manifest has schema %d; this build reads %d", family, rel.Schema, ManifestSchema)
	}
	if rel.Family != family {
		return Release{}, fmt.Errorf("the %s manifest names the family %q", family, rel.Family)
	}
	return rel, nil
}

// Fetch downloads one asset, refusing a URL off the feed's origin, anything
// over MaxAsset, and a body whose digest does not match the manifest's.
func (f Feed) Fetch(ctx context.Context, a Asset) ([]byte, error) {
	if err := f.sameOrigin(a.URL); err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, a.URL, nil)
	if err != nil {
		return nil, err
	}
	resp, err := f.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("downloading %s: %d", a.Name, resp.StatusCode)
	}
	b, err := io.ReadAll(io.LimitReader(resp.Body, MaxAsset+1))
	if err != nil {
		return nil, err
	}
	if len(b) > MaxAsset {
		return nil, fmt.Errorf("downloading %s: over the %d byte ceiling", a.Name, MaxAsset)
	}
	if err := checkDigest(a, b); err != nil {
		return nil, err
	}
	return b, nil
}

// base parses the feed base. It must be https; plain http is allowed only
// on a loopback host, which is where a test or capture rig serves one.
func (f Feed) base() (*url.URL, error) {
	u, err := url.Parse(f.Base)
	if err != nil || u.Host == "" {
		return nil, fmt.Errorf("the release feed base %q is not a URL", f.Base)
	}
	switch {
	case u.Scheme == "https":
	case u.Scheme == "http" && isLoopback(u.Hostname()):
	default:
		return nil, fmt.Errorf("the release feed base %q is not https", f.Base)
	}
	return u, nil
}

// sameOrigin refuses a download URL whose scheme or host differs from the
// feed base's.
func (f Feed) sameOrigin(raw string) error {
	base, err := f.base()
	if err != nil {
		return err
	}
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Errorf("the manifest names an unreadable download URL %q", raw)
	}
	if u.Scheme != base.Scheme || !strings.EqualFold(u.Host, base.Host) || u.User != nil {
		return fmt.Errorf("the manifest names a download off the feed's origin (%s), so it was not fetched: %q",
			base.Scheme+"://"+base.Host, raw)
	}
	return nil
}

func checkDigest(a Asset, body []byte) error {
	if a.SHA256 == "" {
		return nil
	}
	want, err := hex.DecodeString(a.SHA256)
	if err != nil || len(want) != sha256.Size {
		return fmt.Errorf("the manifest's sha256 for %s is unreadable", a.Name)
	}
	got := sha256.Sum256(body)
	if !bytes.Equal(got[:], want) {
		return fmt.Errorf("%s does not match the manifest's sha256", a.Name)
	}
	return nil
}

func isLoopback(host string) bool {
	if strings.EqualFold(host, "localhost") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}
