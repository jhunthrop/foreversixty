package config

import (
	"fmt"
	"strings"
	"testing"
)

func coachEnv(extra map[string]string) func(string) string {
	m := map[string]string{
		"COACH_LLM_ENABLED": "true",
		"LLM_MODEL":         "m", "LLM_API_KEY": "k", "LLM_PROVIDERS": "p, q",
		"LLM_PRICE_IN_PER_MTOK": "0.1", "LLM_PRICE_OUT_PER_MTOK": "0.4",
	}
	for k, v := range extra {
		m[k] = v
	}
	return env(m)
}

func TestCoachConfig_Defaults(t *testing.T) {
	c, err := loadCoach(coachEnv(nil))
	if err != nil || !c.Enabled || c.DailySpendCapUSD != 25 || c.DailyReviewsPerAccount != 30 ||
		c.LLMBaseURL != "https://ai-gateway.vercel.sh/v1" || c.Temperature != 0.3 || c.DataDir != "/coach" ||
		len(c.LLMProviders) != 2 || c.LLMPriceInPerMTok != 0.1 || c.LLMPriceOutPerMTok != 0.4 {
		t.Fatalf("%+v %v", c, err)
	}
}

func TestCoachConfig_EnabledNeedsEachRequiredVariable(t *testing.T) {
	for _, name := range []string{"LLM_MODEL", "LLM_API_KEY", "LLM_PROVIDERS", "LLM_PRICE_IN_PER_MTOK", "LLM_PRICE_OUT_PER_MTOK"} {
		_, err := loadCoach(coachEnv(map[string]string{name: ""}))
		if err == nil || !strings.Contains(err.Error(), name) {
			t.Fatalf("%s: %v", name, err)
		}
	}
}

func TestCoachConfig_ErrorNeverEchoesValues(t *testing.T) {
	_, err := loadCoach(coachEnv(map[string]string{"LLM_MODEL": "", "LLM_API_KEY": "sekrit"}))
	if err == nil || strings.Contains(err.Error(), "sekrit") {
		t.Fatal(err)
	}
}

func TestCoachConfig_UnsetIsDisabledAndNeedsNothing(t *testing.T) {
	c, err := loadCoach(env(nil))
	if err != nil || c.Enabled {
		t.Fatalf("%+v %v", c, err)
	}
	c, err = loadCoach(env(map[string]string{"COACH_LLM_ENABLED": "false"}))
	if err != nil || c.Enabled {
		t.Fatalf("%+v %v", c, err)
	}
}

func TestLoad_CoachUnsetDoesNotBreakLoad(t *testing.T) {
	c, err := Load(base(nil))
	if err != nil || c.Coach.Enabled {
		t.Fatalf("%+v %v", c.Coach, err)
	}
}

func TestLoad_CoachEnabledWithoutVarsNamesTheVariable(t *testing.T) {
	_, err := Load(base(map[string]string{"COACH_LLM_ENABLED": "true"}))
	if err == nil || !strings.Contains(err.Error(), "LLM_MODEL") {
		t.Fatal(err)
	}
}

func TestCoachConfig_RejectsBadValues(t *testing.T) {
	for k, v := range map[string]string{
		"COACH_LLM_ENABLED": "maybe", "COACH_DAILY_SPEND_CAP_USD": "x",
		"COACH_DAILY_REVIEWS_PER_ACCOUNT": "-1",
	} {
		if _, err := loadCoach(coachEnv(map[string]string{k: v})); err == nil {
			t.Fatalf("%s=%s accepted", k, v)
		}
	}
	if _, err := loadCoach(coachEnv(map[string]string{"COACH_DAILY_REVIEWS_PER_ACCOUNT": "0"})); err == nil {
		t.Fatal("zero per-account cap accepted while enabled")
	}
}

func TestCoachConfig_RejectsNonFiniteNumbers(t *testing.T) {
	for _, v := range []string{"Inf", "+Inf", "NaN"} {
		if _, err := loadCoach(coachEnv(map[string]string{"COACH_DAILY_SPEND_CAP_USD": v})); err == nil {
			t.Fatalf("%s accepted", v)
		}
	}
}

func TestCoachConfig_PrintingNeverShowsTheKey(t *testing.T) {
	c, err := loadCoach(coachEnv(map[string]string{"LLM_API_KEY": "sekrit"}))
	if err != nil {
		t.Fatal(err)
	}
	for _, out := range []string{fmt.Sprintf("%v", c), fmt.Sprintf("%+v", c), fmt.Sprintf("%#v", c), c.String()} {
		if strings.Contains(out, "sekrit") {
			t.Fatal(out)
		}
	}
}

func TestCoachConfig_ProvidersIsACopy(t *testing.T) {
	c, _ := loadCoach(coachEnv(nil))
	c.Providers()[0] = "changed"
	if c.LLMProviders[0] == "changed" {
		t.Fatal("Providers shares its backing array")
	}
}
