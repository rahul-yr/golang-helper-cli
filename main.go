package main

import (
	"os"
	"rahul-yr/golang-helper-cli/helper"
)

var (
	SA_CREATE_GOLANG_PROJECT = []string{
		"--create-go",
		"-cgo",
	}
	SA_HELP = []string{
		"--help",
		"-h",
	}
)

func argumentsMapper(args []string) {
	for index, arg := range args {
		switch arg {
		case SA_CREATE_GOLANG_PROJECT[0], SA_CREATE_GOLANG_PROJECT[1]:
			helper.CreateGolangProject(args, index)
		case SA_HELP[0], SA_HELP[1]:
			helper.Help()
		}
	}
}

func main() {
	// read all command line arguments
	arguments := os.Args[1:]
	argumentsMapper(arguments)
}
