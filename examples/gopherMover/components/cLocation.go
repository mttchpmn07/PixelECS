package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// LTAG CLocation tag
	LTAG = "location"
)

// CLocation component interface for storing the location of an entity
type CLocation struct {
	tag string
	Loc pixel.Vec
	Z   int
}

// NewCLocation constructs a CLocation component with a given starting x, y, and z.
// The x and y are floating points, while the z is an integers used to determine the drawing order.
func NewCLocation(x, y float64, z int) ecs.Component {
	return &CLocation{
		tag: LTAG,
		Loc: pixel.V(x, y),
		Z:   z,
	}
}

// GetCLocation returns the actual struct implmenting the component for a given entity
func GetCLocation(e *ecs.Entity) (*CLocation, error) {
	comp, err := e.Query(LTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CLocation), nil
}

// Tag getter for tag
func (l *CLocation) Tag() string {
	return l.tag
}

func (l *CLocation) String() string {
	return fmt.Sprintf("%v", l.tag)
}
