package easyscripts

import "strings"

type linePartVariable struct {
	variable Variable
}

func (l linePartVariable) String(ctx *scriptContext) string {
	visitor := &variableStringVisitor{v: l.variable}
	ctx.OsType.Accept(visitor)
	return visitor.answer
}

type variableStringVisitor struct {
	v      Variable
	answer string
}

func (b *variableStringVisitor) VisitWindows() {
	if strings.Contains(b.v.Name(), " ") {
		panic("Windows variables cannot contain spaces")
	}
	b.answer = "%" + b.v.Name() + "%"
}

func (b *variableStringVisitor) VisitLinux() {
	if strings.Contains(b.v.Name(), " ") {
		panic("Linux variables cannot contain spaces")
	}
	b.answer = "$" + b.v.Name()
}

func (b *variableStringVisitor) VisitDarwin() {
	b.VisitLinux()
}
