package main

import (
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

// KeybaordControlSystem
type KeybaordControlSystem struct {
	controlEntities []*ecs.Entity
}

// NewKeybaordControlSystem
func NewKeybaordControlSystem(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &KeybaordControlSystem{
		controlEntities: []*ecs.Entity{},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// Update
func (cs *KeybaordControlSystem) Update(win *pixelgl.Window, dt float64) error {
	for _, e := range cs.controlEntities {
		loc, err := GetLocation(e)
		if err != nil {
			return err
		}
		kin, err := GetKenetics(e)
		if err != nil {
			return err
		}
		sp, err := GetSpriteProperties(e)
		if err != nil {
			return err
		}
		if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
			loc.Loc.X -= kin.Speed * dt
		}
		if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
			loc.Loc.X += kin.Speed * dt

		}
		if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
			loc.Loc.Y -= kin.Speed * dt

		}
		if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
			loc.Loc.Y += kin.Speed * dt

		}
		if win.Pressed(pixelgl.KeyE) {
			sp.Angle -= kin.AngularVelocity * dt
		}
		if win.Pressed(pixelgl.KeyQ) {
			sp.Angle += kin.AngularVelocity * dt
		}
	}

	return nil
}

// Render
func (cs *KeybaordControlSystem) Render(win *pixelgl.Window, dt float64) error {
	return nil
}

// AddEntity
func (cs *KeybaordControlSystem) AddEntity(es ...*ecs.Entity) error {
	cs.controlEntities = append(cs.controlEntities, es...)
	return nil
}

// RemoveEntity
func (cs *KeybaordControlSystem) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(cs.controlEntities, e)
		if err != nil {
			return err
		}
		cs.controlEntities = newEntries
	}
	return nil
}
