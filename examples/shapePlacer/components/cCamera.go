package components

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CCamera struct {
	tag string

	Translation pixel.Matrix
	Pos         pixel.Vec
	Speed       float64
	Zoom        float64
	ZoomSpeed   float64
}

func (cc *CCamera) String() string {
	return "camera"
}

func (cc *CCamera) Tag() string {
	return cc.tag
}

func NewCCamera(pos pixel.Vec, speed, zoom, zoomSpeed float64) ecs.Component {
	return &CCamera{
		tag: CTAG,

		Pos:       pos,
		Speed:     speed,
		Zoom:      zoom,
		ZoomSpeed: zoomSpeed,
	}
}

func GetCCamera(e *ecs.Entity) (*CCamera, error) {
	comp, err := e.Query(CTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCamera), nil
}
