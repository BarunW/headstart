package main

import (
	"io"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type FileHandler struct{}

func NewFileHandler() FileHandler {
	return FileHandler{}
}

func (F *FileHandler) OpenDezzFile(filePath string) (io.ReadWriteCloser, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		slog.Error("[ERROR]: Cannot Open the file", "[DETAILS]: ", err.Error())
		return nil, err
	}
	return file, err
}

func (F *FileHandler) OpenDezzReadFile(filePath string) (io.ReadCloser, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		slog.Error("[ERROR]: Cannot Open the file", "[DETAILS]: ", err.Error())
		return nil, err
	}
	return file, err
}

func (F *FileHandler) GetExecutablePath() (string, error) {
	exPath, err := os.Executable()
	if err != nil {
		slog.Error("Unable to get executable path", "Details", err.Error())
		return "", err
	}

	return exPath, nil
}

func (F *FileHandler) OpenConfigFile(configPath string) (*ini.File, error) {
	return ini.Load(configPath)
}

func (F *FileHandler) GetBashScriptPath() string {
	fh := NewFileHandler()
	exPath, err := fh.GetExecutablePath()
	if err != nil {
		slog.Error("Executable path not found", "Details", err.Error())
		os.Exit(1)
	}
	exPath = path.Join(filepath.Dir(exPath), "cdbash.sh")
	return exPath
}
