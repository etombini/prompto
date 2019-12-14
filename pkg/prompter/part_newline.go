package prompter

//Newline handles configuration to provide printable prompt information
type Newline struct{}

//NewNewline returns a Newline struct
func NewNewline(param map[string]string) Newline {
	return Newline{}
}

//Side returns the side of the Prompter
func (n Newline) Side() string {
	return "right"
}

//Kind return the kind of part
func (n Newline) Kind() string {
	return "Newline"
}

//IsNewline tells if this part requires a newline to be inserted
func (n Newline) IsNewline() bool {
	return true
}

//Prompt return the resulting string and its real length when written
func (n Newline) Prompt() (string, int) {
	return "\n", 0
}
