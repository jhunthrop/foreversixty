package coach

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

func TestRunReviewJob_ReleasesTheLease(t *testing.T) {
	f := newT30(t)
	ctx := context.Background()
	key := f.key(t)
	until := time.Now().Add(5 * time.Minute)
	if ok, _ := f.svc.Store.TakeLease(ctx, key, "holder-1", until); !ok {
		t.Fatal("first lease")
	}
	f.llm.Responses = []llm.Response{{Content: passingDoc(t, f), Model: "m"}}
	if err := RunReviewJob(ctx, f.svc, []string{t30Report, "1", f.furyKey, strconv.FormatInt(f.premium, 10), "holder-1"}); err != nil {
		t.Fatal(err)
	}
	if ok, _ := f.svc.Store.TakeLease(ctx, key, "holder-2", until); !ok {
		t.Fatal("the job deletes its lease when it ends")
	}
}

func TestRunReviewJob_NeverRewritesAStoredModelReview(t *testing.T) {
	f := newT30(t)
	f.store(t, SourceModel, false)
	if err := RunReviewJob(context.Background(), f.svc, []string{t30Report, "1", f.furyKey, strconv.FormatInt(f.premium, 10), "h"}); err != nil {
		t.Fatal(err)
	}
	if len(f.llm.Calls) != 0 {
		t.Fatalf("the model is not called: %d", len(f.llm.Calls))
	}
}

func TestRunReviewJob_RefusesBadArgs(t *testing.T) {
	f := newT30(t)
	for _, args := range [][]string{
		{t30Report, "1", f.furyKey, "7"},
		{t30Report, "x", f.furyKey, "7", "h"},
		{t30Report, "1", f.furyKey, "0", "h"},
		{t30Report, "1", f.furyKey, "7", ""},
	} {
		if err := RunReviewJob(context.Background(), f.svc, args); err == nil {
			t.Errorf("%v: want an error", args)
		}
	}
}
