package systems

import (
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherMover/components"
)

const (
	// SBMTAG const to hold the SRandomWalk tag
	SBMTAG = "randomwalk"
)

// SRandomWalk random walk system
type SRandomWalk struct {
	tag string

	controlEntities []*ecs.Entity
	comps           []string
}

// NewSRandomWalk returns a new random walk system with all given entities attached
func NewSRandomWalk(es ...*ecs.Entity) (ecs.System, error) {
	cSystem := &SRandomWalk{
		tag:             SBMTAG,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.LTAG,
			components.KTAG,
		},
	}
	err := cSystem.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return cSystem, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (rw *SRandomWalk) GetComponents() []string {
	return rw.comps
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Update calculates next state for all components used by the system for each of its associated entities
func (rw *SRandomWalk) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	dt := (*args[1].(*float64))

	for _, e := range rw.controlEntities {
		sp, err := components.GetCProperties(e)
		if err != nil {
			return err
		}
		if !sp.Active {
			continue
		}
		loc, err := components.GetCLocation(e)
		if err != nil {
			return err
		}
		kin, err := components.GetCKenetics(e)
		if err != nil {
			return err
		}
		kin.Acceleration = pixel.V((rand.Float64() - 0.5), (rand.Float64() - 0.5)).Unit().Scaled(kin.AccMag)
		kin.Velocity = kin.Velocity.Add(kin.Acceleration.Scaled(0.5 * dt)).Unit().Scaled(kin.Speed)
		kin.Velocity.X += kin.Acceleration.X * 0.5 * dt
		kin.Velocity.Y += kin.Acceleration.Y * 0.5 * dt
		//fmt.Println(kin.Acceleration, kin.Velocity)
		loc.Loc = loc.Loc.Add(kin.Velocity.Scaled(dt))
	}
	return nil
}

// AddEntity adds any number of entities to the keyboard control system via a variadic function call
func (rw *SRandomWalk) AddEntity(es ...*ecs.Entity) error {
	err := ecs.ValidateEntitySystem(rw, es...)
	if err != nil {
		return err
	}
	rw.controlEntities = append(rw.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from the keyboard control system via a variadic function call
func (rw *SRandomWalk) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(rw.controlEntities, e)
		if err != nil {
			return err
		}
		rw.controlEntities = newEntries
	}
	return nil
}

// Tag returns the tag for this system
func (rw *SRandomWalk) Tag() string {
	return rw.tag
}
