package prompter

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

//Command handles configuration to provide printable prompt information
type Command struct {
	PartCommon
	command string
	timeout uint16
}

//NewCommand return a new Command struct
func NewCommand(param map[string]string) Command {
	timeout, err := strconv.ParseUint(param["timeout"], 10, 16)
	if err != nil || timeout == 0 {
		// return Command{
		// 	PartCommon: NewPartCommon(param),
		// 	command:    param["command"],
		// 	timeout:    10,
		// }
		timeout = 10
	}

	return Command{
		PartCommon: NewPartCommon(param),
		command:    param["command"],
		timeout:    uint16(timeout),
	}
}

//Side returns the side of the Prompter
func (c Command) Side() string {
	if c.side == "" || c.side != "right" {
		return "left"
	}
	return "right"
}

//Kind return the kind of part
func (c Command) Kind() string {
	return "command"
}

//IsNewline tells if this part requires a newline to be inserted
func (c Command) IsNewline() bool {
	return false
}

//Prompt return the resulting string and its real length when written
func (c Command) Prompt() (string, int) {

	expandedCmd := os.ExpandEnv(c.command)

	cmd := exec.Command("/bin/sh", "-c", expandedCmd)

	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()

	var res string
	if err := cmd.Start(); err != nil {
		res = c.before + "ERR_EXEC: " + err.Error() + c.after
		return colorStringANSI(res, c.fgcolor, c.bgcolor, c.font), RealLen(res)
	}

	done := make(chan error)
	go func() {
		done <- cmd.Wait()
	}()

	select {
	case err := <-done:
		if err != nil {
			res = c.before + "ERR: " + stderr.String() + c.after
			break
		}
		res = c.before + stdout.String() + c.after
	case <-time.After(time.Duration(c.timeout) * time.Millisecond):
		res = c.before + "TIMEOUT" + c.after
	}

	res = strings.TrimSpace(res)
	return colorStringANSI(res, c.fgcolor, c.bgcolor, c.font), RealLen(res)
}
