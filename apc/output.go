package apc

import "strings"

type Output struct {
	Foo string
	Bar string
}

func ParseOutput(str string) map[string]string {
	dict := map[string]string{}
	for _, line := range strings.Split(str, "\n") {
		slice := strings.SplitN(line, ":", 2)
		if len(slice) == 2 {
			k := strings.Trim(slice[0], " \t")
			v := strings.Trim(slice[1], " \t")
			dict[k] = v
		}
	}
	return dict
}
