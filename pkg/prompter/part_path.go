package prompter

import (
	"os"
	"os/user"
	"strings"
)

//Path handles configuration to provide printable prompt information
type Path struct {
	mode string `yaml:"mode"`
	// Index     int    `yaml:"index"`
	side      string `yaml:"side"`
	before    string `yaml:"before"`
	after     string `yaml:"after"`
	separator string `yaml:"separator"`
	bgcolor   string `yaml:"bgcolor"`
	fgcolor   string `yaml:"fgcolor"`
	font      string `yaml:"font"`
}

//NewPath returns a Username struct
func NewPath(param map[string]string) Path {
	return Path{
		side:      param["side"],
		before:    param["before"],
		after:     param["after"],
		bgcolor:   param["bgcolor"],
		fgcolor:   param["fgcolor"],
		font:      param["font"],
		separator: param["separator"],
		mode:      param["mode"],
	}
}

//Side returns the side of the Prompter
func (p Path) Side() string {
	if p.side == "" || p.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (p Path) Kind() string {
	return "path"
}

//IsNewline tells if this part requires a newline to be inserted
func (p Path) IsNewline() bool {
	return false
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
func (p Path) Prompt() (string, int) {
	wd, err := os.Getwd()
	if err != nil {
		return "NO CWD", len("NO CWD")
	}

	var path string
	switch p.mode {
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
	if p.separator != "" {
		path = strings.Join(splitted, p.separator)
	}

	path = p.before + path + p.after
	return colorString(path, p.fgcolor, p.bgcolor, p.font), RealLen(path)
}
