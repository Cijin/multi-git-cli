package helpers

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
)

func ConfigureGit() (err error) {
	err = exec.Command("git", "config", "--global", "user.email", "gg@gg.com").Run()
	if err != nil {
		return
	}

	err = exec.Command("git", "config", "--global", "user.name", "Gigi").Run()
	return
}

func CreateDir(baseDir string, name string, initGit bool) (err error) {
	dirName := path.Join(baseDir, name)

	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		return
	}

	if !initGit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}

	defer os.Chdir(currDir)

	os.Chdir(dirName)
	err = exec.Command("git", "init").Run()
	return
}

func AddFiles(baseDir string, dirName string, commit bool, filenames ...string) (err error) {
	dir := path.Join(baseDir, dirName)

	for _, f := range filenames {
		data := []byte("data for" + f)
		err = ioutil.WriteFile(path.Join(dir, f), data, 777)
		if err != nil {
			return
		}
	}

	if !commit {
		return
	}

	currDir, err := os.Getwd()
	if err != nil {
		return
	}

	defer os.Chdir(currDir)
	os.Chdir(dir)
	output, err := exec.Command("git", "add", "-A").CombinedOutput()
	if err != nil {
		return
	}
	fmt.Println("git add:", output)

	output, err = exec.Command("git", "commit", "-m", "added some files...").CombinedOutput()

	fmt.Println("git commit:", string(output))
	fmt.Println("err:", err)
	return
}
