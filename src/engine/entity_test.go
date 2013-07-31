package engine

import (
    "vector"
    "testing"
    "unittest"
)

var (
    simpleMover = MoveableEntity{Entity{0, vector.Vector2{5, 5}}}
)

func TestEntityMove(t *testing.T) {
    simpleMover.Move(vector.Vector2{5, 5})
    unittest.CheckEqual(t, simpleMover.Position, vector.Vector2{10, 10})
}
