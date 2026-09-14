package secret

import (
	"path/filepath"
	"testing"

	"github.com/jhunthrop/foreversixty/companion/internal/config"
	"github.com/zalando/go-keyring"
)

func TestTheKeyringStoreRoundTripsAToken(t *testing.T) {
	keyring.MockInit()
	var s Store = Keyring{}
	if s.Backend() != Keychain {
		t.Errorf("backend = %q", s.Backend())
	}
	got, err := s.Token()
	if err != nil || got != "" {
		t.Fatalf("empty keychain = %q, %v", got, err)
	}
	if err := s.SetToken("fsd_one"); err != nil {
		t.Fatal(err)
	}
	if got, err := s.Token(); err != nil || got != "fsd_one" {
		t.Fatalf("Token = %q, %v", got, err)
	}
	if err := s.Clear(); err != nil {
		t.Fatal(err)
	}
	if err := s.Clear(); err != nil {
		t.Fatalf("clearing twice failed: %v", err)
	}
	if got, _ := s.Token(); got != "" {
		t.Fatalf("Token after Clear = %q", got)
	}
}

func TestTheFileStoreRoundTripsAToken(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	if err := config.Save(path, config.Default()); err != nil {
		t.Fatal(err)
	}
	var s Store = ConfigFile{Path: path}
	if s.Backend() != File {
		t.Errorf("backend = %q", s.Backend())
	}
	if err := s.SetToken("fsd_two"); err != nil {
		t.Fatal(err)
	}
	cfg, err := config.Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.DeviceToken != "fsd_two" {
		t.Fatalf("config device_token = %q", cfg.DeviceToken)
	}
	if got, err := s.Token(); err != nil || got != "fsd_two" {
		t.Fatalf("Token = %q, %v", got, err)
	}
	if err := s.Clear(); err != nil {
		t.Fatal(err)
	}
	if got, _ := s.Token(); got != "" {
		t.Fatalf("Token after Clear = %q", got)
	}
}

func TestMigrateMovesAFileTokenIntoTheKeychain(t *testing.T) {
	keyring.MockInit()
	path := filepath.Join(t.TempDir(), "config.json")
	cfg := config.Default()
	cfg.DeviceToken = "fsd_three"
	if err := config.Save(path, cfg); err != nil {
		t.Fatal(err)
	}
	s := Open(path)
	if s.Backend() != Keychain {
		t.Fatalf("Open chose %q with a working keychain", s.Backend())
	}
	if err := Migrate(s, path); err != nil {
		t.Fatal(err)
	}
	if got, _ := s.Token(); got != "fsd_three" {
		t.Errorf("keychain token = %q", got)
	}
	after, err := config.Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if after.DeviceToken != "" {
		t.Errorf("config still holds %q", after.DeviceToken)
	}
}

func TestMigrateIsANoOpForTheFileStore(t *testing.T) {
	path := filepath.Join(t.TempDir(), "config.json")
	cfg := config.Default()
	cfg.DeviceToken = "fsd_four"
	if err := config.Save(path, cfg); err != nil {
		t.Fatal(err)
	}
	if err := Migrate(ConfigFile{Path: path}, path); err != nil {
		t.Fatal(err)
	}
	after, _ := config.Load(path)
	if after.DeviceToken != "fsd_four" {
		t.Errorf("the file store lost its token: %q", after.DeviceToken)
	}
}
