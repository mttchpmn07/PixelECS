package main

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// SPTAG const to hold the Location tag
	SPTAG = "spritepropertiies"
)

// SpriteProperties implements the component interface for a location system
type SpriteProperties struct {
	tag   string
	Angle float64
	Scale float64
	Frame pixel.Rect
}

// NewSpriteProperties returns a new location component with a given starting x and y
func NewSpriteProperties(angle, scale float64, sr ecs.Component) ecs.Component {
	r := sr.(*Sprite)
	return &SpriteProperties{
		tag:   SPTAG,
		Angle: angle,
		Scale: scale,
		Frame: r.sprite.Frame(),
	}
}

// GetLSpriteProperties returns the actual SpriteProperties struct implmenting the component for a given entity
func GetSpriteProperties(e *ecs.Entity) (*SpriteProperties, error) {
	comp, err := e.Query(SPTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*SpriteProperties), nil
}

// Tag returns the tag for this component
func (sp *SpriteProperties) Tag() string {
	return sp.tag
}

func (sp *SpriteProperties) String() string {
	return fmt.Sprintf("%v", sp.tag)
}
