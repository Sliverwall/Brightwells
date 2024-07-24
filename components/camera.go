package components

import "github.com/hajimehoshi/ebiten/v2"

const CameraComponentID = "CameraComponentID"

type CameraComponent struct {
	X, Y          float64
	Width, Height float64
	Viewport      *ebiten.Image
}
