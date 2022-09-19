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

func createGoModule(name string) error {
	// create a go module
	var stdErr bytes.Buffer
	cmd := exec.Command("go", "mod", "init", name)
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return err
	}
	return nil
}

func updateGoModule() error {
	var stdErr bytes.Buffer
	cmd := exec.Command("go", "mod", "tidy")
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
	log.Println("Clone the public repository")
	// clone a folder from github
	if err := cloneRepo(); err != nil {
		return
	}
	log.Println("Copy the boilerplate codes")
	// copy the files from the cloned folder to the current directory
	if err := copyBoilerPlate(); err != nil {
		return
	}
	log.Println("Cleaning the repo")
	// delete the cloned repo
	if err := deleteRepo(); err != nil {
		return
	}
	log.Println("Creating a go module")
	// create a go module
	if err := createGoModule(args[index+1]); err != nil {
		return
	}
	log.Println("Updating the go module")
	// update the go module
	if err := updateGoModule(); err != nil {
		return
	}
	log.Println("Done")
}
