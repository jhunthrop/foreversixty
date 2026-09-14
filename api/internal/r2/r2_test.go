package r2

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/logs/engine/store"
)

func testClient(t *testing.T) (*Client, *fakeS3) {
	t.Helper()
	fake, srv := newFakeS3(t)
	c, err := New(Config{
		AccessKeyID: "key", SecretAccessKey: "secret",
		Bucket: "foreversixty-logs", Endpoint: srv.URL,
	})
	if err != nil {
		t.Fatal(err)
	}
	return c, fake
}

func TestNewRefusesAnIncompleteConfiguration(t *testing.T) {
	for name, cfg := range map[string]Config{
		"no bucket":   {AccessKeyID: "k", SecretAccessKey: "s", AccountID: "a"},
		"no key":      {Bucket: "b", SecretAccessKey: "s", AccountID: "a"},
		"no secret":   {Bucket: "b", AccessKeyID: "k", AccountID: "a"},
		"no endpoint": {Bucket: "b", AccessKeyID: "k", SecretAccessKey: "s"},
	} {
		if _, err := New(cfg); err == nil {
			t.Errorf("%s should have been refused", name)
		}
	}
}

func TestTheEndpointDefaultsToTheAccountsOwn(t *testing.T) {
	got := Config{AccountID: "1f1504a52d74d323c0a5648a766e0871"}.endpoint()
	want := "https://1f1504a52d74d323c0a5648a766e0871.r2.cloudflarestorage.com"
	if got != want {
		t.Fatalf("endpoint = %q, want %q", got, want)
	}
}

func TestPutCarriesTheCacheHeadersAndGetReadsItBack(t *testing.T) {
	c, fake := testClient(t)
	ctx := context.Background()
	key := "reports/abc/report.json"
	if err := c.Put(ctx, key, []byte(`{"report_id":"abc"}`), store.PutOptions{
		ContentType: "application/json", CacheControl: "public, max-age=5",
	}); err != nil {
		t.Fatal(err)
	}
	body, ok := fake.object(key)
	if !ok || string(body) != `{"report_id":"abc"}` {
		t.Fatalf("stored %q, %v", body, ok)
	}
	if got := fake.header(key, "Cache-Control"); got != "public, max-age=5" {
		t.Fatalf("cache-control = %q", got)
	}
	if got := fake.header(key, "Content-Type"); got != "application/json" {
		t.Fatalf("content-type = %q", got)
	}

	r, err := c.Get(ctx, key)
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	read, err := io.ReadAll(r)
	if err != nil || string(read) != `{"report_id":"abc"}` {
		t.Fatalf("read %q, %v", read, err)
	}

	if err := c.Delete(ctx, key); err != nil {
		t.Fatal(err)
	}
	if _, ok := fake.object(key); ok {
		t.Fatal("the object should be gone")
	}
	if _, err := c.Get(ctx, key); err == nil {
		t.Fatal("reading a deleted object should fail")
	}
}

func TestTheClientIsAStorePutter(t *testing.T) {
	c, _ := testClient(t)
	var p store.Putter = c
	if p == nil {
		t.Fatal("the client must satisfy store.Putter")
	}
}

func TestPutCarriesTheContentEncodingForRawChunks(t *testing.T) {
	c, fake := testClient(t)
	key := "reports/abc/raw/0.zst"
	if err := c.Put(context.Background(), key, []byte("packed"), store.PutOptions{
		ContentType: "application/zstd", CacheControl: "private", ContentEncoding: "zstd",
	}); err != nil {
		t.Fatal(err)
	}
	if got := fake.header(key, "Content-Encoding"); got != "zstd" {
		t.Fatalf("content-encoding = %q", got)
	}
}

func TestPartCountFollowsThe64MiBPartSize(t *testing.T) {
	for size, want := range map[int64]int{0: 1, 1: 1, PartSize: 1, PartSize + 1: 2, 5 * PartSize: 5} {
		if got := PartCount(size); got != want {
			t.Errorf("PartCount(%d) = %d, want %d", size, got, want)
		}
	}
}

func TestMultipartRoundTripThroughSignedURLs(t *testing.T) {
	c, fake := testClient(t)
	ctx := context.Background()
	key := UploadKey("up1")

	uploadID, err := c.StartMultipart(ctx, key, "application/zstd")
	if err != nil {
		t.Fatal(err)
	}
	if uploadID == "" {
		t.Fatal("no upload id")
	}
	parts, err := c.PresignParts(ctx, key, uploadID, 2, URLTTL)
	if err != nil {
		t.Fatal(err)
	}
	if len(parts) != 2 || parts[0].Number != 1 || parts[1].Number != 2 {
		t.Fatalf("parts = %+v", parts)
	}
	for _, p := range parts {
		if !strings.Contains(p.URL, "X-Amz-Signature=") {
			t.Fatalf("part %d is not signed: %s", p.Number, p.URL)
		}
		if !strings.Contains(p.URL, "X-Amz-Expires=3600") {
			t.Fatalf("part %d does not expire in an hour: %s", p.Number, p.URL)
		}
	}

	// Upload through the signed URLs exactly as a browser would.
	var done []CompletedPart
	for i, p := range parts {
		req, err := http.NewRequest(http.MethodPut, p.URL, strings.NewReader(string(rune('a'+i))))
		if err != nil {
			t.Fatal(err)
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		res.Body.Close()
		if res.StatusCode != http.StatusOK {
			t.Fatalf("part %d upload = %d", p.Number, res.StatusCode)
		}
		// A browser reads the ETag off the response, quotes included.
		done = append(done, CompletedPart{Number: p.Number, ETag: res.Header.Get("ETag")})
	}
	if err := c.CompleteMultipart(ctx, key, uploadID, done); err != nil {
		t.Fatal(err)
	}
	body, ok := fake.object(key)
	if !ok || string(body) != "ab" {
		t.Fatalf("assembled %q, %v", body, ok)
	}
}

func TestAbortMultipartThrowsThePartsAway(t *testing.T) {
	c, _ := testClient(t)
	ctx := context.Background()
	key := UploadKey("up2")
	uploadID, err := c.StartMultipart(ctx, key, "application/zstd")
	if err != nil {
		t.Fatal(err)
	}
	if err := c.AbortMultipart(ctx, key, uploadID); err != nil {
		t.Fatal(err)
	}
}

func TestPresignGetIsSignedAndExpires(t *testing.T) {
	c, _ := testClient(t)
	url, err := c.PresignGet(context.Background(), "reports/abc/report.json", AccessTTL)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(url, "X-Amz-Signature=") || !strings.Contains(url, "X-Amz-Expires=600") {
		t.Fatalf("url = %s", url)
	}
	if !strings.Contains(url, "reports/abc/report.json") {
		t.Fatalf("url does not address the object: %s", url)
	}
}

func TestNormalizeETagAlwaysQuotes(t *testing.T) {
	for in, want := range map[string]string{`"abc"`: `"abc"`, "abc": `"abc"`, ` abc `: `"abc"`} {
		if got := normalizeETag(in); got != want {
			t.Errorf("normalizeETag(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestUploadKeyIsTheContractsPath(t *testing.T) {
	if got := UploadKey("abc"); got != "uploads/abc/raw.txt.zst" {
		t.Fatalf("UploadKey = %q", got)
	}
}

func TestSignedURLsExpireWithinTheHour(t *testing.T) {
	if URLTTL != time.Hour || AccessTTL != 10*time.Minute {
		t.Fatalf("TTLs = %v and %v, want the contract's hour and ten minutes", URLTTL, AccessTTL)
	}
}
