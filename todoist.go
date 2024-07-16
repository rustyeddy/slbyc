package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/url"
	"os/user"
	"path/filepath"
)

// todoist provides the authorization info to get access to
// todoist API
type todoist struct {
	Token        string   `json:"token"`
	APIURL       string   `json:"apiurl"`
	ProjectNames []string `json:"project-names"`
}

// projects variable makes a list of all the projects we will
// be tracking in this application
var projects Projects

// newTodoist will create a new instance of todoist
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

// getToken returns the Todoist API token
func (t *todoist) getToken() string {
	return t.Token
}

// getAPIURL will return the URL for the Todoist API
func (t *todoist) getAPIURL() string {
	return t.APIURL
}

// Project represents a project that we will be tracking
// and managing tasks for
type Project struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ParentID     string `json:"parent_id"`
	Order        int    `json:"order"`
	CommentCount int    `json:"comment_count"`
	URL          string `json:"url"`
}

// Projects is an array of the Project pointers we are
// interested in
type Projects []*Project

// GetProjects will return all projects that we will nre getting all tasks
func (t *todoist) GetProjects() (projs *Projects) {

	jbytes := restcli.Get("projects", nil)
	err := json.Unmarshal(jbytes, &projects)
	if err != nil {
		log.Fatalf("Failed to unmarshal projects: %v", err)
	}

	return &projects
}

// GetProject retrieve the project structure of the given name
func (p *Projects) GetProject(name string) (proj *Project) {

	for _, proj = range *p {
		if proj.Name == name {
			return proj
		}
	}
	return nil
}

type Task struct {
	CreatorID string `json:"creator_id"`

	/*
	   "created_at": "2019-12-11T22:36:50.000000Z",
	   "assignee_id": "2671362",
	   "assigner_id": "2671355",
	   "comment_count": 10,
	   "is_completed": false,
	   "content": "Buy Milk",
	   "description": "",
	   "due": {
	       "date": "2016-09-01",
	       "is_recurring": false,
	       "datetime": "2016-09-01T12:00:00.000000Z",
	       "string": "tomorrow at 12",
	       "timezone": "Europe/Moscow"
	   },
	   "duration": {
	        "amount": 15,
	        "unit": "minute"
	   },
	   "id": "2995104339",
	   "labels": ["Food", "Shopping"],
	   "order": 1,
	   "priority": 1,
	   "project_id": "2203306141",
	   "section_id": "7025",
	   "parent_id": "2995104589",
	   "url": "https://todoist.com/showTask?id=2995104339"
	*/
}

type Tasks []*Task

func (p *Project) GetTasks() *Tasks {

	params := &url.Values{}
	params.Add("project_id", p.ID)

	jbytes := restcli.Get("tasks", params)

	var t Tasks
	err := json.Unmarshal(jbytes, &t)
	if err != nil {
		log.Fatalf("Failed to unmarshal projects: %v", err)
	}

	return &t
}
