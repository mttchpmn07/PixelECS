package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// SPTAG CProperties tag
	SPTAG = "properties"
)

// CProperties component for storing properties of an entity
type CProperties struct {
	tag    string
	Angle  float64
	Scale  float64
	Bounds pixel.Rect
	Active bool
	Class  string
}

// NewCProperties constructs a CSpriteProperties component with a given angle, scale, bounds, active flag, and class
func NewCProperties(angle, scale float64, bounds pixel.Rect, active bool, class string) ecs.Component {
	return &CProperties{
		tag:    SPTAG,
		Angle:  angle,
		Scale:  scale,
		Bounds: bounds,
		Active: active,
		Class:  class,
	}
}

// GetCProperties returns the actual struct implmenting the component for a given entity
func GetCProperties(e *ecs.Entity) (*CProperties, error) {
	comp, err := e.Query(SPTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CProperties), nil
}

// Tag getter for tag
func (sp *CProperties) Tag() string {
	return sp.tag
}

func (sp *CProperties) String() string {
	return fmt.Sprintf("%v", sp.tag)
}
