package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// RTAG const to hold the SRenderer tag
	RTAG = "renderer"
)

// SRenderer Sprite Render System
type SRenderer struct {
	tag string

	controlEntities []*ecs.Entity
}

// NewSRenderer returns a new sprite render system with a give list of entities attached via a variadic function call
func NewSRenderer(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &SRenderer{
		tag:             RTAG,
		controlEntities: []*ecs.Entity{},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// Update draws sprite for each associated entity
func (rs *SRenderer) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	for _, e := range rs.controlEntities {
		sr, err := components.GetCSprite(e)
		if err != nil {
			return err
		}
		if !sr.Render {
			continue
		}
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		trans := pixel.IM.Scaled(pixel.ZV, sp.Scale).Rotated(pixel.ZV, sp.Angle)
		sr.Sprite.Draw(win, trans.Moved(loc.Loc))
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (rs *SRenderer) AddEntity(es ...*ecs.Entity) error {
	rs.controlEntities = append(rs.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (rs *SRenderer) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(rs.controlEntities, e)
		if err != nil {
			return err
		}
		rs.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (rs *SRenderer) Tag() string {
	return rs.tag
}
