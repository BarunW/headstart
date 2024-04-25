package main

import(
    "testing"
    "os"
    "fmt"
)

func createTempDir() string{
    d, err := os.MkdirTemp("", "testDir")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    return d
}

func createTempFile() *os.File{
    f, err := os.CreateTemp("", "testFile")
    if err != nil{
        fmt.Println(err)
        os.Exit(1)
    }
    return f
}

func TestDirectory(t *testing.T){
    d := createTempDir()
    f := createTempFile()
    defer func(){
        os.RemoveAll(d)
        os.Remove(f.Name())
    }()

    err := isFolder(d)
    if err != nil{
        t.Fatal(err)
    }
    
    // check the file is dir or not
    err = isFolder(f.Name()) 
    if err == nil{
        t.Fatal(fmt.Errorf("%s", "File should not be a dir"))
    }
}

func TestIsFile(t *testing.T){
    d := createTempDir()
    f := createTempFile()
    defer func(){
        os.RemoveAll(d)
        os.Remove(f.Name())
    }()

    err := isFile(f.Name())
    if err != nil{
        t.Fatal(err)
    }
 
    // check the file is dir or not
    err = isFile(d) 
    if err == nil{
        t.Fatal(fmt.Errorf("%s", "dir should not be a file"))
    }
}


