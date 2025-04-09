package kratos

import (
	"net/http"
	"os"
)

type KratosClient struct {
	PublicURL string
	AdminURL  string
	Http      *http.Client
}

func NewKratosClient() *KratosClient {
	return &KratosClient{
		PublicURL: os.Getenv("KRATOS_PUBLIC_URL"),
		AdminURL:  os.Getenv("KRATOS_ADMIN_URL"),
		Http:      &http.Client{},
	}
}
