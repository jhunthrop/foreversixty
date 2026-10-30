package config

import (
	"strings"
	"testing"
)

func TestAppConfig_DarkByDefault(t *testing.T) {
	c, err := loadApp(env(map[string]string{}))
	if err != nil || c.Enabled || strings.Join(c.WebCheckoutRegions, ",") != "US" ||
		c.MinVersionIOS != "1.0.0" || c.MinVersionAndroid != "1.0.0" {
		t.Fatalf("%+v %v", c, err)
	}
}

func TestAppConfig_ReadsEachVariable(t *testing.T) {
	c, err := loadApp(env(map[string]string{"APP_API": "true", "APP_WEB_CHECKOUT_REGIONS": "us, GB",
		"APP_MIN_VERSION_IOS": "1.2.0", "APP_MIN_VERSION_ANDROID": "1.10.3"}))
	if err != nil || !c.Enabled || strings.Join(c.WebCheckoutRegions, ",") != "US,GB" ||
		c.MinVersionIOS != "1.2.0" || c.MinVersionAndroid != "1.10.3" {
		t.Fatalf("%+v %v", c, err)
	}
	if c, _ := loadApp(env(map[string]string{"APP_WEB_CHECKOUT_REGIONS": "none"})); len(c.WebCheckoutRegions) != 0 {
		t.Fatalf("none = %v", c.WebCheckoutRegions)
	}
}

func TestAppConfig_RefusesMalformedValues(t *testing.T) {
	for name, value := range map[string]string{
		"APP_WEB_CHECKOUT_REGIONS": "USA", "APP_MIN_VERSION_IOS": "1.0", "APP_MIN_VERSION_ANDROID": "v1.0.0",
	} {
		if _, err := loadApp(env(map[string]string{name: value})); err == nil || !strings.Contains(err.Error(), name) {
			t.Errorf("%s=%s: %v", name, value, err)
		}
	}
}
