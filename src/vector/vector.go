package vector

type Vector2 struct {
    X int
    Y int
}

// TODO: Look into handling overflows.

func AddVector(a, b Vector2) Vector2 {
    return Vector2{a.X + b.X, a.Y + b.Y}
}

func ScalarMulVector(a Vector2, m int) Vector2 {
    return MulVector(a, Vector2{m, m})
}

func SubVector(a, b Vector2) Vector2 {
    return AddVector(a, ScalarMulVector(b, -1))
}

func MulVector(a, b Vector2) Vector2 {
    return Vector2{a.X * b.X, a.Y * b.Y}
}

func DivVector(a, b Vector2) Vector2 {
    return Vector2{a.X / b.X, a.Y / b.Y}
}

func DotProduct(a, b Vector2) int {
    return SumVector(MulVector(a, b))
}

func SumVector(a Vector2) int {
    return a.X + a.Y
}

func (v *Vector2) mutateCall(x Vector2, f func(Vector2, Vector2) Vector2) {
    newVector := f(*v, x)
    v.X = newVector.X
    v.Y = newVector.Y
}

func (v *Vector2) Add(x Vector2) {
    v.mutateCall(x, AddVector)
}

func (v *Vector2) Sub(x Vector2) {
    v.mutateCall(x, SubVector)
}

func (v *Vector2) Mul(x Vector2) {
    v.mutateCall(x, MulVector)
}

func (v *Vector2) Div(x Vector2) {
    v.mutateCall(x, DivVector)
}

func (v *Vector2) Values() (int, int) {
    return v.X, v.Y
}
