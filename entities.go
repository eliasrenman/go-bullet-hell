package main

import (
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

// Entity is an object in space,
// it has a position and a velocity
type Entity struct {
	Position Vector
	Velocity Vector
}

// Move moves the entity by a given vector relative to its position
func (entity *Entity) Move(vector Vector) {
	entity.Position.X += vector.X
	entity.Position.Y += vector.Y
}

// GameObject is an interface for all game objects.
// Game objects are spawned using the Spawn function, which will call the Start method.
type GameObject interface {
	Start()
	Update()
	Die()
	Draw(image *ebiten.Image)
}

// All background objects that should not interact with neither the CharcterObjects nor the BulletObjects.
var BackgroundObjects = make(map[GameObject]struct{})

// All objects that are characters, like the player and enemies.
var CharacterObjects = make(map[GameObject]struct{})

// BulletObjects is a set of all bullet objects
var BulletObjects = make(map[GameObject]struct{})
var mu sync.Mutex

// The spawn queue is used to spawn new background objects in the next frame, to avoid concurrent map writes
var backgroundSpawnQueue = make(map[GameObject]struct{})

// The spawn queue is used to spawn new character objects in the next frame, to avoid concurrent map writes
var characterSpawnQueue = make(map[GameObject]struct{})

// The spawn queue is used to spawn new bullet objects in the next frame, to avoid concurrent map writes
var bulletSpawnQueue = make(map[GameObject]struct{})

const (
	CharacterQueue string = "character"
	BulletQueue    string = "bullet"
)

// Spawn creates a new copy of a game object
func Spawn[T GameObject](obj T, queue string) T {
	mu.Lock()

	switch queue {
	case CharacterQueue:
		characterSpawnQueue[obj] = struct{}{}
	case BulletQueue:
		bulletSpawnQueue[obj] = struct{}{}
	}
	mu.Unlock()

	obj.Start()
	return obj
}

// SpawnObjects spawns all bullet objects in the spawn queue, and clears the queue.
// This function is called at the end of each frame to avoid concurrent map writes.
func SpawnObjects() {
	mu.Lock()
	for obj := range bulletSpawnQueue {
		BulletObjects[obj] = struct{}{}
	}
	bulletSpawnQueue = make(map[GameObject]struct{})
	for obj := range characterSpawnQueue {
		CharacterObjects[obj] = struct{}{}
	}
	characterSpawnQueue = make(map[GameObject]struct{})
	for obj := range backgroundSpawnQueue {
		CharacterObjects[obj] = struct{}{}
	}
	backgroundSpawnQueue = make(map[GameObject]struct{})
	mu.Unlock()
}

// Destroy a game object
func Destroy(obj GameObject) {
	obj.Die()
	delete(BulletObjects, obj)
}

// Bullet is an Entity with additional values for Damage, Size, Speed and Direction
type Bullet struct {
	Entity
	Hitbox Collidable
	Owner  *Entity
	Damage int

	// Read-only! Use SetAngularVelocity to set the speed
	Speed float64
	// Read-only! Use SetAngularVelocity to set the direction
	Direction  float64
	BulletType int
}

// SetAngularVelocity sets the velocity of the bullet given a speed and a direction
func (b *Bullet) SetAngularVelocity(speed float64, direction float64) {
	b.Speed = speed
	b.Direction = direction
	b.Velocity = VectorFromAngle(direction).ScaledBy(speed)
}

// Shoot spawns a bullet at a given position, with a given speed and direction
func (entity *Entity) Shoot(position Vector, direction float64, speed float64, offset float64, bulletType int) {

	// This offests the inital position based on the direction of the bullet.
	position.Add(VectorFromAngle(direction).ScaledBy(offset))

	bullet := Spawn(&Bullet{
		Entity:     Entity{Position: position},
		Owner:      entity,
		BulletType: bulletType,
		Hitbox:     getBulletHitbox(bulletType),
	}, BulletQueue)

	bullet.SetAngularVelocity(speed, direction)
}

func getBulletHitbox(bulletType int) CircleHitbox {
	switch bulletType {
	case BulletSmallYellow:
		return CircleHitbox{Radius: 4}
	default:
		return CircleHitbox{Radius: 4}
	}
}

// Start is called when the bullet is spawned
func (b *Bullet) Start() {}

// Update is called every game tick. 60 times per second
func (b *Bullet) Update() {
	b.Move(b.Velocity)
	if b.Position.Y < 0 || b.Position.Y > ScreenSize.Y {
		Destroy(b)
	}
}

var bulletImage = LoadImage("bullets/bullet_0.png", OriginCenter)
var bulletImage1 = LoadImage("bullets/bullet_1.png", OriginCenter)

const (
	BulletSmallBlue   = 0
	BulletSmallYellow = 1
)

// Draw is called every frame to draw the bullet
func (b *Bullet) Draw(screen *ebiten.Image) {
	var bImage *Image
	switch b.BulletType {
	case BulletSmallYellow:
		bImage = bulletImage1
	default:
		bImage = bulletImage
	}

	bImage.Draw(screen, b.Position, Vector{X: 1, Y: 1}, 0)
}

// Die is called when the bullet is destroyed
func (b *Bullet) Die() {
}
