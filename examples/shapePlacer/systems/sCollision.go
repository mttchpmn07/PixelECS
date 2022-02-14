package systems

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SCollision stores information for batch renderer system
type SCollision struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger

	callbacks map[string]func(contents interface{})
}

// NewSCollision constructs a SCollision from a varidact list of entities
func NewSCollision(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	c := &SCollision{
		tag:             CLTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		callbacks:       map[string]func(contents interface{}){},
	}
	err := c.AddEntity(es...)
	c.initSCollisionCallbacks()

	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (c *SCollision) GetComponents() []string {
	return c.comps
}

func (c *SCollision) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))

	for _, e1 := range c.controlEntities {
		s1, err := components.GetCollisionShape(e1)
		if err != nil {
			return err
		}
		if !s1.Active {
			continue
		}
		for _, e2 := range c.controlEntities {
			if e1 == e2 {
				continue
			}
			s2, err := components.GetCollisionShape(e2)
			if err != nil {
				return err
			}
			if !s2.Active {
				continue
			}
			if hits, _ := s1.Collides(s2); hits {
				rp1, err := components.GetCRenderProperties(e1)
				if err != nil {
					return err
				}
				rp2, err := components.GetCRenderProperties(e2)
				if err != nil {
					return err
				}
				rp1.Color = pixel.RGB(1, 0, 0)
				rp2.Color = pixel.RGB(1, 0, 0)
			}
		}
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (c *SCollision) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(c, es...)
	if err != nil {
		return err
	}
	c.controlEntities = append(c.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (c *SCollision) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(c.controlEntities, e)
		if err != nil {
			return err
		}
		c.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (c *SCollision) Tag() string {
	return c.tag
}

func (c *SCollision) HandleBroadcast(key string, content interface{}) {
	c.callbacks[key](content)
}

func (c *SCollision) initSCollisionCallbacks() {
	c.callbacks["addShape"] = c.addShape
	c.callbacks["clearShapes"] = c.clearShapes
	for key := range c.callbacks {
		c.m.Subscribe(key, c)
	}
}

func (c *SCollision) addShape(content interface{}) {
	c.AddEntity(content.(*ecs.Entity))
}

func (c *SCollision) clearShapes(content interface{}) {
	for len(c.controlEntities) > 0 {
		c.RemoveEntity(c.controlEntities[0])
	}
}
