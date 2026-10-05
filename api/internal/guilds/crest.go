// api/internal/guilds/crest.go
//
// A guild's uploaded crest (docs/contracts/2026-10-05-guild-crest-api.md): officers and the
// claimed guild's leader manage it, a moderator may remove it, and it is served publicly -
// the header ring falls back to the faction logo whenever none is set.
package guilds

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/jackc/pgx/v5"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/httpx"
	"github.com/jhunthrop/foreversixty/api/internal/imagex"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// maxCrestUploadBytes bounds the whole multipart body a crest PUT may send: the
// contract's own 2 MiB limit on the image field, plus headroom for the multipart
// boundary and headers around it. A file right at the limit still parses whole; a
// wildly oversized one is cut off by the reader itself well before imagex ever sees
// it, rather than read in full first.
const maxCrestUploadBytes = imagex.MaxBytes + (1 << 20)

// crestCacheControl is the contract's own header for the served crest object: public,
// a year, and immutable, since a new upload always lives at a new key (its content
// hash), never this one rewritten in place.
const crestCacheControl = "public, max-age=31536000, immutable"

// crestKeyHashChars is how many hex characters of the processed crest's own sha256
// name its object key - the contract's "sha256-prefix": enough to make a collision
// astronomically unlikely, short enough to keep the key and the crest_url's ?v= tidy.
const crestKeyHashChars = 16

// Objects is the object store the crest routes write to, read from, and delete from.
// *r2.Client satisfies it structurally; tests use an in-memory fake.
type Objects interface {
	store.Putter
	Get(ctx context.Context, key string) (io.ReadCloser, error)
	Delete(ctx context.Context, key string) error
}

// crestKeyFor is the contract's own object key shape for guildID's crest, named from
// the processed WebP's own content hash so a new upload is always a new key.
func crestKeyFor(guildID int64, sha256Hex string) string {
	return fmt.Sprintf("guilds/%d/crest/%s.webp", guildID, sha256Hex[:crestKeyHashChars])
}

// crestVersion is the "?v=" the contract's crest_url carries: the key's own hash
// segment, read back out of it rather than recomputed, so CrestURL never needs the
// original sha256 - only whatever key a guild's row already carries.
func crestVersion(key string) string {
	base := path.Base(key)
	return strings.TrimSuffix(base, path.Ext(base))
}

// CrestURL builds the contract's crest_url for guildID from its stored crest_key -
// nil when crestKey is nil, exactly mirroring the "`crest_url ?? faction logo`" rule
// the contract asks every reader of a guild mark to share. apiBaseURL is this
// service's own public base URL (config.APIBaseURL); an empty one (a harness with no
// such config) yields a origin-relative URL rather than a malformed absolute one.
func CrestURL(apiBaseURL string, guildID int64, crestKey *string) *string {
	if crestKey == nil {
		return nil
	}
	url := fmt.Sprintf("%s/v1/guilds/%d/crest.webp?v=%s", apiBaseURL, guildID, crestVersion(*crestKey))
	return &url
}

// crestETag is the strong ETag GET .../crest.webp answers with: the object is
// content-hash-addressed by its own key, so the key's hash segment alone is already
// a correct, strong validator - no need to hash the body again.
func crestETag(key string) string { return `"` + crestVersion(key) + `"` }

// CrestKey reads guildID's stored crest object key, nil when none is set.
func (s *Store) CrestKey(ctx context.Context, guildID int64) (*string, error) {
	var key *string
	err := s.Pool.QueryRow(ctx, `select crest_key from guilds where id = $1`, guildID).Scan(&key)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: read crest key for %d: %w", guildID, err)
	}
	return key, nil
}

// SetCrest stores guildID's new crest key and stamps crest_updated_at, reporting the
// key that was previously stored (nil when none was) so the caller can delete that
// now-orphaned R2 object once the new one is safely recorded.
func (s *Store) SetCrest(ctx context.Context, guildID int64, key string) (*string, error) {
	var previous *string
	err := s.Pool.QueryRow(ctx, `
		with prev as (select crest_key from guilds where id = $1)
		update guilds set crest_key = $2, crest_updated_at = now()
		where id = $1
		returning (select crest_key from prev)`, guildID, key).Scan(&previous)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: set crest for %d: %w", guildID, err)
	}
	return previous, nil
}

// ClearCrest clears guildID's crest key, reporting the key that was previously
// stored (nil when none was) so the caller can delete that R2 object.
func (s *Store) ClearCrest(ctx context.Context, guildID int64) (*string, error) {
	var previous *string
	err := s.Pool.QueryRow(ctx, `
		with prev as (select crest_key from guilds where id = $1)
		update guilds set crest_key = null, crest_updated_at = now()
		where id = $1
		returning (select crest_key from prev)`, guildID).Scan(&previous)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("guilds: clear crest for %d: %w", guildID, err)
	}
	return previous, nil
}

// deleteOrphanedCrest removes key from R2, logging rather than failing the request
// on error: the database already reflects the new state (the new crest stored, or
// the crest cleared), so a stray object left behind in R2 is a cleanup miss, never a
// reason to answer the caller as if their request failed.
func (s *Service) deleteOrphanedCrest(ctx context.Context, key string) {
	if err := s.R2.Delete(ctx, key); err != nil {
		s.logger().Error("guilds", "op", "crest_cleanup", "key", key, "err", err)
	}
}

