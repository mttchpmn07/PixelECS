package systems

import (
	"github.com/Tarliton/collision2d"
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/entities"
)

// SShapePlacer stores information for batch renderer system
type SShapePlacer struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger

	shapeQue    *[]*ecs.Entity
	callbacks   map[string]func(contents interface{})
	ClearShapes bool
}

var shapeQue []*ecs.Entity

// NewSShapePlacer constructs a SShapePlacer from a varidact list of entities
func NewSShapePlacer(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	shapeQue = []*ecs.Entity{}
	sp := &SShapePlacer{
		tag:             SPTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		shapeQue:        &shapeQue,
		callbacks:       map[string]func(contents interface{}){},
		ClearShapes:     false,
	}
	err := sp.AddEntity(es...)
	sp.initSShapePlacerCallbacks()

	if err != nil {
		return nil, err
	}
	return sp, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (sp *SShapePlacer) GetComponents() []string {
	return sp.comps
}

func (sp *SShapePlacer) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))

	if sp.ClearShapes {
		sp.ClearShapes = false
		sp.m.Broadcast("clearShapes", nil)
	}

	for _, s := range shapeQue {
		sp.m.Broadcast("addShape", s)
	}
	shapeQue = shapeQue[:0]

	return nil
}

// AddEntity adds any number of entities to this system
func (sp *SShapePlacer) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(sp, es...)
	if err != nil {
		return err
	}
	sp.controlEntities = append(sp.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (sp *SShapePlacer) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(sp.controlEntities, e)
		if err != nil {
			return err
		}
		sp.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (sp *SShapePlacer) Tag() string {
	return sp.tag
}

func (sp *SShapePlacer) HandleBroadcast(key string, content interface{}) {
	sp.callbacks[key](content)
}

func (sp *SShapePlacer) initSShapePlacerCallbacks() {
	sp.callbacks["leftMoustClicked"] = sp.createSquare
	sp.callbacks["rightMoustClicked"] = sp.createTriangle
	sp.callbacks["middleMouseClicked"] = sp.createCircle
	sp.callbacks["spaceKeyPressed"] = sp.clearShapes
	for key := range sp.callbacks {
		sp.m.Subscribe(key, sp)
	}
}

func (sp *SShapePlacer) createSquare(content interface{}) {
	vec := content.(pixel.Vec)

	square, err := entities.NewSquare(collision2d.NewVector(vec.X, vec.Y), 0, 100)
	if err != nil {
		return
	}
	shapeQue = append(shapeQue, square)
}

func (sp *SShapePlacer) createTriangle(content interface{}) {
	vec := content.(pixel.Vec)

	triangle, err := entities.NewTriangle(collision2d.NewVector(vec.X, vec.Y), 0, 100)
	if err != nil {
		return
	}
	shapeQue = append(shapeQue, triangle)
}

func (sp *SShapePlacer) createCircle(content interface{}) {
	vec := content.(pixel.Vec)

	circle, err := entities.NewCircle(collision2d.NewVector(vec.X, vec.Y), 50)
	if err != nil {
		return
	}
	shapeQue = append(shapeQue, circle)
}

func (sp *SShapePlacer) clearShapes(content interface{}) {
	sp.ClearShapes = true
}
