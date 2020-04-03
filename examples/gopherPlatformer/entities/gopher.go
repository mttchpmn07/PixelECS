package entities

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
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

	cLoc := loc.(*components.CLocation)
	poly := components.NewCCollisionPoly(
		cLoc,
		pixel.V(spriteWidth/4, 3*spriteWidth/8),
		pixel.V(spriteWidth/4, -3*spriteWidth/8),
		pixel.V(-spriteWidth/4, -3*spriteWidth/8),
		pixel.V(-spriteWidth/4, 3*spriteWidth/8),
	)
	err = gopher.Add(poly)
	if err != nil {
		return nil, err
	}

	kin := components.NewCKenetics(500, 10, 0, pixel.V(0, 0), pixel.V(0, 0))
	err = gopher.Add(kin)
	if err != nil {
		return nil, err
	}

	leftSeq, err := components.NewSequence(asset, 15, 200, 200, 0, 0, 5, true)
	if err != nil {
		return nil, err
	}

	rightSeq, err := components.NewSequence(asset, 15, 200, 200, 0, 6, 11, true)
	if err != nil {
		return nil, err
	}

	seqMap := map[string]*components.Sequence{
		"right": rightSeq,
		"left":  leftSeq,
	}
	an := components.NewCAnimation(seqMap, "left")
	if err != nil {
		return nil, err
	}
	err = gopher.Add(an)
	if err != nil {
		return nil, err
	}

	bounds := an.(*components.CAnimation).GetCurrentFrame()
	sp := components.NewCProperties(0, spriteWidth/bounds.W(), bounds, true, "gopher")
	err = gopher.Add(sp)
	if err != nil {
		return nil, err
	}

	return gopher, nil
}
