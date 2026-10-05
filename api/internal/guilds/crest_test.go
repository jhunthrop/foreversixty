// api/internal/guilds/crest_test.go
package guilds

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/HugoSmits86/nativewebp"

	"github.com/jhunthrop/foreversixty/api/internal/auth"
	"github.com/jhunthrop/foreversixty/api/internal/imagex"
	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

// memObjects is an in-memory Objects fake: the Put/Get/Delete the crest
// routes need, in a map rather than real R2.
type memObjects struct {
	mu      sync.Mutex
	objects map[string][]byte
	headers map[string]store.PutOptions
	deleted []string
}

func newMemObjects() *memObjects {
	return &memObjects{objects: map[string][]byte{}, headers: map[string]store.PutOptions{}}
}

func (m *memObjects) Put(_ context.Context, key string, body []byte, o store.PutOptions) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.objects[key] = append([]byte{}, body...)
	m.headers[key] = o
	return nil
}

func (m *memObjects) Get(_ context.Context, key string) (io.ReadCloser, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	b, ok := m.objects[key]
	if !ok {
		return nil, errors.New("no such object: " + key)
	}
	return io.NopCloser(bytes.NewReader(b)), nil
}

func (m *memObjects) Delete(_ context.Context, key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.objects, key)
	m.deleted = append(m.deleted, key)
	return nil
}

func (m *memObjects) has(key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	_, ok := m.objects[key]
	return ok
}

// newCrestHarness is newHTTPHarness's own setup, plus a wired R2 fake and an
// APIBaseURL, so putCrest/deleteCrest/getCrest have everything they need.
func newCrestHarness(t *testing.T) (*httpHarness, *memObjects) {
	t.Helper()
	pool := testPool(t)
	accounts := &auth.Store{Pool: pool}
	quiet := slog.New(slog.NewTextHandler(io.Discard, nil))
	objects := newMemObjects()
	store := &Store{Pool: pool, APIBaseURL: "https://api.forever.test"}
	svc := &Service{Store: store, Accounts: accounts, R2: objects, Log: quiet}
	mux := http.NewServeMux()
	Mount(mux, svc, 0)
	h := &httpHarness{t: t, pool: pool, store: store}
	h.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mux.ServeHTTP(w, r.WithContext(auth.WithActor(r.Context(), h.actor)))
	}))
	t.Cleanup(h.server.Close)
	return h, objects
}

// genTestPNG builds a plain w x h PNG - no need for anything fancier here,
// imagex's own tests already cover crop/resize/alpha correctness.
func genTestPNG(t *testing.T, w, h int) []byte {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.NRGBA{uint8(x % 256), uint8(y % 256), 200, 255})
		}
	}
	buf := &bytes.Buffer{}
	if err := png.Encode(buf, img); err != nil {
		t.Fatal(err)
	}
	return buf.Bytes()
}

func (h *httpHarness) putCrestMultipart(t *testing.T, guildPath string, data []byte) *http.Response {
	t.Helper()
	body := &bytes.Buffer{}
	mw := multipart.NewWriter(body)
	part, err := mw.CreateFormFile("image", "crest.png")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := part.Write(data); err != nil {
		t.Fatal(err)
	}
	if err := mw.Close(); err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest(http.MethodPut, h.server.URL+guildPath, body)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", mw.FormDataContentType())
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// seedOfficer seeds a verified officer character and account-level
// membership row for guildID, wiring the harness's actor to it.
func seedOfficer(t *testing.T, h *httpHarness, guildID int64, email, key string) int64 {
	t.Helper()
	uid := seedUser(t, h.pool, email)
	seedCharacter(t, h.pool, guildID, uid, key, "officer", true)
	syncMembership(t, h.pool, guildID, uid)
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}
	return uid
}

