package entities

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/shapePlacer/components"
)

func NewCamera(pos pixel.Vec, speed, zoom, zoomSpeed float64) (*ecs.Entity, error) {
	camera, err := ecs.NewEntity()
	if err != nil {
		return nil, err
	}

	cam := components.NewCCamera(pos, speed, zoom, zoomSpeed)
	err = camera.Add(cam)
	if err != nil {
		return nil, err
	}

	return camera, nil
}
