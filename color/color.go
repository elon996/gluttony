package color

import (
	"github.com/fatih/color"
)
var Red = color.New(color.BgRed).PrintlnFunc()
var Green = color.New(color.FgHiGreen).PrintlnFunc()
var Cyan = color.New(color.FgHiCyan).PrintlnFunc()
var Magenta = color.New(color.FgHiMagenta).PrintlnFunc()
var Yellow = color.New(color.FgHiYellow).PrintlnFunc()
var Blue = color.New(color.FgHiBlue).PrintlnFunc()

var Redf = color.New(color.BgRed).PrintfFunc()
var Greenf = color.New(color.FgHiGreen).PrintfFunc()
var Cyanf = color.New(color.FgHiCyan).PrintfFunc()
var Magentaf = color.New(color.FgHiMagenta).PrintfFunc()
var Yellowf = color.New(color.FgHiYellow).PrintfFunc()
var Bluef = color.New(color.FgHiBlue).PrintfFunc()

