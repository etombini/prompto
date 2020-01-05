package prompter

import (
	"os/user"
)

//Username handles configuration to provide printable prompt information
type Username struct {
	PartCommon
}

//NewUsername returns a Username struct
func NewUsername(param map[string]string) Username {
	return Username{
		PartCommon: NewPartCommon(param),
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
	return colorStringANSI(username, u.fgcolor, u.bgcolor, u.font), RealLen(username)
}
