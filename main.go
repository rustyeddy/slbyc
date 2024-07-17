package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	flag.Parse()

	todoist := newTodoist()
	fmt.Printf("%+v\n", todoist)

	projects := todoist.GetProjects()
	for _, name := range todoist.ProjectNames {
		proj := projects.GetProject(name)
		if proj == nil {
			log.Println("Could not find project: ", name)
		}
		fmt.Printf("project %s: %+v\n", name, proj)

		tasks := proj.GetTasks()
		fmt.Println("Tasks: ")
		for _, t := range *tasks {
			fmt.Printf("%+v\n", t)
		}
	}
}
