package gogame

type Vector2 struct {
    x int
    y int
}

// TODO: Look into handling overflows.

func AddVector(a, b Vector2) Vector2 {
    return Vector2{a.x + b.x, a.y + b.y}
}

func ScalarMulVector(a Vector2, m int) Vector2 {
    return MulVector(a, Vector2{m, m})
}

func SubVector(a, b Vector2) Vector2 {
    return AddVector(a, ScalarMulVector(b, -1))
}

func MulVector(a, b Vector2) Vector2 {
    return Vector2{a.x * b.x, a.y * b.y}
}

func DivVector(a, b Vector2) Vector2 {
    return Vector2{a.x / b.x, a.y / b.y}
}

func DotProduct(a, b Vector2) int {
    return SumVector(MulVector(a, b))
}

func SumVector(a Vector2) int {
    return a.x + a.y
}