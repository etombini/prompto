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

//GetSide returns the side of the Prompter
func (u Username) GetSide() string {
	if u.Side == "" || u.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the resulting string and its real length when written
func (u Username) Prompt() (string, int, error) {
	username := ""
	user, err := user.Current()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not get current user\n")
		username = "ERROR"
	} else {
		username = user.Username
	}

	username = u.Before + username + u.After
	return colorString(username, u.Fgcolor, u.Bgcolor, u.Font), RealLen(username), nil
}
