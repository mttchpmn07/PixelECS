package systems

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SRender stores information for batch renderer system
type SRender struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSRender constructs a SRender from a varidact list of entities
func NewSRender(es ...*ecs.Entity) (ecs.System, error) {
	br := &SRender{
		tag:             SRTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.RPTAG,
			components.STAG,
		},
	}
	err := br.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return br, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (r *SRender) GetComponents() []string {
	return r.comps
}

func pixelPoint(point collision2d.Vector) pixel.Vec {
	return pixel.V(point.X, point.Y)
}

func renderShape(shape collision2d.Polygon) *imdraw.IMDraw {
	poly := imdraw.New(nil)
	for _, p := range shape.Points {
		poly.Push(pixelPoint(p.Add(shape.Pos).Sub(shape.Offset)))
	}
	poly.Polygon(2)
	return poly
}

func (r *SRender) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)

	for _, e := range r.controlEntities {
		rp, err := components.GetCRenderProperties(e)
		if err != nil {
			return err
		}
		if !rp.Active {
			continue
		}
		s, err := components.GetCShape(e)
		if err != nil {
			return err
		}
		poly := renderShape(s.Shape)
		poly.Draw(win)
	}
	return nil
}

// AddEntity adds any number of entities to this system
func (r *SRender) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(r, es...)
	if err != nil {
		return err
	}
	r.controlEntities = append(r.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (r *SRender) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(r.controlEntities, e)
		if err != nil {
			return err
		}
		r.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (r *SRender) Tag() string {
	return r.tag
}