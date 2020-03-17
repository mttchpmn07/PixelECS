package pixelecs

import (
	"fmt"

	"github.com/faiface/pixel"
)

const (
	// LTAG const to hold the Location tag
	LTAG = "location"
)

// Location implements the component interface for a location system
type Location struct {
	tag string
	Loc pixel.Vec
}

// NewLocation returns a new location component with a given starting x and y
func NewLocation(x, y float64) Component {
	return &Location{
		tag: LTAG,
		Loc: pixel.V(x, y),
	}
}

// GetLocation returns the actual Location struct implmenting the component for a given entity
func GetLocation(e *Entity) (*Location, error) {
	comp, err := e.Query(LTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Location), nil
}

// Tag returns the tag for this component
func (l *Location) Tag() string {
	return l.tag
}

func (l *Location) String() string {
	return fmt.Sprintf("%v", l.tag)
}
