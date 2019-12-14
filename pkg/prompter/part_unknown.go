package prompter

import "fmt"

//Unknown handles unknown kind of prompter parts
type Unknown struct {
	side string
	kind string
}

//NewUnknown return a new Unknown struct
func NewUnknown(config map[string]string) Unknown {
	return Unknown{
		side: config["side"],
		kind: config["kind"],
	}
}

//Side return the side of this part
func (u Unknown) Side() string {
	if u.side == "" || u.side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the resulting string and its real length when written
func (u Unknown) Prompt() (string, int) {
	str := fmt.Sprintf("[UNKNOWN %s]", u.kind)
	return str, len(str)
}

//Kind return the kind of part
func (u Unknown) Kind() string {
	return "username"
}

//IsNewline tells if this part requires a newline to be inserted
func (u Unknown) IsNewline() bool {
	return false
}
