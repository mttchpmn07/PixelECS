package main

import (
	"testing"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/pixelecs/core"
)

func TestLocation(t *testing.T) {
	basicEntity, err := ecs.NewEntity()
	if err != nil {
		t.Errorf(err.Error())
	}
	locationComponent := ecs.NewLocation(0, 100)
	basicEntity.Add(locationComponent)
	location, err := ecs.GetLocation(basicEntity)
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

	location, err = ecs.GetLocation(basicEntity)
	if err != nil {
		t.Errorf(err.Error())
	}

	if location.Loc != pixel.V(500, 500) {
		t.Errorf("Location is not persistent")
	}
}
