package systems

import (
	"sort"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"

	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

/*
TODO:
	If optimization in needed it might be possible by checking for oclusion
*/

type drawObject struct {
	Batch       *pixel.Batch
	Spritesheet *pixel.Picture
	Frame       *pixel.Rect
	Loc         *pixel.Vec
	Angle       float64
	Scale       float64
}

func (do *drawObject) render() {
	trans := pixel.IM.Scaled(pixel.ZV, do.Scale).Rotated(pixel.ZV, do.Angle)
	sprite := pixel.NewSprite(*do.Spritesheet, *do.Frame)
	sprite.Draw(do.Batch, trans.Moved(*do.Loc))
}

const (
	// BRTAG SBatchRenderer tag
	BRTAG = "batchrenderer"
)

// SBatchRenderer stores information for batch renderer system
type SBatchRenderer struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSBatchRenderer constructs a SBatchRenderer from a varidact list of entities
func NewSBatchRenderer(es ...*ecs.Entity) (ecs.System, error) {
	br := &SBatchRenderer{
		tag:             BRTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.ATAG,
			components.BATAG,
			components.LTAG,
			components.SPTAG,
		},
	}
	err := br.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return br, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (br *SBatchRenderer) GetComponents() []string {
	return br.comps
}

// Update renders each entity to its batch and draws the batches based on the z value from CLocation
//sudo code of what is happening here
//
//	layers = map layer => []drawObjs
//	batches = set batch
//	for entity e
//		build drawObj(location, rotation, batch, frame)
//		layers[layer].append(drawObj)
//	for layer l in sorted layers
//		for drawObj do in layers[l]
//			do.Render
//		for batch b in batches
//			b.Draw(win)
//			b.Clear()
func (br *SBatchRenderer) Update(args ...interface{}) error {
	win := args[0].(*pixelgl.Window)

	layers := map[int][]*drawObject{}
	var exists = struct{}{}
	batches := map[*pixel.Batch]struct{}{}
	for _, e := range br.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		an, err := components.GetCAnimation(e)
		if err != nil {
			return err
		}
		curFrame := an.GetCurrentFrame()
		ba, err := components.GetCBatchAsset(e)
		if err != nil {
			return err
		}
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		do := &drawObject{
			Batch:       ba.Batch,
			Spritesheet: &ba.Spritesheet,
			Frame:       &curFrame,
			Loc:         &loc.Loc,
			Angle:       sp.Angle,
			Scale:       sp.Scale,
		}
		if _, OK := layers[loc.Z]; !OK {
			layers[loc.Z] = []*drawObject{}
		}
		layers[loc.Z] = append(layers[loc.Z], do)
		if _, c := batches[ba.Batch]; !c {
			batches[ba.Batch] = exists
		}
	}
	keys := make([]int, 0, len(layers))
	for k := range layers {
		keys = append(keys, k)
	}
	sKeys := sort.IntSlice(keys)
	sort.Sort(sKeys)
	for _, k := range sKeys {
		layer := layers[k]
		for _, do := range layer {
			do.render()
		}
		for b := range batches {
			b.Draw(win)
			b.Clear()
		}
	}
	return nil
}

// AddEntity adds any number of entities to this system
func (br *SBatchRenderer) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(br, es...)
	if err != nil {
		return err
	}
	br.controlEntities = append(br.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (br *SBatchRenderer) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(br.controlEntities, e)
		if err != nil {
			return err
		}
		br.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (br *SBatchRenderer) Tag() string {
	return br.tag
}
