package repoManager

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecGitCommand(dir string, command string, ignoreErrors bool) ([]string, error) {
	// get git repositories, if existing
	gitRepos, err := getGitRepos(dir)

	if err != nil {
		return nil, err
	}

	var output []string
	// handle multi word commands ex: "add ."
	gitCommand := strings.Split(command, " ")

	for _, gitRepo := range gitRepos {
		fmt.Printf("Executing command 'git %s' in dir '%s'\n", command, gitRepo)
		cmd := exec.Command("git", gitCommand...)
		cmd.Dir = gitRepo
		out, err := cmd.CombinedOutput()

		output = append(output, string(out))

		// exit if errors and ignoreErrors set to false
		if err != nil && !ignoreErrors {
			return output, err
		}
	}
	return output, nil
}

/*
 * @param filesInfo is the output of ReadDir
 * @return a slice of paths to gitRepos within @param folder
 *
 * Iterates through the list of files checks if they are git repositories
 * & returns a slice of git repositories
 */
func getGitRepos(rootDir string) ([]string, error) {
	var gitRepos []string

	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		// ignore files
		if !file.IsDir() {
			continue
		}

		rootDir, err = filepath.Abs(rootDir)

		if err != nil {
			return gitRepos, err
		}

		childFolder := rootDir + "/" + file.Name()
		_, err := os.Stat(childFolder + "/.git")

		if err != nil {
			log.Fatal(err)
			continue
		}

		gitRepos = append(gitRepos, childFolder)
	}

	return gitRepos, nil
}
