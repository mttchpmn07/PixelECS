package systems

import (
	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
	"github.com/mttchpmn07/PixelECS/gopherPlatformer/components"
)

const (
	// CTTAG SCollisionTracker tag
	CTTAG = "collisiontracker"
	// GRIDDENSITY the number of windows divisions will be GRIDDDENSITY**2
	GRIDDENSITY = 4
)

type collisionManifold struct {
	entities map[string]*ecs.Entity
	//e1, e2          *ecs.Entity
	//prop1, prop2    *components.CProperties
	collisionNormal pixel.Vec
	collisionDepth  float64
}

// SCollisionTracker stores information for the collision tracking system
type SCollisionTracker struct {
	tag string

	collisions      []collisionManifold
	collisionGrid   []components.CollisionShape
	controlEntities []*ecs.Entity
	comps           []string
}

// NewSCollisionTracker constructs a SCollisionTracker from a varidact list of entities
func NewSCollisionTracker(windowWidth, windowHeight float64, es ...*ecs.Entity) (ecs.System, error) {
	grid := []components.CollisionShape{}
	cellWidth := windowWidth / GRIDDENSITY
	cellHeight := windowHeight / GRIDDENSITY
	for i := 0.0; i < windowWidth; i += cellWidth {
		for j := 0.0; j < windowHeight; j += cellHeight {
			anchor := components.NewCLocation(i+cellWidth/2, j+cellHeight/2, 0)
			poly := components.NewPolygon(
				anchor.(*components.CLocation),
				[]pixel.Vec{
					pixel.V(cellWidth/2, cellHeight/2),
					pixel.V(cellWidth/2, -cellHeight/2),
					pixel.V(-cellWidth/2, -cellHeight/2),
					pixel.V(-cellWidth/2, cellHeight/2),
				}...,
			)
			grid = append(grid, poly)
		}
	}

	ct := &SCollisionTracker{
		tag:             CTTAG,
		collisionGrid:   grid,
		controlEntities: []*ecs.Entity{},
		comps: []string{
			components.CSTAG,
			components.SPTAG,
		},
	}
	err := ct.AddEntity(es...)
	if err != nil {
		return nil, err
	}
	return ct, nil
}

// GetComponents returns the nessary components for an entity to be used in this system
func (ct *SCollisionTracker) GetComponents() []string {
	return ct.comps
}

// Update checks for any valid collisions (for the moment it also resolves them, but I'd like to make that a seperate system)
func (ct *SCollisionTracker) Update(args ...interface{}) error {
	//win := args[0].(*pixelgl.Window)
	compareList := [][]*ecs.Entity{}
	//numCompares := 0
	for _, cell := range ct.collisionGrid {
		compares := []*ecs.Entity{}
		//poly := cell.Render()
		//poly.Draw(win)
		for _, e := range ct.controlEntities {
			cs, err := components.GetCCollisionShape(e)
			if err != nil {
				return err
			}
			if cell.Collides(cs.Shape) {
				compares = append(compares, e)
			}
			//numCompares++
		}
		compareList = append(compareList, compares)
	}
	for _, compares := range compareList {
		for i := 0; i < len(compares)-1; i++ {
			for j := i + 1; j < len(compares); j++ {
				e1 := compares[i]
				e2 := compares[j]
				sp1, err := components.GetCProperties(e1)
				if err != nil {
					return err
				}
				sp2, err := components.GetCProperties(e2)
				if err != nil {
					return err
				}
				if !sp1.Active || !sp2.Active {
					continue
				}
				e1CP, err := components.GetCCollisionShape(e1)
				if err != nil {
					return err
				}
				e2CP, err := components.GetCCollisionShape(e2)
				if err != nil {
					return err
				}
				if e1CP.Collides(e2CP) {
					ct.collisions = append(ct.collisions, collisionManifold{
						entities: map[string]*ecs.Entity{
							sp1.Class: e1,
							sp2.Class: e2,
						},
						//e1:              e1,
						//e2:              e2,
						//prop1:           sp1,
						//prop2:           sp2,
						collisionNormal: pixel.V(0, 0),
						collisionDepth:  0,
					})
				}
				//numCompares++
			}
		}
	}
	//fmt.Println(numCompares)
	return ct.Handle()
}

// Handle handles any collisions generated returns any errors
func (ct *SCollisionTracker) Handle() error {
	for _, man := range ct.collisions {
		if _, foundGopher := man.entities["gopher"]; foundGopher {
			if fly, foundFly := man.entities["fly"]; foundFly {
				sp, err := components.GetCProperties(fly)
				if err != nil {
					return err
				}
				sp.Active = false
			}
		}
		if gopher, foundGopher := man.entities["gopher"]; foundGopher {
			if _, foundFly := man.entities["wall"]; foundFly {
				loc, err := components.GetCLocation(gopher)
				if err != nil {
					return err
				}
				loc.Loc = pixel.V(400, 300)
			}
		}
	}
	ct.collisions = ct.collisions[:0]
	return nil
}

// AddEntity adds any number of entities to this system
func (ct *SCollisionTracker) AddEntity(es ...*ecs.Entity) error {
	ct.controlEntities = append(ct.controlEntities, es...)
	return nil
}

// RemoveEntity removes any number of entities from this system
func (ct *SCollisionTracker) RemoveEntity(es ...*ecs.Entity) error {
	for _, e := range es {
		newEntries, err := ecs.StripEntity(ct.controlEntities, e)
		if err != nil {
			return err
		}
		ct.controlEntities = newEntries
	}
	return nil
}

// Tag getter for tag
func (ct *SCollisionTracker) Tag() string {
	return ct.tag
}
