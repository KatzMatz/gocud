package gocud

import "image/color"

var (
	RED      color.RGBA = color.RGBA{255, 75, 0, 255}
	YELLOW   color.RGBA = color.RGBA{255, 241, 0, 255}
	GREEN    color.RGBA = color.RGBA{3, 175, 122, 255}
	BLUE     color.RGBA = color.RGBA{0, 90, 255, 255}
	SKY_BLUE color.RGBA = color.RGBA{77, 196, 255, 255}
	PINK     color.RGBA = color.RGBA{255, 128, 130, 255}
	ORANGE   color.RGBA = color.RGBA{246, 170, 0, 255}
	PURPLE   color.RGBA = color.RGBA{153, 0, 153, 255}
	BROWN    color.RGBA = color.RGBA{128, 64, 0, 255}
)

var ACCENT_SET [9]color.RGBA = [9]color.RGBA{RED, YELLOW, GREEN, BLUE, SKY_BLUE, PINK, ORANGE, PURPLE, BROWN}

// / Resturn Acceent set
// / Order: RED, YELLOW, GREEN, BLUE, SKY_BLUE, PINK, ORANGE, PURPLE, BROWN
func AccentSet(nums int) []color.RGBA {
	if nums > len(ACCENT_SET) {
		return []color.RGBA{}
	}

	return ACCENT_SET[0:nums]
}
