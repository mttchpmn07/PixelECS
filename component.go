package pixelecs

import (
	"github.com/google/uuid"
)

// Component blank component interface (this is where data should live)
type Component interface {
	String() string
	Tag() string
}

type componentStore map[uuid.UUID][]*Component

var components componentStore

func init() {
	components = componentStore{}
}
