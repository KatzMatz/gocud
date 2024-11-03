package gocud

import "image/color"

var (
	// Accent Colors
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

var Accent6Pattern []color.RGBA = []color.RGBA{ORANGE, YELLOW, GREEN, BLUE, SKY_BLUE, BROWN}

var Accent5Pattern []color.RGBA = []color.RGBA{RED, YELLOW, GREEN, BLUE, SKY_BLUE}

var Accent4Patter []color.RGBA = []color.RGBA{RED, YELLOW, GREEN, SKY_BLUE}
