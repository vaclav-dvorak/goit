// This is main package
package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

const semVerRegex string = `([0-9]+\.[0-9]+\.[0-9])`

type repo struct {
	path      string
	name      string
	defBranch string
}

const (
	maxDepth = 2
	basePath = "~/work"
)

var (
	logoStyle = lipgloss.NewStyle().
			Padding(0, 2, 0, 5).
			Foreground(lipgloss.Color("63"))

	infoStyle = lipgloss.NewStyle().
			Width(12).
			Foreground(lipgloss.Color("63"))
	version, sha, date = "(devel)", "foo", "now"
	logo               = []string{ //ascii generator - font "Thin"
		"          o|",
		",---.,---..|---",
		"|   ||   |||",
		"`---|`---'``---'",
		"`---'",
	}
)

func main() {
	if _, err := exec.LookPath("git"); err != nil {
		log.Fatal("we need to have git binary installed")
	}

	out, err := exec.Command("git", "version").Output()
	if err != nil {
		log.Fatal(err)
	}
	re := regexp.MustCompile(semVerRegex)
	match := re.FindStringSubmatch(string(out))
	gitVersion := match[0]
	infoBar := []string{
		infoStyle.Render("goit v:") + version,
		infoStyle.Render("build sha:") + sha,
		infoStyle.Render("build time:") + date,
		infoStyle.Render("git v:") + gitVersion,
	}
	fmt.Println(lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Render(
		lipgloss.JoinHorizontal(lipgloss.Center,
			lipgloss.JoinVertical(lipgloss.Left,
				infoBar...,
			),
			logoStyle.Render(
				lipgloss.JoinVertical(lipgloss.Left,
					logo...,
				),
			),
		),
	))

	repos, err := getRepos()
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("%v", repos)
	log.Infof("%d", len(repos))
}

func getRepos() (repos []repo, err error) {
	baseDepth := strings.Count(basePath, string(os.PathSeparator))
	err = filepath.WalkDir(basePath, func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}

		if !info.IsDir() {
			return nil
		}

		if info.IsDir() && strings.Count(p, string(os.PathSeparator)) > (baseDepth+maxDepth+1) { //* we need +1 depth to be able to found .git directories
			return filepath.SkipDir
		}

		if info.Name() == ".git" {
			dir, _ := path.Split(p)
			repos = append(repos, repo{path: dir, name: strings.Trim(dir[len(basePath):], string(os.PathSeparator)), defBranch: "main"})
		}

		return nil
	})
	return
}
