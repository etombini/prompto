package prompter

//TextIfNotGit handles configuration to provide printable prompt information
type TextIfNotGit struct {
	PartCommon
	content string
}

//NewTextIfNotGit returns a TextIfGit struct
func NewTextIfNotGit(param map[string]string) TextIfNotGit {
	return TextIfNotGit{
		PartCommon: NewPartCommon(param),
		content:    param["content"],
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
		PartCommon: t.PartCommon,
		content:    t.content,
	}
	return text.Prompt()
}
