package systems

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SRender stores information for batch renderer system
type SRender struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger
	callbacks       map[string]func(contents interface{})
}

// NewSRender constructs a SRender from a varidact list of entities
func NewSRender(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	r := &SRender{
		tag:             SRTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.RPTAG,
			components.STAG,
		},
		m:         m,
		callbacks: map[string]func(contents interface{}){},
	}
	err := r.AddEntity(es...)
	r.initSRenderCallbacks()

	if err != nil {
		return nil, err
	}
	return r, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (r *SRender) GetComponents() []string {
	return r.comps
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
		s, err := components.GetCollisionShape(e)
		if err != nil {
			return err
		}
		if !s.Draw {
			continue
		}
		shape := renderCollisionShape(s)
		shape.Draw(win)
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

func (r *SRender) HandleBroadcast(key string, content interface{}) {
	r.callbacks[key](content)
}

func (r *SRender) addShapeCallback(content interface{}) {
	r.AddEntity(content.(*ecs.Entity))
}

func (r *SRender) initSRenderCallbacks() {
	r.callbacks["addShape"] = r.addShapeCallback
	for key := range r.callbacks {
		r.m.Subscribe(key, r)
	}
}

func pixelPoint(point collision2d.Vector) pixel.Vec {
	return pixel.V(point.X, point.Y)
}

func renderCollisionShape(shape *components.CCollisionShape) *imdraw.IMDraw {
	draw := imdraw.New(nil)
	switch shape.Type {
	case "polygon":
		s := shape.Shape.(collision2d.Polygon)
		for _, p := range s.Points {
			draw.Push(pixelPoint(p.Add(s.Pos).Sub(s.Offset)))
		}
		draw.Polygon(2)
	case "circle":
		s := shape.Shape.(collision2d.Circle)
		draw.Push(pixelPoint(s.Pos))
		draw.Circle(s.R, 2)
	}
	return draw
}
