package systems

import (
	"log"

	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
)

// SUsuerControl stores information for batch renderer system
type SUsuerControl struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger
	callbacks       map[string]func(contents interface{})
}

// NewSUsuerControl constructs a SUsuerControl from a varidact list of entities
func NewSUsuerControl(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	uc := &SUsuerControl{
		tag:             UCTAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		callbacks:       map[string]func(contents interface{}){},
	}
	err := uc.AddEntity(es...)

	if err != nil {
		return nil, err
	}
	return uc, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (uc *SUsuerControl) GetComponents() []string {
	return uc.comps
}

var (
	keys     []pixelgl.Button
	keyBools map[pixelgl.Button]bool
)

func init() {
	keys = []pixelgl.Button{
		pixelgl.KeyA,
		pixelgl.KeyW,
		pixelgl.KeyS,
		pixelgl.KeyD,
		pixelgl.KeyLeft,
		pixelgl.KeyRight,
		pixelgl.KeyUp,
		pixelgl.KeyDown,
		pixelgl.KeyQ,
		pixelgl.KeyE,
		pixelgl.KeySpace,
		pixelgl.MouseButton1,
	}
	keyBools = map[pixelgl.Button]bool{}
}

func updateKeys(win *pixelgl.Window) {
	for _, key := range keys {
		keyBools[key] = win.JustPressed(key)
	}
}

func (uc *SUsuerControl) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))
	updateKeys(win)

	if keyBools[pixelgl.MouseButton1] {
		err := uc.m.Broadcast("createShape", win.MousePosition())
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (uc *SUsuerControl) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(uc, es...)
	if err != nil {
		return err
	}
	uc.controlEntities = append(uc.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (uc *SUsuerControl) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(uc.controlEntities, e)
		if err != nil {
			return err
		}
		uc.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (uc *SUsuerControl) Tag() string {
	return uc.tag
}

func (uc *SUsuerControl) HandleBroadcast(key string, content interface{}) {
	uc.callbacks[key](content)
}
