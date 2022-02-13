package systems

import (
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SCamera stores information for batch renderer system
type SCamera struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger

	shapeQue  *[]*ecs.Entity
	callbacks map[string]func(contents interface{})

	moveUp    bool
	moveLeft  bool
	moveDown  bool
	moveRight bool
	zoom      float64
}

// NewSCamera constructs a SCamera from a varidact list of entities
func NewSCamera(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	shapeQue = []*ecs.Entity{}
	c := &SCamera{
		tag:             CTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		shapeQue:        &shapeQue,
		callbacks:       map[string]func(contents interface{}){},
		moveUp:          false,
		moveLeft:        false,
		moveDown:        false,
		moveRight:       false,
		zoom:            0.0,
	}
	err := c.AddEntity(es...)
	c.initSCameraCallbacks()

	if err != nil {
		return nil, err
	}
	return c, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (c *SCamera) GetComponents() []string {
	return c.comps
}

func (c *SCamera) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	dt := (*args[1].(*float64))

	for _, e := range c.controlEntities {
		cam, err := components.GetCCamera(e)
		if err != nil {
			return err
		}
		cam.Translation = pixel.IM.Scaled(cam.Pos, cam.Zoom).Moved(win.Bounds().Center().Sub(cam.Pos))
		win.SetMatrix(cam.Translation)

		if c.moveUp {
			cam.Pos.Y += cam.Speed * dt
			c.moveUp = false
		}
		if c.moveLeft {
			cam.Pos.X -= cam.Speed * dt
			c.moveLeft = false
		}
		if c.moveDown {
			cam.Pos.Y -= cam.Speed * dt
			c.moveDown = false
		}
		if c.moveRight {
			cam.Pos.X += cam.Speed * dt
			c.moveRight = false
		}
		cam.Zoom *= math.Pow(cam.ZoomSpeed, c.zoom)
		c.zoom = 0.0
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (c *SCamera) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(c, es...)
	if err != nil {
		return err
	}
	c.controlEntities = append(c.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (c *SCamera) RemoveEntity(es ...*ecs.Entity) error {
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
func (c *SCamera) Tag() string {
	return c.tag
}

func (c *SCamera) HandleBroadcast(key string, content interface{}) {
	c.callbacks[key](content)
}

func (c *SCamera) initSCameraCallbacks() {
	c.callbacks["upKeyPressed"] = c.camUp
	c.callbacks["leftKeyPressed"] = c.camLeft
	c.callbacks["downKeyPressed"] = c.camDown
	c.callbacks["rightKeyPressed"] = c.camRight
	c.callbacks["mouseScroll"] = c.camZoom
	for key := range c.callbacks {
		c.m.Subscribe(key, c)
	}

}

func (c *SCamera) camUp(content interface{}) {
	c.moveUp = true
}

func (c *SCamera) camLeft(content interface{}) {
	c.moveLeft = true
}

func (c *SCamera) camDown(content interface{}) {
	c.moveDown = true
}

func (c *SCamera) camRight(content interface{}) {
	c.moveRight = true
}

func (c *SCamera) camZoom(content interface{}) {
	c.zoom = content.(float64)
}
