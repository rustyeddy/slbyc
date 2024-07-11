package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type restClient struct {
	todo *todoist
}

var restcli *restClient

func (rc restClient) Get(endpoint string) (jbytes []byte) {
	url := rc.todo.APIURL + endpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Authorization", "Bearer "+rc.todo.getToken())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("HTTP Request failed: %d %s", resp.StatusCode, resp.Status)
	}

	jbytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return jbytes
}
