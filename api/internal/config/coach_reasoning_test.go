package config

import (
	"strings"
	"testing"
)

func TestCoachConfig_Reasoning(t *testing.T) {
	c, err := loadCoach(coachEnv(nil))
	if err != nil || c.LLMReasoning != "off" {
		t.Fatalf("default: %q %v", c.LLMReasoning, err)
	}
	c, err = loadCoach(coachEnv(map[string]string{"LLM_REASONING": "low"}))
	if err != nil || c.LLMReasoning != "low" {
		t.Fatalf("low: %q %v", c.LLMReasoning, err)
	}
	_, err = loadCoach(coachEnv(map[string]string{"LLM_REASONING": "xhigh"}))
	if err == nil || !strings.Contains(err.Error(), "LLM_REASONING") || !strings.Contains(err.Error(), "off, low, medium, high") {
		t.Fatalf("xhigh: %v", err)
	}
}
