package main

import "time"

type Level struct {
	Events []Event
	index  int
}

type Event struct {
	Type     string
	Blocking bool
	Cooldown time.Duration
	Options  map[string]any
}

func (level *Level) Start() {
	for _, event := range level.Events {
		event.Start()
		time.Sleep(event.Cooldown)
	}
}

func (event *Event) Start() {
	var eventFn func(map[string]any)
	switch event.Type {
	case "spawn":
		eventFn = spawnEvent
	}

	if event.Blocking {
		eventFn(event.Options)
	} else {
		go eventFn(event.Options)
	}
}

func spawnEvent(options map[string]any) {
	// TODO: Spawn logic here
}
