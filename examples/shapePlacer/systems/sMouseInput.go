package systems

import (
	"log"

	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SMouseInput stores information for batch renderer system
type SMouseInput struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger
	callbacks       map[string]func(contents interface{})
	keys            []pixelgl.Button
	keyBools        map[pixelgl.Button]bool
}

// NewSMouseInput constructs a SMouseInput from a varidact list of entities
func NewSMouseInput(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	mi := &SMouseInput{
		tag:             MITAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		callbacks:       map[string]func(contents interface{}){},
		keys: []pixelgl.Button{
			pixelgl.MouseButton1,
		},
		keyBools: map[pixelgl.Button]bool{},
	}
	err := mi.AddEntity(es...)

	if err != nil {
		return nil, err
	}
	return mi, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (mi *SMouseInput) GetComponents() []string {
	return mi.comps
}

func (mi *SMouseInput) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))

	for _, key := range mi.keys {
		mi.keyBools[key] = win.JustPressed(key)
	}

	for _, e := range mi.controlEntities {
		cam, err := components.GetCCamera(e)
		if err != nil {
			return err
		}
		if mi.keyBools[pixelgl.MouseButton1] {
			err := mi.m.Broadcast("lefMoustClicked", cam.Translation.Unproject(win.MousePosition()))
			if err != nil {
				log.Println(err)
			}
		}
		if win.MouseScroll().Y != 0 {
			err := mi.m.Broadcast("mouseScroll", win.MouseScroll().Y)
			if err != nil {
				log.Println(err)
			}
		}
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (mi *SMouseInput) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(mi, es...)
	if err != nil {
		return err
	}
	mi.controlEntities = append(mi.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (mi *SMouseInput) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(mi.controlEntities, e)
		if err != nil {
			return err
		}
		mi.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (mi *SMouseInput) Tag() string {
	return mi.tag
}
