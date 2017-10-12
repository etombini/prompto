package prompter

//Text handles configuration to provide printable prompt information
type Text struct {
	Content string `yaml:"content"`
	Index   int    `yaml:"index"`
	Side    string `yaml:"side"`
	Bgcolor string `yaml:"bgcolor"`
	Fgcolor string `yaml:"fgcolor"`
	Font    string `yaml:"font"`
}

//GetSide returns the side of the Prompter
func (t Text) GetSide() string {
	if t.Side == "" || t.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the string to be written on the screen and its proper length
func (t Text) Prompt() (string, int, error) {
	return colorString(t.Content, t.Fgcolor, t.Bgcolor, t.Font), RealLen(t.Content), nil
}
