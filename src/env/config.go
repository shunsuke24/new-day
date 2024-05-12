package env

import (
	"os"
)

const (
	LOCAL = "local"
	DEV   = "dev"
)

type Config struct {
	AppEnv                   string
	CloudFrontAppDataSetting *CloudFrontAppDataSetting
	CognitoSetting           *CognitoSetting
	CookieSetting            *CookieSetting
}

type CloudFrontAppDataSetting struct {
	URL           string
	PublicKeyID   string
	PrivateKeyPEM string
}

type CognitoSetting struct {
	AWSRegion         string
	CognitoClientID   string
	CognitoUserPoolID string
}

type CookieSetting struct {
	CookieDomain string
}

func NewConfig() *Config {
	appEnv := os.Getenv("APP_ENV") // "local" or "dev"
	if appEnv == "" {
		appEnv = LOCAL
	}

	return &Config{
		AppEnv: appEnv,
	}
}
