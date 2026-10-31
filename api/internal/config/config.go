package config

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
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

	// defaultBnetProfileGame is the profile-game segment of a namespace
	// (spec §1 rule 3): the value between "profile-" and "-{region}".
	// Forever's real namespace is one env change on launch day.
	defaultBnetProfileGame = "classic1x"
	// defaultBnetRegions is which regional hosts to try for an account.
	defaultBnetRegions = "us,eu"
	// defaultBnetProbeGames is every game namespace segment the nightly
	// refresh's probe checks (spec §5).
	defaultBnetProbeGames = "classic1x,classic-forever,classicforever,forever,classic60,anniversary"
)

// defaultTrustedProxyHops is used when TRUSTED_PROXY_HOPS is not set. 1
// matches a single reverse proxy (e.g. Cloud Run) sitting directly in front
// of this service.
const defaultTrustedProxyHops = 1

type Config struct {
	Port     string
	MailFrom string

	// ListenHost is the address the HTTP listener binds to (LISTEN_HOST). Empty binds every
	// interface, which is what a container wants; a local run sets 127.0.0.1.
	ListenHost string
	// DevLogin (FS_DEV_LOGIN=1) mounts POST /v1/dev/login, which signs anyone in as a
	// seeded account without a password. It exists only for the local walkthrough stack
	// (tools/dev-stack) and must never ship enabled: Load refuses it unless ListenHost is
	// a loopback address and PublicBaseURL is not the production origin.
	DevLogin bool

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

	// BnetProfileGame, BnetRegions and BnetProbeGames configure the
	// Battle.net character import and its namespace probe (spec §1
	// rule 3, §5): BNET_PROFILE_GAME, BNET_REGIONS, BNET_PROBE_GAMES.
	BnetProfileGame string
	BnetRegions     []string
	BnetProbeGames  []string
	// BnetImportCharacters (BNET_IMPORT_CHARACTERS=true) turns the account
	// character import on. Off by default since 2026-09-27: Blizzard serves
	// no Forever namespace yet, so with only classic1x answering the import
	// filled accounts with Era and Season of Discovery characters that are
	// not Forever characters at all. Off, sign-in still works, the nightly
	// job still probes for the namespace, and nothing is written.
	BnetImportCharacters bool
	// AppSignIn (APP_SIGN_IN=true) mounts the native app sign-in: GET /v1/auth/app/start,
	// POST /v1/auth/app/exchange and POST /v1/auth/app/logout. Off by default; the routes
	// answer 404 until the apps ship. It needs Battle.net configured to do anything.
	AppSignIn bool
	// AppTokenIdle (APP_TOKEN_IDLE_DAYS, default 30) and AppTokenMaxAge
	// (APP_TOKEN_MAX_AGE_DAYS, default 180) are when a native app's device token lapses:
	// after that long unused, or that long after sign-in, whichever comes first. A lapsed
	// token answers 401 session_expired and the app signs in again. Companion tokens do
	// not lapse.
	AppTokenIdle   time.Duration
	AppTokenMaxAge time.Duration

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

	// NotifyEncryptionKey seals the Discord webhook tokens at rest
	// (AES-256-GCM). It is the base64 of NotifyKeyLen bytes in
	// NOTIFY_ENCRYPTION_KEY. Absent, it is nil and the notifications
	// feature is off; present but malformed, Load refuses to start. The
	// value is never put in an error or a log.
	NotifyEncryptionKey []byte
	// DiscordClientID and DiscordClientSecret are the Discord application the
	// one-click connect uses (DISCORD_CLIENT_ID, DISCORD_CLIENT_SECRET). Both
	// set turns connect on, neither leaves it off; one alone fails startup.
	DiscordClientID     string
	DiscordClientSecret string
	// The bot (DISCORD_APPLICATION_ID, DISCORD_PUBLIC_KEY as 64 hex chars,
	// DISCORD_BOT_TOKEN) is all or nothing and rides the client pair above.
	// DiscordCommandsGuildID (DISCORD_COMMANDS_GUILD_ID) is the private test
	// server that receives unreleased commands.
	DiscordApplicationID   string
	DiscordPublicKey       ed25519.PublicKey
	DiscordBotToken        string
	DiscordCommandsGuildID string

	// Coach is the Log Coach's model client and cap settings.
	Coach CoachConfig
	// Push is the apps' notifications: the dark-launch flag and the APNs and FCM keys.
	Push PushConfig
}

