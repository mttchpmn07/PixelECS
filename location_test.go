package pixelecs

import (
	"testing"

	"github.com/faiface/pixel"
)

func TestLocation(t *testing.T) {
	basicEntity, err := NewEntity()
	if err != nil {
		t.Errorf(err.Error())
	}
	locationComponent := NewLocation(0, 100)
	basicEntity.Add(locationComponent)
	location, err := GetLocation(basicEntity)
	if err != nil {
		t.Errorf(err.Error())
	}
	if err != nil {
		t.Errorf(err.Error())
	}

	location.Loc = location.Loc.Add(pixel.V(10, 10))
	if location.Loc != pixel.V(10, 110) {
		t.Errorf("Location did not add correctly")
	}

	location.Loc = pixel.V(500, 500)

	location, err = GetLocation(basicEntity)
	if err != nil {
		t.Errorf(err.Error())
	}

	if location.Loc != pixel.V(500, 500) {
		t.Errorf("Location is not persistent")
	}
}
