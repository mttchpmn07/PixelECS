package pixelecs

import (
	"fmt"
	"image"
	"os"

	// Needed to set a pixel.Picture from a png
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
	// SRTAG const to hold the spriterender tag
	SRTAG = "spriterender"
)

// SpriteRender implements the component interface for a sprite render system
type SpriteRender struct {
	tag string

	Active         bool
	sprite         *pixel.Sprite
	location       *Location
	Transformation pixel.Matrix
	Bounds         pixel.Rect
}

// global SpriteRender storage for the DrawSprites function
var spriteRenders []*SpriteRender

func init() {
	spriteRenders = []*SpriteRender{}
}

// NewSpriteRender returns a new SpriteRender component with a sprite given via filename, an active flag, and a pointer to the associated location
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

// GetSpriteRender returns the actual SpriteRender struct implemnting the component for a given entity
func GetSpriteRender(e *Entity) (*SpriteRender, error) {
	comp, err := e.Query(SRTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*SpriteRender), nil
}

// Tag returns the tag for this component
func (r *SpriteRender) Tag() string {
	return r.tag
}

func (r *SpriteRender) String() string {
	return fmt.Sprintf("%v", r.tag)
}

// Draw draws the sprite to the given window
func (r *SpriteRender) Draw(win *pixelgl.Window) {
	r.sprite.Draw(win, r.Transformation.Moved(r.location.Loc))
}

// DrawSprites draws all active sprites in spriteRenders
func DrawSprites(win *pixelgl.Window) {
	for _, sr := range spriteRenders {
		if sr.Active {
			sr.Draw(win)
		}
	}
}
