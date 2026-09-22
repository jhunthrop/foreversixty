package config

import (
	"fmt"
	"strconv"
	"strings"
)

// defaultMailFrom is used when MAIL_FROM is not set.
const defaultMailFrom = "Forever Sixty <hello@foreversixty.gg>"

// defaultTreeDataDir is where the Docker image places data/builds; see
// api/Dockerfile.
const defaultTreeDataDir = "/data"

// Defaults for the Phase 3 variables. Every one of them has a working
// default so a developer can run the service with the Phase 0 environment
// and simply not have the features that need credentials.
const (
	defaultR2Bucket            = "foreversixty-logs"
	defaultSessionCookieDomain = ".foreversixty.gg"
	defaultParseJobName        = "parse-report"
	defaultParseJobRegion      = "us-east1"
	defaultParseJobProject     = "foreversixty"
	defaultSimJobName          = "sim-run"
	defaultSimJobRegion        = "us-east1"
	defaultSimJobProject       = "foreversixty"
)

// defaultTrustedProxyHops is used when TRUSTED_PROXY_HOPS is not set. 1
// matches a single reverse proxy (e.g. Cloud Run) sitting directly in front
// of this service.
const defaultTrustedProxyHops = 1

type Config struct {
	Port     string
	MailFrom string

	// DatabaseURL is used for normal runtime connections; it may point at a
	// connection pooler (e.g. Neon's PgBouncer endpoint).
	DatabaseURL string
	// MigrateDatabaseURL is used only for running migrations, which need a
	// session-level advisory lock that a transaction-pooled connection
	// cannot hold. Falls back to DatabaseURL when unset, so a direct
	// (non-pooled) database needs no extra configuration.
	MigrateDatabaseURL string

	ResendAPIKey  string
	PublicBaseURL string
	APIBaseURL    string

	// TreeDataDir holds one directory per client build with that build's
	// talent, item, class, race, and combo JSON. Missing or incomplete
	// directories are reported at startup and simply have no data, so the
	// service still serves health, version, and subscribe.
	TreeDataDir string

	// TrustedProxyHops is how many reverse proxies in front of this service
	// are trusted to append to X-Forwarded-For; see httpx.RateLimit. 0
	// ignores X-Forwarded-For entirely and rate-limits by RemoteAddr.
	TrustedProxyHops int

	// R2AccountID, R2AccessKeyID, R2SecretAccessKey and R2Bucket address
	// the Cloudflare R2 bucket that holds every report file. With the id
	// or either credential empty, R2Configured reports false and the
	// routes that need object storage are not mounted.
	R2AccountID       string
	R2AccessKeyID     string
	R2SecretAccessKey string
	R2Bucket          string
	// R2Endpoint overrides the S3 endpoint derived from AccountID. The
	// tests set it; production leaves it empty.
	R2Endpoint string

	// SessionCookieDomain is the Domain attribute of fs_session and
	// fs_csrf. Production is ".foreversixty.gg" so the cookie is sent to
	// both the site and the API; a local run wants it empty, which makes
	// the cookie host-only.
	SessionCookieDomain string

	// BnetClientID, BnetClientSecret and BnetRedirectURL configure
	// Battle.net sign-in. With the id or the secret empty,
	// BattleNetConfigured reports false and only the email magic link is
	// offered.
	BnetClientID     string
	BnetClientSecret string
	BnetRedirectURL  string

	// ParseJobName, ParseJobRegion and ParseJobProject address the Cloud
	// Run job that parses a whole-file upload.
	ParseJobName    string
	ParseJobRegion  string
	ParseJobProject string

	// SimJobName, SimJobRegion and SimJobProject address the Cloud Run
	// job that runs one premium sim. Both simulator jobs run on this
	// image; the nightly validation job is scheduled, not dispatched
	// from here, so it needs no address.
	SimJobName    string
	SimJobRegion  string
	SimJobProject string

	// StripeSecretKey, StripeWebhookSecret and StripeEnvironment configure
	// billing. Absent, billing routes answer 503 rather than the API
	// failing to start (coordinator's lane constraint — no Stripe keys
	// exist yet).
	StripeSecretKey     string
	StripeWebhookSecret string
	// StripeEnvironment is "test" or "live", from STRIPE_ENVIRONMENT. It
	// must be "live" exactly when PublicBaseURL is the production origin
	// (see ValidateStripeKeyEnvironment) — this is a separate value from
	// the key's own prefix so the check has two independent signals to
	// compare, not one value checked against itself.
	StripeEnvironment string
}

