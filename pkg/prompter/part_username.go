package prompter

import (
	"os/user"
)

// Username handles configuration to provide printable prompt information
type Username struct {
	side    string `yaml:"side"`
	before  string `yaml:"before"`
	after   string `yaml:"after"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewUsername returns a Username struct
func NewUsername(param map[string]string) Username {
	return Username{
		side:    param["side"],
		before:  param["before"],
		after:   param["after"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
	}
}

//Side returns the side of the Prompter
func (u Username) Side() string {
	if u.side == "" || u.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (u Username) Kind() string {
	return "username"
}

//IsNewline tells if this part requires a newline to be inserted
func (u Username) IsNewline() bool {
	return false
}

//Prompt return the resulting string and its real length when written
func (u Username) Prompt() (string, int) {
	user, err := user.Current()
	if err != nil {
		return "USERNAME", len("USERNAME")
	}
	username := u.before + user.Username + u.after
	return colorString(username, u.fgcolor, u.bgcolor, u.font), RealLen(username)
}
