package entities

import (
	"math"

	"github.com/Tarliton/collision2d"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

func NewSquare(position collision2d.Vector, angle, sideLength float64) (*ecs.Entity, error) {
	//offset := collision2d.NewVector(sideLength/2, sideLength/2)
	points := []float64{
		0, 0,
		sideLength, 0,
		sideLength, sideLength,
		0, sideLength,
	}
	return NewPolygon(position, collision2d.NewVector(0, 0), angle, points)
}

func NewTriangle(position collision2d.Vector, angle, sideLength float64) (*ecs.Entity, error) {
	height := math.Sqrt(math.Pow(sideLength, 2) - math.Pow(sideLength/2, 2))
	//offset := collision2d.NewVector(sideLength/2, height/2)
	points := []float64{
		0, 0,
		sideLength, 0,
		sideLength / 2, height,
	}
	return NewPolygon(position, collision2d.NewVector(0, 0), angle, points)
}

func NewCircle(position collision2d.Vector, radius float64) (*ecs.Entity, error) {
	circle, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	rp := components.NewCRenderProperties(true)
	err = circle.Add(rp)
	if err != nil {
		return nil, err
	}

	c := collision2d.NewCircle(position, radius)
	s := components.NewCollisionShape(c, "circle", components.HitA, true, true)
	err = circle.Add(s)
	if err != nil {
		return nil, err
	}

	return circle, nil
}

func NewPolygon(position, offset collision2d.Vector, angle float64, points []float64) (*ecs.Entity, error) {
	polygon, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	rp := components.NewCRenderProperties(true)
	err = polygon.Add(rp)
	if err != nil {
		return nil, err
	}

	poly := collision2d.NewPolygon(position, offset, 0, points)
	s := components.NewCollisionShape(poly, "polygon", components.HitA, true, true)
	err = polygon.Add(s)
	if err != nil {
		return nil, err
	}

	return polygon, nil
}
