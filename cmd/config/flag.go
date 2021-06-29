package flag

import (
	"flag"
	"os"
	"strconv"
	"sync"
)

var (
	// Inicialización con Singleton
	onceInit sync.Once
)

// Config  estructura para configuracion a traves de flags
type Config struct {
	HTTPPort      string
	LogLevel      int
	LogTimeFormat string
}

func setFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		return
	}
}

const (
	defaultHTTPPort  = "3000"
	defaultLogLevel  = "0"
	defaultLogFormat = "2006-01-02T15:04:05Z07:00"
)

// EnvValue funcion que obtiene variable de entorno sino será por defecto
func EnvValue(env, defaultValue string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		return defaultValue
	}
	return value

}

// NewFlagConfig funcion que setea los flags a la estructura Config
func NewFlagConfig() *Config {

	config := &Config{}

	var (
		httpPort    = EnvValue("HTTP_PORT", defaultHTTPPort)
		logLevel, _ = strconv.Atoi(EnvValue("LOG_LEVEL", defaultLogLevel))
		logFormat   = EnvValue("LOG_FORMAT", defaultLogFormat)
	)

	onceInit.Do(func() {

		// HTTP flags
		flag.StringVar(&config.HTTPPort, "p", httpPort, "HTTP port para bind")
		flag.StringVar(&config.HTTPPort, "http-port", httpPort, "HTTP port para bind")

		// Logs level flags
		flag.IntVar(&config.LogLevel, "l", logLevel, "Global log level")
		flag.IntVar(&config.LogLevel, "log-level", logLevel, "Global log level")

		// Logs format flags
		flag.StringVar(&config.LogTimeFormat, "f", logFormat, "Global format")
		flag.StringVar(&config.LogTimeFormat, "log-format", logFormat, "Global format")

		setFlag(flag.CommandLine)

		flag.Parse()
	})

	return config

}
