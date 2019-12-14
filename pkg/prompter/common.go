package prompter

var env = make(map[string]string)

type colorStringFunc func(s string, fg string, bg string, font string) string

var colorString colorStringFunc

var shellFuncs = map[string]colorStringFunc{
	"bash": colorStringBash,
	"zsh":  colorStringBlank,
	"fish": colorStringBlank,
}

//SetShell must be used to set the shell we are dealing with and use the appropriate color functions
func SetShell(shell string) {
	f, ok := shellFuncs[shell]
	if ok {
		colorString = f
		return
	}
	colorString = colorStringBlank
}

func colorStringBlank(s string, fg string, bg string, font string) string {
	return s
}

//RealLen returns the real length of a string when written on screen
func RealLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
