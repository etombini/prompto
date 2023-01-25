package prompter

import "os"

// StatusFail handles configuration to provide printable prompt information
type StatusFail struct {
	PartCommon
	content string `yaml:"content"`
}

// NewStatusFail returns a StatusFail struct.
// The "content" key in the map expands environment variables when Prompt() is called. It is not done at
// declaration time to allow other components to benefit from this.
func NewStatusFail(param map[string]string) StatusFail {
	status := os.ExpandEnv("${STATUS}")
	if status == "0" {
		return StatusFail{}
	}
	return StatusFail{
		PartCommon: NewPartCommon(param),
		content:    param["content"],
	}
}

// Side returns the side of the Prompter
func (t StatusFail) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

// Kind return the kind of part
func (t StatusFail) Kind() string {
	return "text"
}

// IsNewline tells if this part requires a newline to be inserted
func (t StatusFail) IsNewline() bool {
	return false
}

// Prompt return the string to be written on the screen and its proper length
func (t StatusFail) Prompt() (string, int) {
	t.content = os.ExpandEnv(t.content)
	return colorStringANSI(t.content, t.fgcolor, t.bgcolor, t.font), RealLen(t.content)
}
