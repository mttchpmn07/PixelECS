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

// GetCurrentSprite returns the current sprite for the current sequence
func (an *CAnimation) GetCurrentSprite() *pixel.Sprite {
	return an.Sequences[an.Current].Sprite()
}

// Sequence struct to hold a sequence in an animation
type Sequence struct {
	textures   []*pixel.Sprite
	frame      int
	SampleRate float64
	loop       bool
}

// NewSequence constructor for the sequence struct
func NewSequence(filenames []string, sampleRate float64, loop bool) (*Sequence, error) {
	textures := []*pixel.Sprite{}
	for _, filename := range filenames {
		pic, err := loadPicture(filename)
		if err != nil {
			return nil, err
		}
		textures = append(textures, pixel.NewSprite(pic, pic.Bounds()))
	}

	return &Sequence{
		textures:   textures,
		SampleRate: sampleRate,
		loop:       loop,
	}, nil
}

// Sprite return the current sprite for the sequence
func (seq *Sequence) Sprite() *pixel.Sprite {
	return seq.textures[seq.frame]
}

// NextFrame advances the sequence to the next frame until it's finished, it loops if loop is true.
// Also returns true if the sequence if finished or false if not
func (seq *Sequence) NextFrame() bool {
	if seq.frame == len(seq.textures)-1 {
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
