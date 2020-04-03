package components

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// ATAG CAnimation tag
	ATAG = "animation"
)

// CAnimation component for an animated entity
type CAnimation struct {
	tag string

	Sequences       map[string]*Sequence
	Current         string
	LastFrameChange time.Time
	Finished        bool
}

// NewCAnimation constructs a CAnimation from a map of sequences
func NewCAnimation(sequences map[string]*Sequence, defaultSequence string) ecs.Component {
	return &CAnimation{
		tag:             ATAG,
		Sequences:       sequences,
		Current:         defaultSequence,
		LastFrameChange: time.Now(),
	}
}

// GetCAnimation returns the actual struct implmenting the component for a given entity
func GetCAnimation(e *ecs.Entity) (*CAnimation, error) {
	comp, err := e.Query(ATAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CAnimation), nil
}

// Tag getter for tag
func (an *CAnimation) Tag() string {
	return an.tag
}

func (an *CAnimation) String() string {
	return fmt.Sprintf("%v", an.tag)
}

// SetSequence changes the current sequence
func (an *CAnimation) SetSequence(name string) {
	if an.Current != name {
		an.Current = name
		an.LastFrameChange = time.Now()
	}
}

// GetCurrentFrame returns the frame of the current animation
func (an *CAnimation) GetCurrentFrame() pixel.Rect {
	return an.Sequences[an.Current].Frame()
}

// Sequence stores the animation details for a given state of an animation
type Sequence struct {
	Frames     []pixel.Rect
	frame      int
	SampleRate float64
	loop       bool
}

// NewSequence constructor for the sequence struct
func NewSequence(asset *CBatchAsset, sampleRate, width, height, padding float64, startFrame, endFrame int, loop bool) (*Sequence, error) {
	var spriteFrames []pixel.Rect
	max := pixel.V(asset.Spritesheet.Bounds().Max.X, asset.Spritesheet.Bounds().Max.Y)
	min := pixel.V(asset.Spritesheet.Bounds().Min.X, asset.Spritesheet.Bounds().Min.Y)
	var frameCount int
	frameCount = 0
	for y := min.Y; y < max.Y; y += height + padding {
		for x := min.X; x < max.X; x += width + padding {
			if frameCount >= startFrame && frameCount <= endFrame {
				spriteFrames = append(spriteFrames, pixel.R(x, y, x+width, y+height))
			}
			frameCount++
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
