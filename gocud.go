package gocud

import "image/color"

type Color struct {
	JpmaCode string // Japan Pain Manufactures Assosiation Code
	r        uint8
	g        uint8
	b        uint8
}

func newColor(jpma string, r, g, b uint8) Color {
	return Color{
		JpmaCode: jpma,
		r:        r,
		g:        g,
		b:        b,
	}
}

func (c *Color) RGB(alpha uint8) color.RGBA {
	return color.RGBA{c.r, c.g, c.b, alpha}
}

func (c Color) CMYK() color.CMYK {
	cc, mm, yy, kk := color.RGBToCMYK(c.r, c.g, c.b)
	return color.CMYK{cc, mm, yy, kk}
}

var (
	// Accent colors
	RED      Color = newColor("J08-50V", 255, 75, 0)
	YELLOW   Color = newColor("J27-85V", 255, 241, 0)
	GREEN    Color = newColor("J40-60V", 3, 175, 122)
	BLUE     Color = newColor("J72-40T", 0, 90, 255)
	SKY_BLUE Color = newColor("J69-70T", 77, 196, 255)
	PINK     Color = newColor("J02-70T", 255, 128, 130)
	ORANGE   Color = newColor("J15-65X", 246, 170, 0)
	PURPLE   Color = newColor("J89-40T", 153, 0, 153)
	BROWN    Color = newColor("J09-30H", 128, 64, 0)
)

var ACCENT_SET [9]Color = [9]Color{RED, YELLOW, GREEN, BLUE, SKY_BLUE, PINK, ORANGE, PURPLE, BROWN}

var Accent6Pattern []Color = []Color{ORANGE, YELLOW, GREEN, BLUE, SKY_BLUE, BROWN}

var Accent5Pattern []Color = []Color{RED, YELLOW, GREEN, BLUE, SKY_BLUE}

var Accent4Pattern []Color = []Color{RED, YELLOW, GREEN, SKY_BLUE}
