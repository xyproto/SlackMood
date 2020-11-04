package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/voxelbrain/goptions"

	"github.com/xyproto/emojimood"
)

type options struct {
	Verbose bool          `goptions:"-v, --verbose, description='Log verbosely'"`
	Help    goptions.Help `goptions:"-h, --help, description='Show help'"`
	Config  string        `goptions:"-c, --config, description='Config Yaml file to use'"`
	Bind    string        `goptions:"-b, --bind, description='Port/Address to bind on, can also be specified with WEB_BIND environment variable'"`

	goptions.Verbs
}

func main() {
	parsedOptions := options{}

	parsedOptions.Config = "./config.yml"
	parsedOptions.Bind = os.Getenv("WEB_BIND")

	goptions.ParseAndFail(&parsedOptions)

	if parsedOptions.Verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}

	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	log.Debug("Logging verbosely!")

	err := emojimood.LoadConfig(parsedOptions.Config)
	if err != nil {
		log.WithFields(log.Fields{
			"configFile": parsedOptions.Config,
			"error":      err,
		}).Error("Could not load config file")
		os.Exit(1)
	}

	err = emojimood.OpenDB()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Could not open database")
		return
	}

	if !collector.Start() {
		os.Exit(1)
	}
	web.Start(parsedOptions.Bind)
}
