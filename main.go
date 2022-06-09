package main

import (
	"os/exec"
	"regexp"

	log "github.com/sirupsen/logrus"
)

const semVerRegex string = `([0-9]+\.[0-9]+\.[0-9])`

func main() {
	if _, err := exec.LookPath("git"); err != nil {
		log.Fatal("we need to have git installed")
	}

	out, err := exec.Command("git", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(semVerRegex)
	match := re.FindStringSubmatch(string(out))
	gitVersion := match[0]
	log.Infof("%s", gitVersion)

}
