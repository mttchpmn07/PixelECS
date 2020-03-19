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

// CSprite component for storing a sprite and its active flag
type CSprite struct {
	tag string

	Active bool
	sprite *pixel.Sprite
}

// NewCSprite returns a new CSprite component with a sprite given via filename, an active flag
func NewCSprite(filename string, active bool) (ecs.Component, error) {
	r := &CSprite{
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

// GetCSprite returns the actual CSprite struct implemnting the component for a given entity
func GetCSprite(e *ecs.Entity) (*CSprite, error) {
	comp, err := e.Query(SRTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CSprite), nil
}

// Tag returns the tag for this component
func (r *CSprite) Tag() string {
	return r.tag
}

func (r *CSprite) String() string {
	return fmt.Sprintf("%v", r.tag)
}
