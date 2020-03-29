package systems

import (
	"time"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// ARTAG const to hold the SRenderer tag
	ARTAG = "animator"
)

// SAnimator Sprite Render System
type SAnimator struct {
	tag string

	controlEntities []*ecs.Entity
}

// NewSAnimator returns a new sprite render system with a give list of entities attached via a variadic function call
func NewSAnimator(es ...*ecs.Entity) (ecs.System, error) {
	ar := &SAnimator{
		tag:             ARTAG,
		controlEntities: []*ecs.Entity{},
	}
	err := ar.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

// Update draws sprite for each associated entity
func (ar *SAnimator) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	for _, e := range ar.controlEntities {
		an, err := components.GetCAnimation(e)
		if err != nil {
			return err
		}
		if !an.Render {
			continue
		}

		seq := an.Sequences[an.Current]
		frameInterval := float64(time.Second) / seq.SampleRate

		if time.Since(an.LastFrameChange) >= time.Duration(frameInterval) {
			an.Finished = seq.NextFrame()
			an.LastFrameChange = time.Now()
		}
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (ar *SAnimator) AddEntity(es ...*ecs.Entity) error {
	ar.controlEntities = append(ar.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (ar *SAnimator) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ar.controlEntities, e)
		if err != nil {
			return err
		}
		ar.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (ar *SAnimator) Tag() string {
	return ar.tag
}
