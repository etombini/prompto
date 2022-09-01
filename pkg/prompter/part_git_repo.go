package prompter

import "path/filepath"

// GitRepo handles configuration to provide printable prompt information
type GitRepo struct {
	PartCommon
}

// NewGitRepo returns a GitRepo struct
func NewGitRepo(param map[string]string) GitRepo {
	return GitRepo{
		PartCommon: NewPartCommon(param),
	}
}

func repo() string {
	if !isGit() {
		return ""
	}
	return filepath.Base(rootDir()) //os.path.basename(path)
}

// Side returns the side of the Prompter
func (g GitRepo) Side() string {
	if g.side == "" || g.side != "right" {
		return "left"
	}
	return "right"
}

// Kind return the kind of part
func (g GitRepo) Kind() string {
	return "gitbranch"
}

// IsNewline tells if this part requires a newline to be inserted
func (g GitRepo) IsNewline() bool {
	return false
}

// Prompt return the resulting string and its real length when written
func (g GitRepo) Prompt() (string, int) {
	if !isGit() || repo() == "" {
		return "", 0
	}

	b := g.before + repo() + g.after

	return colorStringANSI(b, g.fgcolor, g.bgcolor, g.font), RealLen(b)
}
