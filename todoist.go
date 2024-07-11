package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
)

type todoist struct {
	Token       string `json:"token"`
	APIURL      string `json:"apiurl"`
	ProjectName string `json:"project_name"`
}

var projects Projects

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

	restcli = &restClient{
		todo: t,
	}
	return t
}

func (t *todoist) getToken() string {
	return t.Token
}

func (t *todoist) getAPIURL() string {
	return t.APIURL
}

type Project struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ParentID     string `json:"parent_id"`
	Order        int    `json:"order"`
	CommentCount int    `json:"comment_count"`
	URL          string `json:"url"`
}

type Projects []*Project

func (t *todoist) GetProjects() (projs *Projects) {

	jbytes := restcli.Get("projects")

	err := json.Unmarshal(jbytes, &projects)
	if err != nil {
		log.Fatalf("Failed to unmarshal projects: %v", err)
	}

	return &projects
}

func (t *todoist) GetProject(name string) (proj *Project) {

	if len(projects) == 0 {
		t.GetProjects()
	}
	if len(projects) == 0 {
		log.Fatalf("Failed to get projects")
	}

	for _, proj = range projects {
		if proj.Name == name {
			return proj
		}
	}
	return nil
}
