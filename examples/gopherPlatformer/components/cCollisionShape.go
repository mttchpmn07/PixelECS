package components

import (
	"fmt"

	"github.com/faiface/pixel/imdraw"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// CSTAG CCollisionShape tag
	CSTAG = "collisionshape"
)

// CCollisionShape component storing the colliding shape of an entity.
type CCollisionShape struct {
	tag string

	Shape CollisionShape
}

// NewCCollisionShape constructs a CCollisionShape component from a CollisionShape.
func NewCCollisionShape(shape CollisionShape) ecs.Component {
	return &CCollisionShape{
		tag:   CSTAG,
		Shape: shape,
	}
}

// GetCCollisionShape returns the actual struct implmenting the component for a given entity
func GetCCollisionShape(e *ecs.Entity) (*CCollisionShape, error) {
	comp, err := e.Query(CSTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCollisionShape), nil
}

// Tag getter for tag
func (cp *CCollisionShape) Tag() string {
	return cp.tag
}

func (cp *CCollisionShape) String() string {
	return fmt.Sprintf("%v Shape(%v) : %v", cp.tag, cp.Shape.Type(), cp.Shape.Anchor())
}

// Collides checks if two shapes collide
func (cp *CCollisionShape) Collides(other *CCollisionShape) bool {
	return cp.Shape.Collides(other.Shape)
}

// Render returns the IMDraw to be drawn
func (cp *CCollisionShape) Render() *imdraw.IMDraw {
	return cp.Shape.Render()
}
