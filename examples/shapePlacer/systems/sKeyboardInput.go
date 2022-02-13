package systems

import (
	"log"

	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
)

// SKeybaordInput stores information for batch renderer system
type SKeybaordInput struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger
	keys            []pixelgl.Button
	keyBools        map[pixelgl.Button]bool
}

// NewSKeybaordInput constructs a SKeybaordInput from a varidact list of entities
func NewSKeybaordInput(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	ki := &SKeybaordInput{
		tag:             KITAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
		keys: []pixelgl.Button{
			pixelgl.KeyW,
			pixelgl.KeyA,
			pixelgl.KeyS,
			pixelgl.KeyD,
			pixelgl.KeyUp,
			pixelgl.KeyLeft,
			pixelgl.KeyDown,
			pixelgl.KeyRight,
		},
		keyBools: map[pixelgl.Button]bool{},
	}
	err := ki.AddEntity(es...)

	if err != nil {
		return nil, err
	}
	return ki, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ki *SKeybaordInput) GetComponents() []string {
	return ki.comps
}

func (ki *SKeybaordInput) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))

	for _, key := range ki.keys {
		ki.keyBools[key] = win.Pressed(key)
	}

	if ki.keyBools[pixelgl.KeyW] || ki.keyBools[pixelgl.KeyUp] {
		err := ki.m.Broadcast("upKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if ki.keyBools[pixelgl.KeyA] || ki.keyBools[pixelgl.KeyLeft] {
		err := ki.m.Broadcast("leftKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if ki.keyBools[pixelgl.KeyS] || ki.keyBools[pixelgl.KeyDown] {
		err := ki.m.Broadcast("downKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if ki.keyBools[pixelgl.KeyD] || ki.keyBools[pixelgl.KeyRight] {
		err := ki.m.Broadcast("rightKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (ki *SKeybaordInput) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(ki, es...)
	if err != nil {
		return err
	}
	ki.controlEntities = append(ki.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (ki *SKeybaordInput) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ki.controlEntities, e)
		if err != nil {
			return err
		}
		ki.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (ki *SKeybaordInput) Tag() string {
	return ki.tag
}
