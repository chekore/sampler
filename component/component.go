package component

import (
	ui "github.com/gizak/termui/v3"
	"github.com/sqshq/sampler/config"
)

type Component struct {
	Type          config.ComponentType
	Drawable      ui.Drawable
	Title         string
	Position      config.Position
	Size          config.Size
	RefreshRateMs int
}

func (c *Component) Move(x, y int) {
	c.Position.X += x
	c.Position.Y += y
	c.normalize()
}

func (c *Component) Resize(x, y int) {
	c.Size.X += x
	c.Size.Y += y
	c.normalize()
}

func (c *Component) normalize() {
	if c.Position.X < 0 {
		c.Position.X = 0
	}
	if c.Position.Y < 0 {
		c.Position.Y = 0
	}
}