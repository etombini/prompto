package prompter

import (
	"os"
)

//TextIfEnv handles configuration to provide printable prompt information
type TextIfEnv struct {
	env     string `yaml:"env"`
	content string `yaml:"content"`
	side    string `yaml:"side"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewTextIfEnv returns a TextIfEnv struct
func NewTextIfEnv(param map[string]string) TextIfEnv {
	return TextIfEnv{
		side:    param["side"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
		env:     param["env"],
		content: param["content"],
	}
}

//Side returns the side of the Prompter
func (t TextIfEnv) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (t TextIfEnv) Kind() string {
	return "textifenv"
}

//IsNewline tells if this part requires a newline to be inserted
func (t TextIfEnv) IsNewline() bool {
	return false
}

//Prompt return the string to be written on the screen and its proper length
func (t TextIfEnv) Prompt() (string, int) {
	if len(os.Getenv(t.env)) == 0 {
		return "", 0
	}
	text := Text{
		content: t.content,
		side:    t.side,
		bgcolor: t.bgcolor,
		fgcolor: t.fgcolor,
		font:    t.font,
	}
	return text.Prompt()
}
