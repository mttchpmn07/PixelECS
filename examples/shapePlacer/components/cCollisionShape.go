package components

import (
	"github.com/Tarliton/collision2d"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	HitA = 1 << iota
	HitB
	HitC
	HitD
	NoHit = 0
)

type CCollisionShape struct {
	tag string

	Shape  interface{}
	Type   string
	HitMap uint8
	Draw   bool
	Active bool
}

func (cs *CCollisionShape) String() string {
	return "shape"
}

func (cs *CCollisionShape) Tag() string {
	return cs.tag
}

func NewCollisionShape(shape interface{}, t string, hitMap uint8, draw, active bool) ecs.Component {
	return &CCollisionShape{
		tag:    STAG,
		Shape:  shape,
		Type:   t,
		HitMap: hitMap,
		Draw:   draw,
		Active: active,
	}
}

func GetCollisionShape(e *ecs.Entity) (*CCollisionShape, error) {
	comp, err := e.Query(STAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCollisionShape), nil
}

func (shape *CCollisionShape) Collides(other *CCollisionShape) (bool, collision2d.Response) {
	switch shape.Type {
	case "polygon":
		s1 := shape.Shape.(collision2d.Polygon)
		switch other.Type {
		case "polygon":
			s2 := other.Shape.(collision2d.Polygon)
			return collision2d.TestPolygonPolygon(s1, s2)
		case "circle":
			s2 := other.Shape.(collision2d.Circle)
			return collision2d.TestPolygonCircle(s1, s2)
		}
	case "circle":
		s1 := shape.Shape.(collision2d.Circle)
		switch other.Type {
		case "polygon":
			s2 := other.Shape.(collision2d.Polygon)
			return collision2d.TestCirclePolygon(s1, s2)
		case "circle":
			s2 := other.Shape.(collision2d.Circle)
			return collision2d.TestCircleCircle(s1, s2)
		}
	}
	return false, collision2d.NewResponse()
}
