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
	path string
	name string
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
			Width(14).
			Foreground(lipgloss.Color("202"))
	version, date = "(devel)", "now"
	logo          = []string{
		`          o`,
		` __,  __    _|_`,
		`/  | /  \_|  |`,
		`\_/|/\__/ |_/|_/`,
		`  /|`,
		`  \|`,
	}
)

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
	infoBar := []string{
		infoStyle.Render("goit v:") + version,
		infoStyle.Render("build date:") + date,
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
		if info.IsDir() && strings.Count(p, string(os.PathSeparator)) > (baseDepth+maxDepth+1) { // we need +1 depth to be able to found .git directories
			return nil
		}
		if info.Name() == ".git" {
			dir, _ := path.Split(p)
			repos = append(repos, repo{path: dir, name: strings.Trim(dir[len(basePath):], string(os.PathSeparator))})
		}

		return nil
	})
	return
}
