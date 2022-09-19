package helper

import (
	"bytes"
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
	var stdErr bytes.Buffer
	// create a new golang project
	// name := args[index+1]
	// check if dependencies are installed
	log.Println("Checking dependencies")
	if err := checkDependencies(); err != nil {
		// print error message
		log.Println("Dependencies are not installed i.e. go and git")
		log.Println(err)
		return
	}
	log.Println("Clone the repo")
	// clone a folder from github
	cmd := exec.Command("git", "clone", GITHUB_GO_CONFIGS["repo"])
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return
	}
	log.Println("Copying the boilerplate codes")
	// copy the files from the cloned folder to the current directory
	cmd = exec.Command("cp", GITHUB_GO_CONFIGS["create_template"], ".")
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return
	}
}

func Help() {
	log.Println("Help")
}
