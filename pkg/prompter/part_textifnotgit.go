package prompter

//TextIfNotGit handles configuration to provide printable prompt information
type TextIfNotGit struct {
	content string `yaml:"content"`
	side    string `yaml:"side"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewTextIfNotGit returns a TextIfGit struct
func NewTextIfNotGit(param map[string]string) TextIfNotGit {
	return TextIfNotGit{
		content: param["content"],
		side:    param["side"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
	}
}

//Side returns the side of the Prompter
func (t TextIfNotGit) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (t TextIfNotGit) Kind() string {
	return "ifgit"
}

//IsNewline tells if this part requires a newline to be inserted
func (t TextIfNotGit) IsNewline() bool {
	return false
}

//Prompt return the string to be written on the screen and its proper length
func (t TextIfNotGit) Prompt() (string, int) {
	if isGit() {
		return "", 0
	}
	text := Text{
		content: t.content,
		side:    t.side,
		bgcolor: t.bgcolor,
		fgcolor: t.fgcolor,
		font:    t.font,
	}
	return text.Prompt()
}
