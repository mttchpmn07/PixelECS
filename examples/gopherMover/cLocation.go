package main

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// LTAG const to hold the Location tag
	LTAG = "location"
)

// CLocation component interface for storing the location of an entity
type CLocation struct {
	tag string
	Loc pixel.Vec
}

// NewCLocation returns a new CLocation component with a given starting x and y
func NewCLocation(x, y float64) ecs.Component {
	return &CLocation{
		tag: LTAG,
		Loc: pixel.V(x, y),
	}
}

// GetCLocation returns the actual CLocation struct for a given entity
func GetCLocation(e *ecs.Entity) (*CLocation, error) {
	comp, err := e.Query(LTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CLocation), nil
}

// Tag returns the tag for this component
func (l *CLocation) Tag() string {
	return l.tag
}

func (l *CLocation) String() string {
	return fmt.Sprintf("%v", l.tag)
}
