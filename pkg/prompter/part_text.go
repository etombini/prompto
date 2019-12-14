package prompter

//Text handles configuration to provide printable prompt information
type Text struct {
	content string `yaml:"content"`
	side    string `yaml:"side"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewText returns a Text struct
func NewText(param map[string]string) Text {
	return Text{
		side:    param["side"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
		content: param["content"],
	}
}

//Side returns the side of the Prompter
func (t Text) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (t Text) Kind() string {
	return "text"
}

//IsNewline tells if this part requires a newline to be inserted
func (t Text) IsNewline() bool {
	return false
}

//Prompt return the string to be written on the screen and its proper length
func (t Text) Prompt() (string, int) {
	return colorString(t.content, t.fgcolor, t.bgcolor, t.font), RealLen(t.content)
}
