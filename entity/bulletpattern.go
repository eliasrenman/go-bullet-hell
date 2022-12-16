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
	if time.Since(pattern.lastShootTime).Seconds() < pattern.Shoot.Cooldown {
		return
	}
	// Check if the pattern should stop This should probably be moved to another file.
	if time.Since(pattern.startTime).Seconds() > pattern.Duration {
		pattern.isDone = true
		return
	}
	switch pattern.Shoot.Pattern {
	case "circle":
		pattern.spawnCircleBullet()

	}
	pattern.spawnCount++
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

const (
	radius = 3
)

func (pattern *Pattern) spawnCircleBullet() {
	var step = 2. * math.Pi / float64(pattern.Shoot.Count)
	for i := 0.; int(i) < pattern.Shoot.Count; i++ {

		pattern.owner.Shoot(pattern.owner.Position, step*i, radius, 20)
	}

	pattern.lastShootTime = time.Now()
}
