package sdfxshape

import (
	"github.com/deadsy/sdfx/sdf"
	v3 "github.com/deadsy/sdfx/vec/v3"
	. "github.com/stevegt/goadapt"
)

type Shape struct {
	S sdf.SDF3
}

func (in Shape) Translate(x, y, z float64) (out Shape) {
	out = Shape{}
	t := sdf.Translate3d(v3.Vec{X: x, Y: y, Z: z})
	out.S = sdf.Transform3D(in.S, t)
	return
}

func (in Shape) Rotate(x, y, z float64) (out Shape) {
	out = Shape{}
	xr := sdf.RotateX(sdf.DtoR(x))
	yr := sdf.RotateY(sdf.DtoR(y))
	zr := sdf.RotateZ(sdf.DtoR(z))
	out.S = sdf.Transform3D(in.S, xr)
	out.S = sdf.Transform3D(out.S, yr)
	out.S = sdf.Transform3D(out.S, zr)
	return
}

func (pos Shape) Sub(neg Shape) (out Shape) {
	out = Shape{}
	out.S = sdf.Difference3D(pos.S, neg.S)
	return
}

func Union(ins ...Shape) (out Shape) {
	out = Shape{}
	var ss []sdf.SDF3
	for _, in := range ins {
		ss = append(ss, in.S)
	}
	out.S = sdf.Union3D(ss...)
	return
}

func Box(x, y, z, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.S, err = sdf.Box3D(v3.Vec{X: x, Y: y, Z: z}, round)
	Ck(err)
	return
}

func Cylinder(height, radius, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.S, err = sdf.Cylinder3D(height, radius, round)
	Ck(err)
	return
}

func Cone(height, r0, r1, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.S, err = sdf.Cone3D(height, r0, r1, round)
	Ck(err)
	return
}
