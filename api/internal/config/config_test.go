package config

import "testing"

func env(m map[string]string) func(string) string {
	return func(k string) string { return m[k] }
}

func TestLoadRequiresDatabaseURL(t *testing.T) {
	_, err := Load(env(map[string]string{"PORT": "8080"}))
	if err == nil {
		t.Fatal("expected error for missing DATABASE_URL")
	}
}

func TestLoadDefaults(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.Port != "8080" {
		t.Fatalf("default port = %q", c.Port)
	}
	if c.MailFrom != defaultMailFrom {
		t.Fatalf("default MailFrom = %q", c.MailFrom)
	}
	if c.TrustedProxyHops != 1 {
		t.Fatalf("default TrustedProxyHops = %d, want 1", c.TrustedProxyHops)
	}
	if c.MigrateDatabaseURL != c.DatabaseURL {
		t.Fatalf("default MigrateDatabaseURL = %q, want it to fall back to DatabaseURL %q", c.MigrateDatabaseURL, c.DatabaseURL)
	}
}

func TestLoadMailFromOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"MAIL_FROM": "Test <test@example.com>",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.MailFrom != "Test <test@example.com>" {
		t.Fatalf("MailFrom = %q", c.MailFrom)
	}
}

func TestLoadMigrateDatabaseURLOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://pooled", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"MIGRATE_DATABASE_URL": "postgres://direct",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.MigrateDatabaseURL != "postgres://direct" {
		t.Fatalf("MigrateDatabaseURL = %q, want postgres://direct", c.MigrateDatabaseURL)
	}
	if c.DatabaseURL != "postgres://pooled" {
		t.Fatalf("DatabaseURL should be unaffected, got %q", c.DatabaseURL)
	}
}

func TestLoadTrustedProxyHopsOverride(t *testing.T) {
	c, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "0",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.TrustedProxyHops != 0 {
		t.Fatalf("TrustedProxyHops = %d, want 0", c.TrustedProxyHops)
	}
}

func TestLoadTrustedProxyHopsRejectsNonInteger(t *testing.T) {
	_, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "nope",
	}))
	if err == nil {
		t.Fatal("expected error for non-integer TRUSTED_PROXY_HOPS")
	}
}

func TestLoadTrustedProxyHopsRejectsNegative(t *testing.T) {
	_, err := Load(env(map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
		"TRUSTED_PROXY_HOPS": "-1",
	}))
	if err == nil {
		t.Fatal("expected error for negative TRUSTED_PROXY_HOPS")
	}
}

func TestLoadDefaultsTreeDataDir(t *testing.T) {
	env := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}
	c, err := Load(func(k string) string { return env[k] })
	if err != nil || c.TreeDataDir != "/data" {
		t.Fatalf("TreeDataDir = %q err = %v", c.TreeDataDir, err)
	}

	env["TREE_DATA_DIR"] = "../../data/builds"
	c, err = Load(func(k string) string { return env[k] })
	if err != nil || c.TreeDataDir != "../../data/builds" {
		t.Fatalf("TreeDataDir = %q err = %v", c.TreeDataDir, err)
	}
}

// base is the Phase 0 environment every Phase 3 test starts from.
func base(extra map[string]string) func(string) string {
	m := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}
	for k, v := range extra {
		m[k] = v
	}
	return env(m)
}

