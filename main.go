package main

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/charmbracelet/lipgloss"
	log "github.com/sirupsen/logrus"
)

const semVerRegex string = `([0-9]+\.[0-9]+\.[0-9])`

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
	log.Infof("%s", gitVersion)
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
}
