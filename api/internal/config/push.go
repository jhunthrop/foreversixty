package config

import (
	"fmt"
	"strings"
)

// defaultAPNsTopic is the apps' bundle id (plan decision B).
const defaultAPNsTopic = "gg.foreversixty.app"

// PushConfig is the apps' notifications: the flag that ships them dark, and the APNs and FCM
// credentials. A sender exists only when its credentials are all set, so both stay off
// until the owner adds the keys.
type PushConfig struct {
	// Enabled (APP_NOTIFICATIONS_ENABLED=1, default off) mounts /v1/me/notifications and
	// /v1/me/push-tokens, lets a device token answer a raid sign-up, and routes events to push.
	Enabled bool
	// APNsKeyP8 (APNS_KEY_P8) is the .p8 auth key's PEM, from Secret Manager.
	APNsKeyP8  []byte
	APNsKeyID  string // APNS_KEY_ID
	APNsTeamID string // APNS_TEAM_ID
	APNsTopic  string // APNS_TOPIC, default gg.foreversixty.app
	// APNsSandbox (APNS_SANDBOX=1) sends to the development host, for debug builds.
	APNsSandbox bool
	// FCMServiceAccountJSON (FCM_SERVICE_ACCOUNT_JSON) is the Firebase service account key,
	// from Secret Manager; its project_id names the project.
	FCMServiceAccountJSON []byte
}

// String and GoString redact the keys, as Config does.
func (PushConfig) String() string   { return "config.PushConfig{redacted}" }
func (PushConfig) GoString() string { return "config.PushConfig{redacted}" }

// APNsConfigured reports whether the APNs key, key id and team id are all set.
func (p PushConfig) APNsConfigured() bool {
	return len(p.APNsKeyP8) > 0 && p.APNsKeyID != "" && p.APNsTeamID != ""
}

// FCMConfigured reports whether the FCM service account is set.
func (p PushConfig) FCMConfigured() bool { return len(p.FCMServiceAccountJSON) > 0 }

// loadPush reads the push variables. The APNs three come together; the errors name
// variables, never values.
func loadPush(getenv func(string) string) (PushConfig, error) {
	trimmed := func(name string) string { return strings.TrimSpace(getenv(name)) }
	p := PushConfig{
		Enabled:               trimmed("APP_NOTIFICATIONS_ENABLED") == "1",
		APNsKeyP8:             []byte(trimmed("APNS_KEY_P8")),
		APNsKeyID:             trimmed("APNS_KEY_ID"),
		APNsTeamID:            trimmed("APNS_TEAM_ID"),
		APNsTopic:             trimmed("APNS_TOPIC"),
		APNsSandbox:           trimmed("APNS_SANDBOX") == "1",
		FCMServiceAccountJSON: []byte(trimmed("FCM_SERVICE_ACCOUNT_JSON")),
	}
	if p.APNsTopic == "" {
		p.APNsTopic = defaultAPNsTopic
	}
	set := map[string]bool{"APNS_KEY_P8": len(p.APNsKeyP8) > 0, "APNS_KEY_ID": p.APNsKeyID != "", "APNS_TEAM_ID": p.APNsTeamID != ""}
	anySet := set["APNS_KEY_P8"] || set["APNS_KEY_ID"] || set["APNS_TEAM_ID"]
	for _, name := range []string{"APNS_KEY_P8", "APNS_KEY_ID", "APNS_TEAM_ID"} {
		if anySet && !set[name] {
			return PushConfig{}, fmt.Errorf("config: %s is required when any APNs variable is set", name)
		}
	}
	return p, nil
}
