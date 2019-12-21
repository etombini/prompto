package prompter

//TextIfGit handles configuration to provide printable prompt information
type TextIfGit struct {
	PartCommon
	content string `yaml:"content"`
}

//NewTextIfGit returns a TextIfGit struct
func NewTextIfGit(param map[string]string) TextIfGit {
	return TextIfGit{
		PartCommon: NewPartCommon(param),
		content:    param["content"],
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
		PartCommon: t.PartCommon,
		content:    t.content,
	}
	return text.Prompt()
}
