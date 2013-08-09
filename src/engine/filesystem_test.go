package engine

import (
    "testing"
    "unittest"
    "vector"
)

func TestPlayerSaveLoad(t *testing.T) {
    p := MakePlayer(1, "test-player", vector.Vector2{-1, 1})
    p.Token = "waffle"
    err := fileSystem.Save("testpath", p)
    unittest.CheckNil(t, &err)
    p2, err := fileSystem.LoadPlayer("testpath")
    unittest.CheckNil(t, &err)
    unittest.CheckDeepEqual(t, *p, *p2)
}
