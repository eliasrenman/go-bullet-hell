package main

import (
	"fmt"
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
	Options    PatternOptions
	Duration   float64
	Cooldown   float64
	BulletType int
}

type PatternOptions struct {
	Count   int
	Speed   float64
	From    float64
	To      float64
	Stagger float64
	Target  Vector
	Amount  Vector
	Easing  string
}

// Start starts the pattern
func (p *Pattern) Start(entity *Entity) {
	switch p.Type {
	case "shoot_arc":
		go p.shootArcPattern(entity)
	case "move_to":
		go p.moveToPattern(entity)
	case "move_by":
		go p.moveByPattern(entity)
	}
}

func (p *Pattern) shootArcPattern(entity *Entity) {
	for i := 0; i < p.Options.Count; i++ {
		angle := p.Options.From + (p.Options.To-p.Options.From)/float64(p.Options.Count)*float64(i)
		entity.Shoot(entity.Position, angle, p.Options.Speed, 0, p.BulletType, 1)
		time.Sleep(time.Duration(p.Options.Stagger * float64(time.Second.Nanoseconds())))
	}
}

func (p *Pattern) moveToPattern(entity *Entity) {
	easeFn := easings[p.Options.Easing]

	if easeFn == nil {
		panic(fmt.Sprintf("Invalid easing function: %s", p.Options.Easing))
	}

	startTime := time.Now()
	start := entity.Position.Copy()
	distance := p.Options.Target.Distance(entity.Position)

	for i := 0.0; i < distance; i = p.Options.Speed * time.Since(startTime).Seconds() {
		entity.Position = start.Plus(p.Options.Target.Minus(start).ScaledBy(easeFn(i / distance)))
		time.Sleep(time.Duration((1.0 / 60) * time.Second.Seconds()))
	}
}

func (p *Pattern) moveByPattern(entity *Entity) {
	easeFn := easings[p.Options.Easing]

	if easeFn == nil {
		panic(fmt.Sprintf("Invalid easing function: %s", p.Options.Easing))
	}

	target := entity.Position.Plus(p.Options.Amount)
	startTime := time.Now()
	start := entity.Position.Copy()
	distance := target.Distance(entity.Position)

	for i := 0.0; i < distance; i = p.Options.Speed * time.Since(startTime).Seconds() {
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
	if len(schedule.Patterns) == 0 {
		return
	}

	lastPattern := schedule.Patterns[(schedule.index+len(schedule.Patterns)-1)%len(schedule.Patterns)]

	if time.Since(schedule.lastPlayTime).Seconds() > lastPattern.Duration+lastPattern.Cooldown {
		schedule.Patterns[schedule.index].Start(entity)
		schedule.index = (schedule.index + 1) % len(schedule.Patterns)
		schedule.lastPlayTime = time.Now()
	}
}
