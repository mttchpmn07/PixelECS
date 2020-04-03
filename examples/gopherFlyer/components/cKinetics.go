package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// KTAG CKenetics
	KTAG = "kenetics"
)

// CKenetics component for storing kinetic information of an entity
type CKenetics struct {
	tag string

	Speed           float64
	AngularVelocity float64
	AccMag          float64
	Velocity        pixel.Vec
	Acceleration    pixel.Vec
}

// NewCKenetics constructs a CKenetics component from a starting speed, angularVelocity,
// acceleration magnitude, velocity vector, and acceleration vector
func NewCKenetics(speed, av, accMag float64, vel, acc pixel.Vec) ecs.Component {
	return &CKenetics{
		tag:             KTAG,
		Speed:           speed,
		AngularVelocity: av,
		AccMag:          accMag,
		Velocity:        vel,
		Acceleration:    acc,
	}
}

// GetCKenetics returns the struct implmenting the component for a given entity
func GetCKenetics(e *ecs.Entity) (*CKenetics, error) {
	comp, err := e.Query(KTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CKenetics), nil
}

// Tag getter for tag
func (k *CKenetics) Tag() string {
	return k.tag
}

func (k *CKenetics) String() string {
	return fmt.Sprintf("%v", k.tag)
}
