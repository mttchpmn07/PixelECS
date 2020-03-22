package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// BRTAG const to hold the SBatchRenderer tag
	BRTAG = "batchrenderer"
)

// SBatchRenderer Sprite Render System
type SBatchRenderer struct {
	tag string

	controlEntities []*ecs.Entity
}

// NewSBatchRenderer returns a new sprite render system with a give list of entities attached via a variadic function call
func NewSBatchRenderer(es ...*ecs.Entity) (ecs.System, error) {
	ba := &SBatchRenderer{
		tag:             BRTAG,
		controlEntities: []*ecs.Entity{},
	}
	err := ba.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ba, nil
}

// Update draws sprite for each associated entity
func (ba *SBatchRenderer) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)

	// cheap Set so we don't redraw the same batch
	var exists = struct{}{}
	batches := map[*pixel.Batch]struct{}{}
	for _, e := range ba.controlEntities {
		ba, err := components.GetCBatchAsset(e)
		if err != nil {
			return err
		}
		_, c := batches[ba.Batch]
		if !c {
			batches[ba.Batch] = exists
		}
	}
	for b := range batches {
		b.Draw(win)
		b.Clear()
	}
	// for b := range batches {
	// 	b.Clear()
	// }
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (ba *SBatchRenderer) AddEntity(es ...*ecs.Entity) error {
	ba.controlEntities = append(ba.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (ba *SBatchRenderer) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ba.controlEntities, e)
		if err != nil {
			return err
		}
		ba.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (ba *SBatchRenderer) Tag() string {
	return ba.tag
}
