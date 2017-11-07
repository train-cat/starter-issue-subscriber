package main

import (
	"flag"
	"os"

	"github.com/train-cat/starter-issue-subscriber/helper"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	initLogging()
	initConfig()
	helper.InitHelper()
}

func initLogging() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetOutput(os.Stderr)
	log.SetLevel(log.DebugLevel)
}

func initConfig() {
	cfgFile := flag.String("config", "config.json", "config file")
	flag.Parse()

	viper.SetConfigFile(*cfgFile)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Infof("Using config file: %s", viper.ConfigFileUsed())
	} else {
		log.Error(err.Error())
		os.Exit(helper.ExitCodeErrorInitConfig)
	}
}
