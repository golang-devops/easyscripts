package easyscripts

import "strings"

type LineBuilder interface {
	Arg(args ...string) LineBuilder
	Raw(raws ...string) LineBuilder
	Variable(variables ...Variable) LineBuilder

	build(ctx *scriptContext) string
}

func NewLineBuilder() LineBuilder {
	return &lineBuilder{}
}

type lineBuilder struct {
	parts []linePart
}

func (l *lineBuilder) addPart(part linePart) LineBuilder {
	l.parts = append(l.parts, part)
	return l
}

func (l *lineBuilder) Arg(args ...string) LineBuilder {
	for _, a := range args {
		l.addPart(linePartArg(a))
	}
	return l
}

func (l *lineBuilder) Raw(raws ...string) LineBuilder {
	for _, r := range raws {
		l.addPart(linePartRaw(r))
	}
	return l
}

func (l *lineBuilder) Variable(variables ...Variable) LineBuilder {
	for _, v := range variables {
		l.addPart(linePartVariable{v})
	}
	return l
}

func (l *lineBuilder) build(ctx *scriptContext) string {
	strs := []string{}
	for _, p := range l.parts {
		strs = append(strs, p.String(ctx))
	}
	return strings.Join(strs, " ")
}
