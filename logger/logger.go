package inthemiddle

import (
    "github.com/fatih/color"
    "github.com/kr/pretty"
)

func Info(msg interface{}) {
    write("Info", color.FgGreen, msg)
}

func Warning(msg interface{}) {
    write("Warning", color.FgYellow, msg)
}

func Error(msg interface{}) {
    write("Error", color.FgRed, msg)
}

func Debug(msg interface{}) {
    write("Debug", color.FgMagenta, msg)
}

func write(ev string, c color.Attribute, msg interface{}) {
    cFunc := color.New(c).SprintFunc()
    pretty.Printf("%s: %# v\n", cFunc(ev), msg)
}
