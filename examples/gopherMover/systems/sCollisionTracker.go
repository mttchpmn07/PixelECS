package systems

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// CTTAG const to hold the SCollisionTracker tag
	CTTAG = "collisiontracker"
)

// SCollisionTracker Sprite Render System
type SCollisionTracker struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSCollisionTracker returns a new sprite render system with a give list of entities attached via a variadic function call
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

// Update draws sprite for each associated entity
func (ct *SCollisionTracker) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	for _, e1 := range ct.controlEntities {
		for _, e2 := range ct.controlEntities {
			sp, err := components.GetCProperties(e2)
			if err != nil {
				return err
			}
			if e1 == e2 || (e1.Class != "gopher" && e2.Class != "fly") || !sp.Active {
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
				sp.Active = false
			}
		}
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (ct *SCollisionTracker) AddEntity(es ...*ecs.Entity) error {
	ct.controlEntities = append(ct.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
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

// Tag returns the tag for this system
func (ct *SCollisionTracker) Tag() string {
	return ct.tag
}
