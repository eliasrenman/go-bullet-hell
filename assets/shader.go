package assets

import (
	"io"
	"log"

	"github.com/eliasrenman/go-bullet-hell/geometry"
	"github.com/eliasrenman/go-bullet-hell/util"
	"github.com/hajimehoshi/ebiten/v2"
)

type Shader struct {
	*ebiten.Shader
}

func LoadShader(path string) *Shader {
	file, err := Assets.Open("data/" + path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	shader, err := ebiten.NewShader(data)
	if err != nil {
		log.Fatal(err)
	}

	return &Shader{
		Shader: shader,
	}
}

func (shader *Shader) Draw(target *ebiten.Image, position geometry.Point, scale geometry.Size, rotation float64, images []*Image, uniforms map[string]any) {
	op := &ebiten.DrawRectShaderOptions{}
	width := target.Bounds().Dx()
	height := target.Bounds().Dy()

	// If there are images, use the size of the first image.
	// All images must be the same size.
	if images != nil {
		width, height = images[0].Image.Size()
	}

	op.Uniforms = map[string]any{
		"Time":       float32(util.CurrentSeconds()),
		"Resolution": []float32{float32(width), float32(height)},
	}

	for i, image := range images {
		op.Images[i] = image.Image
	}

	for key, value := range uniforms {
		op.Uniforms[key] = value
	}

	TranslateScaleAndRotateImage(&op.GeoM, position, scale, rotation)
	target.DrawRectShader(width, height, shader.Shader, op)
}
