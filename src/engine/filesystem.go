package engine

import (
    "errors"
    "os"
    "io"
    "path/filepath"
    "fmt"
)

type FileSystem struct {
    rootPath string
}

var fileSystem = MakeFileSystem("")

func MakeFileSystem(path string) *FileSystem {
    cwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }
    return &FileSystem{
        rootPath: filepath.Join(cwd, path),
    }
}

func (fs *FileSystem) Load(path string, obj GameReader) error {
    fullPath := filepath.Join(fs.rootPath, path)
    file, err := os.Open(fullPath)
    defer file.Close()
    if err != nil {
        return err
    }
    objSize, err := file.Stat()
    if err != nil {
        return err
    }
    bytes := make([]byte, objSize.Size())
    read, err := file.Read(bytes)
    if read != len(bytes) {
        return errors.New("Read incorrect number of bytes")
    }
    if err != nil && err != io.EOF {
        return err
    }
    fmt.Printf("%s", bytes)
    err = obj.UnmarshalGame(bytes)
    if err != nil {
        return err
    }
    return nil
}

func (fs *FileSystem) Save(path string, obj GameWriter) error {
    fullPath := filepath.Join(fs.rootPath, path)
    file, err := os.Create(fullPath)
    defer file.Close()
    if err != nil {
        return err
    }
    bytes := obj.MarshalGame()
    _, err = file.Write(bytes)
    if err != nil {
        return err
    }
    return nil
}

func (fs *FileSystem) LoadPlayer(path string) (*Player, error) {
    player := new(Player)
    err := fs.Load(path, player)
    return player, err
}

func (fs *FileSystem) LoadGameChunk(path string) (*GameChunk, error) {
    chunk := new(GameChunk)
    err := fs.Load(path, chunk)
    return chunk, err
}
