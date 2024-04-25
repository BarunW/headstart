package main

import (
	"io"
	"log/slog"
	"os"
)


type FileHandler struct{}

func NewFileHandler() FileHandler{
    return FileHandler{}
}

func(F *FileHandler) OpenDezzFile(filePath string) (io.ReadWriteCloser, error) {
    file, err := os.OpenFile(filePath, os.O_CREATE | os.O_RDWR, 0666)
    if err != nil{
        slog.Error("[ERROR]: Cannot Open the file", "[DETAILS]: ", err.Error())
        return  nil, err
    }
    return  file, err
}

func(F *FileHandler) OpenDezzReadFile(filePath string) (io.ReadCloser, error) {
    file, err := os.OpenFile(filePath,  os.O_RDONLY, os.ModePerm )
    if err != nil{
        slog.Error("[ERROR]: Cannot Open the file", "[DETAILS]: ", err.Error())
        return  nil, err
    }    
    return file, err
}

func(F *FileHandler) GetExecutablePath() (string, error){
    exPath, err := os.Executable() 
    if err != nil{
        slog.Error("Unable to get executable path", "Details", err.Error())
        return "", err
    }

    return exPath, nil 
}




