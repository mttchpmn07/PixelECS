package components

import (
	"image/color"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CRenderProperties struct {
	tag string

	Active bool
	Color  color.Color
}

func (rp *CRenderProperties) String() string {
	return "render properties"
}

func (rp *CRenderProperties) Tag() string {
	return rp.tag
}

func NewCRenderProperties(active bool) ecs.Component {
	return &CRenderProperties{
		tag:    RPTAG,
		Active: active,
		Color:  color.White,
	}
}

func GetCRenderProperties(e *ecs.Entity) (*CRenderProperties, error) {
	comp, err := e.Query(RPTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CRenderProperties), nil
}
