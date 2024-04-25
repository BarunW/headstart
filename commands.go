package main

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
    "github.com/BarunW/headstart/assert"
)

type Flag string
type SubCommand string
type FlagsAndDescrip map[Flag]string
type SubCommandsAndDescrip  map[SubCommand]string
type FileType string

const(
    LINK Flag = "link"
    GEN  Flag = "gen"
)

// first arg for subcommand
const(
    FILE FileType = "file"
    DIR  FileType = "dir"
)

const CONFIG_FILE string = "config.ini"


type UserCommand map[SubCommand]string
type HSCommands struct{
    configFilePath string
}

func(hs *HSCommands) createCONFIG_FILE(){
    fh := NewFileHandler() 
    exPath, err := fh.GetExecutablePath()
    if err != nil{
        os.Exit(1)
    }
    exPath = path.Join(filepath.Dir(exPath),CONFIG_FILE)
    var f = func(error){
        if os.IsNotExist(err){ 
            cfg := ini.Empty()
            if err := cfg.SaveTo(exPath); err != nil{
                slog.Error("Unable to setup", "DETAILS", err.Error())
                os.Exit(1)
            }
            hs.configFilePath = path.Join(exPath, CONFIG_FILE)
            return 
        }
        slog.Error("Failed to open the file", "DETAILS", err.Error())
        os.Exit(1)

    }
    _, err = ini.Load(exPath)
    if err != nil{
        f(err)
    }

    fmt.Println(exPath)
    hs.configFilePath = exPath

}

func NewHSCommands() HSCommands{ 
    hs := HSCommands{}
    hs.createCONFIG_FILE()
    return hs
}

func checkLength(expectedLength , theLength int) bool{    
    if  theLength != expectedLength   {
        fmt.Printf("%d arguments expected", expectedLength)
        return false
    }
    return true
}

func isFile(filePath string) error{
    fs, err := os.Stat(filePath)
    if err != nil{
        slog.Error("Please provide a file", filePath, err.Error())
        return err
    }
    isFile := fs.Mode().IsRegular()
    if !isFile{
        return fmt.Errorf("This is not a file")
    }
    return nil
}

func isFolder(dirPath string) error{
    fs, err := os.Stat(dirPath)
    if err != nil{
        slog.Error("Please provide a file", dirPath, err.Error())
        return err
    }
    _isDir := fs.IsDir()
    if !_isDir {
        return fmt.Errorf("This is not a dir")
    }
    return nil
}


// this function produce a side effect that exit on err of failing 
// certain condition
func(hs HSCommands)handleLinkCommand( fType FileType, subcmd... string){
    var(
        cfg *ini.File
        err error
        section *ini.Section
        linkSection *ini.Section
    )
    
    fmt.Println("CONFIG PATH", hs.configFilePath)
    cfg, err = ini.Load(hs.configFilePath)
    if err != nil {
        slog.Error("Unable to load the file", "DETAILS", err.Error())
    }

    // Create the neccessary section in the config file
    section = cfg.Section("Commands")
    linkSection = cfg.Section("LINK_TYPE")
   
    // key that associate with a link file
    cmdKey := subcmd[1] 
    
    // check the input key is already exist or !exist
    // if exist throw error
    if section.Key(cmdKey).String() != ""{
        slog.Error("This command already link", "Fix:", "User different command")
        os.Exit(1)
    }
    
    // !exist link with the input command key 
    section.NewKey(cmdKey, subcmd[0])
    linkSection.NewKey(cmdKey, string(fType))
     
    if err := cfg.SaveTo(hs.configFilePath); err != nil{
        slog.Error("Failed to save the file", "DETAILS", err.Error())
        os.Exit(1)
    }
    return
}

func(hs HSCommands) handlGenCommand(subcmd... string){
    fh := NewFileHandler()

    cfg, err := ini.Load(hs.configFilePath)
    if err != nil{
        slog.Error("Failed to Load config file", "DETAILS", err.Error())
        os.Exit(1)
    }
    
    // cmdkey is command key that is linked  in the config file
    cmdKey := subcmd[0]   
    key := cfg.Section("Commands").Key(cmdKey) 
    if key.Value() == ""{
        slog.Error("Unable to find the command", "DETAILS", "Please link command to use gen")
        os.Exit(1)
    }
    
    ext := path.Ext(key.Value()) 
    writeToFile := subcmd[1]+ext
    

    f2W, err := fh.OpenDezzFile(writeToFile)
    if err != nil{
        slog.Error("Unable to create the input file", "DETAILS", err.Error())
        os.Exit(1)
    }

    f2R, err := fh.OpenDezzReadFile(key.Value())
    if err != nil{
        slog.Error("Unable to get the file that links to the command", "DETAILS", err.Error())
        os.Exit(1)
    }


    defer func() {
        if err := f2W.Close(); err != nil{
            panic(err)
        }

        if err := f2R.Close(); err != nil{
            panic(err)
        }
    }()

    n, err := io.Copy(f2W, f2R) 
    fmt.Println(n)
    if n == 0 || err != nil{
        slog.Error("Failed to write the file", "DETAILS", err.Error())
        os.Exit(1)
    }
 
}

func(hc HSCommands) processCommandWithSubcmd( cmd Flag , subcmd... string){
    switch cmd{
    case LINK:
        if isValid := checkLength(2, len(subcmd)); !isValid{
            assert.Assert("There should be 2 subcommands", "user error", false)
            return
        } 
        if err := isFile(subcmd[0]); err == nil{
            hc.handleLinkCommand(FILE, subcmd...)
        } else if dErr := isFolder(subcmd[0]); dErr == nil{
            hc.handleLinkCommand(DIR, subcmd...)
        }
    case GEN:
        if isValid := checkLength(2, len(subcmd)); !isValid{
            assert.Assert("There should be 3 subcommands", "user error", false)
            return
        } 
        hc.handlGenCommand(subcmd...)
    }

}

func(hc *HSCommands) CommandsForDisplay() FlagsAndDescrip { 
    var c FlagsAndDescrip = make(FlagsAndDescrip)

    // add the command in the lookup table
    c[LINK] = "link with either module/dir/folder or file"  
    c[GEN]  = "do code gen" 

    return c
}

func(hs *HSCommands) SubCommandsForDisplay () SubCommandsAndDescrip { 
    var s SubCommandsAndDescrip = make(SubCommandsAndDescrip)
 
    // add the command in the lookup table
        
    return s 
}

