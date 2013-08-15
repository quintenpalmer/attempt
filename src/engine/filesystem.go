package engine

import (
    "cgl.tideland.biz/applog"
    "errors"
    "os"
    "io"
    "path/filepath"
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
    fs := &FileSystem{
        rootPath: filepath.Join(cwd, path),
    }
    if res, _ := fs.FileExists("players"); !res {
        err := fs.Mkdir("players")
        // Panic on these errors since there's not much point in recoving
        // from an improperly setup filesystem
        if err != nil { panic(err) }
    }
    if res, _ := fs.FileExists("world"); !res {
        err := fs.Mkdir("world")
        // Panic on these errors since there's not much point in recoving
        // from an improperly setup filesystem
        if err != nil { panic(err) }
    }
    applog.Debugf("Finished setting up directories")
    return fs
}

func (fs *FileSystem) FileExists(path string) (bool, error) {
    fullPath := filepath.Join(fs.rootPath, path)
    _, err := os.Stat(fullPath)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func (fs *FileSystem) Mkdir(path string) error {
    fullPath := filepath.Join(fs.rootPath, path)
    err := os.Mkdir(fullPath, 0777) // full file permissions
    return err
}

func (fs *FileSystem) Load(path string, obj GameReader) error {
    fullPath := filepath.Join(fs.rootPath, path)
    file, err := os.Open(fullPath)
    applog.Debugf("Loading from path: %s", fullPath)
    if err != nil {
        return err
    }
    defer file.Close()
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
    err = obj.UnmarshalGame(bytes)
    if err != nil {
        return err
    }
    return nil
}

func (fs *FileSystem) Save(path string, obj GameWriter) error {
    fullPath := filepath.Join(fs.rootPath, path)
    file, err := os.Create(fullPath)
    applog.Debugf("Saving to path: %s", fullPath)
    if err != nil {
        return err
    }
    defer file.Close()
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
