package sheet

import (
	"errors"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// tileCacheEntries is the tiles' in-process LRU (spec Part C): 2,048 tiles of about 10 KB.
const tileCacheEntries = 2048

// TileItems is what the tiles route reads from the item data: whether a build is known, and an
// item's icon and quality in it.
type TileItems interface {
	KnownBuild(build string) bool
	TileItem(build string, itemID int) (icon string, quality int, ok bool)
}

// Tiles serves GET /v1/discord/tiles/{build}/{item}.png.
type Tiles struct {
	items TileItems
	icons func(build string) IconSource
	log   *slog.Logger
	cache *lru
}

// NewTiles serves tiles from items, reading each build's icons from {dataDir}/{build}/icons.
func NewTiles(dataDir string, items TileItems, log *slog.Logger) *Tiles {
	return newTiles(items, func(build string) IconSource {
		return NewFileIcons(filepath.Join(dataDir, build, "icons"))
	}, log)
}

func newTiles(items TileItems, icons func(string) IconSource, log *slog.Logger) *Tiles {
	if log == nil {
		log = slog.Default()
	}
	return &Tiles{items: items, icons: icons, log: log, cache: newLRU(tileCacheEntries, 0)}
}

// Mount registers the tiles route.
func (t *Tiles) Mount(mux *http.ServeMux) {
	mux.HandleFunc("GET "+TilePath+"{build}/{file}", t.serve)
}

func (t *Tiles) serve(w http.ResponseWriter, r *http.Request) {
	build := r.PathValue("build")
	id, ok := tileItemID(r.PathValue("file"))
	if !ok || !t.items.KnownBuild(build) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such tile", nil)
		return
	}
	tile := t.lookup(build, id)
	png, hit := t.cache.get(tile.key)
	if !hit {
		var err error
		png, err = t.render(build, &tile)
		if err != nil {
			t.log.Error("discord tile", "op", "tile", "build", build, "item", id, "err", err)
			httpx.WriteError(w, r, http.StatusInternalServerError, "internal", "could not draw the tile", nil)
			return
		}
		if tile.cacheable {
			t.cache.put(tile.key, png)
		}
	}
	if tile.provisional {
		httpx.CachePublic(w, provisionalTileMaxAge, 0)
	} else {
		httpx.CachePublicImmutable(w)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(png)))
	w.WriteHeader(http.StatusOK)
	if r.Method != http.MethodHead {
		_, _ = w.Write(png)
	}
}

// provisionalTileMaxAge is how long a question-mark tile may be kept: an item the build does
// not know yet, or whose icon file is missing, may be drawn properly after the next data
// deploy, so it is never immutable.
const provisionalTileMaxAge = time.Hour

// tileSpec is what one request draws and where it is cached.
type tileSpec struct {
	key         string // the LRU key: build/id, or build/? for every item the build lacks
	icon        string
	quality     int
	provisional bool // the question mark: a short cache, not immutable
	cacheable   bool // kept in the LRU; false for a known item whose icon file is missing
}

// lookup resolves a request to its tile. Every item the build does not know shares one key,
// build/?, so random ids cannot churn the LRU.
func (t *Tiles) lookup(build string, id int) tileSpec {
	icon, quality, known := t.items.TileItem(build, id)
	if !known {
		return tileSpec{key: build + "/?", icon: QuestionMark, quality: -1, provisional: true, cacheable: true}
	}
	return tileSpec{key: build + "/" + strconv.Itoa(id), icon: icon, quality: quality, cacheable: true}
}

// render draws one tile. A known item whose icon file is missing draws the question mark in
// its quality ring (logged), provisional and kept out of the LRU, so it heals on the next deploy.
func (t *Tiles) render(build string, tile *tileSpec) ([]byte, error) {
	src := t.icons(build)
	img, err := src.Icon(tile.icon)
	if errors.Is(err, ErrNoIcon) && tile.icon != QuestionMark {
		t.log.Warn("discord tile", "op", "sheet_icon_missing", "build", build, "icon", tile.icon)
		tile.provisional, tile.cacheable = true, false
		img, err = src.Icon(QuestionMark)
	}
	if err != nil {
		return nil, err
	}
	return Tile(img, tile.quality)
}

// tileItemID reads "{digits}.png"; anything else is no tile.
func tileItemID(file string) (int, bool) {
	digits, ok := strings.CutSuffix(file, ".png")
	if !ok || digits == "" || len(digits) > 9 {
		return 0, false
	}
	for _, c := range digits {
		if c < '0' || c > '9' {
			return 0, false
		}
	}
	id, err := strconv.Atoi(digits)
	return id, err == nil && id > 0
}

// TreesItems is TileItems over the loaded item tables: a build is known when it is loaded,
// and an item is found in any class's table.
type TreesItems struct {
	Data *trees.Data
}

// KnownBuild reports whether build is a loaded build folder.
func (t TreesItems) KnownBuild(build string) bool {
	if t.Data == nil {
		return false
	}
	_, ok := t.Data.Build(build)
	return ok
}

// TileItem is the item's icon and quality from the first class table that holds it.
func (t TreesItems) TileItem(build string, itemID int) (string, int, bool) {
	if t.Data == nil {
		return "", 0, false
	}
	b, ok := t.Data.Build(build)
	if !ok {
		return "", 0, false
	}
	for _, c := range b.Classes() {
		if it, ok := b.Item(c.ID, itemID); ok && it.Icon != "" {
			return it.Icon, it.Quality, true
		}
	}
	return "", 0, false
}
