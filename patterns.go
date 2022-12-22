package main

import (
	"time"
)

type Pattern struct {
	Type    string
	Options map[string]interface{}

	Duration time.Duration
	Cooldown time.Duration
}

func (pattern *Pattern) Start(entity *Entity) {
	switch pattern.Type {
	case "staggeredCircle":
		go pattern.staggeredCirclePattern(entity)

	case "arc":
		go pattern.arcPattern(entity)
	}
}

func (p *Pattern) arcPattern(entity *Entity) {
	count := p.Options["count"].(int)
	speed := p.Options["speed"].(float64)
	from := p.Options["from"].(float64)
	to := p.Options["to"].(float64)

	for i := 0; i < count; i++ {
		angle := from + (to-from)/float64(count)*float64(i)
		entity.Shoot(entity.Position, angle, speed, 0)
	}
}

func (p *Pattern) staggeredCirclePattern(entity *Entity) {
	count := p.Options["count"].(int)
	speed := p.Options["speed"].(float64)
	from := p.Options["from"].(float64)
	to := p.Options["to"].(float64)

	for i := 0; i < count; i++ {
		angle := from + (to-from)/float64(count)*float64(i)
		entity.Shoot(entity.Position, angle, speed, 0)

		time.Sleep(p.Duration / time.Duration(count))
	}
}

type Schedule struct {
	index        int
	lastPlayTime time.Time
	Patterns     []Pattern
}

func (schedule *Schedule) Update(entity *Entity) {
	lastPattern := schedule.Patterns[(schedule.index+len(schedule.Patterns)-1)%len(schedule.Patterns)]

	if time.Since(schedule.lastPlayTime) > lastPattern.Duration+lastPattern.Cooldown {
		schedule.Patterns[schedule.index].Start(entity)
		schedule.index = (schedule.index + 1) % len(schedule.Patterns)
		schedule.lastPlayTime = time.Now()
	}
}
