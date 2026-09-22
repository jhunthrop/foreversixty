// api/internal/dataaddon/upload_test.go
package dataaddon

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/api/option"
)

func TestFakeUploaderRecordsWhatWasUploaded(t *testing.T) {
	f := &FakeUploader{}
	if err := f.Upload(context.Background(), "my-bucket", "data-addon/Data.lua", []byte("x")); err != nil {
		t.Fatal(err)
	}
	if len(f.Calls) != 1 || f.Calls[0].Bucket != "my-bucket" || f.Calls[0].Object != "data-addon/Data.lua" {
		t.Fatalf("calls = %+v", f.Calls)
	}
	if string(f.Calls[0].Data) != "x" {
		t.Errorf("data = %q", f.Calls[0].Data)
	}
}

func TestFakeUploaderReportsItsConfiguredFailure(t *testing.T) {
	f := &FakeUploader{Err: context.DeadlineExceeded}
	if err := f.Upload(context.Background(), "b", "o", nil); err == nil {
		t.Fatal("the configured error should surface")
	}
	if len(f.Calls) != 0 {
		t.Fatal("a failed upload records nothing")
	}
}

func TestGCSUploadPostsTheObjectToTheStorageJSONAPI(t *testing.T) {
	var gotMethod, gotPath string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotMethod = r.Method
		gotPath = r.URL.Path
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"name": "data-addon/Data.lua"})
	}))
	defer srv.Close()

	g, err := NewGCSWith(context.Background(), option.WithEndpoint(srv.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Upload(context.Background(), "my-bucket", "data-addon/Data.lua", []byte("ForeverSixtyData = {}")); err != nil {
		t.Fatal(err)
	}
	if gotMethod != http.MethodPost {
		t.Errorf("method = %q, want POST (Objects.Insert(...).Media(...) is a multipart POST, not a PUT)", gotMethod)
	}
	const wantPath = "/upload/storage/v1/b/my-bucket/o"
	if gotPath != wantPath {
		t.Errorf("path = %q, want %q (the bucket belongs in the path, not a name= query param)", gotPath, wantPath)
	}
	if !strings.Contains(string(gotBody), "ForeverSixtyData") {
		// The storage/v1 client sends a multipart body (metadata + media);
		// this only asserts the payload made it across, not the exact
		// multipart framing.
		t.Errorf("body = %q, want it to contain the uploaded content", gotBody)
	}
}
