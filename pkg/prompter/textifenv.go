package prompter

import (
	"os"
)

//TextIfEnv handles configuration to provide printable prompt information
type TextIfEnv struct {
	Env     string `yaml:"env"`
	Content string `yaml:"content"`
	Index   int    `yaml:"index"`
	Side    string `yaml:"side"`
	Bgcolor string `yaml:"bgcolor"`
	Fgcolor string `yaml:"fgcolor"`
	Font    string `yaml:"font"`
}

//GetSide returns the side of the Prompter
func (t TextIfEnv) GetSide() string {
	if t.Side == "" || t.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the string to be written on the screen and its proper length
func (t TextIfEnv) Prompt() (string, int, error) {
	if len(os.Getenv(t.Env)) == 0 {
		return "", 0, nil
	}
	text := Text{t.Content, t.Index, t.Side, t.Bgcolor, t.Fgcolor, t.Font}
	return text.Prompt()
}
