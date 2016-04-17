package easyscripts

import (
	"path/filepath"

	"strings"
)

type linePartRelativeArg string

func (l linePartRelativeArg) String(ctx *scriptContext) string {
	s := filepath.Join(ctx.BaseDir, string(l))
	if strings.Contains(s, " ") {
		return `"` + s + `"`
	}
	return s
}
