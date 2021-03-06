package engine

import (
    "vector"
)

/*
This file contains type definitions for the high level entity definitions.
*/

type Entity struct {
    Id uint
    Position vector.Vector2
}

type Moveable interface {
    Move(v vector.Vector2)
}

type MoveableEntity struct {
    Entity
}

func (me *MoveableEntity) Move(v vector.Vector2) {
    me.Position.Add(v)
}

type Named struct {
    Name string
}
