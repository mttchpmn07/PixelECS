package systems

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// CTTAG SCollisionTracker tag
	CTTAG = "collisiontracker"
)

// SCollisionTracker stores information for the collision tracking system
type SCollisionTracker struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSCollisionTracker constructs a SCollisionTracker from a varidact list of entities
func NewSCollisionTracker(es ...*ecs.Entity) (ecs.System, error) {
	ct := &SCollisionTracker{
		tag:             CTTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.CPTAG,
			components.SPTAG,
		},
	}
	err := ct.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ct *SCollisionTracker) GetComponents() []string {
	return ct.comps
}

// Update checks for any valid collisions (for the moment it also resolves them, but I'd like to make that a seperate system)
func (ct *SCollisionTracker) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	for _, e1 := range ct.controlEntities {
		for _, e2 := range ct.controlEntities {
			if e1 == e2 {
				continue
			}
			sp1, err := components.GetCProperties(e1)
			if err != nil {
				return err
			}
			sp2, err := components.GetCProperties(e2)
			if err != nil {
				return err
			}
			if sp1.Class != "gopher" || sp2.Class != "fly" || !sp1.Active || !sp2.Active {
				continue
			}
			e1CP, err := components.GetCCollisionPoly(e1)
			if err != nil {
				return err
			}
			e2CP, err := components.GetCCollisionPoly(e2)
			if err != nil {
				return err
			}
			if e1CP.Collides(e2CP) {
				sp2.Active = false
			}
		}
	}
	return nil
}

// AddEntity adds any number of entities to this system
func (ct *SCollisionTracker) AddEntity(es ...*ecs.Entity) error {
	ct.controlEntities = append(ct.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (ct *SCollisionTracker) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ct.controlEntities, e)
		if err != nil {
			return err
		}
		ct.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (ct *SCollisionTracker) Tag() string {
	return ct.tag
}
