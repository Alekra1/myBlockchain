package main

import (
	"fmt"
	"os"

	"github.com/lukesampson/figlet/figletlib"
)

func Print(message string) {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Couldn't get working directory lol")
		return
	}
	fontsDirectory := cwd + "/fonts"
	f, err := figletlib.GetFontByName(fontsDirectory, "smslant")
	if err != nil {
		fmt.Println("Could not find that font!")
		return
	}
	figletlib.PrintMsg(message, f, 80, f.Settings(), "left")
}
