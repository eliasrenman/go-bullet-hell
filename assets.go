package main

import (
	"embed"
	"image"
	"io"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets
var Assets embed.FS

type Background struct {
	Image    *Image
	Velocity Vector
	offset   Vector
}

func (background *Background) Update() {
	offset := background.Velocity.ScaledBy((1. / 100) / 60)
	background.offset.Add(offset)
}

func (background *Background) Draw(screen *ebiten.Image) {
	background.Image.DrawTiled(screen, Vector{}, Vector{X: 2, Y: 2}, 0, background.offset)

}

type Font struct {
	font.Face
}

func LoadFont(path string, op opentype.FaceOptions) *Font {
	data, err := Assets.ReadFile("assets/" + path)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(data)
	if err != nil {
		log.Fatal(err)
	}

	if op.Size == 0 {
		op = opentype.FaceOptions{
			Size:    12,
			DPI:     72,
			Hinting: font.HintingFull,
		}
	}

	fnt, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}

	return &Font{fnt}
}

var (
	OriginTop         = Vector{X: 0.5, Y: 0}
	OriginTopLeft     = Vector{X: 0, Y: 0}
	OriginTopRight    = Vector{X: 1, Y: 0}
	OriginCenter      = Vector{X: 0.5, Y: 0.5}
	OriginLeft        = Vector{X: 0, Y: 0.5}
	OriginRight       = Vector{X: 1, Y: 0.5}
	OriginBottom      = Vector{X: 0.5, Y: 1}
	OriginBottomLeft  = Vector{X: 0, Y: 1}
	OriginBottomRight = Vector{X: 1, Y: 1}
)

type Image struct {
	*ebiten.Image
	Size   Vector
	Origin Vector
}

func LoadImage(path string, origin Vector) *Image {
	data, err := Assets.Open("assets/" + path)
	if err != nil {
		log.Fatal(err)
	}

	dataImg, _, err := image.Decode(data)
	if err != nil {
		log.Fatal(err)
	}

	image := ebiten.NewImageFromImage(dataImg)

	width, height := image.Size()
	return &Image{
		Image: image,
		Size: Vector{
			X: float64(width),
			Y: float64(height),
		},
		Origin: origin,
	}
}

func TranslateScaleAndRotateImage(geom *ebiten.GeoM, position Vector, scale Vector, rotation float64) {
	geom.Translate(position.X, position.Y)
	geom.Scale(float64(scale.X), float64(scale.Y))
	geom.Rotate(rotation)
}

func (image *Image) Draw(target *ebiten.Image, position Vector, scale Vector, rotation float64) {
	op := &ebiten.DrawImageOptions{}
	position.Subtract(image.Origin.Dot(image.Size))
	TranslateScaleAndRotateImage(&op.GeoM, position, scale, rotation)
	target.DrawImage(image.Image, op)
}

var tilingShader = LoadShader("shaders/tile.go")

func (image *Image) DrawTiled(target *ebiten.Image, position Vector, scale Vector, rotation float64, offset Vector) {
	images := []*Image{image}
	uniforms := map[string]any{
		"Offset": []float32{float32(offset.X), float32(offset.Y)},
	}
	tilingShader.Draw(target, position, scale, rotation, images, uniforms)
}

type Shader struct {
	*ebiten.Shader
}

func LoadShader(path string) *Shader {
	file, err := Assets.Open("assets/" + path)
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

func (shader *Shader) Draw(target *ebiten.Image, position Vector, scale Vector, rotation float64, images []*Image, uniforms map[string]any) {
	op := &ebiten.DrawRectShaderOptions{}
	width := target.Bounds().Dx()
	height := target.Bounds().Dy()

	// If there are images, use the size of the first image.
	// All images must be the same size.
	if images != nil {
		width, height = images[0].Image.Size()
	}

	op.Uniforms = map[string]any{
		"Time":       float32(CurrentSeconds()),
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
