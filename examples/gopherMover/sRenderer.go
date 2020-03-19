package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

// SRenderer Sprite Render System
type SRenderer struct {
	controlEntities []*ecs.Entity
}

// NewSRenderer returns a new sprite render system with a give list of entities attached via a variadic function call
func NewSRenderer(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &SRenderer{
		controlEntities: []*ecs.Entity{},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// Update draws sprite for each associated entity
func (rs *SRenderer) Update(win *pixelgl.Window, dt float64) error {
	for _, e := range rs.controlEntities {
		sr, err := GetCSprite(e)
		if err != nil {
			return err
		}
		if !sr.Active {
			continue
		}
		loc, err := GetCLocation(e)
		if err != nil {
			return err
		}
		sp, err := GetCSpriteProperties(e)
		if err != nil {
			return err
		}
		trans := pixel.IM.Scaled(pixel.ZV, sp.Scale).Rotated(pixel.ZV, sp.Angle)
		sr.sprite.Draw(win, trans.Moved(loc.Loc))
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
