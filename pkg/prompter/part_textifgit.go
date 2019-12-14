package prompter

//TextIfGit handles configuration to provide printable prompt information
type TextIfGit struct {
	content string `yaml:"content"`
	side    string `yaml:"side"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewTextIfGit returns a TextIfGit struct
func NewTextIfGit(param map[string]string) TextIfGit {
	return TextIfGit{
		content: param["content"],
		side:    param["side"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
	}
}

//Side returns the side of the Prompter
func (t TextIfGit) Side() string {
	if t.side == "" || t.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (t TextIfGit) Kind() string {
	return "ifgit"
}

//IsNewline tells if this part requires a newline to be inserted
func (t TextIfGit) IsNewline() bool {
	return false
}

//Prompt return the string to be written on the screen and its proper length
func (t TextIfGit) Prompt() (string, int) {
	if !isGit() {
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
