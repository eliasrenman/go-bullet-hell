package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Bullets struct {
	playerX         *int16
	playerY         *int16
	bullets         []*Bullet
	framesPerBullet uint8
	cooldown        uint8
	image           *ebiten.Image
	bulletSize      uint8
}

func (bullets *Bullets) Draw(screen *ebiten.Image) {
	for _, bullet := range bullets.bullets {

		x, y := float64(bullet.x)+float64(PLAYFIELD_OFFSET)+float64(bullets.bulletSize/2), float64(bullet.y)+float64(PLAYFIELD_OFFSET)+float64(bullets.bulletSize/2)

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
		bullets.bullets = append(bullets.bullets, &Bullet{
			x:          *bullets.playerX,
			y:          *bullets.playerY,
			directions: []int8{0, -1},
		})
		bullets.cooldown = bullets.framesPerBullet
	}
	// And decrease cooldown if it is not already at 0
	if bullets.cooldown > 0 {
		bullets.cooldown--
	}
	// Future: Check colision

}

type Bullet struct {
	x          int16
	y          int16
	directions []int8
}

func (bullet *Bullet) Update() {
	bullet.updateLocation()
}

func (bullet *Bullet) updateLocation() {

	// Set the apporpriate delta depending on if the slow movement is enabled
	var delta int16 = regularBulletSpeed

	// Check X direction
	if bullet.directions[0] < 0 {
		bullet.x += -delta
	} else if bullet.directions[0] > 0 {
		bullet.x += delta
	}
	// Check Y Direction
	if bullet.directions[1] < 0 {
		bullet.y += -delta
	} else if bullet.directions[1] > 0 {
		bullet.y += delta
	}
}

func (bullet *Bullet) isOutOfBounds() bool {
	// Check if bullet is outside of the bounds
	return bullet.x < 0 || bullet.y < 0 || bullet.x > PLAYFIELD_X_MAX || bullet.y > PLAYFIELD_Y_MAX
}
