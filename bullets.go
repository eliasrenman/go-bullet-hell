package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullets struct {
	playerX          int
	playerY          int
	bullets          []Bullet
	framesPerBullet  uint8
	cooldown         uint8
	image            *ebiten.Image
	bulletSize       uint
	defaultDirection []int8
	defaultDelta     int16
}

func (bullets *Bullets) Draw(screen *ebiten.Image) {
	for _, bullet := range bullets.bullets {
		x, y := normalizeCoords(bullet.x, bullet.y)
		x -= float64(bullet.size / 2)
		y -= float64(bullet.size / 2)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(bullets.image, op)
	}
}

func (bullets *Bullets) Update(input *Input) {
	i := 0 // output index
	// create a zero-length slice with the same underlying array

	tmp := bullets.bullets[:0]
	for _, bullet := range bullets.bullets {
		// Update location of existing bullet(s)
		bullet.Update()
		// Remove bullets that are outside the playing field
		if !bullet.isOutOfBounds() {
			// copy and increment index
			// bullets.bullets[i] = bullet
			i++
			tmp = append(tmp, bullet)

		}
	}

	bullets.bullets = tmp

	// Check if we should add new bullets
	if bullets.cooldown == 0 && input.shootingRegularGun {
		bullets.Spawn(bullets.playerX, bullets.playerY, bullets.bulletSize, 0, float64(bullets.defaultDelta))
		bullets.cooldown = bullets.framesPerBullet
	}
	// And decrease cooldown if it is not already at 0
	if bullets.cooldown > 0 {
		bullets.cooldown--
	}
	// Future: Check colision

}

func (bullets *Bullets) Spawn(x int, y int, size uint, direction float64, speed float64) {
	velocity := [2]float64{
		math.Cos(direction) * speed,
		math.Sin(direction) * speed,
	}

	bullets.bullets = append(bullets.bullets, Bullet{
		x:        x,
		y:        y,
		velocity: velocity,
		size:     size,
	})
}

type Bullet struct {
	x        int
	y        int
	velocity [2]float64
	size     uint
}

func (bullet *Bullet) Update() {
	bullet.updateLocation()
}

func (bullet *Bullet) updateLocation() {
	bullet.x += int(bullet.velocity[0])
	bullet.y += int(bullet.velocity[1])
}

func (bullet *Bullet) isOutOfBounds() bool {
	return bullet.x < -int(bullet.size) ||
		bullet.y < -int(bullet.size) ||
		bullet.x > PLAYFIELD_WIDTH ||
		bullet.y > PLAYFIELD_HEIGHT
}
