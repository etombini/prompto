package prompter

//RealLen returns the real length of a string when written on screen
func RealLen(s string) int {
	count := 0
	for range s {
		count++
	}
	return count
}
