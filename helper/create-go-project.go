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
	// check folder is exist or not
	cmd := exec.Command("ls", "/"+GITHUB_GO_CONFIGS["main_folder"])
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		// delete the folder
		cmd_delete := exec.Command("rm", "-rf", "/"+GITHUB_GO_CONFIGS["main_folder"])
		cmd_delete.Stderr = &stdErr
		err = cmd_delete.Run()
		if err != nil {
			log.Println("Error while deleting the folder")
			log.Println(err)
			return
		}
	}
	log.Println("Clone a repo")
	// clone a folder from github
	cmd = exec.Command("git", "clone", GITHUB_GO_CONFIGS["repo"])
	cmd.Stderr = &stdErr
	err = cmd.Run()
	if err != nil {
		log.Println(stdErr.String())
		return
	}
	log.Println("Copying the boilerplate codes")
	// copy the files from the cloned folder to the current directory
	cmd = exec.Command("cp", "-r", GITHUB_GO_CONFIGS["create_template_folder"]+"/.", ".")
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
