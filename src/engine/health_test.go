package engine

import (
    "testing"
    "unittest"
)

func TestHealthIsDead(t *testing.T) {
    unittest.Check(t, MakeHealth(0).IsDead())
    unittest.CheckFalse(t, MakeHealth(5).IsDead())
}

func TestHealthDoDamage(t *testing.T) {
    h := MakeHealth(10)
    h.DoDamage(5)
    unittest.CheckEqual(t, h.CurHealth, uint(5))
    h.DoDamage(7)
    unittest.CheckEqual(t, h.CurHealth, uint(0))
    unittest.Check(t, h.IsDead())
}

func TestHealthHeal(t *testing.T) {
    h := MakeHealth(10)
    unittest.CheckEqual(t, h.CurHealth, uint(10))
    h.Heal(10)
    unittest.CheckEqual(t, h.CurHealth, uint(10))
    h.DoDamage(10)
    unittest.CheckEqual(t, h.CurHealth, uint(0))
    unittest.Check(t, h.IsDead())
    h.Heal(5)
    unittest.CheckEqual(t, h.CurHealth, uint(5))
}
