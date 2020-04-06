package systems

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
)

const (
	// PTAG SPhysics tag
	PTAG = "physics"
)

// SPhysics stores information for animation system
type SPhysics struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSPhysics constructs a SPhysics from a varidact list of entities
func NewSPhysics(es ...*ecs.Entity) (ecs.System, error) {
	ph := &SPhysics{
		tag:             PTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.LTAG,
			components.KTAG,
		},
	}
	err := ph.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ph, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ph *SPhysics) GetComponents() []string {
	return ph.comps
}

// Update updates the state of each animation
func (ph *SPhysics) Update(args ...interface{}) error {
	dt := (*args[1].(*float64))
	for _, e := range ph.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		kin, err := components.GetCKenetics(e)
		if err != nil {
			return err
		}
		kin.Velocity = kin.Velocity.Add(kin.Acceleration.Scaled(dt)).Unit().Scaled(kin.Speed)
		loc.Loc = loc.Loc.Add(kin.Velocity.Scaled(dt))
	}
	return nil
}

// AddEntity adds any number of entities to this system
func (ph *SPhysics) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(ph, es...)
	if err != nil {
		return err
	}
	ph.controlEntities = append(ph.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (ph *SPhysics) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ph.controlEntities, e)
		if err != nil {
			return err
		}
		ph.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (ph *SPhysics) Tag() string {
	return ph.tag
}
