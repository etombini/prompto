package prompter

//PartCommon is the common datastructure to be set up when constructing a
//prompter part
type PartCommon struct {
	side    string `yaml:"side"`
	before  string `yaml:"before"`
	after   string `yaml:"after"`
	bgcolor string `yaml:"bgcolor"`
	fgcolor string `yaml:"fgcolor"`
	font    string `yaml:"font"`
}

//NewPartCommon returns a PartCommon
func NewPartCommon(param map[string]string) PartCommon {
	return PartCommon{
		side:    param["side"],
		before:  param["before"],
		after:   param["after"],
		bgcolor: param["bgcolor"],
		fgcolor: param["fgcolor"],
		font:    param["font"],
	}
}
