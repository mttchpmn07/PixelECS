package systems

import (
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// SKCTAG SKeyboardController tag
	SKCTAG = "keyboardcontroller"
)

// SKeyboardController Keyboard Control system
type SKeyboardController struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSKeyboardController constructs a SKeyboardController from a varidact list of entities
func NewSKeyboardController(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &SKeyboardController{
		tag:             SKCTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.ATAG,
			components.LTAG,
			components.KTAG,
			components.SPTAG,
		},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (cs *SKeyboardController) GetComponents() []string {
	return cs.comps
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

// Update updates the state of each entity that is controlled by the keyboard
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

// AddEntity adds any number of entities to this system
func (cs *SKeyboardController) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(cs, es...)
	if err != nil {
		return err
	}
	cs.controlEntities = append(cs.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
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

// Tag getter for tag
func (cs *SKeyboardController) Tag() string {
	return cs.tag
}
