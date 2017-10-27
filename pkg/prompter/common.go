package prompter

var env = make(map[string]string)

type colorStringFunc func(s string, fg string, bg string, font string) string

var colorString colorStringFunc

//SetShell must be used to set the shell we are dealing with and use the appropriate color functions
func SetShell(shell string) {
	if shell == "bash" {
		colorString = colorStringBash
		return
	}
	if shell == "zsh" {
		colorString = colorStringBlank
		return
	}
	if shell == "fish" {
		colorString = colorStringBlank
		return
	}
	colorString = colorStringBlank
	return
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
