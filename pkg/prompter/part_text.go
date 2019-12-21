package prompter

//Text handles configuration to provide printable prompt information
type Text struct {
	PartCommon
	content string `yaml:"content"`
}

//NewText returns a Text struct
func NewText(param map[string]string) Text {
	return Text{
		PartCommon: NewPartCommon(param),
		content:    param["content"],
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
	return colorStringANSI(t.content, t.fgcolor, t.bgcolor, t.font), RealLen(t.content)
}