func TestPutCrestOfficerUploadSucceeds(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-officer@example.com", "us/hardcore/crestofficer")

	res := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200))
	if res.StatusCode != http.StatusOK {
		t.Fatalf("officer upload = %d, want 200", res.StatusCode)
	}
	var data struct {
		CrestURL string `json:"crest_url"`
	}
	h.data(res, &data)
	if !strings.Contains(data.CrestURL, fmt.Sprintf("/v1/guilds/%d/crest.webp?v=", gid)) {
		t.Fatalf("crest_url = %q, want the contract's own shape", data.CrestURL)
	}
	if !strings.HasPrefix(data.CrestURL, "https://api.forever.test") {
		t.Fatalf("crest_url = %q, want it built against the configured API base URL", data.CrestURL)
	}
}

func TestGetCrestServesA256WebPWithImmutableCacheHeaders(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-get@example.com", "us/hardcore/crestget")

	put := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 300, 150))
	if put.StatusCode != http.StatusOK {
		t.Fatalf("upload = %d, want 200", put.StatusCode)
	}

	res, err := http.Get(h.server.URL + fmt.Sprintf("/v1/guilds/%d/crest.webp", gid))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("GET crest.webp = %d, want 200", res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "image/webp" {
		t.Fatalf("Content-Type = %q, want image/webp", ct)
	}
	if cc := res.Header.Get("Cache-Control"); !strings.Contains(cc, "immutable") || !strings.Contains(cc, "max-age=31536000") {
		t.Fatalf("Cache-Control = %q, want immutable, max-age=31536000", cc)
	}
	if res.Header.Get("ETag") == "" {
		t.Fatal("want an ETag")
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	img, err := nativewebp.Decode(bytes.NewReader(body))
	if err != nil {
		t.Fatalf("served body did not decode as WebP: %v", err)
	}
	b := img.Bounds()
	if b.Dx() != imagex.CrestSize || b.Dy() != imagex.CrestSize {
		t.Fatalf("served crest = %dx%d, want %dx%d", b.Dx(), b.Dy(), imagex.CrestSize, imagex.CrestSize)
	}
}

func TestGetCrestIs404WhenUnset(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")

	res, err := http.Get(h.server.URL + fmt.Sprintf("/v1/guilds/%d/crest.webp", gid))
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("GET crest.webp on an unset crest = %d, want 404", res.StatusCode)
	}
}

func TestHomeCarriesTheCrestURL(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-home@example.com", "us/hardcore/cresthome")

	put := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200))
	if put.StatusCode != http.StatusOK {
		t.Fatalf("upload = %d, want 200", put.StatusCode)
	}
	var uploaded struct {
		CrestURL string `json:"crest_url"`
	}
	h.data(put, &uploaded)

	res := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var home struct {
		Guild struct {
			CrestURL *string `json:"crest_url"`
		} `json:"guild"`
	}
	h.data(res, &home)
	if home.Guild.CrestURL == nil || *home.Guild.CrestURL != uploaded.CrestURL {
		t.Fatalf("home guild.crest_url = %v, want %q", home.Guild.CrestURL, uploaded.CrestURL)
	}
}

func TestPutCrestReplacesAndDeletesTheOldObject(t *testing.T) {
	h, objects := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-replace@example.com", "us/hardcore/crestreplace")

	first := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200))
	var firstData struct {
		CrestURL string `json:"crest_url"`
	}
	h.data(first, &firstData)
	firstKey := crestKeyFromURL(t, firstData.CrestURL)
	if !objects.has(firstKey) {
		t.Fatal("the first object should be stored")
	}

	second := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 400, 160))
	var secondData struct {
		CrestURL string `json:"crest_url"`
	}
	h.data(second, &secondData)
	secondKey := crestKeyFromURL(t, secondData.CrestURL)

	if firstKey == secondKey {
		t.Fatal("two different uploads should not collide on the same key")
	}
	if objects.has(firstKey) {
		t.Fatal("replacing the crest should delete the old object")
	}
	if !objects.has(secondKey) {
		t.Fatal("the new object should be stored")
	}
}

func TestPutCrestRefusesAPlainMember(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	uid := seedUser(t, h.pool, "crest-member@example.com")
	seedCharacter(t, h.pool, gid, uid, "us/hardcore/crestmember", "member", true)
	syncMembership(t, h.pool, gid, uid)
	h.actor = auth.Actor{UserID: uid, Role: "user", Method: "session"}

	res := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200))
	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("a plain member uploading = %d, want 403", res.StatusCode)
	}
}

