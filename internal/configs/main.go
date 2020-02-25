package configs

import "os"

type Config struct {
	Port string
	Mode string
	TgToken string
	FrontendUrl string
}

func New() *Config  {
	return &Config{
		Port: getEnv("PORT", ":3000"),
		Mode: getEnv("MODE", "development"),
		TgToken: getEnv("TG_TOKEN", "no-token"),
		FrontendUrl: getEnv("FRONTEND_URL", "*"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

