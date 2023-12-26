package main

import(
    "os"
    "log/slog"
    "fmt"
    "io"
    "io/fs"
    "path/filepath"
)

type Creator struct {}

type createFile func(string) 

func NewCreator() *Creator {
    return &Creator{}
}


func(c *Creator) ProcessRawSetup(lang, frmWrk, projectName string,) error{

    pwd, err := os.Getwd()
    if err != nil && !os.IsExist(err){
        fmt.Println("unable to get working dir")
        return err
    }

    src, err := os.Executable()
    if err != nil {
        slog.Error("unable to get the exec path", "desc", err.Error())
        os.Exit(1)
    }

    srcDir := filepath.Dir(src)

    // Build the source path
    src = filepath.Join(srcDir,"framework", lang, frmWrk)

    // Build the project path
    projectPath := filepath.Join(pwd, projectName)    


    mkDir(projectPath) 
    return walkAndCopy(src, projectPath)
    
}
func mkDir(name string){

    err := os.MkdirAll(name, os.ModePerm)
    if err != nil{
        panic(err)
    }
}

func walkAndCopy(src, dest string) error{
    _, err:= os.Stat(src)
    if err != nil{
        fmt.Println("Unable to get stat for src dir")
        return err
    }

    filesystem := os.DirFS(src)
    
    return fs.WalkDir(filesystem, ".", func(path string, d fs.DirEntry, err error) error {
        if err != nil{
            fmt.Println("error at the filesystem")
            return err 
        }
        
        if d.IsDir(){
            mkDir(filepath.Join(dest,path))
        } else { 
            // open the source file
            srcFile, err := os.OpenFile(filepath.Join(src,path), os.O_RDONLY, os.ModePerm)
            if err != nil {
                fmt.Println("error while  opening src file")
                return err
            }
            defer srcFile.Close()

            // create the file in destination dir
            destFile, err := os.Create(filepath.Join(dest,path))
            if err != nil {
                fmt.Println("error while creating file")
                return err
            }
            defer destFile.Close()

            _, err = io.Copy(destFile, srcFile)
            if err != nil {
                fmt.Println("error while copying the content")
                return err
            }

            destFile.Sync()
        }
        return nil
    })

}



