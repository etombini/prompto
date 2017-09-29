package prompter

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
)

//Prompter is an interface allowing any struct to provide its parts of a prompt
//It returns a string (the prompt part per se), the length on the screen (numbers of columns occupied)
//and an error
type Prompter interface {
	Prompt() (string, int, error)
	GetSide() string
}

//Config is handling the top level YAML configuration which is a list of Lines
type Config struct {
	Lines []Line `yaml:"lines"`
}

//Line is composed of a list of PromptItem
type Line struct {
	Items []PromptItem `yaml:"line"`
}

//PromptItem is meant to be used as a union, only one attribute is to be set per declaration
type PromptItem struct {
	Username     Username     `yaml:"username"`
	Hostname     Hostname     `yaml:"hostname"`
	Path         Path         `yaml:"path"`
	Git          Git          `yaml:"git"`
	Text         Text         `yaml:"text"`
	TextIfGit    TextIfGit    `yaml:"ifgit"`
	TextIfNotGit TextIfNotGit `yaml:"ifnotgit"`
}

//GetPrompt returns PromptItem first non nil attribute as a Prompter interface
func (p *PromptItem) GetPrompt() Prompter {
	if p.Username != (Username{}) {
		return p.Username
	}
	if p.Hostname != (Hostname{}) {
		return p.Hostname
	}
	if p.Path != (Path{}) {
		return p.Path
	}
	if p.Git != (Git{}) {
		return p.Git
	}
	if p.Text != (Text{}) {
		return p.Text
	}
	if p.TextIfGit != (TextIfGit{}) {
		return p.TextIfGit
	}
	if p.TextIfNotGit != (TextIfNotGit{}) {
		return p.TextIfNotGit
	}
	return nil
}

func (l Line) String() string {
	prompters := make([]Prompter, 0)
	for _, v := range l.Items {
		prompt := v
		prompters = append(prompters, prompt.GetPrompt())
	}

	var left string
	var leftLength int
	var right string
	var rightLength int

	for _, p := range prompters {
		prompt, len, err := p.Prompt()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can not get prompt from '%T'\n", p)
			os.Exit(1)
		}
		if p.GetSide() == "left" {
			left += prompt
			left += Font["reset"]
			leftLength += len
		} else {
			right += prompt
			right += Font["reset"]
			rightLength += len
		}

	}

	if rightLength > 0 {
		right += "\n"
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

	line := left + padding + right
	return line
}
