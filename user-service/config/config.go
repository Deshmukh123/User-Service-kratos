package config

import "os"

type Config struct {
	KratosPublicURL string
	KratosAdminURL  string
	Port            string
}

func LoadConfig() Config {
	return Config{
		KratosPublicURL: os.Getenv("KRATOS_PUBLIC_URL"),
		KratosAdminURL:  os.Getenv("KRATOS_ADMIN_URL"),
		Port:            os.Getenv("PORT"),
	}
}
