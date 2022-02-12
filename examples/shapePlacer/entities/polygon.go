package entities

import (
	"github.com/Tarliton/collision2d"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

func NewSquare(position collision2d.Vector, angle, sideLength float64) (*ecs.Entity, error) {
	//offset := collision2d.NewVector(0, 0)
	offset := collision2d.NewVector(sideLength/2, sideLength/2)
	//offset := collision2d.NewVector(sideLength, sideLength)
	points := []float64{0, 0, sideLength, 0, sideLength, sideLength, 0, sideLength}
	return NewPolygon(position, offset, angle, points)
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
	s := components.NewCShape(poly)
	err = polygon.Add(s)
	if err != nil {
		return nil, err
	}

	return polygon, nil
}
