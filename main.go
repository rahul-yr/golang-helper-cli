package main

import (
	"os"
	"rahul-yr/golang-helper-cli/helper"
)

func argumentsMapper(args []string) {
	for index, arg := range args {
		switch arg {
		case helper.SA_CREATE_GOLANG_PROJECT[0], helper.SA_CREATE_GOLANG_PROJECT[1]:
			helper.CreateGolangProject(args, index)
		case helper.SA_HELP[0], helper.SA_HELP[1]:
			helper.Help()
		case helper.SA_ENABLE_ADDONS[0]:
			helper.EnableAddons(args[index+1])
		}
	}
}

func main() {
	// read all command line arguments
	arguments := os.Args[1:]
	argumentsMapper(arguments)
}
