package main

import (
	"fmt"
	"image"
	"os"

	// Needed to set a pixel.Picture from a png
	_ "image/png"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

const (
	// SRTAG const to hold the sprite tag
	SRTAG = "sprite"
)

// Sprite implements the component interface for a sprite render system
type Sprite struct {
	tag string

	Active bool
	sprite *pixel.Sprite
}

// NewSprite returns a new Sprite component with a sprite given via filename, an active flag, and a pointer to the associated location
func NewSprite(filename string, active bool) (ecs.Component, error) {
	r := &Sprite{
		tag:    SRTAG,
		Active: active,
	}
	pic, err := loadPicture(filename)
	if err != nil {
		return nil, err
	}
	r.sprite = pixel.NewSprite(pic, pic.Bounds())
	return r, nil
}

// GetSprite returns the actual Sprite struct implemnting the component for a given entity
func GetSprite(e *ecs.Entity) (*Sprite, error) {
	comp, err := e.Query(SRTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*Sprite), nil
}

// Tag returns the tag for this component
func (r *Sprite) Tag() string {
	return r.tag
}

func (r *Sprite) String() string {
	return fmt.Sprintf("%v", r.tag)
}
