package config

import (
	"fmt"
	"strings"
	"testing"
)

func pushEnv(extra map[string]string) func(string) string {
	m := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "https://foreversixty.gg", "API_BASE_URL": "https://api.foreversixty.gg",
	}
	for k, v := range extra {
		m[k] = v
	}
	return env(m)
}

func TestPushIsOffAndUnconfiguredByDefault(t *testing.T) {
	c, err := Load(pushEnv(nil))
	if err != nil {
		t.Fatal(err)
	}
	if c.Push.Enabled || c.Push.APNsConfigured() || c.Push.FCMConfigured() || c.Push.APNsTopic != "gg.foreversixty.app" {
		t.Fatalf("%+v", c.Push)
	}
}

func TestPushReadsItsVariables(t *testing.T) {
	c, err := Load(pushEnv(map[string]string{
		"APP_NOTIFICATIONS_ENABLED": "1", "APNS_KEY_P8": "pem-secret", "APNS_KEY_ID": "K", "APNS_TEAM_ID": "T",
		"APNS_SANDBOX": "1", "FCM_SERVICE_ACCOUNT_JSON": `{"project_id":"p"}`,
	}))
	if err != nil {
		t.Fatal(err)
	}
	if !c.Push.Enabled || !c.Push.APNsConfigured() || !c.Push.FCMConfigured() || !c.Push.APNsSandbox {
		t.Fatalf("enabled %v", c.Push.Enabled)
	}
}

func TestPushConfigNeverPrintsTheKeys(t *testing.T) {
	p := PushConfig{APNsKeyP8: []byte("pem-secret"), FCMServiceAccountJSON: []byte("sa-secret")}
	for _, s := range []string{fmt.Sprintf("%v", p), fmt.Sprintf("%+v", p), fmt.Sprintf("%#v", p)} {
		if strings.Contains(s, "secret") {
			t.Fatalf("printed %s", s)
		}
	}
}

func TestAPNsVariablesComeTogether(t *testing.T) {
	_, err := Load(pushEnv(map[string]string{"APNS_KEY_P8": "pem-secret-value", "APNS_KEY_ID": "K"}))
	if err == nil || !strings.Contains(err.Error(), "APNS_TEAM_ID") || strings.Contains(err.Error(), "pem-secret-value") {
		t.Fatalf("err %v", err)
	}
}
