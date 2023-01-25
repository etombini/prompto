package prompter

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

// Part is an interface allowing any struct to provide its parts of a prompt
// It returns a string (the prompt part per se), the length on the screen (numbers of columns occupied)
// and an error
type Part interface {
	Prompt() (string, int)
	Side() string
	Kind() string
	IsNewline() bool
}

// NewPart return a part of a prompter
func NewPart(config map[string]string) Part {
	kind, ok := config["kind"]
	if !ok {
		return NewUnknown(map[string]string{"kind": "undefined"})
	}

	switch kind {
	case "text":
		return NewText(config)
	case "textifenv":
		return NewTextIfEnv(config)
	case "username":
		return NewUsername(config)
	case "hostname":
		return NewHostname(config)
	case "path":
		return NewPath(config)
	case "newline":
		return NewNewline(config)
	case "ifgit":
		return NewTextIfGit(config)
	case "ifnotgit":
		return NewTextIfNotGit(config)
	case "gitrepo":
		return NewGitRepo(config)
	case "gitbranch":
		return NewGitBranch(config)
	case "gittag":
		return NewGitTag(config)
	case "command":
		return NewCommand(config)
	case "status_success":
		return NewStatusSuccess(config)
	case "status_fail":
		return NewStatusFail(config)
	default:
		return NewUnknown(config)
	}
}

// Prompter is the structure handling the final delivery of the prompt string
type Prompter struct {
	parts [][]Part
}

// New returns a new Prompter
func New(config []map[string]string) Prompter {
	p := Prompter{
		parts: make([][]Part, 0),
	}

	line := 0
	p.parts = append(p.parts, make([]Part, 0))

	for _, cfg := range config {
		np := NewPart(cfg)
		p.parts[line] = append(p.parts[line], np)
		if np.IsNewline() {
			line++
			p.parts = append(p.parts, make([]Part, 0))
		}
	}
	return p
}

func (p Prompter) String() string {

	lines := make([]string, 0)

	for _, parts := range p.parts {

		var left bytes.Buffer
		var leftLength int
		var right bytes.Buffer
		var rightLength int

		for _, p := range parts {
			prompt, len := p.Prompt()
			if p.Side() == "left" {
				left.WriteString(prompt)
				leftLength += len
			} else {
				right.WriteString(prompt)
				rightLength += len
			}
		}

		columns, _, err := terminal.GetSize(0)
		if err != nil {
			fmt.Printf("ERROR COLUMS: %s\n", err.Error())
			columns = 204
		}

		var padding string
		if columns > leftLength+rightLength && rightLength > 0 {
			padding = strings.Repeat(" ", columns-(leftLength+rightLength))
		}

		var line bytes.Buffer
		line.WriteString(left.String())
		line.WriteString(padding)
		line.WriteString(right.String())
		lines = append(lines, line.String())
	}
	var res bytes.Buffer

	for _, line := range lines {
		res.WriteString(line)
	}
	return res.String()
}