// String and GoString redact the whole Config so a stray %v or %+v never prints the
// bot token, the Stripe keys or the database URLs.
func (Config) String() string   { return "config.Config{redacted}" }
func (Config) GoString() string { return "config.Config{redacted}" }

// isDigits reports whether s is empty or only ASCII digits.
func isDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// DiscordBotConfigured reports whether the application id, public key and bot
// token are all set.
func (c Config) DiscordBotConfigured() bool {
	return c.DiscordApplicationID != "" && len(c.DiscordPublicKey) == ed25519.PublicKeySize && c.DiscordBotToken != ""
}

// NotifyKeyLen is the byte length of NotifyEncryptionKey.
const NotifyKeyLen = 32

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
	c.ListenHost = getenv("LISTEN_HOST")
	c.DevLogin = getenv("FS_DEV_LOGIN") == "1"
	if c.DevLogin {
		if !isLoopbackHost(c.ListenHost) {
			return Config{}, fmt.Errorf("config: FS_DEV_LOGIN=1 needs LISTEN_HOST to be a loopback address (127.0.0.1 or ::1), got %q", c.ListenHost)
		}
		if c.PublicBaseURL == productionBaseURL {
			return Config{}, fmt.Errorf("config: FS_DEV_LOGIN=1 is refused when PUBLIC_BASE_URL is %s", productionBaseURL)
		}
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
	c.BnetProfileGame = getenv("BNET_PROFILE_GAME")
	if c.BnetProfileGame == "" {
		c.BnetProfileGame = defaultBnetProfileGame
	}
	c.BnetImportCharacters = getenv("BNET_IMPORT_CHARACTERS") == "true"
	c.AppSignIn = getenv("APP_SIGN_IN") == "true"
	if err := loadAppTokenLifetimes(&c, getenv); err != nil {
		return Config{}, err
	}
	c.BnetRegions = splitCommaList(getenv("BNET_REGIONS"), defaultBnetRegions)
	c.BnetProbeGames = splitCommaList(getenv("BNET_PROBE_GAMES"), defaultBnetProbeGames)
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

	key, err := decodeNotifyKey(getenv("NOTIFY_ENCRYPTION_KEY"))
	if err != nil {
		return c, err
	}
	c.NotifyEncryptionKey = key

	c.DiscordClientID, c.DiscordClientSecret = getenv("DISCORD_CLIENT_ID"), getenv("DISCORD_CLIENT_SECRET")
	if err := checkDiscordPair(c.DiscordClientID, c.DiscordClientSecret); err != nil {
		return c, err
	}
	if err := loadDiscordBot(&c, getenv); err != nil {
		return c, err
	}
	if c.Coach, err = loadCoach(getenv); err != nil {
		return c, err
	}
	if c.Push, err = loadPush(getenv); err != nil {
		return c, err
	}

	return c, nil
}

// loadDiscordBot reads the bot's variables into c and checks them as a set.
func loadDiscordBot(c *Config, getenv func(string) string) error {
	trimmed := func(name string) string { return strings.TrimSpace(getenv(name)) }
	c.DiscordApplicationID = trimmed("DISCORD_APPLICATION_ID")
	c.DiscordBotToken = trimmed("DISCORD_BOT_TOKEN")
	c.DiscordCommandsGuildID = trimmed("DISCORD_COMMANDS_GUILD_ID")
	if !isDigits(c.DiscordCommandsGuildID) {
		return errors.New("config: DISCORD_COMMANDS_GUILD_ID must be a Discord server id (digits only)")
	}
	rawKey := trimmed("DISCORD_PUBLIC_KEY")
	if err := checkDiscordBot(*c, rawKey != ""); err != nil {
		return err
	}
	key, err := decodeDiscordPublicKey(rawKey)
	if err != nil {
		return err
	}
	c.DiscordPublicKey = key
	return nil
}

