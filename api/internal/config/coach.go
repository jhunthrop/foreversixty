package config

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jhunthrop/foreversixty/api/internal/llm"
)

const (
	defaultCoachSpendCapUSD   = 25.0
	defaultCoachReviewsPerDay = 30
	defaultCoachDataDir       = "/coach"
	defaultLLMBaseURL         = "https://ai-gateway.vercel.sh/v1"
	defaultCoachTemperature   = 0.3

	// The per-account caps are about a sixth of each plan's monthly price ($6, $15, $25). A
	// review costs about $0.0005 at the pinned prices, so they allow thousands of reviews
	// and only stop abuse.
	defaultMonthlyCapPremiumUSD   = 1.00
	defaultMonthlyCapGuildUSD     = 2.50
	defaultMonthlyCapRaidGuildUSD = 4.00
	defaultDailyCapPerAccountUSD  = 0.15
)

// AccountSpendCaps bound the model spend one account may cause: a calendar month (UTC) by
// its plan tier, and a UTC day.
type AccountSpendCaps struct {
	MonthlyPremiumUSD   float64 // COACH_MONTHLY_SPEND_CAP_PREMIUM_USD
	MonthlyGuildUSD     float64 // COACH_MONTHLY_SPEND_CAP_GUILD_USD
	MonthlyRaidGuildUSD float64 // COACH_MONTHLY_SPEND_CAP_RAID_GUILD_USD
	DailyUSD            float64 // COACH_DAILY_SPEND_CAP_PER_ACCOUNT_USD
}

// CoachConfig is the Log Coach's model client and cap settings.
type CoachConfig struct {
	// Enabled (COACH_LLM_ENABLED, default false until the coach ships) is the kill switch.
	Enabled bool
	// DailySpendCapUSD (COACH_DAILY_SPEND_CAP_USD) bounds a UTC day's recorded spend.
	DailySpendCapUSD float64
	// DailyReviewsPerAccount (COACH_DAILY_REVIEWS_PER_ACCOUNT) bounds generations per account per UTC day.
	DailyReviewsPerAccount int
	// AccountSpendCaps bound each account's spend by month (per plan tier) and day.
	AccountSpendCaps AccountSpendCaps
	DataDir          string // COACH_DATA_DIR
	LLMBaseURL       string // LLM_BASE_URL
	LLMModel         string // LLM_MODEL
	LLMAPIKey        string // LLM_API_KEY, from Secret Manager
	LLMProviders     []string
	// LLMPriceInPerMTok and LLMPriceOutPerMTok are USD per million tokens.
	LLMPriceInPerMTok, LLMPriceOutPerMTok float64
	Temperature                           float64 // COACH_TEMPERATURE
	// LLMReasoning (LLM_REASONING, default llm.DefaultReasoning) is one of llm.ReasoningLevels.
	LLMReasoning string
}

// Providers returns a copy of LLMProviders, so callers cannot alter the config.
func (c CoachConfig) Providers() []string { return slices.Clone(c.LLMProviders) }

// String describes the config without the API key.
func (c CoachConfig) String() string {
	return fmt.Sprintf("config.CoachConfig{Enabled:%t Model:%q BaseURL:%q Providers:%v SpendCapUSD:%v PerAccount:%d APIKey:[redacted]}",
		c.Enabled, c.LLMModel, c.LLMBaseURL, c.LLMProviders, c.DailySpendCapUSD, c.DailyReviewsPerAccount)
}

// GoString matches String under %#v.
func (c CoachConfig) GoString() string { return c.String() }

