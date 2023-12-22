package main

import (
	"fmt"
	"testing"
    "log"
)

func TestAnsciiRemoveRegex(t *testing.T){
    s := Red + "Aloha PHP" + Reset
    rs := RemoveANSI(s)
    if "Aloha PHP" == rs{
        fmt.Println("True")    
        return
    }
    log.Fatal("Error")
}
