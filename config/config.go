package config

var (
	// WindowSize stores the current window size
	WindowSize = struct {
		Width  int
		Height int
	}{
		Width:  600,
		Height: 600,
	}

	ResolutionSize = struct {
		Width  int
		Height int
	}{
		Width:  190,
		Height: 120,
	}

	TileSize float64 = 16
)
