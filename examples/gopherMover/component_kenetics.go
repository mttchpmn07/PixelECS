package main

import (
	"fmt"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// KTAG const to hold the Location tag
	KTAG = "kenetics"
)

// Kenetics implements the component interface for a location system
type Kenetics struct {
	tag string

	Speed           float64
	AngularVelocity float64
}

// NewKenetics returns a new Kenetics component with a given starting speed and angularVelocity
func NewKenetics(speed, av float64) ecs.Component {
	return &Kenetics{
		tag:             KTAG,
		Speed:           speed,
		AngularVelocity: av,
	}
}

// GetKenetics returns the actual Kenetics struct implmenting the component for a given entity
func GetKenetics(e *ecs.Entity) (*Kenetics, error) {
	comp, err := e.Query(KTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Kenetics), nil
}

// Tag returns the tag for this component
func (k *Kenetics) Tag() string {
	return k.tag
}

func (k *Kenetics) String() string {
	return fmt.Sprintf("%v", k.tag)
}
