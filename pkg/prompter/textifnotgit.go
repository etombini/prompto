package prompter

//TextIfNotGit handles configuration to provide printable prompt information
type TextIfNotGit struct {
	Content string `yaml:"content"`
	Index   int    `yaml:"index"`
	Side    string `yaml:"side"`
	Bgcolor string `yaml:"bgcolor"`
	Fgcolor string `yaml:"fgcolor"`
	Font    string `yaml:"font"`
}

//GetIndex returns the index of the Prompter
func (t TextIfNotGit) GetIndex() int {
	return t.Index
}

//GetSide returns the side of the Prompter
func (t TextIfNotGit) GetSide() string {
	if t.Side == "" || t.Side != "right" {
		return "left"
	}
	return "right"
}

//Prompt return the string to be written on the screen and its proper length
func (t TextIfNotGit) Prompt() (string, int, error) {
	if isGit() {
		return "", 0, nil
	}
	text := Text{t.Content, t.Index, t.Side, t.Bgcolor, t.Fgcolor, t.Font}
	return text.Prompt()
}
