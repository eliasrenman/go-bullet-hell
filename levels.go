package main

import (
	"fmt"
	"time"
)

type Level struct {
	Events []Event
}

type Event struct {
	Type     string
	Blocking bool
	Cooldown float64
	Options  EventOptions
}

type EventOptions struct {
	Enemies []EnemySpawnOptions
}

type EnemySpawnOptions struct {
	Enemy    string
	Position Vector
}

func (level *Level) Start() {
	for _, event := range level.Events {
		event.Start()
		time.Sleep(time.Duration(event.Cooldown * float64(time.Second)))
	}
}

func (event *Event) Start() {
	fmt.Println("Starting event", event.Type)

	var eventFn func()
	switch event.Type {
	case "spawn":
		eventFn = event.spawnEvent
	}

	if event.Blocking {
		eventFn()
	} else {
		go eventFn()
	}
}

func (e *Event) spawnEvent() {
	fmt.Printf("Spawning enemies: %v\n", e.Options.Enemies)

	for _, options := range e.Options.Enemies {
		enemy := LoadEnemy(options.Enemy)
		instance := enemy.Spawn()
		fmt.Printf("%v spawned at %v", instance, options.Position)
		instance.Position = options.Position
	}
}
