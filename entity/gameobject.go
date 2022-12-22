package entity

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject interface {
	Start()
	Update()
	Die()
	Draw(image *ebiten.Image)
}

// Workaround for Set types
// empty structs consume 0 bytes of memory, so a mapping between T->struct{}
// is equivalent to what most other languages would call a "Set"
var GameObjects = make(map[GameObject]struct{})

// The spawn queue is used to spawn new game objects in the next frame, to avoid concurrent map writes
var spawnQueue = make(map[GameObject]struct{})
var mu sync.RWMutex

// Spawn a new copy of a game object
func Spawn[T GameObject](obj T) T {
	spawnQueue[obj] = struct{}{}
	obj.Start()
	return obj
}

func EachGameObject(cb func(GameObject)) {
	mu.RLock()
	for obj := range GameObjects {
		cb(obj)
	}
	mu.RUnlock()
}

// Spawn all game objects in the spawn queue, and clear the queue
func SpawnGameObjects() {
	mu.Lock()
	for obj := range spawnQueue {
		GameObjects[obj] = struct{}{}
	}
	spawnQueue = make(map[GameObject]struct{})
	mu.Unlock()
}

// Destroy a game object
func Destroy(obj GameObject) {
	obj.Die()
	delete(GameObjects, obj)
}