func TestPutCrestRejectsAnOversizeFileWithTheExactMessage(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-oversize@example.com", "us/hardcore/crestoversize")

	oversize := make([]byte, imagex.MaxBytes+100*1024)
	res := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), oversize)
	if res.StatusCode != http.StatusBadRequest {
		t.Fatalf("oversize upload = %d, want 400", res.StatusCode)
	}
	var env struct {
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
	defer res.Body.Close()
	if err := json.NewDecoder(res.Body).Decode(&env); err != nil {
		t.Fatal(err)
	}
	if env.Error.Code != "invalid" {
		t.Fatalf("code = %q, want invalid", env.Error.Code)
	}
	if !strings.Contains(env.Error.Message, "the limit is 2 MB") {
		t.Fatalf("message = %q, want it to state the 2 MB limit", env.Error.Message)
	}
}

func TestDeleteCrestClearsTheURLAndDeletesTheObject(t *testing.T) {
	h, objects := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-delete@example.com", "us/hardcore/crestdelete")

	put := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200))
	var uploaded struct {
		CrestURL string `json:"crest_url"`
	}
	h.data(put, &uploaded)
	key := crestKeyFromURL(t, uploaded.CrestURL)

	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/crest", gid), "")
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("delete = %d, want 204", res.StatusCode)
	}
	if objects.has(key) {
		t.Fatal("delete should remove the R2 object")
	}

	home := h.do(http.MethodGet, fmt.Sprintf("/v1/guilds/%d/home", gid), "")
	var homeData struct {
		Guild struct {
			CrestURL *string `json:"crest_url"`
		} `json:"guild"`
	}
	h.data(home, &homeData)
	if homeData.Guild.CrestURL != nil {
		t.Fatalf("crest_url after delete = %v, want null", homeData.Guild.CrestURL)
	}

	get, err := http.Get(h.server.URL + fmt.Sprintf("/v1/guilds/%d/crest.webp", gid))
	if err != nil {
		t.Fatal(err)
	}
	defer get.Body.Close()
	if get.StatusCode != http.StatusNotFound {
		t.Fatalf("GET crest.webp after delete = %d, want 404", get.StatusCode)
	}
}

func TestDeleteCrestAllowsAModeratorOnAnyGuild(t *testing.T) {
	h, _ := newCrestHarness(t)
	gid := seedGuild(t, h.pool, "Forever")
	seedOfficer(t, h, gid, "crest-mod-owner@example.com", "us/hardcore/crestmodowner")
	if put := h.putCrestMultipart(t, fmt.Sprintf("/v1/guilds/%d/crest", gid), genTestPNG(t, 200, 200)); put.StatusCode != http.StatusOK {
		t.Fatalf("setup upload = %d, want 200", put.StatusCode)
	}

	modID := seedUser(t, h.pool, "crest-moderator@example.com")
	h.actor = auth.Actor{UserID: modID, Role: "moderator", Method: "session"}

	res := h.do(http.MethodDelete, fmt.Sprintf("/v1/guilds/%d/crest", gid), "")
	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("moderator delete = %d, want 204", res.StatusCode)
	}
}

// crestKeyFromURL recovers the R2 object key from a crest_url the handler
// returned, by re-deriving it from the ?v= hash segment the contract
// carries - the same segment crestKeyFor names the object with.
func crestKeyFromURL(t *testing.T, url string) string {
	t.Helper()
	_, v, ok := strings.Cut(url, "?v=")
	if !ok {
		t.Fatalf("crest_url = %q, want a ?v= query", url)
	}
	idx := strings.LastIndex(url, "/v1/guilds/")
	if idx < 0 {
		t.Fatalf("crest_url = %q, want /v1/guilds/", url)
	}
	rest := url[idx+len("/v1/guilds/"):]
	id, _, _ := strings.Cut(rest, "/")
	return "guilds/" + id + "/crest/" + v + ".webp"
}
