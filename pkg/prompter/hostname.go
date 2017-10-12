package prompter

import (
	"fmt"
	"os"
	"strings"
)

//Hostname handles configuration to provide printable prompt information
type Hostname struct {
	Mode    string `yaml:"mode"`
	Index   int    `yaml:"index"`
	Side    string `yaml:"side"`
	Before  string `yaml:"before"`
	After   string `yaml:"after"`
	Bgcolor string `yaml:"bgcolor"`
	Fgcolor string `yaml:"fgcolor"`
	Font    string `yaml:"font"`
}

//GetSide returns the side of the Prompter
func (h Hostname) GetSide() string {
	if h.Side == "" || h.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the resulting string and its real length when written
func (h Hostname) Prompt() (string, int, error) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can not get current user\n")
		hostname = "ERROR"
	}
	hostname = strings.Split(hostname, ".")[0]
	hostname = h.Before + hostname + h.After

	return colorString(hostname, h.Fgcolor, h.Bgcolor, h.Font), RealLen(hostname), nil
}
