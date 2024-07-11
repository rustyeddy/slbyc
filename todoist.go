package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
)

type todoist struct {
	Token string `json:"token"`
}

func newTodoist() (t *todoist) {

	user, err := user.Current()
	if err != nil {
		log.Fatal("Can not open user directory")
	}

	var configfile string
	if config.ConfigPath == "" {
		configfile = user.HomeDir
	}
	configfile = filepath.Join(configfile, ".config/todoist.json")
	configstr, err := ioutil.ReadFile(configfile)

	if err != nil {
		log.Fatalf("Failed to read config %s: %v", configstr, err)
	}

	t = &todoist{}
	err = json.Unmarshal([]byte(configstr), t)
	if err != nil {
		log.Fatalf("Failed to decode config JSON %v", err)
	}

	return t
}

func (t *todoist) getToken() string {
	return t.Token
}
