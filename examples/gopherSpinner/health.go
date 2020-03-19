package main

import (
	"fmt"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// HTAG const to hold the Health tag
	HTAG = "health"
)

// Health implements the component interface for a health system
type Health struct {
	tag   string
	Value float64
	Max   float64
}

// NewHealth returns a new health component with a given current and max health
func NewHealth(value, max float64) ecs.Component {
	return &Health{
		tag:   HTAG,
		Value: value,
		Max:   max,
	}
}

// GetHealth returns the actual Health struct implemnting the component for a given entity
func GetHealth(e *ecs.Entity) (*Health, error) {
	comp, err := e.Query(HTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Health), nil
}

// Tag returns the tag for this component
func (h *Health) Tag() string {
	return h.tag
}

func (h *Health) String() string {
	return fmt.Sprintf("%v/%v", h.Value, h.Max)
}

// Increase increases the health by a given amount up to the max health
func (h *Health) Increase(value float64) {
	h.Value += value
	if h.Value >= h.Max {
		h.Value = h.Max
	}
}

// Decrease decreases the health by a given amount down to 0
func (h *Health) Decrease(value float64) {
	h.Value -= value
	if h.Value < 0 {
		h.Value = 0
	}
}
