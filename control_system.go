package pixelecs

import (
	"github.com/faiface/pixel/pixelgl"
)

// ControlSystem
type ControlSystem struct {
	controlEntities []*Entity

	userVelocity float64
}

// NewControlSystem
func NewControlSystem(velocity float64, es ...*Entity) System {
	cSystem := &ControlSystem{
		controlEntities: []*Entity{},
		userVelocity:    velocity,
	}
	cSystem.AddEntity(es...)
	return cSystem
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
func (cs *ControlSystem) AddEntity(es ...*Entity) error {
	for _, e := range es {
		cs.controlEntities = append(cs.controlEntities, e)
	}
	return nil
}

// RemoveEntity
func (cs *ControlSystem) RemoveEntity(es ...*Entity) error {
	for _, e := range es {
		newEntries, err := removeEntity(cs.controlEntities, e)
		if err != nil {
			return err
		}
		cs.controlEntities = newEntries
	}
	return nil
}
