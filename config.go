package main

import "os"

// Settings holds the current settings of this service
var Settings FlipperConfig

// FlipperConfig defines all the configurable settings for the Flipper API
type FlipperConfig struct {
	FlipperProxServiceListener string
	FlipperAppServiceHost      string
	FlipperAppServicePort      string
	FlipperAPIServiceHost      string
	FlipperAPIServicePort      string
	CertbotChallengePrompt     string
	CertbotChallengeResponse   string
}

// InitializeSettings takes care of reading environment variables into memory
func InitializeSettings() error {
	Settings = FlipperConfig{
		FlipperProxServiceListener: os.Getenv("FLIPPERPROX_SERVICE_LISTENER"),
		FlipperAppServiceHost:      os.Getenv("FLIPPERAPP_SERVICE_HOST"),
		FlipperAppServicePort:      os.Getenv("FLIPPERAPP_SERVICE_PORT"),
		FlipperAPIServiceHost:      os.Getenv("FLIPPERAPI_SERVICE_HOST"),
		FlipperAPIServicePort:      os.Getenv("FLIPPERAPI_SERVICE_PORT"),
		CertbotChallengePrompt:     os.Getenv("FLIPPER_CERTBOT_CHALLENGE_PROMPT"),
		CertbotChallengeResponse:   os.Getenv("FLIPPER_CERTBOT_CHALLENGE_RESPONSE"),
	}
	return nil
}
