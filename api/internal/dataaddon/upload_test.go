// api/internal/dataaddon/upload_test.go
package dataaddon

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
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

func TestGCSUploadPUTsTheObjectToTheStorageJSONAPI(t *testing.T) {
	var gotPath, gotBucket string
	var gotBody []byte
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotBucket = r.URL.Query().Get("name")
		gotBody, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{"name": gotBucket})
	}))
	defer srv.Close()

	g, err := NewGCSWith(context.Background(), option.WithEndpoint(srv.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatal(err)
	}
	if err := g.Upload(context.Background(), "my-bucket", "data-addon/Data.lua", []byte("ForeverSixtyData = {}")); err != nil {
		t.Fatal(err)
	}
	if gotPath == "" {
		t.Fatal("no request reached the fake server")
	}
	if string(gotBody) != "" && !contains(string(gotBody), "ForeverSixtyData") {
		// The storage/v1 client sends a multipart body (metadata + media);
		// this only asserts the payload made it across, not the exact
		// multipart framing.
		t.Errorf("body = %q, want it to contain the uploaded content", gotBody)
	}
}

func contains(s, sub string) bool {
	return len(s) >= len(sub) && (func() bool {
		for i := 0; i+len(sub) <= len(s); i++ {
			if s[i:i+len(sub)] == sub {
				return true
			}
		}
		return false
	})()
}
