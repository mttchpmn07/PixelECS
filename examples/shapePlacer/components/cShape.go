package components

import (
	"github.com/Tarliton/collision2d"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CShape struct {
	tag string

	Shape collision2d.Polygon
}

func (cs *CShape) String() string {
	return "shape"
}

func (cs *CShape) Tag() string {
	return cs.tag
}

func NewCShape(shape collision2d.Polygon) ecs.Component {
	return &CShape{
		tag:   STAG,
		Shape: shape,
	}
}

func GetCShape(e *ecs.Entity) (*CShape, error) {
	comp, err := e.Query(STAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CShape), nil
}
