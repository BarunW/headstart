package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

type OpenWithTextEditor struct{
   Data 
}

func NewOpenWithTexEditor(d Data) OpenWithTextEditor{
   return OpenWithTextEditor{d} 
}

func(ow OpenWithTextEditor) Open(textEditor string, cmdKey string){
    fileOrDirPath := ow.Data.GetValFromCommandsSection(cmdKey) 
    if fileOrDirPath == ""{
        slog.Error("There is no file or dir link to this command key", "[Command-Key]", cmdKey)
        os.Exit(1)
    }
    cmd := exec.Command(textEditor, fileOrDirPath)
    // Set the standard input, output, and error streams
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    if err := cmd.Start(); err != nil {
        slog.Error("Unable to start the text editor", "DETAILS", err.Error())
        os.Exit(1)
    }

    if err := cmd.Wait(); err != nil {
        slog.Error("Unable to wait the text editor", "DETAILS", err.Error())
        os.Exit(1)
    }
    fmt.Println("DONE")
}
