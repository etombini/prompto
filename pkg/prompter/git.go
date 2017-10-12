package prompter

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/src-d/go-git.v4/plumbing"

	git "gopkg.in/src-d/go-git.v4"
)

//Git handles configuration to provide printable prompt information
type Git struct {
	Index int    `yaml:"index"`
	Side  string `yaml:"side"`

	Before  string `yaml:"before"`
	After   string `yaml:"after"`
	Fgcolor string `yaml:"fgcolor"`
	Bgcolor string `yaml:"bgcolor"`
	Font    string `yaml:"font"`

	BranchBefore  string `yaml:"branch_before"`
	BranchAfter   string `yaml:"branch_after"`
	BranchFgcolor string `yaml:"branch_fgcolor"`
	BranchBgcolor string `yaml:"branch_bgcolor"`

	TagBefore  string `yaml:"tag_before"`
	TagAfter   string `yaml:"tag_after"`
	TagFgcolor string `yaml:"tag_fgcolor"`
	TagBgcolor string `yaml:"tag_bgcolor"`

	StatusCleanText    string `yaml:"status_clean_text"`
	StatusCleanFgColor string `yaml:"status_clean_fgcolor"`
	StatusCleanBgColor string `yaml:"status_clean_bgcolor"`

	StatusDirtyText    string `yaml:"status_dirty_text"`
	StatusDirtyFgColor string `yaml:"status_dirty_fgcolor"`
	StatusDirtyBgColor string `yaml:"status_dirty_bgcolor"`

	// ModifiedBefore string `yaml:"modified_before"`
	// ModifiedAfter  string `yaml:"modified_after"`
	// // ModifiedFgColor string `yaml:"modified_fgcolor"`
	// // ModifiedBgColor string `yaml:"modified_bgcolor"`

	// StagedBefore string `yaml:"staged_before"`
	// StagedAfter  string `yaml:"staged_after"`
	// // StagedFgColor string `yaml:"staged_fgcolor"`
	// StagedBgColor string `yaml:"staged_bgcolor"`
}

var gitInfo = make(map[string]string)

func rootDir() string {
	if dir, ok := gitInfo["root"]; ok {
		return dir
	}
	path, err := os.Getwd()
	if err != nil {
		gitInfo["root"] = ""
		return gitInfo["root"]
	}

	found := false

	for !found && len(path) > 0 {
		dirInfo, err := os.Stat(filepath.Join(path, "/.git"))
		if err == nil && dirInfo.IsDir() {
			found = true
			break
		}
		splitted := strings.Split(path, "/")
		splitted = splitted[0 : len(splitted)-2]
		path = strings.Join(splitted, "/")
	}
	if !found {
		gitInfo["root"] = ""
		return gitInfo["root"]
	}
	gitInfo["root"] = path
	return gitInfo["root"]
}

func isGit() bool {
	if rootDir() != "" {
		return true
	}
	return false
}

func branch() string {
	if !isGit() {
		return ""
	}

	path := rootDir()

	headFile := filepath.Join(path, ".git/HEAD")
	headBuf, err := ioutil.ReadFile(headFile)
	if err != nil {
		return ""
	}
	head := string(headBuf)

	if !strings.HasPrefix(head, "ref: refs/heads/") {
		return string(head[:10])
	}

	branch := strings.TrimPrefix(head, "ref: refs/heads/")
	if strings.HasSuffix(branch, "\n") {
		branch = strings.TrimSuffix(branch, "\n")
	}

	return branch
}

func isBranch() bool {
	if branch() != "" {
		return true
	}
	return false
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

func isClean() bool {
	if !isGit() {
		return true
	}

	path := rootDir()
	repo, err := git.PlainOpen(path)
	if err != nil {
		return true
	}

	wt, err := repo.Worktree()
	if err != nil {
		return true
	}

	s, err := wt.Status()
	if err != nil {
		return true
	}

	return s.IsClean()
}

//GetSide returns the side of the Prompter
func (g Git) GetSide() string {
	if g.Side == "" || g.Side != "right" {
		return "left"
	}
	return "right"
}

func (g Git) branchPrompt() (string, int, error) {
	if !isGit() {
		return "", 0, nil
	}

	fgColor := g.BranchFgcolor
	bgColor := g.BranchBgcolor

	if fgColor == "" {
		fgColor = g.Fgcolor
	}
	if bgColor == "" {
		bgColor = g.Bgcolor
	}
	b := ""
	if b = branch(); b != "" {
		b = g.BranchBefore + b + g.BranchAfter
	}

	return colorString(b, fgColor, bgColor, g.Font), RealLen(b), nil
}

func (g Git) tagPrompt() (string, int, error) {
	if !isGit() {
		return "", 0, nil
	}

	fgColor := g.TagFgcolor
	bgColor := g.TagBgcolor

	if fgColor == "" {
		fgColor = g.Fgcolor
	}
	if bgColor == "" {
		bgColor = g.Bgcolor
	}
	t := ""
	if t = tag(); t != "" {
		t = g.TagBefore + t + g.TagAfter
	}

	return colorString(t, fgColor, bgColor, g.Font), RealLen(t), nil
}

func (g Git) cleanPrompt() (string, int, error) {
	if !isGit() {
		return "", 0, nil
	}

	clean := g.StatusCleanText
	fgColor := g.StatusCleanFgColor
	bgColor := g.StatusCleanBgColor
	if !isClean() {
		clean = g.StatusDirtyText
		fgColor = g.StatusDirtyFgColor
		bgColor = g.StatusDirtyBgColor
	}

	if fgColor == "" {
		fgColor = g.Fgcolor
	}
	if bgColor == "" {
		bgColor = g.Bgcolor
	}

	return colorString(clean, fgColor, bgColor, g.Font), RealLen(clean), nil
}

//Prompt return the resulting string and its real length when written
func (g Git) Prompt() (string, int, error) {
	if !isGit() {
		return "", 0, nil
	}

	before := colorString(g.Before, g.Fgcolor, g.Bgcolor, g.Font)
	after := colorString(g.After, g.Fgcolor, g.Bgcolor, g.Font)

	bp, bl, err := g.branchPrompt()
	if err != nil {
		bp = ""
		bl = 0
	}

	cp, cl, err := g.cleanPrompt()
	if err != nil {
		cp = ""
		cl = 0
	}

	tp, tl, err := g.tagPrompt()
	if err != nil {
		tp = ""
		tl = 0
	}

	return before + bp + cp + tp + after, RealLen(g.Before) + bl + cl + tl + RealLen(g.After), nil
}
