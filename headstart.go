package main

import(
    "fmt"
)

type HeadStart struct {
    Name
    Commands
    Usage string

    Temp *Templates
}

type Name string
type Commands map[string]string


func NewHeadStart() *HeadStart{

    temp := NewTemplates()
    hs := &HeadStart{}
    
    // get the commands
    c := hs.commands()
   
    // assgined the fields
    hs.Temp = temp
    hs.Name = "hs"
    hs.Commands = c 
    hs.setUsage()
    
    
    return hs
}

func(hs *HeadStart) commands () Commands { 
    var c Commands = make(Commands)
    
    // add the command in the lookup table
    c["create-api"] = Cyan + "\tcreate-api" + Reset + " --" +"Generate the api code-base\n"
    c["create-auth"] = Yellow + "\tcreate-api" + Reset + " --" + "Generate the api code-base\n"
    
    return c 
}

func(hs *HeadStart) setUsage(){
    avilableCommands := ""
    for _, value := range hs.Commands{
        avilableCommands += value 
    }

    s := fmt.Sprintf("Usage : %s hs <command> \ncommands: \n%s",White, avilableCommands)
    hs.Usage = s
}


