package entities

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

// NewGopher creates a new gopher with a given sprite loaded from a png file and a starting x and y
func NewGopher(asset string, x, y float64) (*ecs.Entity, error) {
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

	sp := components.NewCSpriteProperties(0, 1, sr)
	err = gopher.Add(sp)
	if err != nil {
		return nil, err
	}
	sprop, _ := components.GetCSpriteProperties(gopher)
	sprop.Scale = 150 / sprop.Frame.W()

	return gopher, nil
}
