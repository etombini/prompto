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

//GetIndex returns the index of the Prompter
func (t Text) GetIndex() int {
	return t.Index
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
	prompt := ""
	if fgcolor, ok := ForegroundColor16[t.Fgcolor]; ok {
		prompt += fgcolor
	}
	if bgcolor, ok := BackgroundColor16[t.Bgcolor]; ok {
		prompt += bgcolor
	}
	if font, ok := Font[t.Font]; ok {
		prompt += font
	}

	prompt += t.Content

	return prompt, RealLen(t.Content), nil
}
