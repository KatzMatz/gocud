// Package gocud provides Japanese Paint Manufacturers Association (JPMA) color definitions
// and utilities for color conversion and manipulation.
package gocud

import "image/color"

// Color represents a color with its JPMA code and RGB values.
type Color struct {
	JpmaCode string // Japan Paint Manufacturers Association Code
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

// RGB returns the color as color.RGBA with the specified alpha value.
func (c *Color) RGB(alpha uint8) color.RGBA {
	return color.RGBA{c.r, c.g, c.b, alpha}
}

// CMYK returns the color as color.CMYK.
func (c Color) CMYK() color.CMYK {
	cc, mm, yy, kk := color.RGBToCMYK(c.r, c.g, c.b)
	return color.CMYK{cc, mm, yy, kk}
}

var (
	// Standard accent colors with JPMA codes
	RED      Color = newColor("J08-50V", 255, 75, 0)   // Vivid red
	YELLOW   Color = newColor("J27-85V", 255, 241, 0)  // Bright yellow
	GREEN    Color = newColor("J40-60V", 3, 175, 122)  // Forest green
	BLUE     Color = newColor("J72-40T", 0, 90, 255)   // Deep blue
	SKY_BLUE Color = newColor("J69-70T", 77, 196, 255) // Light blue
	PINK     Color = newColor("J02-70T", 255, 128, 130) // Soft pink
	ORANGE   Color = newColor("J15-65X", 246, 170, 0)  // Bright orange
	PURPLE   Color = newColor("J89-40T", 153, 0, 153)  // Dark purple
	BROWN    Color = newColor("J09-30H", 128, 64, 0)   // Earth brown
)

// AccentSet returns all available accent colors as an immutable slice.
func AccentSet() []Color {
	return []Color{RED, YELLOW, GREEN, BLUE, SKY_BLUE, PINK, ORANGE, PURPLE, BROWN}
}

// Accent6Pattern returns a balanced 6-color palette suitable for charts and visualizations.
func Accent6Pattern() []Color {
	return []Color{ORANGE, YELLOW, GREEN, BLUE, SKY_BLUE, BROWN}
}

// Accent5Pattern returns a balanced 5-color palette for medium complexity data.
func Accent5Pattern() []Color {
	return []Color{RED, YELLOW, GREEN, BLUE, SKY_BLUE}
}

// Accent4Pattern returns a minimal 4-color palette for simple visualizations.
func Accent4Pattern() []Color {
	return []Color{RED, YELLOW, GREEN, SKY_BLUE}
}
