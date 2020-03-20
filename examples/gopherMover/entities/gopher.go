package entities

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

// NewGopher creates a new gopher with a given sprite loaded from a png file and a starting x and y
func NewGopher(asset string, x, y, width float64) (*ecs.Entity, error) {
	gopher, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	loc := components.NewCLocation(x, y)
	err = gopher.Add(loc)
	if err != nil {
		return nil, err
	}

	kin := components.NewCKenetics(500, 10)
	err = gopher.Add(kin)
	if err != nil {
		return nil, err
	}

	sr, err := components.NewCSprite(asset, true)
	if err != nil {
		return nil, err
	}
	err = gopher.Add(sr)
	if err != nil {
		return nil, err
	}
	spriteFrame := sr.(*components.CSprite).Sprite.Frame()

	prop := components.NewCProperties(0, width/spriteFrame.W(), spriteFrame)
	err = gopher.Add(prop)
	if err != nil {
		return nil, err
	}

	return gopher, nil
}
