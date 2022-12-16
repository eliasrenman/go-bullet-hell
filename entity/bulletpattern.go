package entity

import (
	"math"
	"time"
)

type Pattern struct {
	Shoot ShootingPattern `json:"shoot"`
	Move  MovePattern     `json:"move"`

	Cooldown float64 `json:"cooldown"`
	Duration float64 `json:"duration"`
	paused   bool

	lastShootTime time.Time
	startTime     time.Time
	owner         *Entity
	spawnCount    int
	isDone        bool
}

func (pattern *Pattern) Start(entity *Entity) {
	pattern.owner = entity
	pattern.startTime = time.Now()
}

// func (pattern *Pattern) Cancel() {}
// func (pattern *Pattern) Pause()  {}
// func (pattern *Pattern) Play()   {}

func (pattern *Pattern) Update() {
	// if cooldown has passed.
	if pattern.canShoot() {
		return
	}
	// Check if the pattern should stop This should probably be moved to another file.
	// if time.Since(pattern.startTime).Seconds() > pattern.Duration {
	// 	pattern.isDone = true
	// 	return
	// }
	pattern.SpawnBullets()
	pattern.spawnCount++
}

func (pattern *Pattern) canShoot() bool {
	return time.Since(pattern.lastShootTime).Seconds() < pattern.Shoot.Cooldown
}

type ShootingPattern struct {
	BulletSpeed int `json:"bulletSpeed"`
	// Something like circle, aim, static, split.
	Pattern    string  `json:"pattern"`
	BulletType string  `json:"bulletType"`
	Cooldown   float64 `json:"cooldown"`
	Count      int     `json:"count"`
}

type MovePattern struct {
}

var CirclePattern = Pattern{
	Shoot: ShootingPattern{
		BulletSpeed: 1,
		Cooldown:    0.5,
		Pattern:     "circle",
		BulletType:  "regular",
		Count:       50,
	},
	Move:     MovePattern{},
	Cooldown: 5,
	Duration: 10,
}

var StaggeredCirclePatternInstance = StaggeredCirclePattern{
	Pattern: Pattern{
		Shoot: ShootingPattern{
			BulletSpeed: 2,
			Cooldown:    0.2,
			Pattern:     "circle",
			BulletType:  "regular",
			Count:       10,
		},
		Move:     MovePattern{},
		Cooldown: 5,
		Duration: 10,
	},
	bulletOffset:      0,
	BulletOffsetLimit: 16,
	Reverse:           true,
}

const (
	radius = 3
)

func (pattern *Pattern) SpawnBullets() {
	var step = 2. * math.Pi / float64(pattern.Shoot.Count)
	for i := 0.; int(i) < pattern.Shoot.Count; i++ {

		pattern.owner.Shoot(pattern.owner.Position, step*i, radius, 20)
	}

	pattern.lastShootTime = time.Now()
}

type StaggeredCirclePattern struct {
	Pattern
	bulletOffset      int
	BulletOffsetLimit int
	Reverse           bool
}

func (pattern *StaggeredCirclePattern) Update() {
	// if cooldown has passed.
	if pattern.canShoot() {
		return
	}
	pattern.SpawnBullets()
	pattern.updateOffset()
	pattern.spawnCount++
}

func (pattern *StaggeredCirclePattern) updateOffset() {
	if pattern.Reverse {
		if pattern.bulletOffset != 0 {
			pattern.bulletOffset--
		} else {
			pattern.bulletOffset = pattern.BulletOffsetLimit
		}

	} else {
		if pattern.bulletOffset != pattern.BulletOffsetLimit {
			pattern.bulletOffset++
		} else {
			pattern.bulletOffset = 0

		}
	}

}

func (pattern *StaggeredCirclePattern) SpawnBullets() {
	count := float64(pattern.Shoot.Count * pattern.BulletOffsetLimit)
	var step = 2. * math.Pi / count

	for i := 0.; i < count; i++ {
		if math.Mod(i, float64(pattern.bulletOffset)) == 1 {
			pattern.owner.Shoot(pattern.owner.Position, step*(i), radius, 20)
		}
	}

	pattern.lastShootTime = time.Now()
}
