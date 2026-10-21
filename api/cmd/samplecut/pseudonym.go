package main

import (
	"fmt"
	"hash/fnv"
	"regexp"
	"strings"
)

// The realm every pseudonym lives on, in the form the client writes into a unit: the
// report's own realm, with the trailing hyphen.
const sampleRealm = "ClassicBetaPvP"

// sampleServer is the server id every pseudonymised GUID carries, the one the fixture uses.
const sampleServer = "4184"

var firstNames = []string{
	"Bow", "Aldric", "Brenna", "Caelum", "Dorin", "Elowen", "Fenwick", "Garrow", "Hesper", "Isolde",
	"Joren", "Kestrel", "Lyra", "Marek", "Nessa", "Orrin", "Perrin", "Quill", "Rhosyn", "Soren",
	"Tamsin", "Ulric", "Vesper", "Wynne", "Xander", "Ysolt", "Zephyr", "Alwen", "Bram", "Cedra",
	"Doran", "Embry", "Farrin", "Gwyn", "Halric", "Ilsa", "Jasper", "Kell", "Lorne", "Maren",
	"Nolan", "Oriel", "Pell", "Quinlan", "Rowan", "Sable", "Torin", "Una", "Vance", "Wren",
	"Yarrow", "Zora", "Ashby", "Briar", "Corwin", "Dara", "Eamon", "Fable", "Gareth", "Hollis",
	"Ivor", "Juno", "Kerrin", "Lowen", "Merrin", "Nym", "Osric", "Penn", "Rook", "Sedge",
	"Thane", "Umber", "Vell", "Wick", "Yewen", "Zane", "Alder", "Birch", "Cobb", "Dunmor",
}

var lastNames = []string{
	"Jackzon", "Ashgrove", "Blackwell", "Crowfoot", "Dunmere", "Embergale", "Fairwind", "Grimtoft", "Hollowell", "Ironbark",
	"Jessamy", "Kingsfold", "Larkspur", "Marrowdale", "Nightbloom", "Oakhaven", "Pennywhistle", "Quickwater", "Ravensmoor", "Stonebrook",
	"Thistledown", "Underhill", "Valewood", "Whitlock", "Yarrowby", "Zephyrine", "Amberley", "Brightmoor", "Cinderfell", "Dawnrider",
	"Elmstead", "Frostmantle", "Glimmerwick", "Hartsong", "Ivywood", "Jadecroft", "Kestrelby", "Lowmere", "Mistral", "Northcott",
	"Oldfield", "Pinecrest", "Redmayne", "Silverlode", "Tanglewood", "Umbermoor", "Vinegrove", "Wolfsbane", "Yewdale", "Ziptoe",
	"Applegarth", "Briarwood", "Copperfield", "Duskmantle", "Everhart", "Fernhollow", "Goldcrest", "Hawthorne", "Inkwell", "Juniper",
	"Knollwood", "Lanternby", "Mossgate", "Nettlebed", "Orchardson", "Prickett", "Quarrel", "Rushmoor", "Saltmarsh", "Tallow",
	"Umberfield", "Vexley", "Windlass", "Yellowby", "Zinnia", "Alderwick", "Barrowby", "Cragmoor", "Dellwater", "Eastmere",
}

// playerRef is a player's GUID, optionally followed by the quoted unit name that comes next
// in a combat log's source and destination fields.
var playerRef = regexp.MustCompile(`Player-\d+-[0-9A-Fa-f]{8}(,"[^"]*")?`)

// Pseudonymiser rewrites a combat log so no real player survives in it: every player GUID
// becomes a stable synthetic one and every player name a stable fantasy two-word name on the
// report's own realm. The mapping depends only on the seed and the log's own text, so the
// same log cut twice is byte-identical.
type Pseudonymiser struct {
	seed  string
	guids map[string]string
	units map[string]string // by original GUID: one pseudonym per player, even if two share a name
	names map[string]bool
	taken map[string]bool
	// Originals is every original unit name seen, for the leak check.
	Originals []string
}

// NewPseudonymiser starts an empty mapping under seed.
func NewPseudonymiser(seed string) *Pseudonymiser {
	return &Pseudonymiser{seed: seed, guids: map[string]string{}, units: map[string]string{}, names: map[string]bool{}, taken: map[string]bool{}}
}

func (p *Pseudonymiser) hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(p.seed + "\x00" + s))
	return h.Sum64()
}

func (p *Pseudonymiser) guid(orig string) string {
	if g, ok := p.guids[orig]; ok {
		return g
	}
	g := fmt.Sprintf("Player-%s-%08X", sampleServer, len(p.guids)+1)
	p.guids[orig] = g
	return g
}

// unit maps a player's original unit name ("Name-Realm-") to a pseudonym in the client's own
// form ("First-Last-ClassicBetaPvP-"). The pseudonym belongs to the GUID: two players that
// shared a name in the log still get two names, and no two players share a pseudonym (a
// collision walks to the next free pair).
func (p *Pseudonymiser) unit(guid, orig string) string {
	if orig == "" || orig == "Unknown" || orig == "nil" {
		return orig
	}
	if u, ok := p.units[guid]; ok {
		return u
	}
	if !p.names[orig] {
		p.names[orig] = true
		p.Originals = append(p.Originals, orig)
	}
	h := p.hash(guid + "|" + orig)
	n := uint64(len(firstNames) * len(lastNames))
	for i := uint64(0); ; i++ {
		k := (h + i) % n
		first, last := firstNames[k%uint64(len(firstNames))], lastNames[k/uint64(len(firstNames))]
		u := first + "-" + last + "-" + sampleRealm + "-"
		if !p.taken[u] {
			p.taken[u] = true
			p.units[guid] = u
			return u
		}
	}
}

// Line rewrites one log line.
func (p *Pseudonymiser) Line(line string) string {
	return playerRef.ReplaceAllStringFunc(line, func(m string) string {
		guid, name, hasName := strings.Cut(m, ",")
		out := p.guid(guid)
		if hasName {
			out += `,"` + p.unit(guid, strings.Trim(name, `"`)) + `"`
		}
		return out
	})
}

// Log rewrites a whole log, line by line, keeping every other byte.
func (p *Pseudonymiser) Log(b []byte) []byte {
	lines := strings.SplitAfter(string(b), "\n")
	var sb strings.Builder
	sb.Grow(len(b))
	for _, l := range lines {
		sb.WriteString(p.Line(l))
	}
	return []byte(sb.String())
}

// BaseNames is the distinct bare name words (without the realm) among the originals, for the
// leak check.
func (p *Pseudonymiser) BaseNames() []string {
	seen := map[string]bool{}
	var out []string
	for _, o := range p.Originals {
		parts := strings.Split(strings.TrimSuffix(o, "-"), "-")
		for _, w := range parts[:max(1, len(parts)-1)] {
			if len(w) >= 3 && !seen[w] {
				seen[w] = true
				out = append(out, w)
			}
		}
	}
	return out
}
