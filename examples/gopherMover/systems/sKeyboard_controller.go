package systems

import (
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// SKCTAG const to hold the SKeyboardController tag
	SKCTAG = "keyboardcontroller"
)

// SKeyboardController Keyboard Control system
type SKeyboardController struct {
	tag string

	controlEntities []*ecs.Entity
}

// NewSKeyboardController returns a new Keyboard Control System with all given entities attached
func NewSKeyboardController(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &SKeyboardController{
		tag:             SKCTAG,
		controlEntities: []*ecs.Entity{},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

var (
	keys     []pixelgl.Button
	keyBools map[pixelgl.Button]bool
)

func init() {
	keys = []pixelgl.Button{
		pixelgl.KeyA,
		pixelgl.KeyW,
		pixelgl.KeyS,
		pixelgl.KeyD,
		pixelgl.KeyLeft,
		pixelgl.KeyRight,
		pixelgl.KeyUp,
		pixelgl.KeyDown,
		pixelgl.KeyQ,
		pixelgl.KeyE,
		pixelgl.KeySpace,
	}
	keyBools = map[pixelgl.Button]bool{}
}

func updateKeys(win *pixelgl.Window) {
	for _, key := range keys {
		keyBools[key] = win.Pressed(key)
	}

}

// Update calculates next state for all components used by the system for each of its associated entities
func (cs *SKeyboardController) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	dt := (*args[1].(*float64))
	updateKeys(win)

	for _, e := range cs.controlEntities {
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		kin, err := components.GetCKenetics(e)
		if err != nil {
			return err
		}
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		an, err := components.GetCAnimation(e)
		if err != nil {
			return err
		}
		if keyBools[pixelgl.KeyA] || keyBools[pixelgl.KeyLeft] {
			if an.Current != "left" {
				an.SetSequence("left")
				sp.Angle *= -1
			}
			loc.Loc.X -= kin.Speed * dt
		}
		if keyBools[pixelgl.KeyD] || keyBools[pixelgl.KeyRight] {
			if an.Current != "right" {
				an.SetSequence("right")
				sp.Angle *= -1
			}
			loc.Loc.X += kin.Speed * dt
		}
		if keyBools[pixelgl.KeyS] || keyBools[pixelgl.KeyDown] {
			loc.Loc.Y -= kin.Speed * dt
		}
		if keyBools[pixelgl.KeyW] || keyBools[pixelgl.KeyUp] {
			loc.Loc.Y += kin.Speed * dt
		}
		if keyBools[pixelgl.KeyE] {
			sp.Angle -= kin.AngularVelocity * dt
		}
		if keyBools[pixelgl.KeyQ] {
			sp.Angle += kin.AngularVelocity * dt
		}
		if keyBools[pixelgl.KeySpace] {
			//
		}
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (cs *SKeyboardController) AddEntity(es ...*ecs.Entity) error {
	cs.controlEntities = append(cs.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (cs *SKeyboardController) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(cs.controlEntities, e)
		if err != nil {
			return err
		}
		cs.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (cs *SKeyboardController) Tag() string {
	return cs.tag
}
