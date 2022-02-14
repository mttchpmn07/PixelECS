package systems

import (
	"log"

	"github.com/faiface/pixel/pixelgl"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/messenger"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

// SUserInput stores information for batch renderer system
type SUserInput struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
	m               messenger.Messenger
}

// NewSUserInput constructs a SUserInput from a varidact list of entities
func NewSUserInput(m messenger.Messenger, es ...*ecs.Entity) (ecs.System, error) {
	ui := &SUserInput{
		tag:             KITAG,
		controlEntities: []*ecs.Entity{},
		comps:           []string{},
		m:               m,
	}
	err := ui.AddEntity(es...)

	if err != nil {
		return nil, err
	}
	return ui, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ui *SUserInput) GetComponents() []string {
	return ui.comps
}

func (ui *SUserInput) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)
	//dt := (*args[1].(*float64))

	cam, err := components.GetCCamera(ui.controlEntities[0])
	if err != nil {
		return err
	}
	if win.JustPressed(pixelgl.MouseButton1) {
		err := ui.m.Broadcast("leftMoustClicked", cam.Translation.Unproject(win.MousePosition()))
		if err != nil {
			log.Println(err)
		}
	}
	if win.JustPressed(pixelgl.MouseButton2) {
		err := ui.m.Broadcast("rightMoustClicked", cam.Translation.Unproject(win.MousePosition()))
		if err != nil {
			log.Println(err)
		}
	}
	if win.JustPressed(pixelgl.MouseButton3) {
		err := ui.m.Broadcast("middleMouseClicked", cam.Translation.Unproject(win.MousePosition()))
		if err != nil {
			log.Println(err)
		}
	}
	if win.MouseScroll().Y != 0 {
		err := ui.m.Broadcast("mouseScroll", win.MouseScroll().Y)
		if err != nil {
			log.Println(err)
		}
	}
	if win.Pressed(pixelgl.KeyW) || win.Pressed(pixelgl.KeyUp) {
		err := ui.m.Broadcast("upKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if win.Pressed(pixelgl.KeyA) || win.Pressed(pixelgl.KeyLeft) {
		err := ui.m.Broadcast("leftKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if win.Pressed(pixelgl.KeyS) || win.Pressed(pixelgl.KeyDown) {
		err := ui.m.Broadcast("downKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}
	if win.Pressed(pixelgl.KeyD) || win.Pressed(pixelgl.KeyRight) {
		err := ui.m.Broadcast("rightKeyPressed", nil)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

// AddEntity adds any number of entities to this system
func (ui *SUserInput) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(ui, es...)
	if err != nil {
		return err
	}
	ui.controlEntities = append(ui.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (ui *SUserInput) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ui.controlEntities, e)
		if err != nil {
			return err
		}
		ui.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (ui *SUserInput) Tag() string {
	return ui.tag
}
