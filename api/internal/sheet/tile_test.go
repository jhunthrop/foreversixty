package sheet

import (
	"image"
	"image/color"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

func TestTileURL(t *testing.T) {
	got := TileURL("https://api.foreversixty.gg/", "1.60.1.70658", 22423)
	if got != "https://api.foreversixty.gg/v1/discord/tiles/1.60.1.70658/22423.png" {
		t.Fatal(got)
	}
}

// ringPixel is the middle of the tile's top ring edge, clear of the rounded corners.
var ringPixel = image.Pt(TileSide/2, 1)

func TestTileRingPerQuality(t *testing.T) {
	icon, err := testIcons().Icon("inv_helmet_24")
	if err != nil {
		t.Fatal(err)
	}
	for q := -1; q < len(qualityRing); q++ {
		png, err := Tile(icon, q)
		if err != nil {
			t.Fatal(err)
		}
		img := decodePNG(t, png)
		if img.Bounds() != image.Rect(0, 0, TileSide, TileSide) {
			t.Fatalf("tile %v", img.Bounds())
		}
		want := colorEmptyRing
		if q >= 0 {
			want = qualityRing[q]
		}
		if !closeColor(img.At(ringPixel.X, ringPixel.Y), want) {
			t.Errorf("quality %d ring %v, want %v", q, img.At(ringPixel.X, ringPixel.Y), want)
		}
		if _, _, _, a := img.At(0, 0).RGBA(); a != 0 {
			t.Errorf("quality %d: the corner is not transparent", q)
		}
		// The icon sits 1:1 inside the ring.
		if !closeColor(img.At(tileRing+32, tileRing+32), icon.At(32, 32)) {
			t.Errorf("quality %d: the icon is not drawn 1:1", q)
		}
	}
}

// fakeItems knows one build and two items.
type fakeItems struct{ lookups atomic.Int32 }

func (f *fakeItems) KnownBuild(build string) bool { return build == "1.60.1.70658" }

func (f *fakeItems) TileItem(_ string, id int) (string, int, bool) {
	f.lookups.Add(1)
	switch id {
	case 16846:
		return "inv_helmet_24", 4, true
	case 99:
		return "inv_no_such_icon", 3, true
	}
	return "", 0, false
}

func tileServer(t *testing.T) (*httptest.Server, *fakeItems) {
	t.Helper()
	items := &fakeItems{}
	tiles := newTiles(items, func(string) IconSource { return testIcons() }, nil)
	mux := http.NewServeMux()
	tiles.Mount(mux)
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv, items
}

func get(t *testing.T, url string) *http.Response {
	t.Helper()
	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = res.Body.Close() })
	return res
}

