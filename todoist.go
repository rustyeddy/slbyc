package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"path/filepath"
)

type todoist struct {
	Token  string `json:"token"`
	APIURL string `json:"apiurl"`
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

func (t *todoist) getAPIURL() string {
	return t.APIURL
}

func (t *todoist) GetProjects() {
	url := t.APIURL + "projects"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+t.getToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP Request failed: %d %s", resp.StatusCode, resp.Status)
	}

	projects, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("projects: %s\n", projects)
}
