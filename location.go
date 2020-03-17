package pixelecs

import (
	"fmt"

	"github.com/faiface/pixel"
)

const (
	LTAG = "Location"
)

// Location
type Location struct {
	tag string
	Loc pixel.Vec
}

// NewLocation
func NewLocation(x, y float64) Component {
	return &Location{
		tag: LTAG,
		Loc: pixel.V(x, y),
	}
}

// GetLocation
func GetLocation(e *Entity) (*Location, error) {
	comp, err := e.Query(LTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Location), nil
}

func (l *Location) String() string {
	return fmt.Sprintf("%v", l.tag)
}

// Tag
func (l *Location) Tag() string {
	return l.tag
}
