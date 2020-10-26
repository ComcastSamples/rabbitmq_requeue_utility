package main

import (
	"github.com/fatih/color"
)

const (
	FgBlack   color.Attribute = iota + 30 // 30 (Black)
	FgRed                                 // 31 (Red)
	FgYellow                              // 32 (Green)
	FgMagenta                             // 33 (Yellow)
	FgCyan                                // 34 (Blue)
	FgGreen                               // 34 (Blue)
	FgBlue                                // 35 (Magenta)
	FgWhite                               // 37 (White)
)

var (
	red     = color.New(color.FgRed).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	blue    = color.New(color.FgBlue, color.Bold).SprintFunc()
	white   = color.New(color.FgWhite).SprintFunc()
)