func (s *Service) putCrest(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrLeader(r, guildID)
	if err != nil {
		s.fail(w, r, "crest_put", err, "could not save that crest just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild to change its crest", nil)
		return
	}
	if frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator()); err != nil {
		s.fail(w, r, "crest_put", err, "could not save that crest just now")
		return
	} else if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	if s.R2 == nil {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable",
			"guild crests are not configured on this deployment", nil)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxCrestUploadBytes)
	if err := r.ParseMultipartForm(maxCrestUploadBytes); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"that upload is too large or is not valid multipart form data", nil)
		return
	}
	defer func() {
		if r.MultipartForm != nil {
			_ = r.MultipartForm.RemoveAll()
		}
	}()
	file, _, err := r.FormFile("image")
	if err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid",
			"send the image as a multipart field named \"image\"", nil)
		return
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		s.fail(w, r, "crest_put", err, "could not save that crest just now")
		return
	}

	processed, procErr := imagex.ProcessCrest(data)
	var verr *imagex.ValidationError
	if errors.As(procErr, &verr) {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", verr.Message, nil)
		return
	}
	if procErr != nil {
		s.fail(w, r, "crest_put", procErr, "could not process that image")
		return
	}

	key := crestKeyFor(guildID, processed.SHA256)
	if err := s.R2.Put(r.Context(), key, processed.WebP, store.PutOptions{
		ContentType: "image/webp", CacheControl: crestCacheControl,
	}); err != nil {
		s.fail(w, r, "crest_put", err, "could not save that crest just now")
		return
	}
	previous, err := s.Store.SetCrest(r.Context(), guildID, key)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "crest_put", err, "could not save that crest just now")
		return
	}
	if previous != nil && *previous != key {
		s.deleteOrphanedCrest(r.Context(), *previous)
	}
	s.logger().Info("guilds", "op", "crest_put", "guild_id", guildID, "user_id", actor.UserID)
	httpx.WriteOK(w, r, http.StatusOK, map[string]any{"crest_url": CrestURL(s.Store.APIBaseURL, guildID, &key)})
}

func (s *Service) deleteCrest(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	actor := auth.ActorFrom(r.Context())
	allowed, err := s.verifiedOfficerOrModerator(r, guildID)
	if err != nil {
		s.fail(w, r, "crest_delete", err, "could not remove that crest just now")
		return
	}
	if !allowed {
		httpx.WriteError(w, r, http.StatusForbidden, "forbidden",
			"you must be a verified officer of this guild, or a moderator, to remove its crest", nil)
		return
	}
	if frozen, err := s.freezeCheck(r.Context(), guildID, actor.IsModerator()); err != nil {
		s.fail(w, r, "crest_delete", err, "could not remove that crest just now")
		return
	} else if frozen {
		httpx.WriteError(w, r, http.StatusConflict, "claim_contested",
			"this guild's claim is contested; officer actions are frozen until a moderator resolves it", nil)
		return
	}
	if s.R2 == nil {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable",
			"guild crests are not configured on this deployment", nil)
		return
	}
	previous, err := s.Store.ClearCrest(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "crest_delete", err, "could not remove that crest just now")
		return
	}
	if previous != nil {
		s.deleteOrphanedCrest(r.Context(), *previous)
	}
	s.logger().Info("guilds", "op", "crest_delete", "guild_id", guildID, "user_id", actor.UserID)
	w.WriteHeader(http.StatusNoContent)
}

// getCrest serves a guild's crest image directly - no envelope, no session: every
// place that shows a guild mark (the header ring, eventually the data addon) reads
// this unauthenticated, exactly like the faction logo it falls back to.
func (s *Service) getCrest(w http.ResponseWriter, r *http.Request) {
	guildID, ok := guildIDFrom(r)
	if !ok {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if s.R2 == nil {
		httpx.WriteError(w, r, http.StatusServiceUnavailable, "unavailable",
			"guild crests are not configured on this deployment", nil)
		return
	}
	key, err := s.Store.CrestKey(r.Context(), guildID)
	if errors.Is(err, ErrNotFound) {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "no such guild", nil)
		return
	}
	if err != nil {
		s.fail(w, r, "crest_get", err, "could not load that crest just now")
		return
	}
	if key == nil {
		httpx.WriteError(w, r, http.StatusNotFound, "not_found", "this guild has no crest", nil)
		return
	}
	etag := crestETag(*key)
	if inm := r.Header.Get("If-None-Match"); inm != "" && inm == etag {
		w.Header().Set("Cache-Control", crestCacheControl)
		w.Header().Set("ETag", etag)
		w.WriteHeader(http.StatusNotModified)
		return
	}
	body, err := s.R2.Get(r.Context(), *key)
	if err != nil {
		s.fail(w, r, "crest_get", err, "could not load that crest just now")
		return
	}
	defer body.Close()
	w.Header().Set("Content-Type", "image/webp")
	w.Header().Set("Cache-Control", crestCacheControl)
	w.Header().Set("ETag", etag)
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, body)
}
