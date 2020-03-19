package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// SPTAG const to hold the Location tag
	SPTAG = "spritepropertiies"
)

// CSpriteProperties component for storing properties of a sprite
type CSpriteProperties struct {
	tag   string
	Angle float64
	Scale float64
	Frame pixel.Rect
}

// NewCSpriteProperties returns a new CSpriteProperties component with a given angle, scale, and sprite bounds (requires input of sprite component)
func NewCSpriteProperties(angle, scale float64, sr ecs.Component) ecs.Component {
	r := sr.(*CSprite)
	return &CSpriteProperties{
		tag:   SPTAG,
		Angle: angle,
		Scale: scale,
		Frame: r.Sprite.Frame(),
	}
}

// GetCSpriteProperties returns the actual CSpriteProperties struct for a given entity
func GetCSpriteProperties(e *ecs.Entity) (*CSpriteProperties, error) {
	comp, err := e.Query(SPTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CSpriteProperties), nil
}

// Tag returns the tag for this component
func (sp *CSpriteProperties) Tag() string {
	return sp.tag
}

func (sp *CSpriteProperties) String() string {
	return fmt.Sprintf("%v", sp.tag)
}
