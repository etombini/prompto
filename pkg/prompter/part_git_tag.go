package prompter

import (
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

//GitTag handles configuration to provide printable prompt information
type GitTag struct {
	PartCommon
}

//NewGitTag returns a GitTag struct
func NewGitTag(param map[string]string) GitTag {
	return GitTag{
		PartCommon: NewPartCommon(param),
	}
}

func tag() string {
	if !isGit() {
		return ""
	}
	if t, ok := gitInfo["tag"]; ok {
		return t
	}

	path := rootDir()
	repo, err := git.PlainOpen(path)
	if err != nil {
		gitInfo["tag"] = ""
		return ""
	}
	head, err := repo.Head()
	if err != nil {
		gitInfo["tag"] = ""
		return ""
	}

	tt, _ := repo.Tags()
	tagref := make(map[string]string)

	tt.ForEach(func(p *plumbing.Reference) error {
		to, err := repo.TagObject(p.Hash())
		if err != nil {
			tagref[p.Hash().String()] = strings.TrimPrefix(p.Name().String(), "refs/tags/")
			return nil
		}
		tagref[to.Target.String()] = to.Name
		return nil
	})

	if len(tagref) == 0 {
		gitInfo["tag"] = ""
		return ""
	}

	headCommit, err := repo.CommitObject(head.Hash())
	var theTag string
	var ok bool

	for headCommit.NumParents() > 0 {
		if theTag, ok = tagref[headCommit.Hash.String()]; ok {
			break
		}
		theTag = ""
		headCommit, err = repo.CommitObject(headCommit.ParentHashes[0])
	}
	gitInfo["tag"] = theTag
	return theTag
}

func isTag() bool {
	if tag() != "" {
		return true
	}
	return false
}

//Side returns the side of the Prompter
func (g GitTag) Side() string {
	if g.side == "" || g.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (g GitTag) Kind() string {
	return "gitbranch"
}

//IsNewline tells if this part requires a newline to be inserted
func (g GitTag) IsNewline() bool {
	return false
}

//Prompt return the resulting string and its real length when written
func (g GitTag) Prompt() (string, int) {
	if !isGit() || tag() == "" {
		return "", 0
	}
	t := g.before + tag() + g.after
	return colorStringANSI(t, g.fgcolor, g.bgcolor, g.font), RealLen(t)
}
