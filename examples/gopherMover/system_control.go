package main

import (
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

// ControlSystem
type ControlSystem struct {
	controlEntities []*ecs.Entity

	userVelocity float64
}

// NewControlSystem
func NewControlSystem(velocity float64, es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &ControlSystem{
		controlEntities: []*ecs.Entity{},
		userVelocity:    velocity,
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// Update
func (cs *ControlSystem) Update(win *pixelgl.Window, dt float64) error {
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		for _, e := range cs.controlEntities {
			loc, err := GetLocation(e)
			if err != nil {
				return err
			}
			loc.Loc.X -= cs.userVelocity * dt
		}
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		for _, e := range cs.controlEntities {
			loc, err := GetLocation(e)
			if err != nil {
				return err
			}
			loc.Loc.X += cs.userVelocity * dt
		}
	}
	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		for _, e := range cs.controlEntities {
			loc, err := GetLocation(e)
			if err != nil {
				return err
			}
			loc.Loc.Y -= cs.userVelocity * dt
		}
	}
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		for _, e := range cs.controlEntities {
			loc, err := GetLocation(e)
			if err != nil {
				return err
			}
			loc.Loc.Y += cs.userVelocity * dt
		}
	}

	return nil
}

// Render
func (cs *ControlSystem) Render(win *pixelgl.Window, dt float64) error {
	return nil
}

// AddEntity
func (cs *ControlSystem) AddEntity(es ...*ecs.Entity) error {
	cs.controlEntities = append(cs.controlEntities, es...)
	return nil
}

// RemoveEntity
func (cs *ControlSystem) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(cs.controlEntities, e)
		if err != nil {
			return err
		}
		cs.controlEntities = newEntries
	}
	return nil
}
