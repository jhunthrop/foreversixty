package jobs

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/api/option"
)

func TestFakeRecordsWhatWasRun(t *testing.T) {
	var r Runner = &Fake{}
	if err := r.Run(context.Background(), "parse-report", "abc"); err != nil {
		t.Fatal(err)
	}
	f := r.(*Fake)
	if got := f.Ran(); len(got) != 1 || got[0][0] != "parse-report" || got[0][1] != "abc" {
		t.Fatalf("ran = %v", got)
	}
}

func TestFakeReportsItsConfiguredFailure(t *testing.T) {
	f := &Fake{Err: errors.New("no such job")}
	if err := f.Run(context.Background(), "parse-report", "abc"); err == nil {
		t.Fatal("the configured error should surface")
	}
	if len(f.Ran()) != 0 {
		t.Fatal("a failed run records nothing")
	}
}

func TestNewCloudRunAddressesTheRegionalJob(t *testing.T) {
	// No credentials are needed to build the client: nothing is dialled
	// until Run. On a machine with no application default credentials
	// this reports that, which is still the constructor working.
	c, err := NewCloudRun(context.Background(), "foreversixty", "us-east1", "parse-report")
	if err != nil {
		if !strings.Contains(err.Error(), "credential") && !strings.Contains(err.Error(), "default") {
			t.Fatalf("unexpected error: %v", err)
		}
		t.Skip("no application default credentials on this machine")
	}
	want := "projects/foreversixty/locations/us-east1/jobs/parse-report"
	if c.name != want {
		t.Fatalf("job name = %q, want %q", c.name, want)
	}
}

func TestRunExecutesTheJobWithTheArgumentsAsOverrides(t *testing.T) {
	var got struct {
		path string
		body []byte
	}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.path = r.URL.Path
		got.body, _ = io.ReadAll(r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"name":"operations/1"}`))
	}))
	defer srv.Close()

	c, err := NewCloudRunWith(context.Background(), "foreversixty", "us-east1", "parse-report",
		option.WithEndpoint(srv.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Run(context.Background(), "parse-report", "abc"); err != nil {
		t.Fatal(err)
	}
	if want := "/v2/projects/foreversixty/locations/us-east1/jobs/parse-report:run"; got.path != want {
		t.Fatalf("path = %q, want %q", got.path, want)
	}
	if !strings.Contains(string(got.body), `"parse-report"`) || !strings.Contains(string(got.body), `"abc"`) {
		t.Fatalf("body = %s, want the arguments as a container override", got.body)
	}
}

func TestRunReportsAFailedExecution(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"error":{"code":403,"message":"permission denied"}}`))
	}))
	defer srv.Close()
	c, err := NewCloudRunWith(context.Background(), "p", "us-east1", "parse-report",
		option.WithEndpoint(srv.URL), option.WithoutAuthentication())
	if err != nil {
		t.Fatal(err)
	}
	if err := c.Run(context.Background(), "parse-report", "abc"); err == nil {
		t.Fatal("a 403 from Cloud Run must surface")
	}
}
