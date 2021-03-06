package entities

import (
	"math/rand"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherFlyer/components"
)

// NewFly creates a new animated fly
func NewFly(winWidth, winHeight, spriteWidth float64, asset *components.CBatchAsset) (*ecs.Entity, error) {
	fly, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	err = fly.Add(asset)
	if err != nil {
		return nil, err
	}

	x := winWidth * rand.Float64()
	y := winHeight * rand.Float64()
	loc := components.NewCLocation(x, y, rand.Intn(10))
	err = fly.Add(loc)
	if err != nil {
		return nil, err
	}

	cLoc := loc.(*components.CLocation)
	poly := components.NewCCollisionPoly(
		cLoc,
		pixel.V(spriteWidth/2, spriteWidth/2),
		pixel.V(spriteWidth/2, -spriteWidth/2),
		pixel.V(-spriteWidth/2, -spriteWidth/2),
		pixel.V(-spriteWidth/2, spriteWidth/2),
	)
	err = fly.Add(poly)
	if err != nil {
		return nil, err
	}

	vel := pixel.V(rand.Float64()-0.5, rand.Float64()-0.5).Unit()
	kin := components.NewCKenetics(100, 10, 1000, vel.Scaled(100), pixel.V(0, 0))
	err = fly.Add(kin)
	if err != nil {
		return nil, err
	}

	seq, err := components.NewSequence(asset, 10, 105, 105, 0, 0, 1, true)
	if err != nil {
		return nil, err
	}
	seqMap := map[string]*components.Sequence{
		"fly": seq,
	}
	an := components.NewCAnimation(seqMap, "fly")
	err = fly.Add(an)
	if err != nil {
		return nil, err
	}

	bounds := an.(*components.CAnimation).GetCurrentFrame()
	sp := components.NewCProperties(0, spriteWidth/bounds.W(), bounds, true, "fly")
	err = fly.Add(sp)
	if err != nil {
		return nil, err
	}

	return fly, nil
}