func TestTilesRoute(t *testing.T) {
	srv, _ := tileServer(t)
	res := get(t, srv.URL+TilePath+"1.60.1.70658/16846.png")
	if res.StatusCode != http.StatusOK || res.Header.Get("Content-Type") != "image/png" {
		t.Fatalf("status %d type %q", res.StatusCode, res.Header.Get("Content-Type"))
	}
	if cc := res.Header.Get("Cache-Control"); cc != "public, max-age=31536000, immutable" {
		t.Fatalf("Cache-Control %q", cc)
	}
	img, _, err := image.Decode(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	if img.Bounds().Dx() != TileSide || !closeColor(img.At(ringPixel.X, ringPixel.Y), qualityRing[4]) {
		t.Fatalf("tile %v ring %v", img.Bounds(), img.At(ringPixel.X, ringPixel.Y))
	}
}

func TestTilesRouteQuestionMark(t *testing.T) {
	srv, _ := tileServer(t)
	question, err := testIcons().Icon(QuestionMark)
	if err != nil {
		t.Fatal(err)
	}
	for _, id := range []int{12345, 99} { // unknown to the build; known with a missing icon
		res := get(t, srv.URL+TilePath+"1.60.1.70658/"+strconv.Itoa(id)+".png")
		img, _, err := image.Decode(res.Body)
		if err != nil || res.StatusCode != http.StatusOK {
			t.Fatalf("%d: status %d err %v", id, res.StatusCode, err)
		}
		if !closeColor(img.At(tileRing+32, tileRing+32), question.At(32, 32)) {
			t.Errorf("%d: not the question mark", id)
		}
		wantRing := colorEmptyRing
		if id == 99 {
			wantRing = qualityRing[3]
		}
		if !closeColor(img.At(ringPixel.X, ringPixel.Y), wantRing) {
			t.Errorf("%d: ring %v", id, img.At(ringPixel.X, ringPixel.Y))
		}
	}
}

func TestTilesRoute404s(t *testing.T) {
	srv, _ := tileServer(t)
	for _, path := range []string{
		"9.9.9.9/16846.png", "1.60.1.70658/16846", "1.60.1.70658/abc.png", "1.60.1.70658/-1.png",
		"1.60.1.70658/.png", "1.60.1.70658/0.png", "..%2F..%2Fetc/16846.png", "1.60.1.70658/16846.png/x",
		"1.60.1.70658/1234567890.png",
	} {
		res := get(t, srv.URL+TilePath+path)
		if res.StatusCode != http.StatusNotFound {
			t.Errorf("%s: status %d", path, res.StatusCode)
		}
		if strings.Contains(res.Header.Get("Cache-Control"), "immutable") {
			t.Errorf("%s: a 404 is marked immutable", path)
		}
	}
}

func TestTilesRouteCaches(t *testing.T) {
	srv, items := tileServer(t)
	for range 3 {
		if res := get(t, srv.URL+TilePath+"1.60.1.70658/16846.png"); res.StatusCode != http.StatusOK {
			t.Fatal(res.StatusCode)
		}
	}
	if n := items.lookups.Load(); n != 1 {
		t.Fatalf("drawn %d times, want once", n)
	}
}

func TestLRU(t *testing.T) {
	c := newLRU(2, 0)
	c.put("a", []byte("1"))
	c.put("b", []byte("2"))
	if _, ok := c.get("a"); !ok { // a is now the most recent
		t.Fatal("a missing")
	}
	c.put("c", []byte("3"))
	if _, ok := c.get("b"); ok {
		t.Fatal("b should be evicted as the least recently used")
	}
	if _, ok := c.get("a"); !ok || c.len() != 2 {
		t.Fatalf("a kept %v, len %d", ok, c.len())
	}
	byBytes := newLRU(10, 4)
	byBytes.put("x", []byte("12"))
	byBytes.put("y", []byte("34"))
	byBytes.put("z", []byte("5"))
	if _, ok := byBytes.get("x"); ok || byBytes.len() != 2 {
		t.Fatalf("byte bound: x kept %v, len %d", ok, byBytes.len())
	}
	byBytes.put("big", []byte("123456"))
	if _, ok := byBytes.get("big"); ok {
		t.Fatal("a value over the byte bound was kept")
	}
	byBytes.put("y", []byte("6"))
	if v, _ := byBytes.get("y"); string(v) != "6" {
		t.Fatalf("replace: %q", v)
	}
	if newLRU(tileCacheEntries, 0).maxEntries != 2048 {
		t.Fatal("the tile cache holds 2,048")
	}
}

func TestQualityDots(t *testing.T) {
	dots, err := QualityDots()
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"fs_q_poor", "fs_q_common", "fs_q_uncommon", "fs_q_rare", "fs_q_epic", "fs_q_legendary"}
	if len(dots) != len(want) {
		t.Fatalf("%d dots", len(dots))
	}
	for q, dot := range dots {
		if dot.Name != want[q] || QualityEmojiName(q) != want[q] {
			t.Errorf("quality %d named %q", q, dot.Name)
		}
		img := decodePNG(t, dot.PNG)
		if img.Bounds() != image.Rect(0, 0, dotSide, dotSide) {
			t.Fatalf("%s is %v", dot.Name, img.Bounds())
		}
		if !closeColor(img.At(dotSide/2, dotSide/2), qualityRing[q]) {
			t.Errorf("%s centre %v", dot.Name, img.At(dotSide/2, dotSide/2))
		}
		if _, _, _, a := img.At(2, 2).RGBA(); a != 0 {
			t.Errorf("%s corner not transparent", dot.Name)
		}
		edge := color.NRGBAModel.Convert(img.At(dotSide/2, dotSide/2-dotRadius-1)).(color.NRGBA)
		if edge.R > 40 || edge.A < 80 || edge.A > 150 {
			t.Errorf("%s outline %v, want black near 45%%", dot.Name, edge)
		}
	}
	if QualityEmojiName(-1) != "" || QualityEmojiName(6) != "" {
		t.Fatal("an unknown quality has no dot")
	}
}

func TestTreesItems(t *testing.T) {
	data, err := trees.LoadFixture()
	if err != nil {
		t.Fatal(err)
	}
	items := TreesItems{Data: data}
	if !items.KnownBuild("test-1") || items.KnownBuild("9.9.9.9") || (TreesItems{}).KnownBuild("test-1") {
		t.Fatal("known builds")
	}
	icon, quality, ok := items.TileItem("test-1", 12640)
	if !ok || icon != "inv_helmet_29" || quality != 4 {
		t.Fatalf("Lionheart Helm: %q %d %v", icon, quality, ok)
	}
	if _, _, ok := items.TileItem("test-1", 1); ok {
		t.Fatal("an unknown item was found")
	}
}
