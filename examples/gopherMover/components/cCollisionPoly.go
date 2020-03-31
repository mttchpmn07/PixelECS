package components

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	ecs "github.com/mttchpmn07/PixelECS/core"
)

const (
	// KTAG const to hold the Location tag
	CPTAG = "collisionpoly"
)

// CKenetics component for storing kinetic information of an entity
type CCollisionPoly struct {
	tag string

	Points      []*pixel.Vec
	UniqueEdges edgeSlice
}

// NewCKenetics returns a new CKenetics component with a given starting speed and angularVelocity
func NewCCollisionPoly(points ...pixel.Vec) ecs.Component {
	var angles []float64
	var uniqueEdges []edge
	var pointPtrs []*pixel.Vec
FINDNORMALS:
	for i := range points {
		pointPtrs = append(pointPtrs, &points[i])
		var pointA, pointB *pixel.Vec
		pointA = &points[i]
		if i == len(points)-1 {
			pointB = &points[0]
		} else {
			pointB = &points[i+1]
		}
		angle := pointA.Sub(*pointB).Angle()
		for _, a := range angles {
			if a == angle || a == angle-math.Pi || a == angle+math.Pi {
				continue FINDNORMALS
			}
		}
		angles = append(angles, angle)
		uniqueEdges = append(uniqueEdges, edge{
			pointA: pointA,
			pointB: pointB,
		})
	}
	return &CCollisionPoly{
		tag:         CPTAG,
		Points:      pointPtrs,
		UniqueEdges: uniqueEdges,
	}
}

// GetCCollisionPoly returns the actual CCollisionPoly struct implmenting the component for a given entity
func GetCCollisionPoly(e *ecs.Entity) (*CCollisionPoly, error) {
	comp, err := e.Query(KTAG)
	if err != nil {
		return nil, err
	}
	return comp.(*CCollisionPoly), nil
}

// Tag returns the tag for this component
func (cp *CCollisionPoly) Tag() string {
	return cp.tag
}

func (cp *CCollisionPoly) String() string {
	var points []pixel.Vec
	for _, ptr := range cp.Points {
		points = append(points, *ptr)
	}
	return fmt.Sprintf("%v Poly(%v) : %v", cp.tag, points, cp.UniqueEdges.NormalAxes())
}

type edge struct {
	pointA, pointB *pixel.Vec
}

type edgeSlice []edge

func (ed *edge) String() string {
	return fmt.Sprintf("{<%v>,<%v>}", *ed.pointA, *ed.pointB)
}

func (eds *edgeSlice) String() string {
	edges := ""
	for _, ed := range *eds {
		edges += fmt.Sprintf("%s ", ed.String())
	}
	return edges
}

func (eds edgeSlice) NormalAxes() []pixel.Vec {
	var normals []pixel.Vec
	for _, ed := range eds {
		angle := ed.pointA.Sub(*ed.pointB).Angle() + math.Pi/2
		normal := pixel.Unit(angle)
		normals = append(normals, normal)
	}
	return normals
}

func (cp *CCollisionPoly) Collides(other *CCollisionPoly) bool {
	// Check for collision
	return false
}
