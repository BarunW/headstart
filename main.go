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
    
    if len(os.Args) == 2{
        hs.LinkOutPut(os.Args[1])        
        return
    }

	hs.ShowAllCommands()
}
