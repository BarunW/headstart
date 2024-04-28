package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/BarunW/headstart/assert"
	"gopkg.in/ini.v1"
)

func createTempDir(t *testing.T ) string{
    d, err := os.MkdirTemp("", "testDir")
    if err != nil{
        assert.Assert(err.Error(), "createTempDir()", true)
        t.Fail()
    }
    return d
}

func createTempFile(t *testing.T) *os.File{
    f, err := os.CreateTemp("", "testFile")
    if err != nil{
        assert.Assert(err.Error(), "createTempFile()", true)
        t.Fail()
    }
    return f
}

func TestDirectory(t *testing.T){
    d := createTempDir(t)
    f := createTempFile(t)
    defer func(){
        os.RemoveAll(d)
        os.Remove(f.Name())
    }()

    err := isFolder(d)
    if err != nil{
        fmt.Println(err)
        t.Fail()
    }
    
    // check the file is dir or not
    err = isFolder(f.Name()) 
    if err == nil{
        assert.Assert("File should not be a dir", "FAILED at TestDirectory()", true)
        t.Fail()
    }
}

func TestIsFile(t *testing.T){
    d := createTempDir(t)
    f := createTempFile(t)
    defer func(){
        os.RemoveAll(d)
        os.Remove(f.Name())
    }()

    err := isFile(f.Name())
    if err != nil{
        assert.Assert(err.Error(), "Failed At TestIsFIle()", true)
        t.Failed()
    }
 
    // check the file is dir or not
    err = isFile(d) 
    if err == nil{
        assert.Assert("dir should not be a file", "TestIsFile()", true)
        t.Fail()
    }
}


func checkTheLink( t *testing.T, key, expectedValue string){
    wd, err := os.Getwd()
    if err != nil{
        fmt.Println("Unable to get current working dir", err.Error())
        t.Fail()
    }
    configPath := filepath.Join(wd, "bin", CONFIG_FILE)
    
    cfg, err := ini.Load(configPath)
    if err != nil{
        assert.Assert(err.Error(), "checkTheLink()", true)
        t.Fail()
    }
    commandSection := cfg.Section(SECTION_COMMANDS)
    if !commandSection.HasKey(key){ 
        assert.Assert("Command Key not found", "checkTheLink()", true)
        t.Fail()
    }
    value := commandSection.Key(key).Value() 
    if value == "" || value != expectedValue{
        assert.Assert("Expected Value doesn't match with value in config", "checkTheLink()", true)
        t.Fail()
    } 


//    cfg.Section(SECTION_COMMANDS).DeleteKey(key)
//    
//    // delete the key 
//    cmdSectionKey = cfg.Section(SECTION_COMMANDS).Key(key)
//    value = cmdSectionKey.Value()
//    if value != ""{
//        assert.Assert("Failed to delete the key", "checkTheLink()", true)
//        t.Fail()
//    }     
}

func TestHandleLinkCommand(t *testing.T){
    hs := NewHSCommands() 
    cmdKey := "gomod"
    file := "go.mod"
    hs.processCommandWithSubcmd(LINK, file, cmdKey)    
    
//    expectedValue, err := filepath.Abs(file)
//    fmt.Println(expectedValue)
//    if err != nil{
//        assert.Assert(err.Error(), "TestHanleLinkCommand()->expectedValue", true)
//        t.Fail()
//    }

    //checkTheLink(t, cmdKey, expectedValue)
}


