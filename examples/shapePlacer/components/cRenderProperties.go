package components

import (
	ecs "github.com/mttchpmn07/PixelECS/core"
)

type CRenderProperties struct {
	tag string

	Active bool
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
	}
}

func GetCRenderProperties(e *ecs.Entity) (*CRenderProperties, error) {
	comp, err := e.Query(RPTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CRenderProperties), nil
}
