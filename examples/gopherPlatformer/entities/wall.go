package entities

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
)

// NewWall creates a new animated fly
func NewWall(x, y, width, height float64) (*ecs.Entity, error) {
	wall, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}
	loc := components.NewCLocation(x, y, 10)
	err = wall.Add(loc)
	if err != nil {
		return nil, err
	}

	cLoc := loc.(*components.CLocation)
	poly := components.NewPolygon(
		cLoc,
		pixel.V(-width/2, height/2),
		pixel.V(width/2, height/2),
		pixel.V(width/2, -height/2),
		pixel.V(-width/2, -height/2),
	)
	collision := components.NewCCollisionShape(poly)
	err = wall.Add(collision)
	if err != nil {
		return nil, err
	}

	sp := components.NewCProperties(0, 1, pixel.Rect{}, true, "wall")
	err = wall.Add(sp)
	if err != nil {
		return nil, err
	}
	return wall, nil
}
