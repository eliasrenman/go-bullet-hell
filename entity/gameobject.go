package entity

import "github.com/hajimehoshi/ebiten/v2"

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

// Spawn a new copy of a game object
func Spawn[T GameObject](obj T) T {
	GameObjects[obj] = struct{}{}
	obj.Start()
	return obj
}

// Destroy a game object
func Destroy(obj GameObject) {
	obj.Die()
	delete(GameObjects, obj)
}
