package main

import(
    "fmt"
)

type Proceed struct {
   creator *Creator 
}

func NewProceed() *Proceed{
    c := NewCreator()
    return &Proceed{
        creator : c, 
    }
}

func (p *Proceed) ShowFrameWorkOpt(lang, projectName string) {
    r := NewRender() 
    temp := NewTemplates() 
    goMp := temp.GOFrameWork()

    switch lang {
    case Js:
        return 
    case Golang:
        s := r.RenderOptions(goMp)
        p.createProjDir(lang, s, projectName)
    default:
        return
    }
}

func (p *Proceed) createProjDir(lang, selectedFrmWrk, projectName string){
    if selectedFrmWrk == RawDawgGo {
        err := p.creator.ProcessRawSetup("go", "stdLib", projectName)
        if err != nil{
            fmt.Println(err)
            return
        }
    }
}


