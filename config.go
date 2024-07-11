package main

import "flag"

type configuration struct {
	ConfigPath string
}

var config configuration

func init() {
	flag.StringVar(&config.ConfigPath, "config", "", "Config path defaults to nil")
}
