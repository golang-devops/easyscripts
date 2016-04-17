package easyscripts

import (
	"fmt"

	"github.com/go-zero-boilerplate/osvisitors"
)

type Builder interface {
	RawLine(rawLine string) Builder
	Variables(variables ...Variable) Builder
	LineBuilder(lineBuilder LineBuilder) Builder

	BuildLines(baseDir string, osType osvisitors.OsType) []string
}

func NewBuilder() Builder {
	return &builder{}
}

type builder struct {
	variables    []Variable
	lineBuilders []LineBuilder
}

func (b *builder) RawLine(rawLine string) Builder {
	return b.LineBuilder(NewLineBuilder().Raw(rawLine))
}

func (b *builder) Variables(variables ...Variable) Builder {
	b.variables = append(b.variables, variables...)
	return b
}

func (b *builder) LineBuilder(lineBuilder LineBuilder) Builder {
	b.lineBuilders = append(b.lineBuilders, lineBuilder)
	return b
}

func (b *builder) BuildLines(baseDir string, osType osvisitors.OsType) []string {
	ctx := &scriptContext{
		BaseDir: baseDir,
		OsType:  osType,
	}

	visitor := &buildLinesVisitor{ctx: ctx, v: b.variables, l: b.lineBuilders}
	osType.Accept(visitor)
	return visitor.answer
}

type buildLinesVisitor struct {
	ctx *scriptContext
	v   []Variable
	l   []LineBuilder

	answer []string
}

func (b *buildLinesVisitor) VisitWindows() {
	b.answer = append(b.answer, "@echo off")
	for _, v := range b.v {
		b.answer = append(b.answer, fmt.Sprintf(`SET %s=%s`, v.Name(), v.Value()))
	}
	for _, l := range b.l {
		b.answer = append(b.answer, l.build(b.ctx))
	}
}

func (b *buildLinesVisitor) VisitLinux() {
	b.answer = append(b.answer, "#!/bin/bash")
	for _, v := range b.v {
		b.answer = append(b.answer, fmt.Sprintf(`%s=%s`, v.Name(), v.Value()))
	}
	for _, l := range b.l {
		b.answer = append(b.answer, l.build(b.ctx))
	}
}

func (b *buildLinesVisitor) VisitDarwin() {
	b.VisitLinux()
}
