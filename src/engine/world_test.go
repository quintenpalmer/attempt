package engine

import (
    "testing"
    "unittest"
)

var (
    w = MakeWorld(10, 10)
)

func TextGetNextId(t *testing.T) {
    unittest.CheckEqual(t, w.getNextId(), uint(0))
    unittest.CheckEqual(t, w.getNextId(), uint(1))
    unittest.CheckEqual(t, w.getNextId(), uint(2))
    unittest.CheckEqual(t, w.getNextId(), uint(3))
}
