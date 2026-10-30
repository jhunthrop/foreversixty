package sheet

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"golang.org/x/sync/singleflight"
)

// The sheets' in-process LRU (spec Part C): 32 sheets, at most 48 MB.
const (
	sheetCacheEntries = 32
	sheetCacheBytes   = 48 << 20
)

// Renderer renders sheets through an LRU keyed by the SHA-256 of the model's canonical JSON,
// so two shares of the same answer draw once, even when they arrive together (singleflight). It
// is safe for concurrent use.
type Renderer struct {
	icons  IconSource
	cache  *lru
	flight singleflight.Group
}

// NewRenderer renders with icons.
func NewRenderer(icons IconSource) *Renderer {
	return &Renderer{icons: icons, cache: newLRU(sheetCacheEntries, sheetCacheBytes)}
}

// Render is the sheet's PNG, from the cache when the same model was drawn before.
func (r *Renderer) Render(s Sheet) ([]byte, error) {
	key, err := sheetKey(s)
	if err != nil {
		return nil, err
	}
	if png, ok := r.cache.get(key); ok {
		return png, nil
	}
	v, err, _ := r.flight.Do(key, func() (any, error) {
		png, err := Render(s, r.icons)
		if err != nil {
			return nil, err
		}
		r.cache.put(key, png)
		return png, nil
	})
	if err != nil {
		return nil, err
	}
	return v.([]byte), nil
}

func sheetKey(s Sheet) (string, error) {
	raw, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("sheet: key: %w", err)
	}
	sum := sha256.Sum256(raw)
	return hex.EncodeToString(sum[:]), nil
}
