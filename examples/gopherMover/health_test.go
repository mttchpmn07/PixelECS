package main

import (
	"testing"

	ecs "github.com/mttchpmn07/PixelECS/core"
)

func TestHealth(t *testing.T) {
	basicEntity, err := ecs.NewEntity()
	if err != nil {
		panic(err)
	}

	healthComponent := NewHealth(0, 100)
	basicEntity.Add(healthComponent)

	health, err := GetHealth(basicEntity)
	if err != nil {
		t.Errorf("failed to get health from entity: %v", err.Error())
	}

	health.Increase(150)
	if health.Value > 100 {
		t.Errorf("Health.Increase increased health above max")
	}
	if health.Value != 100 {
		t.Errorf("Health.Increase didn't correctly increase health")
	}

	health.Decrease(150)
	if health.Value < 0 {
		t.Errorf("Health.Decrease decreased health below 0")
	}
	if health.Value != 0 {
		t.Errorf("Health.Decrease didn't correctly decrease health")
	}
	health.Increase(75)

	health, err = GetHealth(basicEntity)
	if err != nil {
		t.Errorf("failed to get health from entity: %v", err.Error())
	}
	if health.Value != 75 {
		t.Errorf("Health is not persistent")
	}

}