func Load(getenv func(string) string) (Config, error) {
	c := Config{
		Port:               getenv("PORT"),
		MailFrom:           getenv("MAIL_FROM"),
		DatabaseURL:        getenv("DATABASE_URL"),
		MigrateDatabaseURL: getenv("MIGRATE_DATABASE_URL"),
		ResendAPIKey:       getenv("RESEND_API_KEY"),
		PublicBaseURL:      getenv("PUBLIC_BASE_URL"),
		APIBaseURL:         getenv("API_BASE_URL"),
	}
	if c.Port == "" {
		c.Port = "8080"
	}
	if c.MailFrom == "" {
		c.MailFrom = defaultMailFrom
	}
	c.TreeDataDir = getenv("TREE_DATA_DIR")
	if c.TreeDataDir == "" {
		c.TreeDataDir = defaultTreeDataDir
	}
	for name, v := range map[string]string{
		"DATABASE_URL": c.DatabaseURL, "RESEND_API_KEY": c.ResendAPIKey,
		"PUBLIC_BASE_URL": c.PublicBaseURL, "API_BASE_URL": c.APIBaseURL,
	} {
		if v == "" {
			return Config{}, fmt.Errorf("config: %s is required", name)
		}
	}
	if c.MigrateDatabaseURL == "" {
		c.MigrateDatabaseURL = c.DatabaseURL
	}

	c.TrustedProxyHops = defaultTrustedProxyHops
	if v := getenv("TRUSTED_PROXY_HOPS"); v != "" {
		hops, err := strconv.Atoi(v)
		if err != nil || hops < 0 {
			return Config{}, fmt.Errorf("config: TRUSTED_PROXY_HOPS must be an integer >= 0, got %q", v)
		}
		c.TrustedProxyHops = hops
	}

	c.R2AccountID = getenv("R2_ACCOUNT_ID")
	c.R2AccessKeyID = getenv("R2_ACCESS_KEY_ID")
	c.R2SecretAccessKey = getenv("R2_SECRET_ACCESS_KEY")
	c.R2Endpoint = getenv("R2_ENDPOINT")
	c.BnetClientID = getenv("BNET_CLIENT_ID")
	c.BnetClientSecret = getenv("BNET_CLIENT_SECRET")
	c.BnetRedirectURL = getenv("BNET_REDIRECT_URL")
	for _, d := range []struct {
		dst *string
		env string
		def string
	}{
		{&c.R2Bucket, "R2_BUCKET", defaultR2Bucket},
		{&c.ParseJobName, "PARSE_JOB_NAME", defaultParseJobName},
		{&c.ParseJobRegion, "PARSE_JOB_REGION", defaultParseJobRegion},
		{&c.ParseJobProject, "PARSE_JOB_PROJECT", defaultParseJobProject},
		{&c.SimJobName, "SIM_JOB_NAME", defaultSimJobName},
		{&c.SimJobRegion, "SIM_JOB_REGION", defaultSimJobRegion},
		{&c.SimJobProject, "SIM_JOB_PROJECT", defaultSimJobProject},
	} {
		*d.dst = getenv(d.env)
		if *d.dst == "" {
			*d.dst = d.def
		}
	}
	// SESSION_COOKIE_DOMAIN is handled apart from the loop above: "none"
	// asks for a host-only cookie, which is an empty Domain attribute and
	// so cannot be spelled with an empty environment variable.
	c.SessionCookieDomain = getenv("SESSION_COOKIE_DOMAIN")
	switch c.SessionCookieDomain {
	case "":
		c.SessionCookieDomain = defaultSessionCookieDomain
	case "none":
		c.SessionCookieDomain = ""
	}

	c.StripeSecretKey = getenv("STRIPE_SECRET_KEY")
	c.StripeWebhookSecret = getenv("STRIPE_WEBHOOK_SECRET")
	c.StripeEnvironment = getenv("STRIPE_ENVIRONMENT")

	return c, nil
}

// R2Configured reports whether the object store can be reached. Without it
// the ingest, upload, and report-file routes are not mounted.
func (c Config) R2Configured() bool {
	return c.R2AccountID != "" && c.R2AccessKeyID != "" && c.R2SecretAccessKey != "" && c.R2Bucket != ""
}

// BattleNetConfigured reports whether Battle.net sign-in can be offered.
func (c Config) BattleNetConfigured() bool {
	return c.BnetClientID != "" && c.BnetClientSecret != "" && c.BnetRedirectURL != ""
}

// productionBaseURL is the one PublicBaseURL value ValidateStripeKeyEnvironment treats as
// production (spec §3).
const productionBaseURL = "https://foreversixty.gg"

// StripeConfigured reports whether billing can be offered: both the
// secret key and the webhook secret are present. Absent, the billing
// routes answer 503 rather than the deployment failing to start.
func (c Config) StripeConfigured() bool {
	return c.StripeSecretKey != "" && c.StripeWebhookSecret != ""
}

// ValidateStripeKeyEnvironment is the spec §3 startup check, scoped by
// the coordinator's lane constraint that the API must start with no
// Stripe keys at all: a completely absent StripeSecretKey is always a
// no-op (the billing routes answer 503 instead — Task 8). Once a key is
// present, it must only be a live key where PublicBaseURL is the
// production origin, and STRIPE_ENVIRONMENT must be "live" exactly then
// too.
func (c Config) ValidateStripeKeyEnvironment() error {
	if c.StripeSecretKey == "" {
		return nil
	}
	wantLive := c.PublicBaseURL == productionBaseURL
	gotLive := c.StripeEnvironment == "live"
	if wantLive != gotLive {
		if wantLive {
			return fmt.Errorf("config: STRIPE_ENVIRONMENT must be \"live\" when PUBLIC_BASE_URL is %s", productionBaseURL)
		}
		return fmt.Errorf("config: STRIPE_ENVIRONMENT must not be \"live\" unless PUBLIC_BASE_URL is %s", productionBaseURL)
	}
	isLiveKey := strings.HasPrefix(c.StripeSecretKey, "sk_live_") || strings.HasPrefix(c.StripeSecretKey, "rk_live_")
	isTestKey := strings.HasPrefix(c.StripeSecretKey, "sk_test_") || strings.HasPrefix(c.StripeSecretKey, "rk_test_")
	if !isLiveKey && !isTestKey {
		return fmt.Errorf("config: STRIPE_SECRET_KEY does not look like a Stripe secret or restricted key")
	}
	if gotLive && !isLiveKey {
		return fmt.Errorf("config: STRIPE_ENVIRONMENT is \"live\" but STRIPE_SECRET_KEY is a test key")
	}
	if !gotLive && isLiveKey {
		return fmt.Errorf("config: STRIPE_ENVIRONMENT is %q but STRIPE_SECRET_KEY is a live key", c.StripeEnvironment)
	}
	return nil
}
