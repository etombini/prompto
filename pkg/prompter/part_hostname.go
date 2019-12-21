package prompter

import (
	"os"
	"strings"
)

//Hostname handles configuration to provide printable prompt information
type Hostname struct {
	PartCommon
}

//NewHostname returns a Hostname struct
func NewHostname(param map[string]string) Hostname {
	return Hostname{
		PartCommon: NewPartCommon(param),
	}
}

//Side returns the side of the Prompter
func (h Hostname) Side() string {
	if h.side == "" || h.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (h Hostname) Kind() string {
	return "hostname"
}

//IsNewline tells if this part requires a newline to be inserted
func (h Hostname) IsNewline() bool {
	return false
}

//Prompt return the resulting string and its real length when written
func (h Hostname) Prompt() (string, int) {
	hostname, err := os.Hostname()
	if err != nil {
		return "HOSTNAME", len("HOSTNAME")
	}
	hostname = strings.Split(hostname, ".")[0]
	hostname = h.before + hostname + h.after

	return colorStringANSI(hostname, h.fgcolor, h.bgcolor, h.font), RealLen(hostname)
}
