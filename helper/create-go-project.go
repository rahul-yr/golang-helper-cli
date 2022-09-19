package helper

import (
	"log"
	"os/exec"
)

func checkDependencies() error {
	// check if dependencies are installed
	// check if go is installed or not using exec.LookPath
	_, err := exec.LookPath("go")
	if err != nil {
		return err
	}
	// check if git is installed or not using exec.LookPath
	_, err = exec.LookPath("git")

	return err
}

func CreateGolangProject(args []string, index int) {
	// create a new golang project
	// name := args[index+1]
	// check if dependencies are installed
	if err := checkDependencies(); err != nil {
		// print error message
		log.Println("Dependencies are not installed i.e. go and git")
	}
	// clone a folder from github
	cmd := exec.Command("git", "clone", GITHUB_GO_CONFIGS["repo"])
	err := cmd.Run()
	if err != nil {
		log.Println(err)
	}
	// copy the files from the cloned folder to the current directory
	cmd = exec.Command("cp", GITHUB_GO_CONFIGS["create_template"], ".")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
}

func Help() {
	log.Println("Help")
}
