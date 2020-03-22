package components

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// ATAG const to hold the CAnimation tag
	ATAG = "animation"
)

// CAnimation component for an animated entity
type CAnimation struct {
	tag string

	Sequences       map[string]*Sequence
	Current         string
	LastFrameChange time.Time
	Finished        bool
	Render          bool
}

// NewCAnimation constructor for CAnimation
func NewCAnimation(sequences map[string]*Sequence, defaultSequence string, render bool) ecs.Component {
	return &CAnimation{
		tag:             ATAG,
		Sequences:       sequences,
		Current:         defaultSequence,
		LastFrameChange: time.Now(),
		Render:          render,
	}
}

// GetCAnimation returns the actual CAnimation struct implmenting the component for a given entity
func GetCAnimation(e *ecs.Entity) (*CAnimation, error) {
	comp, err := e.Query(ATAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CAnimation), nil
}

// Tag returns the tag for this component
func (an *CAnimation) Tag() string {
	return an.tag
}

func (an *CAnimation) String() string {
	return fmt.Sprintf("%v", an.tag)
}

// SetSequence changes the current sequence
func (an *CAnimation) SetSequence(name string) {
	an.Current = name
	an.LastFrameChange = time.Now()
}

// GetCurrentFrame returns the current sprite for the current sequence
func (an *CAnimation) GetCurrentFrame() pixel.Rect {
	return an.Sequences[an.Current].Frame()
}

// Sequence struct to hold a sequence in an animation
type Sequence struct {
	Frames     []pixel.Rect
	frame      int
	SampleRate float64
	loop       bool
}

// NewSequence constructor for the sequence struct
func NewSequence(asset *CBatchAsset, sampleRate, width, height, padding float64, loop bool) (*Sequence, error) {
	var spriteFrames []pixel.Rect
	max := pixel.V(asset.Spritesheet.Bounds().Max.X, asset.Spritesheet.Bounds().Max.Y)
	min := pixel.V(asset.Spritesheet.Bounds().Min.X, asset.Spritesheet.Bounds().Min.Y)
	for x := min.X; x < max.X; x += width + padding {
		for y := min.Y; y < max.Y; y += height + padding {
			spriteFrames = append(spriteFrames, pixel.R(x, y, x+width, y+height))
		}
	}

	return &Sequence{
		Frames:     spriteFrames,
		SampleRate: sampleRate,
		loop:       loop,
	}, nil
}

// Frame return the current frame for the sequence
func (seq *Sequence) Frame() pixel.Rect {
	return seq.Frames[seq.frame]
}

// NextFrame advances the sequence to the next frame until it's finished, it loops if loop is true.
// Also returns true if the sequence if finished or false if not
func (seq *Sequence) NextFrame() bool {
	if seq.frame == len(seq.Frames)-1 {
		if seq.loop {
			seq.frame = 0
		} else {
			return true
		}
	} else {
		seq.frame++
	}

	return false
}
