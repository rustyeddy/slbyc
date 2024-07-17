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

func TestProjects(t *testing.T) {
	todo := newTodoist()

	projs := todo.GetProjects()
	if len(*projs) == 0 {
		t.Errorf("Expected projects (> 0) got (0)")
	}

	name := todo.ProjectName
	proj := todo.GetProject(name)
	if proj == nil {
		t.Errorf("Expected project (%s) got nil", name)
	} else if proj.Name != name {
		t.Errorf("Expected project (%s) got (%s)", name, proj.Name)
	}

}
