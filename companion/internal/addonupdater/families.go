// companion/internal/addonupdater/families.go
// Package addonupdater keeps the Forever Sixty addon and its data addon
// current in every game folder the companion knows. It is the companion's
// own updater (internal/updater) pointed at two more release families:
// the same GitHub releases, the same minisign key compiled into the
// binary, and nothing written that the key has not vouched for.
//
// Two rules shape it. The game is never running while a folder under
// Interface/AddOns changes (a release that arrives mid-session waits in
// the companion's home), and only the three folders the releases ship
// are ever written or removed, each swapped in by a rename that can be
// undone. SavedVariables are never touched.
package addonupdater

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// bom is the byte order mark an editor may put at the start of a TOC.
const bom = "\xef\xbb\xbf"

// Family is one kind of release and the folders it installs.
type Family struct {
	// Key names the family in file names and the log.
	Key string
	// TagPrefix is what a release tag of this family starts with.
	TagPrefix string
	// Folders are the top-level folders of the release's zip, which are
	// the only folders ever written under AddOns. Folders[0] carries the
	// version the family is judged by.
	Folders []string
	// Dates says the version is a date (2026.10.21), compared as one.
	Dates bool
}

// The two families.
var (
	// Addon is ForeverSixty and the inbox addon beside it, released as
	// addon-v<semver>.
	Addon = Family{Key: "addon", TagPrefix: "addon-v", Folders: []string{"ForeverSixty", "ForeverSixtyInbox"}}
	// Data is the nightly ratings addon, released as data-v<yyyy.mm.dd>.
	Data = Family{Key: "data", TagPrefix: "data-v", Folders: []string{"ForeverSixtyData"}, Dates: true}
)

// Families is every family, addon first: the data addon is only ever
// installed beside a main addon.
var Families = []Family{Addon, Data}

// Main is the folder whose TOC carries the family's version.
func (f Family) Main() string { return f.Folders[0] }

// AddonsDir is where an install keeps its addons.
func AddonsDir(install string) string {
	return filepath.Join(install, "Interface", "AddOns")
}

// NormalizeVersion turns what a TOC or a tag says into a bare version:
// "addon-v2.0.16", "data-v2026.10.21" and "v2.0.16" all lose their prefix.
func NormalizeVersion(raw string) string {
	v := strings.TrimSpace(strings.TrimPrefix(raw, bom))
	for _, p := range []string{Addon.TagPrefix, Data.TagPrefix, "v", "V"} {
		if strings.HasPrefix(v, p) {
			v = strings.TrimPrefix(v, p)
			break
		}
	}
	return v
}

// VersionOfTag is the version a release tag of this family names.
func (f Family) VersionOfTag(tag string) string {
	return NormalizeVersion(strings.TrimPrefix(tag, f.TagPrefix))
}

// ReadVersion reads "## Version" from <addons>/<folder>/<folder>.toc. found
// is false when there is no such TOC; version is empty when the TOC has no
// readable version line.
func ReadVersion(addonsDir, folder string) (version string, found bool) {
	f, err := os.Open(filepath.Join(addonsDir, folder, folder+".toc"))
	if err != nil {
		return "", false
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 4096), 64<<10)
	first := true
	for sc.Scan() {
		line := sc.Text()
		if first {
			line = strings.TrimPrefix(line, bom)
			first = false
		}
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "##") {
			continue
		}
		key, val, ok := strings.Cut(strings.TrimSpace(strings.TrimPrefix(line, "##")), ":")
		if ok && strings.EqualFold(strings.TrimSpace(key), "Version") {
			return NormalizeVersion(val), true
		}
	}
	return "", true
}

// Newer reports whether candidate is a later version of the family than
// installed. A date family compares as dates; the addon compares as dotted
// integers. An installed version that cannot be read is older than any
// candidate; a candidate that cannot be read is never newer.
func (f Family) Newer(candidate, installed string) bool {
	c, i := NormalizeVersion(candidate), NormalizeVersion(installed)
	if c == "" {
		return false
	}
	if f.Dates {
		cd, ok := parseDate(c)
		if !ok {
			return false
		}
		id, ok := parseDate(i)
		if !ok {
			return true
		}
		return cd.After(id)
	}
	if i == "" {
		return true
	}
	return compareDotted(c, i) > 0
}

// parseDate reads yyyy.mm.dd (or with dashes) as a date.
func parseDate(v string) (time.Time, bool) {
	parts := strings.FieldsFunc(v, func(r rune) bool { return r == '.' || r == '-' })
	if len(parts) != 3 {
		return time.Time{}, false
	}
	var n [3]int
	for k, p := range parts {
		x, err := strconv.Atoi(p)
		if err != nil || x <= 0 {
			return time.Time{}, false
		}
		n[k] = x
	}
	if n[0] < 2000 || n[1] > 12 || n[2] > 31 {
		return time.Time{}, false
	}
	return time.Date(n[0], time.Month(n[1]), n[2], 0, 0, 0, 0, time.UTC), true
}

func compareDotted(a, b string) int {
	as, bs := strings.Split(a, "."), strings.Split(b, ".")
	for i := range max(len(as), len(bs)) {
		x, y := part(as, i), part(bs, i)
		if x != y {
			if x < y {
				return -1
			}
			return 1
		}
	}
	return 0
}

func part(s []string, i int) int {
	if i >= len(s) {
		return 0
	}
	n, err := strconv.Atoi(strings.SplitN(s[i], "-", 2)[0])
	if err != nil {
		return 0
	}
	return n
}
