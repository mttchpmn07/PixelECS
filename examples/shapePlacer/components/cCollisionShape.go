package components

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CCollisionShape struct {
	tag string

	Shape interface{}
	Type  string
	Draw  bool
}

func (cs *CCollisionShape) String() string {
	return "shape"
}

func (cs *CCollisionShape) Tag() string {
	return cs.tag
}

func NewCollisionShape(shape interface{}, t string, draw bool) ecs.Component {
	return &CCollisionShape{
		tag:   STAG,
		Shape: shape,
		Type:  t,
		Draw:  draw,
	}
}

func GetCollisionShape(e *ecs.Entity) (*CCollisionShape, error) {
	comp, err := e.Query(STAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCollisionShape), nil
}
