package core

import (
	"fmt"
)

// System comment here
type System interface {
	Update(args ...interface{}) error
	AddEntity(es ...*Entity) error
	RemoveEntity(es ...*Entity) error
	Tag() string
}

// TODO consider using something like a priority que to force execution of updates in a specific order
type systemRegister []System

var systems systemRegister

func init() {
	systems = systemRegister{}
}

// RegisterSystem adds a system to the global systems pool
func RegisterSystem(sys System) error {
	_, err := GetSystem(sys.Tag())
	if err == nil {
		return fmt.Errorf("system of type %v already registered", sys.Tag())
	}
	systems = append(systems, sys)
	return nil
}

// RemoveSystem removes a system from the global systems pool
func RemoveSystem(sys System) error {
	var idx int = -1
	for i, s := range systems {
		if sys == s {
			idx = i
		}
	}
	if idx == -1 {
		return fmt.Errorf("system not registered")
	}
	l := len(systems)
	systems[idx] = systems[l-1]
	systems[l-1] = nil
	systems = systems[:l-1]
	return nil
}

// UpdateSystems all systems in the systems pool
func UpdateSystems(args ...interface{}) error {
	for _, s := range systems {
		err := s.Update(args...)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetSystem returns the system from the system pool with a given tag
func GetSystem(tag string) (System, error) {
	for _, s := range systems {
		if s.Tag() == tag {
			return s, nil
		}
	}
	return nil, fmt.Errorf("system of type %v not registers", tag)
}

// StripEntity helper function that removes a *Entity from a []*Entity and returns the new list
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
	ce[l-1] = nil
	ce = ce[:l-1]
	return ce, nil
}
