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

	Anchor      *CLocation
	Points      []pixel.Vec
	UniqueEdges edgeSlice
}

// NewCCollisionPoly returns a new CKenetics component with a given starting speed and angularVelocity
func NewCCollisionPoly(anchor *CLocation, points ...pixel.Vec) ecs.Component {
	var angles []float64
	var uniqueEdges []edge
	var pointA, pointB pixel.Vec
	var pointAptr, pointBptr *pixel.Vec
FINDNORMALS:
	for i := range points {
		pointA = points[i]
		pointAptr = &points[i]
		if i == len(points)-1 {
			pointB = points[0]
			pointBptr = &points[0]
		} else {
			pointB = points[i+1]
			pointBptr = &points[i+1]
		}
		angle := pointA.Sub(pointB).Angle()
		for _, a := range angles {
			if a == angle || a == angle-math.Pi || a == angle+math.Pi {
				continue FINDNORMALS
			}
		}
		angles = append(angles, angle)
		uniqueEdges = append(uniqueEdges, edge{
			pointA: pointAptr,
			pointB: pointBptr,
		})
	}
	return &CCollisionPoly{
		tag:         CPTAG,
		Anchor:      anchor,
		Points:      points,
		UniqueEdges: uniqueEdges,
	}
}

// GetCCollisionPoly returns the actual CCollisionPoly struct implmenting the component for a given entity
func GetCCollisionPoly(e *ecs.Entity) (*CCollisionPoly, error) {
	comp, err := e.Query(CPTAG)
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
	return fmt.Sprintf("%v Poly(%v) : %v @ %v", cp.tag, cp.Points, cp.NormalAxes(), cp.Anchor)
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

// NormalAxes returns all unique normal axes
func (cp *CCollisionPoly) NormalAxes() []pixel.Vec {
	var normals []pixel.Vec
	for _, ed := range cp.UniqueEdges {
		angle := ed.pointA.Sub(*ed.pointB).Angle() + math.Pi/2
		normal := pixel.Unit(angle)
		normals = append(normals, normal)
	}
	return normals
}

func minMax(slice []float64) (min, max float64) {
	min = math.Inf(1)
	max = math.Inf(-1)
	for _, num := range slice {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

// Collides check for a collision between two CCollisionPolys
func (cp *CCollisionPoly) Collides(other *CCollisionPoly) bool {
	normals := cp.NormalAxes()
	normals = append(normals, other.NormalAxes()...)

	for _, axis := range normals {
		cpProjections := []float64{}
		otherProjections := []float64{}
		for _, pointA := range cp.Points {
			cpProjections = append(cpProjections, pointA.Add(cp.Anchor.Loc).Dot(axis))
		}
		for _, pointB := range other.Points {
			otherProjections = append(otherProjections, pointB.Add(other.Anchor.Loc).Dot(axis))
		}
		cpMin, cpMax := minMax(cpProjections)
		otherMin, otherMax := minMax(otherProjections)
		if cpMax < otherMin || otherMax < cpMin {
			return false
		}
	}
	return true
}
