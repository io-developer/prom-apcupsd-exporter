package apc

import "strings"

type Output struct {
	Raw    string
	Parsed map[string]string
}

func NewOutput(raw string) *Output {
	return &Output{Raw: raw}
}

func (self *Output) Parse() {
	dict := map[string]string{}
	for _, line := range strings.Split(self.Raw, "\n") {
		slice := strings.SplitN(line, ":", 2)
		if len(slice) == 2 {
			k := strings.Trim(slice[0], " \t")
			v := strings.Trim(slice[1], " \t")
			dict[k] = v
		}
	}
	self.Parsed = dict
}
