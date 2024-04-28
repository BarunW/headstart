package main

import (
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path"
	"path/filepath"
	"sync"
	"time"

	"github.com/BarunW/headstart/assert"
	"gopkg.in/ini.v1"
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
// ALL sections
const(
    SECTION_COMMANDS = "Commands"
    SECTION_TYPES    = "LINK_TYPE"
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

//    fmt.Println(exPath)
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
        commandSection *ini.Section
        linkSection *ini.Section
        linkPath string
    )
    linkPath, err = filepath.Abs(subcmd[0])
    if err != nil{
        slog.Error("Unable to convert to abs path", "Details", err.Error()) 
        os.Exit(1) 
    }

//    fmt.Println("LINK PATH", linkPath)
    
    cfg, err = ini.Load(hs.configFilePath)
    if err != nil {
        slog.Error("Unable to load the file", "DETAILS", err.Error())
        os.Exit(1)
    }

    // Create the neccessary section in the config file
    commandSection = cfg.Section(SECTION_COMMANDS)
    linkSection = cfg.Section(SECTION_TYPES)
   
    // key that associate with a link file
    cmdKey := subcmd[1] 
    
    // check the input key is already exist or !exist
    // if exist throw error
    if commandSection.HasKey(cmdKey){
        slog.Error("This key is already link")
        os.Exit(1)
    }

    // !exist link with the input command key 
    commandSection.NewKey(cmdKey, linkPath)
    linkSection.NewKey(cmdKey, string(fType))
     
    if err := cfg.SaveTo(hs.configFilePath); err != nil{
        slog.Error("Failed to save the file", "DETAILS", err.Error())
        os.Exit(1)
    }
    fmt.Println("Sucessfully Linked")
    return
}

// this function have a side effect of panic if file aren't able to close
func handleFileGeneration(desPath string, srcPath string) error{
    fh := NewFileHandler()

    f2W, err := fh.OpenDezzFile(desPath)
    if err != nil{
        slog.Error("Unable to create the input file", "DETAILS", err.Error())
        return err
    }

    f2R, err := fh.OpenDezzReadFile(srcPath)
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

    _, err = io.Copy(f2W, f2R) 
    if  err != nil{
        slog.Error("Failed to write the file", "DETAILS", err.Error())
        return err
    }
    
    return nil
}

func handleDirGeneration(srcPath string, destPath string) error{
    now := time.Now()
    fileSystem := os.DirFS(srcPath)
    wg := sync.WaitGroup{}

    var HandleFileFn = func(src fs.File, des os.File){
        defer func(){ src.Close(); des.Close() }()
        _, err := io.Copy(&des, src)
        if err != nil{
            wg.Done()
            return
        }  
        wg.Done()
        
    }
  
    err := fs.WalkDir(fileSystem, ".", func(fpath string, d fs.DirEntry, err error) error {
        if d.IsDir(){
            if err := os.Mkdir(path.Join(destPath, fpath), os.ModePerm); err != nil{
                slog.Error("Unable to create the dir", "DETAILS", err.Error())
                return err
            }

        }
        if d.Type().IsRegular(){
            src,  err := fileSystem.Open(fpath) 
            if err != nil{ 
                slog.Error("Unable to open the source file", "DETAILS", err.Error())
                return err
            }
             
            des, err := os.Create(path.Join(destPath, fpath))
            if err != nil{
                slog.Error("Unable to create the destination file", "DETAILS", err.Error())
                return err
            }
            wg.Add(1)
            go HandleFileFn(src, *des) 
        }
        return nil
    })

    wg.Wait()
    if err != nil{
        slog.Error("Unable to execute the command", "DETATILS", err.Error())
        return err
    }
    fmt.Println(time.Since(now))
    return nil
}

func(hs HSCommands) handlGenCommand(subcmd... string){
    cfg, err := ini.Load(hs.configFilePath)
    if err != nil{
        slog.Error("Failed to Load config file", "DETAILS", err.Error())
        os.Exit(1)
    }
    
    // cmdkey is command key that is linked  in the config file 
    cmdKey := subcmd[0]   
    commandSectionKey := cfg.Section(SECTION_COMMANDS).Key(cmdKey) 
    typeSectionKey := cfg.Section(SECTION_TYPES).Key(cmdKey) 
    if commandSectionKey.Value() == ""{ 
        slog.Error("Unable to find the command", "DETAILS", "Please link command to use gen")
        os.Exit(1)
    }
    
    switch FileType(typeSectionKey.Value()){
        case FILE:
            // strip out the extension of the file
            ext := path.Ext(commandSectionKey.Value()) 
            destPath := subcmd[1]+ext

            srcPath := commandSectionKey.Value() 
            if err := handleFileGeneration(destPath, srcPath); err != nil{
                os.Exit(1)
            }
        case DIR:
            handleDirGeneration(commandSectionKey.Value(), subcmd[1]) 
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


func(hs *HSCommands) CommandsForDisplay() FlagsAndDescrip { 
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

