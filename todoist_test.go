package main

import "testing"

func TestToken(t *testing.T) {
	todo := newTodoist()

	token := todo.getToken()
	if token == "" {
		t.Error("Expected a token string but got the empty string")
	}
}

func TestAPIURL(t *testing.T) {
	todo := newTodoist()

	apiurl := todo.getAPIURL()
	if apiurl == "" {
		t.Error("Expected a token string but got the empty string")
	}
}
