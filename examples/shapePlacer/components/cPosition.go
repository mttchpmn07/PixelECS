package components

import (
	"github.com/Tarliton/collision2d"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CPosition struct {
	tag string

	Coord collision2d.Vector
	Z     int
}

func (cp *CPosition) String() string {
	return "position"
}

func (cp *CPosition) Tag() string {
	return cp.tag
}

func NewCPosition() ecs.Component {
	return &CPosition{
		tag: PTAG,
	}
}

func GetCPosition(e *ecs.Entity) (*CPosition, error) {
	comp, err := e.Query(PTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CPosition), nil
}
