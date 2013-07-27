package vector

import (
    "testing"
    "unittest"
)

var (
    zero = Vector2{0, 0}
    one = Vector2{1, 1}
    two = Vector2{2, 2}
    three_four = Vector2{3, 4}
    six_eight = Vector2{6, 8}
    nine_sixteen = Vector2{9, 16}
)

// Test for CheckEqual/NotEqual -- Should go in another file eventually.
func TestEqual(t *testing.T) {
    unittest.CheckEqual(t, zero, zero)
    unittest.CheckEqual(t, three_four, three_four)
    unittest.CheckNotEqual(t, zero, three_four)
}

func TestAddVector(t *testing.T) {
    unittest.CheckEqual(t, AddVector(zero, zero), zero)
    unittest.CheckEqual(t, AddVector(zero, three_four), three_four)
    unittest.CheckEqual(t, AddVector(three_four, three_four), six_eight)
}

func TestSubVector(t *testing.T) {
    unittest.CheckEqual(t, SubVector(zero, zero), zero)
    unittest.CheckEqual(t, SubVector(six_eight, three_four), three_four)
    unittest.CheckEqual(t, SubVector(three_four, three_four), zero)
}

func TestScalarMulVector(t *testing.T) {
    unittest.CheckEqual(t, ScalarMulVector(three_four, 1), three_four)
    unittest.CheckEqual(t, ScalarMulVector(three_four, 0), zero)
    unittest.CheckEqual(t, ScalarMulVector(three_four, 2), six_eight)
}

func TestMulVector(t *testing.T) {
    unittest.CheckEqual(t, MulVector(three_four, zero), zero)
    unittest.CheckEqual(t, MulVector(three_four, three_four), nine_sixteen)
}

func TestDivVector(t *testing.T) {
    unittest.CheckEqual(t, DivVector(zero, three_four), zero)
    unittest.CheckEqual(t, DivVector(three_four, three_four), one)
    unittest.CheckEqual(t, DivVector(six_eight, three_four), two)
}

func TestSumVector(t *testing.T) {
    unittest.CheckEqual(t, SumVector(zero), 0)
    unittest.CheckEqual(t, SumVector(three_four), 7)
    unittest.CheckEqual(t, SumVector(six_eight), 14)
}

func TestVectorAdd(t *testing.T) {
    v := Vector2{0, 1}
    v.Add(Vector2{3, 5})
    unittest.CheckEqual(t, v, Vector2{3, 6})
}

func TestVectorSub(t *testing.T) {
    v := Vector2{5, 5}
    v.Sub(Vector2{3, 7})
    unittest.CheckEqual(t, v, Vector2{2, -2})
}

func TestVectorMul(t *testing.T) {
    v := Vector2{5, 5}
    v.Mul(v)
    unittest.CheckEqual(t, v, Vector2{25, 25})
}

func TestVectorDiv(t *testing.T) {
    v := Vector2{10, 10}
    v.Div(Vector2{5, 5})
    unittest.CheckEqual(t, v, Vector2{2, 2})
}