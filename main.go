package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func main() {
    args := os.Args
    hs := NewHSCommands()   
    if len(os.Args) >= 3{
        hs.processCommandWithSubcmd(Flag(args[1]), args[2:]...) 
    }
    wd, err := os.Getwd()
    if err != nil{
        fmt.Println(err)
    }
    fmt.Println(filepath.Base(os.Args[1]), os.Args[1])
    fmt.Println(wd)
    fmt.Println(args[0])
}
