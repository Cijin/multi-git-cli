package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	repomanager "github.com/multiGit/pkg/repoManager"
)

const envKey = "ROOTDIR"

var rootDir string

func init() {
	rootDir = os.Getenv(envKey)

	if rootDir == "" {
		panic(fmt.Sprintf("Goenv missing key: %s", envKey))
	}
}

func main() {
	command := flag.String("command", "", "The git command to be executed. Ex: ls")
	ignoreErrors := flag.Bool("ignoreErrors", false, "Pass true if errors are to be ignored. Not ignored by default")

	flag.Parse()
	if *command == "" {
		panic("command flag cannot be empty")
	}

	// execute command on root directory
	outputs, err := repomanager.ExecGitCommand(rootDir, *command, *ignoreErrors)

	if err != nil {
		log.Fatal(err)
	}

	for _, output := range outputs {
		fmt.Println(output)
	}

	fmt.Println("Done.")
}
