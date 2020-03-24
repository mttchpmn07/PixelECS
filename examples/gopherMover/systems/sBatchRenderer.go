package systems

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

/*
Issue here... can't render different sprites at different z with batch without splitting some up...


I got it... I can redraw and clear the same batch for each layer and still be >>> more effecient then multiple single draw commands

1

You could do something like this:

Go through list of sprites from back to front

1.1 sprite already rendered -> next

1.2 check if 2d sprite extends collide with any sprite in the current batch or with any sprite coming before this one that is not yet marked as rendered (partly overlapping)

TRUE: continue with next sprite; rendering the current sprite with the current batch may result in sorting issues

FALSE: add sprite to batch and mark as rendered (eg: flag)

render and clear current batches

still any sprites not flagged as rendered? Resume at 1.
This will lead to an error free result with heavily reduced draw calls. Note: For multiple materials you need to maintain multiple sprite batches during this loop but the algorithm basically stays the same.
*/

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
