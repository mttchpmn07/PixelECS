package systems

import (
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// CPRTAG const to hold the SCollisionTracker tag
	CPRTAG = "collisionpolyrenderer"
)

// SCollisionPolyRenderer Sprite Render System
type SCollisionPolyRenderer struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSCollisionPolyRenderer returns a new sprite render system with a give list of entities attached via a variadic function call
func NewSCollisionPolyRenderer(es ...*ecs.Entity) (ecs.System, error) {
	cpr := &SCollisionPolyRenderer{
		tag:             CPRTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{components.CPTAG},
	}
	err := cpr.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cpr, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (cpr *SCollisionPolyRenderer) GetComponents() []string {
	return cpr.comps
}

// Update draws sprite for each associated entity
func (cpr *SCollisionPolyRenderer) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	for _, e := range cpr.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		cp, err := components.GetCCollisionPoly(e)
		if err != nil {
			return err
		}
		poly := imdraw.New(nil)
		for _, p := range cp.Points {
			poly.Push(p.Add(cp.Anchor.Loc))
		}
		poly.Polygon(2)
		poly.Draw(win)
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (cpr *SCollisionPolyRenderer) AddEntity(es ...*ecs.Entity) error {
	cpr.controlEntities = append(cpr.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (cpr *SCollisionPolyRenderer) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(cpr.controlEntities, e)
		if err != nil {
			return err
		}
		cpr.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (cpr *SCollisionPolyRenderer) Tag() string {
	return cpr.tag
}
