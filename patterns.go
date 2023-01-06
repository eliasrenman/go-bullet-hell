package main

import (
	"fmt"
	"math"
	"time"

	"github.com/fogleman/ease"
)

var easings = map[string]func(float64) float64{
	"linear_in":   ease.Linear,
	"linear_out":  ease.Linear,
	"linear":      ease.Linear,
	"quad_in":     ease.InQuad,
	"quad_out":    ease.OutQuad,
	"quad":        ease.InOutQuad,
	"cube_in":     ease.InCubic,
	"cube_out":    ease.OutCubic,
	"cube":        ease.InOutCubic,
	"quart_in":    ease.InQuart,
	"quart_out":   ease.OutQuart,
	"quart":       ease.InOutQuart,
	"quint_in":    ease.InQuint,
	"quint_out":   ease.OutQuint,
	"quint":       ease.InOutQuint,
	"sine_in":     ease.InSine,
	"sine_out":    ease.OutSine,
	"sine":        ease.InOutSine,
	"expo_in":     ease.InExpo,
	"expo_out":    ease.OutExpo,
	"expo":        ease.InOutExpo,
	"circ_in":     ease.InCirc,
	"circ_out":    ease.OutCirc,
	"circ":        ease.InOutCirc,
	"back_in":     ease.InBack,
	"back_out":    ease.OutBack,
	"back":        ease.InOutBack,
	"elastic_in":  ease.InElastic,
	"elastic_out": ease.OutElastic,
	"elastic":     ease.InOutElastic,
	"bounce_in":   ease.InBounce,
	"bounce_out":  ease.OutBounce,
	"bounce":      ease.InOutBounce,
}

// Pattern is a specific behavior an enemy can perform.
// This can be a bullet pattern or a movement pattern, animations, etc.
type Pattern struct {
	Type       string
	Options    map[string]any
	Duration   time.Duration
	Cooldown   time.Duration
	BulletType int
}

// Option returns the value of a pattern option, or a default value if it isn't provided
func (p *Pattern) Option(key string, defaultValue any) any {
	if value, ok := p.Options[key]; ok {
		return value
	}
	return defaultValue
}

// Start starts the pattern
func (p *Pattern) Start(entity *Entity) {
	switch p.Type {
	case "shoot_arc":
		go p.shootArcPattern(entity)
	case "move_to":
		go p.moveToPattern(entity)
	}
}

func (p *Pattern) shootArcPattern(entity *Entity) {
	count := p.Option("count", 1).(int)
	speed := p.Option("speed", 1.0).(float64)
	from := p.Option("from", 0.0).(float64)
	to := p.Option("to", 2*math.Pi).(float64)
	stagger := p.Option("stagger", 0.0).(float64)

	for i := 0; i < count; i++ {
		angle := from + (to-from)/float64(count)*float64(i)
		entity.Shoot(entity.Position, angle, speed, 0, p.BulletType, 1)
		time.Sleep(time.Duration(stagger * float64(time.Second.Nanoseconds())))
	}
}

func (p *Pattern) moveToPattern(entity *Entity) {
	target := p.Option("target", Vector{}).(Vector)
	speed := p.Option("speed", 1.0).(float64)
	easing := p.Option("easing", "linear").(string)

	easeFn := easings[easing]

	if easeFn == nil {
		panic(fmt.Sprintf("Invalid easing function: %s", easing))
	}

	startTime := time.Now()
	start := entity.Position.Copy()
	distance := target.Distance(entity.Position)

	for i := 0.0; i < distance; i = speed * time.Since(startTime).Seconds() {
		entity.Position = start.Plus(target.Minus(start).ScaledBy(easeFn(i / distance)))
		time.Sleep(time.Duration((1.0 / 60) * time.Second.Seconds()))
	}
}

// Schedule is a queue of patterns that will be played in order, looping back to the beginning when it reaches the end
type Schedule struct {
	index        int
	lastPlayTime time.Time
	Patterns     []Pattern
}

// Update updates the schedule, and should be called every game tick
func (schedule *Schedule) Update(entity *Entity) {
	lastPattern := schedule.Patterns[(schedule.index+len(schedule.Patterns)-1)%len(schedule.Patterns)]

	if time.Since(schedule.lastPlayTime) > lastPattern.Duration+lastPattern.Cooldown {
		schedule.Patterns[schedule.index].Start(entity)
		schedule.index = (schedule.index + 1) % len(schedule.Patterns)
		schedule.lastPlayTime = time.Now()
	}
}
