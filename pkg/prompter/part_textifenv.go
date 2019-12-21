package prompter

import (
	"os"
)

//TextIfEnv handles configuration to provide printable prompt information
type TextIfEnv struct {
	PartCommon
	env     string `yaml:"env"`
	content string `yaml:"content"`
}

//NewTextIfEnv returns a TextIfEnv struct
func NewTextIfEnv(param map[string]string) TextIfEnv {
	return TextIfEnv{
		PartCommon: NewPartCommon(param),
		env:        param["env"],
		content:    param["content"],
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
		PartCommon: t.PartCommon,
		content:    t.content,
	}
	return text.Prompt()
}
