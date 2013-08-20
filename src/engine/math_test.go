package engine

import (
    "testing"
    "unittest"
)

func TestMax(t *testing.T) {
    unittest.CheckEqual(t, Max(3, 5), uint(5))
    unittest.CheckEqual(t, Max(5, 3), uint(5))
    unittest.CheckEqual(t, Max(5, 5), uint(5))
}

func TestMin(t *testing.T) {
    unittest.CheckEqual(t, Min(3, 5), uint(3))
    unittest.CheckEqual(t, Min(5, 3), uint(3))
    unittest.CheckEqual(t, Min(5, 5), uint(5))
}

func TestSubNoWrap(t *testing.T) {
    unittest.CheckEqual(t, SubtractNoWrap(5, 7), uint(0))
    unittest.CheckEqual(t, SubtractNoWrap(7, 5), uint(2))
}
