package common

import (
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	DATABASE_URL   = ""
	WsAddr         = "0.0.0.0:9998"
	CurrentVersion = ""
)

type EnvConfigSpec struct {
	Debug                   bool          `envconfig:"DEBUG" default:"false"`
	DebugRequests           bool          `envconfig:"DEBUG_REQUESTS" default:"false"`
	DebugSQL                bool          `envconfig:"DEBUG_SQL" default:"false"`
	DebugWS                 bool          `envconfig:"DEBUG_WS" default:"false"`
	UIUsername              string        `envconfig:"UI_USERNAME" required:"false"`
	UIPassword              string        `envconfig:"UI_PASSWORD" required:"false"`
	DisableAnalytics        bool          `envconfig:"DISABLE_ANALYTICS" default:"false"`
	DatabaseURL             string        `envconfig:"DATABASE_URL" required:"false" default:""`
	DatabaseMaxIdleConns    int           `envconfig:"DATABASE_MAX_IDLE_CONNS" required:"false" default:2`
	DatabaseMaxOpenConns    int           `envconfig:"DATABASE_MAX_OPEN_CONNS" required:"false" default:0`
	DatabaseConnMaxLifetime time.Duration `envconfig:"DATABASE_CONN_MAX_LIFETIME" required:"false" default:0`
}

var EnvConfig EnvConfigSpec

func init() {
	godotenv.Load()

	err := envconfig.Process("", &EnvConfig)
	if err != nil {
		Log.Fatal(err.Error())
	}
}
