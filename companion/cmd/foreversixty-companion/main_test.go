package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/instance"
	"github.com/jhunthrop/foreversixty/companion/internal/paths"
)

// A second launch shows the first one's window and exits cleanly, without
// starting a pipeline of its own.
func TestASecondLaunchOpensTheRunningOneAndExits(t *testing.T) {
	home := t.TempDir()
	t.Setenv(paths.HomeEnv, home)
	first, err := instance.Acquire(home)
	if err != nil {
		t.Fatal(err)
	}
	defer first.Release()
	if err := first.PublishURL("http://127.0.0.1:4242/secret/"); err != nil {
		t.Fatal(err)
	}
	var opened []string
	openURL = func(u string) error { opened = append(opened, u); return nil }
	defer func() { openURL = defaultOpen }()

	out, err := os.CreateTemp(t.TempDir(), "out")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()
	if err := run(nil, out); err != nil {
		t.Fatalf("the second launch failed: %v", err)
	}
	if len(opened) != 1 || opened[0] != "http://127.0.0.1:4242/secret/" {
		t.Fatalf("opened %v", opened)
	}
	log, _ := os.ReadFile(filepath.Join(home, "logs", "companion.log"))
	if !strings.Contains(string(log), "already running") {
		t.Errorf("the log does not say so:\n%s", log)
	}
	if strings.Contains(string(log), "the companion is running") {
		t.Error("the second launch started the companion")
	}
}
