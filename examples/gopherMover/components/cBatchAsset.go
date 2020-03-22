package components

import (
	"fmt"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// BATAG const to hold the Location tag
	BATAG = "batchasset"
)

// CBatchAsset component for storing a sprite sheet and draw batch
type CBatchAsset struct {
	tag         string
	Spritesheet pixel.Picture
	Batch       *pixel.Batch
}

// NewCBatchAsset returns a new CBatchAsset component with a given starting x and y
func NewCBatchAsset(filename string) (ecs.Component, error) {
	spritesheet, err := loadPicture(filename)
	if err != nil {
		return nil, err
	}
	batch := pixel.NewBatch(&pixel.TrianglesData{}, spritesheet)
	return &CBatchAsset{
		tag:         BATAG,
		Spritesheet: spritesheet,
		Batch:       batch,
	}, nil
}

// GetCBatchAsset returns the actual CBatchAsset struct for a given entity
func GetCBatchAsset(e *ecs.Entity) (*CBatchAsset, error) {
	comp, err := e.Query(BATAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CBatchAsset), nil
}

// Tag returns the tag for this component
func (ba *CBatchAsset) Tag() string {
	return ba.tag
}

func (ba *CBatchAsset) String() string {
	return fmt.Sprintf("%v", ba.tag)
}
