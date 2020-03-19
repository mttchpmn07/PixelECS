package components

import (
	"fmt"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// KTAG const to hold the Location tag
	KTAG = "kenetics"
)

// CKenetics component for storing kinetic information of an entity
type CKenetics struct {
	tag string

	Speed           float64
	AngularVelocity float64
}

// NewCKenetics returns a new CKenetics component with a given starting speed and angularVelocity
func NewCKenetics(speed, av float64) ecs.Component {
	return &CKenetics{
		tag:             KTAG,
		Speed:           speed,
		AngularVelocity: av,
	}
}

// GetCKenetics returns the actual CKenetics struct implmenting the component for a given entity
func GetCKenetics(e *ecs.Entity) (*CKenetics, error) {
	comp, err := e.Query(KTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CKenetics), nil
}

// Tag returns the tag for this component
func (k *CKenetics) Tag() string {
	return k.tag
}

func (k *CKenetics) String() string {
	return fmt.Sprintf("%v", k.tag)
}
