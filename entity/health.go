package entity

type Health struct {
	maxHitpoint     int
	currentHitpoint int
}

func NewHealth(maxHitpoint int) *Health {
	return &Health{
		maxHitpoint:     maxHitpoint,
		currentHitpoint: maxHitpoint,
	}
}
func (h *Health) TakeDamage(bullet *Bullet) {
	// Make sure that the currentHitpoint is not less than 0
	if (h.currentHitpoint - bullet.Damage) <= 0 {
		h.currentHitpoint -= bullet.Damage
		return
	}
	h.currentHitpoint = 0

}
