package engine

import (
    "os"
    "testing"
    "unittest"
    "vector"
)

func TestPlayerSaveLoad(t *testing.T) {
    p := MakePlayer(1, "test-player", vector.Vector2{-1, 1})
    p.Token = "waffle"
    defer func() {
        os.Remove("testpath")
    }()
    err := fileSystem.Save("testpath", p)
    unittest.CheckNil(t, &err)
    p2, err := fileSystem.LoadPlayer("testpath")
    unittest.CheckNil(t, &err)
    unittest.CheckDeepEqual(t, *p, *p2)
}

func TestFileExists(t *testing.T) {
    exists, _ := fileSystem.FileExists("testpath")
    unittest.CheckFalse(t, exists)

    p := MakePlayer(1, "test-player", vector.Vector2{-1, 1})
    p.Token = "waffle"
    defer func() {
        os.Remove("testpath")
    }()
    err := fileSystem.Save("testpath", p)
    unittest.CheckNil(t, &err)
    exists, _ = fileSystem.FileExists("testpath")
    unittest.Check(t, exists)
}

func TestMkdir(t *testing.T) {
    exists, _ := fileSystem.FileExists("testdir")
    unittest.CheckFalse(t, exists)

    err := fileSystem.Mkdir("testdir")
    unittest.CheckNil(t, &err)
    defer func() {
        os.Remove("testdir")
    }()
    exists, _ = fileSystem.FileExists("testdir")
    unittest.Check(t, exists)
}
