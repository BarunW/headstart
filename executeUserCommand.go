package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

type ExecuteCommand struct {
	Data
}

func NewExecuteCommand(d Data) ExecuteCommand {
	return ExecuteCommand{d}
}

func (ow ExecuteCommand) ChangeDir(path, key string) {
	fType := ow.Data.GetTypeFromLinkTypeSection(key)
	if fType == string(DIR) {
		cmd := exec.Command("bash", "-c", "cd /home/dbarun/Desktop && ls")
		// Execute the command
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}
		// Print the output of the command
		fmt.Println(string(output))
	}

}

func (ow ExecuteCommand) OpenTextEditor(textEditor string, cmdKey string) {
	fileOrDirPath := ow.Data.GetValFromCommandsSection(cmdKey)
	if fileOrDirPath == "" {
		slog.Error("There is no file or dir link to this command key", "[Command-Key]", cmdKey)
		os.Exit(1)
	}

	ow.ChangeDir(fileOrDirPath, cmdKey)

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

}
