package components

import ecs "github.com/mttchpmn07/PixelECS/core"

type CCamera struct {
	tag string
}

func (cc *CCamera) String() string {
	return "camera"
}

func (cc *CCamera) Tag() string {
	return cc.tag
}

func NewCCamera() ecs.Component {
	return &CCamera{
		tag: CTAG,
	}
}

func GetCCamera(e *ecs.Entity) (*CCamera, error) {
	comp, err := e.Query(CTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCamera), nil
}
