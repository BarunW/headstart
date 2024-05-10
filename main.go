package main

import (
	"os"
)

func main() {
	args := os.Args
	hs := NewHSCommands()
	if len(os.Args) >= 3 {
		hs.processCommandWithSubcmd(Command(args[1]), args[2:]...)
		return
	}
	hs.ShowAllCommands()
}
