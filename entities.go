package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// An Entity is an object in space
// it has a position and a velocity
type Entity struct {
	Position Vector
	Velocity Vector
}

func (entity *Entity) Move(vector Vector) {
	entity.Position.X += vector.X
	entity.Position.Y += vector.Y
}

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

// Bullets are Entities with additional values for Damage, Size, Speed and Direction
type Bullet struct {
	Entity
	Owner  *Entity
	Damage int

	// Read-only! Use SetAngularVelocity to set the speed
	Speed float64
	// Read-only! Use SetAngularVelocity to set the direction
	Direction float64
}

func (b *Bullet) SetAngularVelocity(speed float64, direction float64) {
	b.Speed = speed
	b.Direction = direction
	b.Velocity = VectorFromAngle(direction).ScaledBy(speed)
}

func (owner *Entity) Shoot(position Vector, direction float64, speed float64, offset float64) {

	// This offests the inital position based on the direction of the bullet.
	position.Add(VectorFromAngle(direction).ScaledBy(offset))

	bullet := Spawn(&Bullet{
		Entity: Entity{Position: position},
		Owner:  owner,
	})
	bullet.SetAngularVelocity(speed, direction)
}

func (b *Bullet) Start() {}

func (b *Bullet) Update() {
	b.Move(b.Velocity)
	if b.Position.Y < 0 || b.Position.Y > ScreenSize.Y {
		Destroy(b)
	}
}

var bulletImage = LoadImage("bullets/bullet.png", OriginCenter)

func (b *Bullet) Draw(screen *ebiten.Image) {
	bulletImage.Draw(screen, b.Position, Vector{X: 1, Y: 1}, 0)
}

func (b *Bullet) Die() {
}
