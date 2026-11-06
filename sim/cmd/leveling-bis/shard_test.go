package main

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestApplyBISShardKeepsWrittenOrder(t *testing.T) {
	written := []string{"druid-balance", "priest-holy", "warrior-fury"}
	got, err := applyBISShard(written, bisShard{Specs: []string{"warrior-fury", "druid-balance"}, Reason: "r"})
	if err != nil {
		t.Fatal(err)
	}
	if want := []string{"druid-balance", "warrior-fury"}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestApplyBISShardRefusesAnUnwrittenSpec(t *testing.T) {
	_, err := applyBISShard([]string{"warrior-fury"}, bisShard{Specs: []string{"warior-fury"}, Reason: "r"})
	if err == nil || !strings.Contains(err.Error(), "warior-fury") {
		t.Errorf("err = %v, want it to name the unknown spec", err)
	}
}

func writeShard(t *testing.T, body string) string {
	t.Helper()
	root := t.TempDir()
	path := filepath.Join(root, bisShardFile)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestLoadBISShard(t *testing.T) {
	if _, ok, err := loadBISShard(t.TempDir()); ok || err != nil {
		t.Errorf("no file: ok=%v err=%v, want no shard and no error", ok, err)
	}
	shard, ok, err := loadBISShard(writeShard(t, `{"specs":["warrior-fury"],"reason":"split"}`))
	if !ok || err != nil || !reflect.DeepEqual(shard.Specs, []string{"warrior-fury"}) {
		t.Errorf("shard=%v ok=%v err=%v", shard, ok, err)
	}
	for _, body := range []string{`{"specs":[],"reason":"split"}`, `{"specs":["warrior-fury"]}`, `{`} {
		if _, _, err := loadBISShard(writeShard(t, body)); err == nil {
			t.Errorf("%s: want an error", body)
		}
	}
}

// Without a shard file -all ranks every written spec (the fixture root).
func TestSpecsToRankWithoutAShardIsEveryWrittenSpec(t *testing.T) {
	got, err := specsToRank(repoRootFixture)
	if err != nil {
		t.Fatal(err)
	}
	want, _ := writtenSpecs(repoRootFixture)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
