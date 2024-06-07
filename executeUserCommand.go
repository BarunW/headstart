package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
)

type ExecuteCommand struct {
	Data
}

func NewExecuteCommand(d Data) ExecuteCommand {
	return ExecuteCommand{d}
}

func (ow ExecuteCommand) OpenTextEditor(userInput string, cmdKey string) {
    var cmd *exec.Cmd
	fileOrDirPath := ow.Data.GetValFromCommandsSection(cmdKey)
	if fileOrDirPath == "" {
		slog.Error("There is no file or dir link to this command key", "[Command-Key]", cmdKey)
		os.Exit(1)
	}
    
    ext := filepath.Ext(fileOrDirPath)
    if userInput == "exec" && ext == ".sh" || ext == "" { 
        fmt.Println(userInput)
        cmd = exec.Command(fileOrDirPath) 
    } else {
        cmd = exec.Command(userInput, fileOrDirPath)
    }
	// Set the standard input, output, and error streams
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		slog.Error("Unable to open the file", "DETAILS", err.Error())
		os.Exit(1)
	}

	if err := cmd.Wait(); err != nil {
		slog.Error("Unable to wait the file", "DETAILS", err.Error())
		os.Exit(1)
	}

}