// checkDiscordBot requires the bot's three values together, on top of the
// client id and secret, and the application id to equal the client id. The
// errors name variables, never values.
func checkDiscordBot(c Config, keySet bool) error {
	anySet := c.DiscordApplicationID != "" || c.DiscordBotToken != "" || keySet
	if !anySet {
		return nil
	}
	for _, v := range []struct {
		name string
		set  bool
	}{
		{"DISCORD_APPLICATION_ID", c.DiscordApplicationID != ""},
		{"DISCORD_PUBLIC_KEY", keySet},
		{"DISCORD_BOT_TOKEN", c.DiscordBotToken != ""},
	} {
		if !v.set {
			return fmt.Errorf("config: %s is required when any Discord bot variable is set", v.name)
		}
	}
	if c.DiscordClientID == "" || c.DiscordClientSecret == "" {
		return errors.New("config: DISCORD_CLIENT_ID and DISCORD_CLIENT_SECRET are required when the Discord bot is set")
	}
	if c.DiscordApplicationID != c.DiscordClientID {
		return errors.New("config: DISCORD_APPLICATION_ID must equal DISCORD_CLIENT_ID")
	}
	return nil
}

// decodeDiscordPublicKey turns DISCORD_PUBLIC_KEY (hex of a 32-byte Ed25519
// key) into a key; empty means unset (nil).
func decodeDiscordPublicKey(raw string) (ed25519.PublicKey, error) {
	if raw == "" {
		return nil, nil
	}
	key, err := hex.DecodeString(raw)
	if err != nil || len(key) != ed25519.PublicKeySize {
		return nil, errors.New("config: DISCORD_PUBLIC_KEY must be 64 hex characters (32 bytes)")
	}
	return ed25519.PublicKey(key), nil
}

// checkDiscordPair requires the Discord application's two values together.
// The error names the missing variable, never a value.
func checkDiscordPair(id, secret string) error {
	switch {
	case id != "" && secret == "":
		return errors.New("config: DISCORD_CLIENT_SECRET is required when DISCORD_CLIENT_ID is set")
	case id == "" && secret != "":
		return errors.New("config: DISCORD_CLIENT_ID is required when DISCORD_CLIENT_SECRET is set")
	}
	return nil
}

// decodeNotifyKey turns NOTIFY_ENCRYPTION_KEY into key bytes: empty means
// the feature is off (nil); anything else must be base64 of NotifyKeyLen
// bytes. The error names the variable, never its value.
func decodeNotifyKey(raw string) ([]byte, error) {
	if raw == "" {
		return nil, nil
	}
	key, err := base64.StdEncoding.DecodeString(raw)
	if err != nil || len(key) != NotifyKeyLen {
		return nil, fmt.Errorf("config: NOTIFY_ENCRYPTION_KEY must be base64 of %d bytes", NotifyKeyLen)
	}
	return key, nil
}

// The app token lifetimes' defaults, in days.
const (
	defaultAppTokenIdleDays   = 30
	defaultAppTokenMaxAgeDays = 180
)

func loadAppTokenLifetimes(c *Config, getenv func(string) string) error {
	idle, err := positiveDays(getenv, "APP_TOKEN_IDLE_DAYS", defaultAppTokenIdleDays)
	if err != nil {
		return err
	}
	maxAge, err := positiveDays(getenv, "APP_TOKEN_MAX_AGE_DAYS", defaultAppTokenMaxAgeDays)
	if err != nil {
		return err
	}
	if idle > maxAge {
		return fmt.Errorf("config: APP_TOKEN_IDLE_DAYS (%d) must not exceed APP_TOKEN_MAX_AGE_DAYS (%d)", idle, maxAge)
	}
	const day = 24 * time.Hour
	c.AppTokenIdle, c.AppTokenMaxAge = time.Duration(idle)*day, time.Duration(maxAge)*day
	return nil
}

// positiveDays reads a whole number of days above zero, or def when unset.
func positiveDays(getenv func(string) string, name string, def int) (int, error) {
	v := getenv(name)
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 1 {
		return 0, fmt.Errorf("config: %s must be a whole number of days >= 1, got %q", name, v)
	}
	return n, nil
}

// isLoopbackHost reports whether host is a literal loopback IP. "localhost" and the empty
// string are deliberately not accepted: the first depends on name resolution and the second
// binds every interface, and the dev sign-in must not depend on either.
func isLoopbackHost(host string) bool {
	ip := net.ParseIP(strings.Trim(host, "[]"))
	return ip != nil && ip.IsLoopback()
}

// splitCommaList splits a comma-separated env value into its trimmed,
// non-empty parts, falling back to def (itself comma-separated) when v is
// empty — the shape BNET_REGIONS and BNET_PROBE_GAMES both use.
func splitCommaList(v, def string) []string {
	if v == "" {
		v = def
	}
	var out []string
	for _, part := range strings.Split(v, ",") {
		if trimmed := strings.TrimSpace(part); trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
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
