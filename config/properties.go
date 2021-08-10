package config

const (
	// Hostname represents the URL entrypoint of our application
	Hostname = "https://api.phoenixnap.com/bmc/v1/"
	// TokenURL represents the URL of the OpenID Connect provider from where we can retrieve a token
	TokenURL = "https://auth.phoenixnap.com/auth/realms/BMC/protocol/openid-connect/token"
)
