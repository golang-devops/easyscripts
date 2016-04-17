package easyscripts

import "strings"

type linePartArg string

func (l linePartArg) String(ctx *scriptContext) string {
	s := string(l)
	if strings.Contains(s, " ") {
		return `"` + s + `"`
	}
	return s
}
