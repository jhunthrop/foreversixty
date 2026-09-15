package r2

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"sync"
	"testing"
)

// fakeS3 is the smallest S3 an R2 client needs: PUT, GET, DELETE, and
// the three multipart calls, path style, no authentication check beyond
// requiring a signature to be present.
type fakeS3 struct {
	mu      sync.Mutex
	objects map[string][]byte
	headers map[string]http.Header
	parts   map[string]map[int][]byte
	nextID  int
	signed  []string
}

func newFakeS3(t *testing.T) (*fakeS3, *httptest.Server) {
	t.Helper()
	f := &fakeS3{objects: map[string][]byte{}, headers: map[string]http.Header{}, parts: map[string]map[int][]byte{}}
	srv := httptest.NewServer(f)
	t.Cleanup(srv.Close)
	return f, srv
}

func (f *fakeS3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// The path is /<bucket>/<key...> because the client is path style.
	key := strings.TrimPrefix(r.URL.Path, "/")
	_, key, _ = strings.Cut(key, "/")
	q := r.URL.Query()
	if q.Get("X-Amz-Signature") != "" {
		f.signed = append(f.signed, r.URL.String())
	}

	switch {
	case r.Method == http.MethodPost && q.Has("uploads"):
		f.nextID++
		id := fmt.Sprintf("upload-%d", f.nextID)
		f.parts[id] = map[int][]byte{}
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<InitiateMultipartUploadResult><Bucket>b</Bucket><Key>%s</Key><UploadId>%s</UploadId></InitiateMultipartUploadResult>`, key, id)
	case r.Method == http.MethodPut && q.Get("uploadId") != "":
		body, _ := io.ReadAll(r.Body)
		var n int
		fmt.Sscanf(q.Get("partNumber"), "%d", &n)
		f.parts[q.Get("uploadId")][n] = body
		w.Header().Set("ETag", fmt.Sprintf("%q", fmt.Sprintf("etag-%d", n)))
		w.WriteHeader(http.StatusOK)
	case r.Method == http.MethodPost && q.Get("uploadId") != "":
		id := q.Get("uploadId")
		numbers := make([]int, 0, len(f.parts[id]))
		for n := range f.parts[id] {
			numbers = append(numbers, n)
		}
		sort.Ints(numbers)
		var joined []byte
		for _, n := range numbers {
			joined = append(joined, f.parts[id][n]...)
		}
		f.objects[key] = joined
		delete(f.parts, id)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintf(w, `<?xml version="1.0" encoding="UTF-8"?>
<CompleteMultipartUploadResult><Bucket>b</Bucket><Key>%s</Key><ETag>"whole"</ETag></CompleteMultipartUploadResult>`, key)
	case r.Method == http.MethodDelete && q.Get("uploadId") != "":
		delete(f.parts, q.Get("uploadId"))
		w.WriteHeader(http.StatusNoContent)
	case r.Method == http.MethodPut:
		body, _ := io.ReadAll(r.Body)
		f.objects[key] = body
		f.headers[key] = r.Header.Clone()
		w.WriteHeader(http.StatusOK)
	case r.Method == http.MethodGet:
		body, ok := f.objects[key]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_, _ = w.Write(body)
	case r.Method == http.MethodDelete:
		delete(f.objects, key)
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (f *fakeS3) object(key string) ([]byte, bool) {
	f.mu.Lock()
	defer f.mu.Unlock()
	b, ok := f.objects[key]
	return b, ok
}

func (f *fakeS3) header(key, name string) string {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.headers[key].Get(name)
}
