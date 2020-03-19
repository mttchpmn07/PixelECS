package core

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
)

// System comment here
type System interface {
	Update(win *pixelgl.Window, dt float64) error
	AddEntity(es ...*Entity) error
	RemoveEntity(es ...*Entity) error
}

func StripEntity(ce []*Entity, e *Entity) ([]*Entity, error) {
	idx := -1
	for i, ent := range ce {
		if ent == e {
			idx = i
			break
		}
	}
	if idx == -1 {
		return []*Entity{}, fmt.Errorf("entity %v not found", e.ID)
	}
	l := len(ce)
	ce[idx] = ce[l-1]
	ce = ce[:l-1]
	return ce, nil
}
