package config

import "fmt"

type Config struct {
	Port          string
	DatabaseURL   string
	ResendAPIKey  string
	PublicBaseURL string
	APIBaseURL    string
}

func Load(getenv func(string) string) (Config, error) {
	c := Config{
		Port:          getenv("PORT"),
		DatabaseURL:   getenv("DATABASE_URL"),
		ResendAPIKey:  getenv("RESEND_API_KEY"),
		PublicBaseURL: getenv("PUBLIC_BASE_URL"),
		APIBaseURL:    getenv("API_BASE_URL"),
	}
	if c.Port == "" {
		c.Port = "8080"
	}
	for name, v := range map[string]string{
		"DATABASE_URL": c.DatabaseURL, "RESEND_API_KEY": c.ResendAPIKey,
		"PUBLIC_BASE_URL": c.PublicBaseURL, "API_BASE_URL": c.APIBaseURL,
	} {
		if v == "" {
			return Config{}, fmt.Errorf("config: %s is required", name)
		}
	}
	return c, nil
}
