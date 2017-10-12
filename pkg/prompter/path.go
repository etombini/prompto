package prompter

import (
	"errors"
	"os"
	"os/user"
	"strings"
)

//Path handles configuration to provide printable prompt information
type Path struct {
	Mode      string `yaml:"mode"`
	Index     int    `yaml:"index"`
	Side      string `yaml:"side"`
	Before    string `yaml:"before"`
	After     string `yaml:"after"`
	Separator string `yaml:"separator"`
	Bgcolor   string `yaml:"bgcolor"`
	Fgcolor   string `yaml:"fgcolor"`
	Font      string `yaml:"font"`
}

//GetSide returns the side of the Prompter
func (p Path) GetSide() string {
	if p.Side == "" || p.Side != "right" {
		return "left"
	}
	return "right"
}

var mode = map[string]bool{
	"default": true,
	"short":   true,
	"elided":  true,
	"initial": true,
}

func pathDefault(path string) string {
	return path
}

func pathShort(path string) string {
	user, err := user.Current()
	if err != nil {
		return path
	}
	short := strings.Replace(path, user.HomeDir, "~", 1)
	return short
}

func pathElided(path string) string {
	short := pathShort(path)

	splitted := strings.Split(short, "/")
	if len(splitted) < 4 {
		return short
	}
	shorten := append(splitted[0:2], "...", splitted[len(splitted)-1])
	elided := strings.Join(shorten, "/")
	return elided
}

func pathInitial(path string) string {
	short := pathShort(path)

	splitted := strings.Split(short, "/")
	for i, value := range splitted {
		if i < len(splitted)-1 {
			initial := value[0:1]
			splitted[i] = initial
		}
	}
	initial := strings.Join(splitted, "/")
	return initial
}

//Prompt return the resulting string and its real length when written
func (p Path) Prompt() (string, int, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "NO CWD", len("NO CWD"), errors.New("Can not get current working dir")
	}

	var path string
	switch p.Mode {
	case "short":
		path = pathShort(wd)
	case "elided":
		path = pathElided(wd)
	case "initial":
		path = pathInitial(wd)
	default:
		path = pathDefault(wd)
	}

	splitted := strings.Split(path, "/")
	if p.Separator != "" {
		path = strings.Join(splitted, p.Separator)
	}

	path = p.Before + path + p.After
	return colorString(path, p.Fgcolor, p.Bgcolor, p.Font), RealLen(path), nil
}
