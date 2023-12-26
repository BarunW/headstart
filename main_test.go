package main

import (
	"fmt"
	"testing"
    "os"
)

func TestAnsiRemoveRegex(t *testing.T){
    s := Red + "Aloha PHP" + Reset
    rs := RemoveANSI(s)
    if "Aloha PHP" == rs{
        fmt.Println("True")    
        return
    }
    fmt.Println("Failed to parse")
    os.Exit(1)
}



