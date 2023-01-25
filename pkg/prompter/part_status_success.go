package prompter

import "os"

// StatusSuccess handles configuration to provide printable prompt information
type StatusSuccess struct {
	PartCommon
	content string `yaml:"content"`
}

// NewStatusSuccess returns a StatusSuccess struct.
// The "content" key in the map expands environment variables when Prompt() is called. It is not done at
// declaration time to allow other components to benefit from this.
func NewStatusSuccess(param map[string]string) StatusSuccess {
	status := os.ExpandEnv("${STATUS}")
	if status != "0" {
		return StatusSuccess{}
	}
	return StatusSuccess{
		PartCommon: NewPartCommon(param),
		content:    param["content"],
	}
}

// Side returns the side of the Prompter
func (t StatusSuccess) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

// Kind return the kind of part
func (t StatusSuccess) Kind() string {
	return "text"
}

// IsNewline tells if this part requires a newline to be inserted
func (t StatusSuccess) IsNewline() bool {
	return false
}

// Prompt return the string to be written on the screen and its proper length
func (t StatusSuccess) Prompt() (string, int) {
	t.content = os.ExpandEnv(t.content)
	return colorStringANSI(t.content, t.fgcolor, t.bgcolor, t.font), RealLen(t.content)
}