// loadCoach reads the coach settings. While the coach is enabled the model,
// key, providers and prices are required; errors name variables, never values.
func loadCoach(getenv func(string) string) (CoachConfig, error) {
	c := CoachConfig{
		DataDir:      envOr(getenv, "COACH_DATA_DIR", defaultCoachDataDir),
		LLMBaseURL:   envOr(getenv, "LLM_BASE_URL", defaultLLMBaseURL),
		LLMModel:     strings.TrimSpace(getenv("LLM_MODEL")),
		LLMAPIKey:    strings.TrimSpace(getenv("LLM_API_KEY")),
		LLMProviders: splitCommaList(getenv("LLM_PROVIDERS"), ""),
		LLMReasoning: envOr(getenv, "LLM_REASONING", llm.DefaultReasoning),
	}
	if !slices.Contains(llm.ReasoningLevels, c.LLMReasoning) {
		return CoachConfig{}, fmt.Errorf("config: LLM_REASONING must be one of %s", strings.Join(llm.ReasoningLevels, ", "))
	}
	var err error
	if c.Enabled, err = parseBoolDefault(getenv, "COACH_LLM_ENABLED", false); err != nil {
		return CoachConfig{}, err
	}
	numbers := []struct {
		env string
		def float64
		dst *float64
	}{
		{"COACH_DAILY_SPEND_CAP_USD", defaultCoachSpendCapUSD, &c.DailySpendCapUSD},
		{"COACH_TEMPERATURE", defaultCoachTemperature, &c.Temperature},
		{"LLM_PRICE_IN_PER_MTOK", 0, &c.LLMPriceInPerMTok},
		{"LLM_PRICE_OUT_PER_MTOK", 0, &c.LLMPriceOutPerMTok},
		{"COACH_MONTHLY_SPEND_CAP_PREMIUM_USD", defaultMonthlyCapPremiumUSD, &c.AccountSpendCaps.MonthlyPremiumUSD},
		{"COACH_MONTHLY_SPEND_CAP_GUILD_USD", defaultMonthlyCapGuildUSD, &c.AccountSpendCaps.MonthlyGuildUSD},
		{"COACH_MONTHLY_SPEND_CAP_RAID_GUILD_USD", defaultMonthlyCapRaidGuildUSD, &c.AccountSpendCaps.MonthlyRaidGuildUSD},
		{"COACH_DAILY_SPEND_CAP_PER_ACCOUNT_USD", defaultDailyCapPerAccountUSD, &c.AccountSpendCaps.DailyUSD},
	}
	for _, n := range numbers {
		if *n.dst, err = parseFloatDefault(getenv, n.env, n.def); err != nil {
			return CoachConfig{}, err
		}
	}
	if c.DailyReviewsPerAccount, err = parseIntDefault(getenv, "COACH_DAILY_REVIEWS_PER_ACCOUNT", defaultCoachReviewsPerDay); err != nil {
		return CoachConfig{}, err
	}
	if c.Enabled {
		return c, c.checkEnabled()
	}
	return c, nil
}

// checkEnabled names the first missing or non-positive variable an enabled coach needs.
func (c CoachConfig) checkEnabled() error {
	for _, r := range []struct {
		env string
		ok  bool
	}{
		{"LLM_MODEL", c.LLMModel != ""},
		{"LLM_API_KEY", c.LLMAPIKey != ""},
		{"LLM_PROVIDERS", len(c.LLMProviders) > 0},
		{"LLM_PRICE_IN_PER_MTOK", c.LLMPriceInPerMTok > 0},
		{"LLM_PRICE_OUT_PER_MTOK", c.LLMPriceOutPerMTok > 0},
		{"COACH_DAILY_SPEND_CAP_USD", c.DailySpendCapUSD > 0},
		{"COACH_DAILY_REVIEWS_PER_ACCOUNT", c.DailyReviewsPerAccount > 0},
		{"COACH_MONTHLY_SPEND_CAP_PREMIUM_USD", c.AccountSpendCaps.MonthlyPremiumUSD > 0},
		{"COACH_MONTHLY_SPEND_CAP_GUILD_USD", c.AccountSpendCaps.MonthlyGuildUSD > 0},
		{"COACH_MONTHLY_SPEND_CAP_RAID_GUILD_USD", c.AccountSpendCaps.MonthlyRaidGuildUSD > 0},
		{"COACH_DAILY_SPEND_CAP_PER_ACCOUNT_USD", c.AccountSpendCaps.DailyUSD > 0},
	} {
		if !r.ok {
			return fmt.Errorf("config: %s is required (non-empty, positive) while COACH_LLM_ENABLED is on", r.env)
		}
	}
	return nil
}

func envOr(getenv func(string) string, name, def string) string {
	if v := strings.TrimSpace(getenv(name)); v != "" {
		return v
	}
	return def
}

func parseBoolDefault(getenv func(string) string, name string, def bool) (bool, error) {
	v := strings.TrimSpace(getenv(name))
	if v == "" {
		return def, nil
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false, fmt.Errorf("config: %s must be true or false", name)
	}
	return b, nil
}

func parseFloatDefault(getenv func(string) string, name string, def float64) (float64, error) {
	v := strings.TrimSpace(getenv(name))
	if v == "" {
		return def, nil
	}
	f, err := strconv.ParseFloat(v, 64)
	if err != nil || f < 0 || math.IsInf(f, 0) || math.IsNaN(f) {
		return 0, fmt.Errorf("config: %s must be a finite, non-negative number", name)
	}
	return f, nil
}

func parseIntDefault(getenv func(string) string, name string, def int) (int, error) {
	v := strings.TrimSpace(getenv(name))
	if v == "" {
		return def, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil || n < 0 {
		return 0, fmt.Errorf("config: %s must be a non-negative whole number", name)
	}
	return n, nil
}