func TestPhase3DefaultsAreUsableWithoutAnyNewVariables(t *testing.T) {
	c, err := Load(base(nil))
	if err != nil {
		t.Fatal(err)
	}
	if c.R2Bucket != defaultR2Bucket {
		t.Errorf("bucket = %q", c.R2Bucket)
	}
	if c.SessionCookieDomain != defaultSessionCookieDomain {
		t.Errorf("cookie domain = %q", c.SessionCookieDomain)
	}
	if c.ParseJobName != "parse-report" || c.ParseJobRegion != "us-east1" || c.ParseJobProject != "foreversixty" {
		t.Errorf("parse job = %s/%s/%s", c.ParseJobProject, c.ParseJobRegion, c.ParseJobName)
	}
	if c.SimJobName != "sim-run" || c.SimJobRegion != "us-east1" || c.SimJobProject != "foreversixty" {
		t.Errorf("sim job = %s/%s/%s", c.SimJobProject, c.SimJobRegion, c.SimJobName)
	}
	if c.R2Configured() {
		t.Error("R2 must not read as configured with no credentials")
	}
	if c.BattleNetConfigured() {
		t.Error("Battle.net must not read as configured with no client")
	}
}

func TestR2AndBattleNetReadAsConfiguredOnlyWhenComplete(t *testing.T) {
	full := map[string]string{
		"R2_ACCOUNT_ID": "acct", "R2_ACCESS_KEY_ID": "key", "R2_SECRET_ACCESS_KEY": "secret",
		"BNET_CLIENT_ID": "id", "BNET_CLIENT_SECRET": "secret",
		"BNET_REDIRECT_URL": "https://api.foreversixty.gg/v1/auth/battlenet/callback",
	}
	c, err := Load(base(full))
	if err != nil {
		t.Fatal(err)
	}
	if !c.R2Configured() || !c.BattleNetConfigured() {
		t.Fatalf("config = %+v, want both configured", c)
	}
	for _, missing := range []string{"R2_ACCOUNT_ID", "R2_ACCESS_KEY_ID", "R2_SECRET_ACCESS_KEY"} {
		partial := map[string]string{}
		for k, v := range full {
			partial[k] = v
		}
		delete(partial, missing)
		c, err := Load(base(partial))
		if err != nil {
			t.Fatal(err)
		}
		if c.R2Configured() {
			t.Errorf("R2 reads as configured without %s", missing)
		}
	}
	for _, missing := range []string{"BNET_CLIENT_ID", "BNET_CLIENT_SECRET", "BNET_REDIRECT_URL"} {
		partial := map[string]string{}
		for k, v := range full {
			partial[k] = v
		}
		delete(partial, missing)
		c, err := Load(base(partial))
		if err != nil {
			t.Fatal(err)
		}
		if c.BattleNetConfigured() {
			t.Errorf("Battle.net reads as configured without %s", missing)
		}
	}
}

func TestSessionCookieDomainCanBeTurnedOff(t *testing.T) {
	c, err := Load(base(map[string]string{"SESSION_COOKIE_DOMAIN": "none"}))
	if err != nil {
		t.Fatal(err)
	}
	if c.SessionCookieDomain != "" {
		t.Fatalf("domain = %q, want a host-only cookie", c.SessionCookieDomain)
	}
	c, err = Load(base(map[string]string{"SESSION_COOKIE_DOMAIN": ".example.test"}))
	if err != nil {
		t.Fatal(err)
	}
	if c.SessionCookieDomain != ".example.test" {
		t.Fatalf("domain = %q", c.SessionCookieDomain)
	}
}

func TestTheParseJobAndBucketCanBeOverridden(t *testing.T) {
	c, err := Load(base(map[string]string{
		"R2_BUCKET": "other-bucket", "R2_ENDPOINT": "http://127.0.0.1:9000",
		"PARSE_JOB_NAME": "parse-staging", "PARSE_JOB_REGION": "europe-west1",
		"PARSE_JOB_PROJECT": "staging",
	}))
	if err != nil {
		t.Fatal(err)
	}
	if c.R2Bucket != "other-bucket" || c.R2Endpoint != "http://127.0.0.1:9000" {
		t.Fatalf("r2 = %+v", c)
	}
	if c.ParseJobName != "parse-staging" || c.ParseJobRegion != "europe-west1" || c.ParseJobProject != "staging" {
		t.Fatalf("parse job = %s/%s/%s", c.ParseJobProject, c.ParseJobRegion, c.ParseJobName)
	}
}
