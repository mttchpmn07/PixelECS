package entities

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

func buildAnimations() (ecs.Component, error) {
	seq, err := components.NewSequence("assets/bug.png", 10, 105, 105, 0, true)
	if err != nil {
		return nil, err
	}

	seqMap := map[string]*components.Sequence{
		"fly": seq,
	}
	an := components.NewCAnimation(seqMap, "fly", true)
	return an, nil
}

// NewFly creates a new animated fly
func NewFly(x, y, width float64) (*ecs.Entity, error) {
	fly, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	loc := components.NewCLocation(x, y)
	err = fly.Add(loc)
	if err != nil {
		return nil, err
	}

	kin := components.NewCKenetics(500, 10)
	err = fly.Add(kin)
	if err != nil {
		return nil, err
	}

	an, err := buildAnimations()
	if err != nil {
		return nil, err
	}
	err = fly.Add(an)
	if err != nil {
		return nil, err
	}

	bounds := an.(*components.CAnimation).GetCurrentSprite().Frame()
	sp := components.NewCProperties(0, width/bounds.W(), bounds)
	err = fly.Add(sp)
	if err != nil {
		return nil, err
	}

	return fly, nil
}
