package core

import (
	"fmt"

	"github.com/google/uuid"
)

// Entity basic entity struct (litterally just an id to find components with)
type Entity struct {
	ID uuid.UUID
}

// NewEntity returns a blank entity with a randomly generated ID
func NewEntity() (*Entity, error) {
	ID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	// Check and handle rare case of a collision
	_, exists := components[ID]
	for exists {
		ID, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}
		_, exists = components[ID]
	}
	components[ID] = []*Component{}
	return &Entity{
		ID: ID,
	}, nil
}

// Query queries the components map for a specifified component for this entity
func (e *Entity) Query(tag string) (Component, error) {
	for _, comp := range components[e.ID] {
		if (*comp).Tag() == tag {
			return (*comp), nil
		}
	}
	return nil, fmt.Errorf("component %v not attached to entity %v", tag, e)
}

// Add adds component to an entity in the components map
func (e *Entity) Add(c Component) error {
	for _, comp := range components[e.ID] {
		if c == (*comp) {
			return fmt.Errorf("duplicate component %v on entity %v", c.Tag(), e)
		}
	}
	components[e.ID] = append(components[e.ID], &c)
	return nil
}

// Remove removes a component from an entity in the components map
func (e *Entity) Remove(c Component) error {
	var idx int = -1
	for i, comp := range components[e.ID] {
		if c == (*comp) {
			idx = i
		}
	}
	if idx == -1 {
		return fmt.Errorf("component %v not attached to entity %v", c.Tag(), e)
	}
	length := len(components[e.ID])
	components[e.ID][idx] = components[e.ID][length-1]
	components[e.ID][length-1] = nil
	components[e.ID] = components[e.ID][:length-1]
	return nil
}

// GetID returns the entity uuid
func (e *Entity) GetID() uuid.UUID {
	return e.ID
}

// Delete removes and entity completly from the components map
func (e *Entity) Delete() {
	delete(components, e.ID)
}

func (e *Entity) String() string {
	return fmt.Sprintf("{ %v }", e.ID)
}
