package config

import (
	"strings"
	"testing"
)

func TestAppConfig_DarkByDefault(t *testing.T) {
	c, err := loadApp(env(map[string]string{}))
	if err != nil || c.Enabled || c.MinVersionIOS != "1.0.0" || c.MinVersionAndroid != "1.0.0" {
		t.Fatalf("%+v %v", c, err)
	}
	// Decision L: iOS links in the US storefront, Android nowhere, and no feature is switched.
	if strings.Join(c.WebCheckoutRegions["ios"], ",") != "US" || c.WebCheckoutRegions["android"] == nil ||
		len(c.WebCheckoutRegions["android"]) != 0 || c.Features == nil || len(c.Features) != 0 {
		t.Fatalf("defaults = %+v", c)
	}
}

func TestAppConfig_ReadsEachVariable(t *testing.T) {
	c, err := loadApp(env(map[string]string{"APP_API": "true",
		"APP_WEB_CHECKOUT_REGIONS_IOS": "us, GB", "APP_WEB_CHECKOUT_REGIONS_ANDROID": "US",
		"APP_MIN_VERSION_IOS": "1.2.0", "APP_MIN_VERSION_ANDROID": "1.10.3", "APP_FEATURES": "raids, sign_in"}))
	if err != nil || !c.Enabled || strings.Join(c.WebCheckoutRegions["ios"], ",") != "US,GB" ||
		strings.Join(c.WebCheckoutRegions["android"], ",") != "US" ||
		c.MinVersionIOS != "1.2.0" || c.MinVersionAndroid != "1.10.3" || strings.Join(c.Features, ",") != "raids,sign_in" {
		t.Fatalf("%+v %v", c, err)
	}
	if c, _ := loadApp(env(map[string]string{"APP_WEB_CHECKOUT_REGIONS_IOS": "none"})); len(c.WebCheckoutRegions["ios"]) != 0 {
		t.Fatalf("none = %v", c.WebCheckoutRegions)
	}
}

func TestAppConfig_RefusesMalformedValues(t *testing.T) {
	for name, value := range map[string]string{
		"APP_WEB_CHECKOUT_REGIONS_IOS": "USA", "APP_WEB_CHECKOUT_REGIONS_ANDROID": "u1",
		"APP_MIN_VERSION_IOS": "1.0", "APP_MIN_VERSION_ANDROID": "v1.0.0", "APP_FEATURES": "Raids",
	} {
		if _, err := loadApp(env(map[string]string{name: value})); err == nil || !strings.Contains(err.Error(), name) {
			t.Errorf("%s=%s: %v", name, value, err)
		}
	}
}
