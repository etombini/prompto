package prompter

import (
	"fmt"
	"strconv"
)

func colorStringANSI(s string, fg string, bg string, font string) string {
	return ansiForegroundColor(fg) + ansiBackgroundColor(bg) + s + ansiFont["reset"]
}

func ansiForegroundColor(code string) string {
	c, err := strconv.Atoi(code)
	if err == nil && c >= 0 && c < 256 {
		return fmt.Sprintf("\\[\\e[38;5;%dm\\]", c)
	}
	if color, ok := ansiForegroundColor16[code]; ok {
		return color
	}
	return ""
}

func ansiBackgroundColor(code string) string {
	c, err := strconv.Atoi(code)
	if err == nil && c >= 0 && c < 256 {
		return fmt.Sprintf("\\[\\e[48;5;%sm\\]", code)
	}
	if color, ok := ansiBackgroundColor16[code]; ok {
		return color
	}
	return ""
}

//ansiForegroundColor16 is a map associating a 16 color name with its ANSI escape sequence
var ansiForegroundColor16 = map[string]string{
	"default":       "\\[\\e[39m\\]",
	"black":         "\\[\\e[30m\\]",
	"red":           "\\[\\e[31m\\]",
	"green":         "\\[\\e[32m\\]",
	"yellow":        "\\[\\e[33m\\]",
	"blue":          "\\[\\e[34m\\]",
	"magenta":       "\\[\\e[35m\\]",
	"cyan":          "\\[\\e[36m\\]",
	"light-gray":    "\\[\\e[37m\\]",
	"dark-gray":     "\\[\\e[90m\\]",
	"light-red":     "\\[\\e[91m\\]",
	"light-green":   "\\[\\e[92m\\]",
	"light-yellow":  "\\[\\e[93m\\]",
	"light-blue":    "\\[\\e[94m\\]",
	"light-magenta": "\\[\\e[95m\\]",
	"light-cyan":    "\\[\\e[96m\\]",
	"white":         "\\[\\e[97m\\]",
}

//ansiBackgroundColor16 is a map associating a 16 color name with its ANSI escape sequence
var ansiBackgroundColor16 = map[string]string{
	"default":       "\\[\\e[49m\\]",
	"black":         "\\[\\e[40m\\]",
	"red":           "\\[\\e[41m\\]",
	"green":         "\\[\\e[42m\\]",
	"yellow":        "\\[\\e[43m\\]",
	"blue":          "\\[\\e[44m\\]",
	"magenta":       "\\[\\e[45m\\]",
	"cyan":          "\\[\\e[46m\\]",
	"light-gray":    "\\[\\e[47m\\]",
	"dark-gray":     "\\[\\e[100m\\]",
	"light-red":     "\\[\\e[101m\\]",
	"light-green":   "\\[\\e[102m\\]",
	"light-yellow":  "\\[\\e[103m\\]",
	"light-blue":    "\\[\\e[104m\\]",
	"light-magenta": "\\[\\e[105m\\]",
	"light-cyan":    "\\[\\e[106m\\]",
	"white":         "\\[\\e[107m\\]",
}

var ansiReset = "\\[\\e[0m\\]"

//ansiFont is a map associating font format with its ANSI escape sequence to
var ansiFont = map[string]string{
	"default":   "\\[\\e[22m\\]",
	"reset":     "\\[\\e[0m\\]",
	"bold":      "\\[\\e[1m\\]",
	"low":       "\\[\\e[2m\\]",
	"underline": "\\[\\e[4m\\]",
	"blink":     "\\[\\e[5m\\]",
}
