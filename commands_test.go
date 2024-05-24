package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/BarunW/headstart/assert"
)

func createTempDir(t *testing.T) string {
	d, err := os.MkdirTemp("", "testDir")
	if err != nil {
		assert.Assert(err.Error(), "createTempDir()", true)
		t.Fail()
	}
	return d
}

func createTempFile(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "testFile")
	if err != nil {
		assert.Assert(err.Error(), "createTempFile()", true)
		t.Fail()
	}
	return f
}

func TestDirectory(t *testing.T) {
	d := createTempDir(t)
	f := createTempFile(t)
	defer func() {
		os.RemoveAll(d)
		os.Remove(f.Name())
	}()
    
    // Testing isFolder() -> func
    // 
	err := isFolder(d)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

    // Testing again isFolder ->func 
	// check the file is dir or not
	err = isFolder(f.Name())
	if err == nil {
		assert.Assert("File should not be a dir", "FAILED at TestDirectory()", true)
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	d := createTempDir(t)
	f := createTempFile(t)
	defer func() {
		os.RemoveAll(d)
		os.Remove(f.Name())
	}()
    
    // Tesing isFile() -> func()
	err := isFile(f.Name())
	if err != nil {
		assert.Assert(err.Error(), "Failed At TestIsFIle()", true)
		t.Failed()
	}
    // again isFile() -> func() 
	// to check the file is dir or not
	err = isFile(d)
	if err == nil {
		assert.Assert("dir should not be a file", "TestIsFile()", true)
		t.Fail()
	}
}


func TestHandleLinkCommand(t *testing.T) {
	hs := NewHSCommands("config_test.ini")
	cmdKey := "gomod"
	file := "go.mod"
	hs.processCommandWithSubcmd(LINK, file, cmdKey)

    expectedValue, err := filepath.Abs(file)
    fmt.Println(expectedValue)
    if err != nil{
        assert.Assert(err.Error(), "TestHanleLinkCommand()->expectedValue", true)
        t.Fail()
    }
 
}
