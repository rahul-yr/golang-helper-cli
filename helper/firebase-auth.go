package helper

import (
	"bytes"
	"log"
	"os/exec"
)

func copyFirebaseBoilerPlate() error {
	// copy the boilerplate codes
	var stdErr bytes.Buffer
	cmd := exec.Command("cp", GITHUB_GO_CONFIGS["firebase-auth"], "./boilerplate/")
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return err
	}
	return nil
}

func EnableFirebaseAuth() {
	log.Println("Clone the public repository")
	// clone a folder from github
	if err := cloneRepo(); err != nil {
		return
	}
	log.Println("Copy the boilerplate codes")
	// copy the files from the cloned folder to the current directory
	if err := copyFirebaseBoilerPlate(); err != nil {
		return
	}
	log.Println("Delete the public repository")
	// delete the public repository
	if err := deleteRepo(); err != nil {
		return
	}
	log.Println("Updating the go module")
	// update the go module
	if err := updateGoModule(); err != nil {
		return
	}
}
