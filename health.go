package main

type Health struct {
	MaxHitPoints int
	HitPoints    int
}

func (health *Health) TakeDamage(bullet *Bullet) {
	health.HitPoints -= bullet.Damage
	if health.HitPoints < 0 {
		health.HitPoints = 0
	}
	if health.HitPoints > health.MaxHitPoints {
		health.HitPoints = health.MaxHitPoints
	}

}
