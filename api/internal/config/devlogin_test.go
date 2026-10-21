package config

import "testing"

// devLoginEnv is a valid local environment with the dev sign-in asked for; each test
// changes one thing.
func devLoginEnv(extra map[string]string) func(string) string {
	m := map[string]string{
		"DATABASE_URL": "postgres://x", "RESEND_API_KEY": "k",
		"PUBLIC_BASE_URL": "http://localhost:4321", "API_BASE_URL": "http://localhost:8080",
		"FS_DEV_LOGIN": "1", "LISTEN_HOST": "127.0.0.1",
	}
	for k, v := range extra {
		m[k] = v
	}
	return env(m)
}

// The dev sign-in signs anyone in without a password, so it must be impossible to enable
// on a listener another machine can reach.
func TestDevLoginNeedsALoopbackListener(t *testing.T) {
	c, err := Load(devLoginEnv(nil))
	if err != nil {
		t.Fatal(err)
	}
	if !c.DevLogin {
		t.Fatal("FS_DEV_LOGIN=1 on 127.0.0.1 should enable the dev sign-in")
	}
	if _, err := Load(devLoginEnv(map[string]string{"LISTEN_HOST": "::1"})); err != nil {
		t.Fatalf("::1 is loopback: %v", err)
	}
	for _, host := range []string{"", "0.0.0.0", "localhost", "192.168.1.5", "::"} {
		if _, err := Load(devLoginEnv(map[string]string{"LISTEN_HOST": host})); err == nil {
			t.Errorf("LISTEN_HOST=%q with FS_DEV_LOGIN=1 should be refused", host)
		}
	}
}

func TestDevLoginIsRefusedForTheProductionOrigin(t *testing.T) {
	_, err := Load(devLoginEnv(map[string]string{"PUBLIC_BASE_URL": productionBaseURL}))
	if err == nil {
		t.Fatal("FS_DEV_LOGIN=1 must be refused when PUBLIC_BASE_URL is production")
	}
}

// Off is the default, and anything but the literal 1 is off.
func TestDevLoginIsOffUnlessExactlyOne(t *testing.T) {
	for _, v := range []string{"", "0", "true", "yes"} {
		c, err := Load(devLoginEnv(map[string]string{"FS_DEV_LOGIN": v, "LISTEN_HOST": ""}))
		if err != nil {
			t.Fatalf("FS_DEV_LOGIN=%q: %v", v, err)
		}
		if c.DevLogin {
			t.Errorf("FS_DEV_LOGIN=%q enabled the dev sign-in", v)
		}
	}
}
