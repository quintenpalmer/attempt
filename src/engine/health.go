package engine

type Health struct {
    CurHealth uint
    MaxHealth uint
}

func MakeHealth(max uint) *Health {
    return &Health {
        max,
        max,
    }
}

func (health *Health) IsDead() bool {
    return health.CurHealth == 0
}

func (health *Health) DoDamage(damage uint) {
    health.CurHealth = SubtractNoWrap(health.CurHealth, damage)
}

func (health *Health) Heal(amt uint) {
    health.CurHealth = Min(health.CurHealth + amt, health.MaxHealth)
}
