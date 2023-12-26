package main

import (
)

type Templates struct {
    Options
}

type Options struct{}


// programming Languages
const (
    Js string = Yellow + "JavaScript" + Reset 
    Ts string = Blue + "TypeScript" + Reset
    Golang string = Cyan + "Go" + Reset
    Rust string = Orange256 + "Rust" + Reset
)

//FrameWork for Go
const (
    RawDawgGo string = Blue256 + "RawDawg" + Reset
)

func NewTemplates() *Templates{
    return &Templates{}
}

func(o *Options) ProgrammingLangWithDesc() map[string]string{
    return map[string]string{
        Js : "[ ask to J_BLOW ]",
        Ts : "[ ts ignore || ask to DHH ]",
        Golang :"[ mid !C ]",
        Rust : "[ C/C++ have better community/Foundation ]" ,

    }
}

func(o *Options) GOFrameWork() map[string]string{
    return map[string]string{
        RawDawgGo  : "[ setup  using standard library ]",    
    }
}















