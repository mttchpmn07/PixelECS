package systems

import (
	"time"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// ARTAG SAnimator tag
	ARTAG = "animator"
)

// SAnimator stores information for animation system
type SAnimator struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSAnimator constructs a SAnimator from a varidact list of entities
func NewSAnimator(es ...*ecs.Entity) (ecs.System, error) {
	ar := &SAnimator{
		tag:             ARTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.ATAG,
			components.SPTAG,
		},
	}
	err := ar.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ar *SAnimator) GetComponents() []string {
	return ar.comps
}

// Update updates the state of each animation
func (ar *SAnimator) Update(args ...interface{}) error {
	for _, e := range ar.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		an, err := components.GetCAnimation(e)
		if err != nil {
			return err
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

// AddEntity adds any number of entities to this system
func (ar *SAnimator) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(ar, es...)
	if err != nil {
		return err
	}
	ar.controlEntities = append(ar.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
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

// Tag getter for tag
func (ar *SAnimator) Tag() string {
	return ar.tag
}
