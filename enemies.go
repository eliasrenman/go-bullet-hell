package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyMetadata struct {
	Name      string
	MaxHealth int
	Type      string
	Schedule  Schedule
}

type Enemy struct {
	EnemyMetadata
	Sprites map[string]*Image
}

type EnemyInstance struct {
	Entity
	Enemy
	Health int
}

var enemyCache = make(map[string]*Enemy)
var mutex sync.Mutex

func LoadEnemy(name string) *Enemy {
	mutex.Lock()
	if enemy, ok := enemyCache[name]; ok {
		return enemy
	}
	mutex.Unlock()

	meta := loadMetadata(name)
	sprites := loadSprites(name)

	enemy := &Enemy{
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
	fmt.Printf("Spawning enemy %s\n", enemy.Name)

	return Spawn(&EnemyInstance{
		Entity: Entity{
			Position: Vector{},
		},
		Enemy:  *enemy,
		Health: enemy.MaxHealth,
	}, CharacterLayer)
}

func (enemy *EnemyInstance) Start() {}

func (enemy *EnemyInstance) Update(game *Game) {
	enemy.Schedule.Update(&enemy.Entity)
}

func (enemy *EnemyInstance) Draw(screen *ebiten.Image) {
	enemy.Sprites["idle"].Draw(screen, enemy.Position, Vector{1, 1}, 0)
}

func (enemy *EnemyInstance) Die() {}

func (enemy *EnemyInstance) TakeDamage(amount int) {
	enemy.Health -= amount
	if enemy.Health <= 0 {
		Destroy(enemy)
	}
}
