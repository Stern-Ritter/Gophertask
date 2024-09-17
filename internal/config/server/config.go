package server

type ServerConfig struct {
	URL               string `env:"ADDRESS"`      // The address and port to run the server
	DatabaseDSN       string `env:"DATABASE_DSN"` // The database DSN
	AuthenticationKey string `env:"AUTH_KEY"`     // The secret key for authentication
	TLSCertPath       string `env:"TLS_CERT"`     // The path to TLS certificate
	TLSKeyPath        string `env:"TLS_KEY"`      // The path to TLS key
	ConfigFile        string `env:"CONFIG"`       // The path to json config file
	ShutdownTimeout   int    // The server shutdown timeout in seconds
	LoggerLvl         string // The logging level
}
