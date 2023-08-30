package appconfig

import (
	"flag"
	"github.com/codingconcepts/env"
	"os"
)

type AppConfig struct {
	DebugMode bool   `default:"false"`
	Host      string `env:"MY_HOME_HOST_NAME" default:"localhost"`
	Port      string `env:"MY_HOME_HOST_PORT" default:"8080"`
}

var envSet = env.Set

func (ac *AppConfig) Init() error {

	// pull values from the environment / set defaults
	if err := envSet(ac); err != nil {
		return err
	}

	// Pull values from flags
	flags := flag.NewFlagSet("plexr", flag.ContinueOnError)
	flags.BoolVar(&ac.DebugMode, "debug", ac.DebugMode, "Run the app in debug mode for more logging and suppressed features")

	// Webservice Flags
	flags.StringVar(&ac.Host, "host", ac.Host, "the host name to run the service on")
	flags.StringVar(&ac.Port, "port", ac.Port, "the host port to run the service on")

	if err := flags.Parse(os.Args[1:]); err != nil {
		return err
	}

	return ac.Validate()
}

// Validate check if we have proper values before continuing
func (ac *AppConfig) Validate() error {

	return nil
}
