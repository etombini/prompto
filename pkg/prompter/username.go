package prompter

import (
	"fmt"
	"os"
	"os/user"
)

// Username handles configuration to provide printable prompt information
type Username struct {
	Index   int    `yaml:"index"`
	Side    string `yaml:"side"`
	Before  string `yaml:"before"`
	After   string `yaml:"after"`
	Bgcolor string `yaml:"bgcolor"`
	Fgcolor string `yaml:"fgcolor"`
	Font    string `yaml:"font"`
}

//GetIndex returns the index of the Prompter
func (u Username) GetIndex() int {
	return u.Index
}

//GetSide returns the side of the Prompter
func (u Username) GetSide() string {
	if u.Side == "" || u.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the resulting string and its real length when written
func (u Username) Prompt() (string, int, error) {
	prompt := ""
	if fgcolor, ok := ForegroundColor16[u.Fgcolor]; ok {
		prompt += fgcolor
	}
	if bgcolor, ok := BackgroundColor16[u.Bgcolor]; ok {
		prompt += bgcolor
	}
	if font, ok := Font[u.Font]; ok {
		prompt += font
	}

	username := ""
	user, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not get current user\n")
		username = "ERROR"
	} else {
		username = user.Username
	}

	prompt += u.Before
	prompt += username
	prompt += u.After

	return prompt, RealLen(u.Before + username + u.After), nil

}
