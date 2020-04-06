package systems

import (
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
)

const (
	// CPRTAG SCollisionTracker tag
	CPRTAG = "collisionpolyrenderer"
)

// SCollisionPolyRenderer stores information for collision poly rendering system
type SCollisionPolyRenderer struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSCollisionPolyRenderer constructs a SCollisionPolyRenderer from a varidact list of entities
func NewSCollisionPolyRenderer(es ...*ecs.Entity) (ecs.System, error) {
	cpr := &SCollisionPolyRenderer{
		tag:             CPRTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{components.CSTAG},
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

// Update draws each active collision polygon
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
		cp, err := components.GetCCollisionShape(e)
		if err != nil {
			return err
		}
		poly := cp.Render()
		poly.Draw(win)
	}
	return nil
}

// AddEntity adds any number of entities to this system
func (cpr *SCollisionPolyRenderer) AddEntity(es ...*ecs.Entity) error {
	cpr.controlEntities = append(cpr.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
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

// Tag getter for tag
func (cpr *SCollisionPolyRenderer) Tag() string {
	return cpr.tag
}
