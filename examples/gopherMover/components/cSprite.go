package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// SRTAG const to hold the sprite tag
	SRTAG = "sprite"
)

// CSprite component for storing a sprite and its active flag
type CSprite struct {
	tag string

	Render     bool
	Sprite     *pixel.Sprite
	Properties *CProperties
}

// NewCSprite returns a new CSprite component with a sprite given via filename, an active flag
func NewCSprite(filename string, active bool) (ecs.Component, error) {
	r := &CSprite{
		tag:    SRTAG,
		Render: active,
	}
	pic, err := loadPicture(filename)
	if err != nil {
		return nil, err
	}
	r.Sprite = pixel.NewSprite(pic, pic.Bounds())
	return r, nil
}

// GetCSprite returns the actual CSprite struct implemnting the component for a given entity
func GetCSprite(e *ecs.Entity) (*CSprite, error) {
	comp, err := e.Query(SRTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CSprite), nil
}

// Tag returns the tag for this component
func (r *CSprite) Tag() string {
	return r.tag
}

func (r *CSprite) String() string {
	return fmt.Sprintf("%v", r.tag)
}
