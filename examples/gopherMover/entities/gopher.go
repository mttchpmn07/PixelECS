package entities

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

// NewGopher creates a new gopher with a given sprite loaded from a png file and a starting x and y
func NewGopher(winWidth, winHeight, spriteWidth float64, asset *components.CBatchAsset) (*ecs.Entity, error) {
	gopher, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	err = gopher.Add(asset)
	if err != nil {
		return nil, err
	}

	loc := components.NewCLocation(winWidth/2, winHeight/2, 5)
	err = gopher.Add(loc)
	if err != nil {
		return nil, err
	}

	kin := components.NewCKenetics(500, 10, 0, pixel.V(0, 0), pixel.V(0, 0))
	err = gopher.Add(kin)
	if err != nil {
		return nil, err
	}

	seq, err := components.NewSequence(asset, 1000,
		asset.Spritesheet.Bounds().W(), asset.Spritesheet.Bounds().H(), 0, true)
	if err != nil {
		return nil, err
	}

	seqMap := map[string]*components.Sequence{
		"walk": seq,
	}
	an := components.NewCAnimation(seqMap, "walk", true)
	if err != nil {
		return nil, err
	}
	err = gopher.Add(an)
	if err != nil {
		return nil, err
	}

	bounds := an.(*components.CAnimation).GetCurrentFrame()
	sp := components.NewCProperties(0, spriteWidth/bounds.W(), bounds)
	err = gopher.Add(sp)
	if err != nil {
		return nil, err
	}

	return gopher, nil
}
