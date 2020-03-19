package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

// RendererSystem
type RendererSystem struct {
	controlEntities []*ecs.Entity
}

// NewRendererSystem
func NewRendererSystem(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &RendererSystem{
		controlEntities: []*ecs.Entity{},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// Update
func (rs *RendererSystem) Update(win *pixelgl.Window, dt float64) error {
	return nil
}

// Render
func (rs *RendererSystem) Render(win *pixelgl.Window, dt float64) error {
	for _, e := range rs.controlEntities {
		sr, err := GetSprite(e)
		if err != nil {
			return err
		}
		if !sr.Active {
			continue
		}
		loc, err := GetLocation(e)
		if err != nil {
			return err
		}
		sp, err := GetSpriteProperties(e)
		if err != nil {
			return err
		}
		trans := pixel.IM.Scaled(pixel.ZV, sp.Scale).Rotated(pixel.ZV, sp.Angle)
		sr.sprite.Draw(win, trans.Moved(loc.Loc))
	}
	return nil
}

// AddEntity
func (rs *RendererSystem) AddEntity(es ...*ecs.Entity) error {
	rs.controlEntities = append(rs.controlEntities, es...)
	return nil
}

// RemoveEntity
func (rs *RendererSystem) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(rs.controlEntities, e)
		if err != nil {
			return err
		}
		rs.controlEntities = newEntries
	}
	return nil
}
