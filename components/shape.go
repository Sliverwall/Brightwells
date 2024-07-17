package components

// ShapeComponentID is the identifier for the ShapeComponent
const ShapeComponentID = "ShapeComponent"
const ShapeColorComponentID = "ShapeColorComponentID"

type ShapeComponent struct {
	Height, Width float64
	Geometry      string
}

type ShapeColorComponent struct {
	// RGBA units for shape's color
	R, G, B, A uint8
}
