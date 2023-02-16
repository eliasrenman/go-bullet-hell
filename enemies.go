package main

import (
	"encoding/json"
	"log"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyMetadata struct {
	Name      string
	MaxHealth int
	Hitbox    map[string]any
	Type      string
	Schedule  Schedule
}

type Enemy struct {
	EnemyMetadata
	Sprites map[string]*Image
}

type EnemyInstance struct {
	*Entity
	Hitbox Collidable
	Enemy
	Health int
}

var enemyCache = make(map[string]Enemy)
var mutex sync.Mutex

func LoadEnemy(name string) Enemy {
	if enemy, ok := enemyCache[name]; ok {
		return enemy
	}

	meta := loadMetadata(name)
	sprites := loadSprites(name)

	enemy := Enemy{
		EnemyMetadata: meta,
		Sprites:       sprites,
	}

	mutex.Lock()
	enemyCache[name] = enemy
	mutex.Unlock()
	return enemy
}

func loadMetadata(name string) EnemyMetadata {
	data, err := Assets.ReadFile("assets/enemies/" + name + "/meta.json")
	if err != nil {
		log.Fatal(err)
	}

	var meta EnemyMetadata
	err = json.Unmarshal(data, &meta)
	if err != nil {
		log.Fatal(err)
	}

	return meta
}

func (enemy *EnemyInstance) loadHitbox() Collidable {
	hitboxData := enemy.EnemyMetadata.Hitbox

	baseHitbox := Hitbox{
		Owner: enemy.Entity,
		Position: Vector{
			X: hitboxData["x"].(float64),
			Y: hitboxData["y"].(float64),
		},
	}

	switch hitboxData["type"] {
	case "circle":
		return &CircleHitbox{
			Hitbox: baseHitbox,
			Radius: hitboxData["radius"].(float64),
		}
	case "rectangle":
		return &RectangleHitbox{
			Hitbox: baseHitbox,
			Size: Vector{
				X: hitboxData["width"].(float64),
				Y: hitboxData["height"].(float64),
			},
		}
	}
	return nil
}

func loadSprites(path string) map[string]*Image {
	// TODO: support spritesheets as animations

	sprites := make(map[string]*Image)
	files, err := Assets.ReadDir("assets/enemies/" + path + "/sprites")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		name := strings.Split(file.Name(), ".")[0]
		sprites[name] = LoadImage("enemies/"+path+"/sprites/"+file.Name(), OriginCenter)
	}

	return sprites
}

func (enemy *Enemy) Spawn() *EnemyInstance {
	return Spawn(&EnemyInstance{
		Entity: &Entity{},
		Enemy:  *enemy,
		Health: enemy.MaxHealth,
	}, CharacterLayer)
}

func (enemy *EnemyInstance) Start() {
	enemy.Hitbox = enemy.loadHitbox()
}

func (enemy *EnemyInstance) Update(game *Game) {
	enemy.Schedule.Update(enemy.Entity)
	enemy.checkBulletCollisions()
}

func (enemy *EnemyInstance) checkBulletCollisions() {
	EachGameObject(func(obj GameObject, layer int) {
		bullet, ok := obj.(*Bullet)
		if ok && bullet.Owner == GetGame().player.Entity && CollidesAt(enemy.Hitbox, enemy.Position, bullet.Hitbox, bullet.Position) {
			enemy.TakeDamage(bullet.Damage)
			Destroy(bullet)
		}
	}, BulletLayer)
}

func (enemy *EnemyInstance) Draw(screen *ebiten.Image) {
	enemy.Sprites["idle"].Draw(screen, enemy.Position, Vector{1, 1}, 0)

	if HitboxesVisible {
		enemy.Hitbox.Draw(screen, enemy.Position)
	}
}

func (enemy *EnemyInstance) Die() {}

func (enemy *EnemyInstance) TakeDamage(amount int) {
	enemy.Health -= amount
	if enemy.Health <= 0 {
		Destroy(enemy)
	}
}
