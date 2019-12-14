package prompter

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

//GitBranch handles configuration to provide printable prompt information
type GitBranch struct {
	before  string `yaml:"before"`
	after   string `yaml:"after"`
	fgcolor string `yaml:"fgcolor"`
	bgcolor string `yaml:"bgcolor"`
	font    string `yaml:"font"`
	side    string
}

//NewGitBranch returns a GitBranch struct
func NewGitBranch(param map[string]string) GitBranch {
	return GitBranch{
		side:    param["side"],
		before:  param["before"],
		after:   param["after"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
	}
}

func branch() string {
	if !isGit() {
		return ""
	}

	path := rootDir()

	headFile := filepath.Join(path, ".git/HEAD")
	headBuf, err := ioutil.ReadFile(headFile)
	if err != nil {
		return ""
	}
	head := string(headBuf)

	if !strings.HasPrefix(head, "ref: refs/heads/") {
		return string(head[:10])
	}

	branch := strings.TrimPrefix(head, "ref: refs/heads/")
	if strings.HasSuffix(branch, "\n") {
		branch = strings.TrimSuffix(branch, "\n")
	}

	return branch
}

func isBranch() bool {
	if branch() != "" {
		return true
	}
	return false
}

//Side returns the side of the Prompter
func (g GitBranch) Side() string {
	if g.side == "" || g.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (g GitBranch) Kind() string {
	return "gitbranch"
}

//IsNewline tells if this part requires a newline to be inserted
func (g GitBranch) IsNewline() bool {
	return false
}

//Prompt return the resulting string and its real length when written
func (g GitBranch) Prompt() (string, int) {
	if !isGit() || branch() == "" {
		return "", 0
	}

	b := g.before + branch() + g.after

	return colorString(b, g.fgcolor, g.bgcolor, g.font), RealLen(b)
}
