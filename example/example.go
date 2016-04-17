package main

import (
	"fmt"
	"strings"

	"github.com/go-zero-boilerplate/osvisitors"

	"github.com/golang-devops/easyscripts"
)

func main() {
	echoTextVariable := easyscripts.NewVariable("ECHO_TEXT", "hello world")

	builder := easyscripts.NewBuilder().
		Variables(echoTextVariable).
		LineBuilder(easyscripts.NewLineBuilder().
			Raw("echo").
			Variable(echoTextVariable))

	windowsLines := builder.BuildLines(`c:\`, osvisitors.WindowsOs)
	fmt.Println(strings.Join(windowsLines, "\n"))
	/*
	   Running this will print out these lines:

	   @echo off
	   SET ECHO_TEXT=hello world
	   echo %ECHO_TEXT%
	*/

	linuxLines := builder.BuildLines(`c:\`, osvisitors.LinuxOs)
	fmt.Println(strings.Join(linuxLines, "\n"))
	/*
		Running this will print out these lines:

		#!/bin/bash
		ECHO_TEXT=hello world
		echo $ECHO_TEXT
	*/
}
