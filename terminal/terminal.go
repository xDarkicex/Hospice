package terminal

import (
	"fmt"
	"strconv"

	"github.com/mgutz/ansi"
)

// color codes http://en.wikipedia.org/wiki/ANSI_escape_code#Colors
//
// var (
// 		// Default wraps colorFunc for the color light green
// 		defaulted = newColor("Default")
// 		// Lime wraps colorFunc for the color light green
// 		lime      = newColor("Lime")
// 		// Red wraps colorFunc for the color red
// 		red       = newColor("Red")
// 		// Blue wraps colorFunc for the color blue
// 		blue      = newColor("Blue")
// 		// Pink wraps colorFunc for the color pink
// 		pink      = newColor("Pink")
// 		// PinkBold wraps colorFunc for the color PinkBold
// 		pinkBold  = newColor("PinkBold")
// 		// LightPink wraps colorFunc for the color LightPink
// 		lightPink = newColor("LightPink")
// )

// Color data struct for colors
type Color struct {
	Style struct {
		Default    func(text string) string
		Blue       func(text string) string
		Red        func(text string) string
		RedBlink   func(text string) string
		Lime       func(text string) string
		Green      func(text string) string
		GreenLight func(text string) string
		Pink       func(text string) string
		PinkLight  func(text string) string
		PinkBold   func(text string) string
		Coral   func(text string) string
		Orange   func(text string) string
	}
	selected func(text string) string
}
var Colors map[int]func(text string) string
func init() {
	Colors = make(map[int]func(text string) string, 0)
	for i := 0; i <= 256; i++ {
		var a string
		a = strconv.FormatInt(int64(i), 10)
		Colors[i] = ansi.ColorFunc(a)
	}
	fmt.Println(Colors[178]("Terminal Package loaded"))
}

// NewTerminalColor Default exposed Creation function returns default white color
// usage:
// var color = NewTerminalColor()
// fmt.Println(color.Style.Red("= TEXT_TO_BE_STYLED =")
func NewTerminalColor() *Color {
	return &Color{
		selected: ansi.ColorFunc("white+h"),
	}
}

func (c *Color) Default(text string) string {
	c.selected = ansi.ColorFunc("white+h")
	return c.write(text)
}

func (c *Color) Red(text string) string {
	c.selected = ansi.ColorFunc("red+h")
	return c.write(text)
}
func (c *Color) RedBlink(text string) string {
	c.selected = ansi.ColorFunc("red+B")
	return c.write(text)
}
func (c *Color) Blue(text string) string {
	c.selected = ansi.ColorFunc("blue+h")
	return c.write(text)
}
func (c *Color) Lime(text string) string {
	c.selected = ansi.ColorFunc("green+h")
	return c.write(text)
}
func (c *Color) Pink(text string) string {
	c.selected = ansi.ColorFunc("199+b")
	return c.write(text)
}
func (c *Color) Orange(text string) string {
	c.selected = ansi.ColorFunc("214+H")
	return c.write(text)
}

func (c *Color) Coral(text string) string {
	c.selected = ansi.ColorFunc("204+h")
	return c.write(text)
}

func (c *Color) GreenLight(text string) string {
	c.selected = ansi.ColorFunc("51+h")
	return c.write(text)
}

func (c *Color) Green(text string) string {
	c.selected = ansi.ColorFunc("10+h")
	return c.write(text)
}
func (c *Color) PinkLight(text string) string {
	c.selected = ansi.ColorFunc("199+b")
	return c.write(text)
}
func (c *Color) PinkBold(text string) string {
	c.selected = ansi.ColorFunc("211+h")
	return c.write(text)
}

func (c *Color) write(in string) string {
	return c.selected(in)
}


