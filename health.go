package pixelecs

import (
	"fmt"
)

const (
	HTAG = "Health"
)

type Health struct {
	tag   string
	Value float64
	Max   float64
}

// NewHealth
func NewHealth(value, max float64) Component {
	return &Health{
		tag:   HTAG,
		Value: value,
		Max:   max,
	}
}

// GetHealth
func GetHealth(e *Entity) (*Health, error) {
	comp, err := e.Query(HTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Health), nil
}

func (h *Health) String() string {
	return fmt.Sprintf("%v/%v", h.Value, h.Max)
}

func (h *Health) Tag() string {
	return h.tag
}

func (h *Health) Increase(value float64) {
	h.Value += value
	if h.Value >= h.Max {
		h.Value = h.Max
	}
}

func (h *Health) Decrease(value float64) {
	h.Value -= value
	if h.Value < 0 {
		h.Value = 0
	}
}
