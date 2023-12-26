package main

import (
	"flag"
	"fmt"
	"os"
)

func main(){
    hs := NewHeadStart()
    name := "" 

    flag.Usage  = func()  {
        os.Stderr.Write([]byte(hs.Usage))
    }

    l := len(os.Args)
    var usrCommand string
    
    if l > 2{
        usrCommand = os.Args[1]
        name = os.Args[2]
        
    }

    if _, ok := hs.Commands[usrCommand]; !ok || l <= 2{
        fmt.Printf("%s", name)
        flag.Usage()
        return
    }
    
    pLangs := hs.Temp.ProgrammingLangWithDesc()    
    
    nR := NewRender()
    slectedLang := nR.RenderOptions(pLangs)

    proceed := NewProceed()
    proceed.ShowFrameWorkOpt(slectedLang, name)
}



