package components

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
)

// CollisionShape generic interface for implementing different types of collisions
type CollisionShape interface {
	Type() string
	Collides(other CollisionShape) bool
	Anchor() *CLocation
	Render() *imdraw.IMDraw
}

type polygon struct {
	anchor      *CLocation
	points      []pixel.Vec
	uniqueEdges edgeSlice
}

// NewPolygon constructs a CCollisionPoly from an anchor CLocation and a list of points.
// The points should be defined either clockwise or counter clockwise around the polygon.
func NewPolygon(anchor *CLocation, points ...pixel.Vec) CollisionShape {
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
	return &polygon{
		anchor:      anchor,
		points:      points,
		uniqueEdges: uniqueEdges,
	}
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

// normalAxes returns all unique normal axis of a CCollisionPoly
func (cp *polygon) normalAxes() []pixel.Vec {
	var normals []pixel.Vec
	for _, ed := range cp.uniqueEdges {
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
func (cp *polygon) Collides(other CollisionShape) bool {
	switch other.Type() {
	case "polygon":
		otherPoly := other.(*polygon)
		normals := cp.normalAxes()
		normals = append(normals, otherPoly.normalAxes()...)

		for _, axis := range normals {
			cpProjections := []float64{}
			otherProjections := []float64{}
			for _, pointA := range cp.points {
				cpProjections = append(cpProjections, pointA.Add(cp.anchor.Loc).Dot(axis))
			}
			for _, pointB := range otherPoly.points {
				otherProjections = append(otherProjections, pointB.Add(otherPoly.anchor.Loc).Dot(axis))
			}
			cpMin, cpMax := minMax(cpProjections)
			otherMin, otherMax := minMax(otherProjections)
			if cpMax < otherMin || otherMax < cpMin {
				return false
			}
		}
		return true
	}
	return false
}

func (cp *polygon) Type() string {
	return "polygon"
}

func (cp *polygon) Anchor() *CLocation {
	return cp.anchor
}

func (cp *polygon) Render() *imdraw.IMDraw {
	poly := imdraw.New(nil)
	for _, p := range cp.points {
		poly.Push(p.Add(cp.anchor.Loc))
	}
	poly.Polygon(2)
	return poly
}
