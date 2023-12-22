package main

import (
)

type Templates struct {
    Options
}

type Options struct{}


const (
    js string = Yellow + "JavaScript" + Reset 
    ts string = Blue + "TypeScript" + Reset
    golang string = Cyan + "GO" + Reset
    rust string = Orange256 + "RUST" + Reset
)

func NewTemplates() *Templates{
    return &Templates{}
}

 //func (o *Options) ProgrammingLang() []string{
 //    // args(color, name, reset)    
 //    return []string{js, ts, golang, rust}
 //}

func(o *Options) ProgrammingLangWithDesc() map[string]string{
    return map[string]string{
        js : "[ ask to J_BLOW ]",
        ts : "[ ts ignore || ask to DHH ]",
        golang :"[ mid !C ]",
        rust : "[ C/C++ have better community/Foundation ]" ,
    }
}















