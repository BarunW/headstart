package main

import (
	"fmt"
	"os"
)


func main() {
    args := os.Args
    hs := NewHSCommands()   
    fmt.Println(hs.configFilePath)
    if len(os.Args) >= 3{
        hs.processCommandWithSubcmd(Flag(args[1]), args[2:]...) 
    }
    fmt.Println(args[0])
}
