package engine

func Max(x, y uint) uint {
    if x < y {
        return y
    } else {
        return x
    }
}

func Min(x, y uint) uint {
    if x > y {
        return y
    } else {
        return x
    }
}

func SubtractNoWrap(x, y uint) uint {
    if x < y {
        return 0
    } else {
        return x - y
    }
}
