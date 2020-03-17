package pixelecs

import (
	"fmt"
	"image"
	"os"

	_ "image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
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
	SRTAG = "SpriteRender"
)

// SpriteRender
type SpriteRender struct {
	tag string

	Active         bool
	sprite         *pixel.Sprite
	location       *Location
	Transformation pixel.Matrix
	Bounds         pixel.Rect
}

var spriteRenders []*SpriteRender

func init() {
	spriteRenders = []*SpriteRender{}
}

// NewSpriteRender
func NewSpriteRender(filename string, active bool, location Component) (Component, error) {
	loc := location.(*Location)
	r := &SpriteRender{
		tag:            SRTAG,
		Active:         active,
		location:       loc,
		Transformation: pixel.IM,
	}
	pic, err := loadPicture(filename)
	if err != nil {
		return r, err
	}
	r.sprite = pixel.NewSprite(pic, pic.Bounds())
	r.Bounds = pic.Bounds()
	spriteRenders = append(spriteRenders, r)
	return r, nil
}

// GetSpriteRender
func GetSpriteRender(e *Entity) (*SpriteRender, error) {
	comp, err := e.Query(SRTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*SpriteRender), nil
}

func (r *SpriteRender) String() string {
	return fmt.Sprintf("%v", r.tag)
}

// Tag
func (r *SpriteRender) Tag() string {
	return r.tag
}

// Draw
func (r *SpriteRender) Draw(win *pixelgl.Window) {
	r.sprite.Draw(win, r.Transformation.Moved(r.location.Loc))
}

// DrawSprites
func DrawSprites(win *pixelgl.Window) {
	for _, sr := range spriteRenders {
		if sr.Active {
			sr.Draw(win)
		}
	}
}
