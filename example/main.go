package main

import (
	"github.com/deadsy/sdfx/render"
	. "github.com/stevegt/goadapt"
	shape "github.com/stevegt/sdfxshape"
)

func main() {
	var x float64 = 30.0
	var y float64 = 40.0
	var z float64 = 100.0

	box := shape.Box(x, y, z/2.0, 1)
	cone := shape.Cone(z, x/2.0, x/4.0, 1)
	pos := shape.Union(box, cone).Rotate(90, 0, 0)

	neg := shape.Cylinder(z, x/5.0, 0).Translate(x/5.0, 0, 0)
	all := pos.Sub(neg)

	Pl("rendering...")
	render.RenderSTL(all.S, 300, "all.stl")
}
