package components

// RenderComponentID is the identifier for the RenderComponent
const RenderComponentID = "RenderComponent"

type RenderComponent struct {
	Layer int // Higher layers are drawn on top of lower layers
}
