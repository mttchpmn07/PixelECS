package pixelecs

import (
	"github.com/google/uuid"
)

// Component
type Component interface {
	String() string
	Tag() string
}

type componentStore map[uuid.UUID][]*Component

var components componentStore

func init() {
	components = componentStore{}
}
