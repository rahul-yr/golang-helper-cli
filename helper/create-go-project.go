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
		log.Println(err)
		return err
	}
	// check if git is installed or not using exec.LookPath
	_, err = exec.LookPath("git")
	if err != nil {
		log.Println(err)
	}
	return err
}

func cloneRepo() error {
	var stdErr bytes.Buffer
	cmd := exec.Command("git", "clone", GITHUB_GO_CONFIGS["repo"])
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return err
	}
	return nil
}

func copyBoilerPlate() error {
	// copy the boilerplate codes
	var stdErr bytes.Buffer
	cmd := exec.Command("cp", "-r", GITHUB_GO_CONFIGS["create_template_folder"]+"/.", ".")
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return err
	}
	return nil
}

func deleteRepo() error {
	var stdErr bytes.Buffer
	cmd := exec.Command("rm", "-rf", GITHUB_GO_CONFIGS["main_folder"]+"/")
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return err
	}
	return nil
}

func CreateGolangProject(args []string, index int) {
	// create a new golang project
	// name := args[index+1]
	// check if dependencies are installed
	log.Println("Checking dependencies")
	if err := checkDependencies(); err != nil {
		// print error message
		return
	}
	log.Println("Clone a repo")
	// clone a folder from github
	if err := cloneRepo(); err != nil {
		return
	}
	log.Println("Copying the boilerplate codes")
	// copy the files from the cloned folder to the current directory
	if err := copyBoilerPlate(); err != nil {
		return
	}
	log.Println("Deleting the cloned repo")
	// delete the cloned repo
	if err := deleteRepo(); err != nil {
		return
	}
	log.Println("Done")
}

func Help() {
	log.Println("Help")
}
