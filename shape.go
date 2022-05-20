package sdfxshape

import (
	"github.com/deadsy/sdfx/sdf"
	. "github.com/stevegt/goadapt"
)

type Shape struct {
	s sdf.SDF3
}

func (in Shape) Translate(x, y, z float64) (out Shape) {
	out = Shape{}
	t := sdf.Translate3d(sdf.V3{X: x, Y: y, Z: z})
	out.s = sdf.Transform3D(in.s, t)
	return
}

func (in Shape) Rotate(x, y, z float64) (out Shape) {
	out = Shape{}
	xr := sdf.RotateX(sdf.DtoR(x))
	yr := sdf.RotateY(sdf.DtoR(y))
	zr := sdf.RotateZ(sdf.DtoR(z))
	out.s = sdf.Transform3D(in.s, xr)
	out.s = sdf.Transform3D(out.s, yr)
	out.s = sdf.Transform3D(out.s, zr)
	return
}

func (pos Shape) Sub(neg Shape) (out Shape) {
	out = Shape{}
	out.s = sdf.Difference3D(pos.s, neg.s)
	return
}

func Union(ins ...Shape) (out Shape) {
	out = Shape{}
	var ss []sdf.SDF3
	for _, in := range ins {
		ss = append(ss, in.s)
	}
	out.s = sdf.Union3D(ss...)
	return
}

func Box(x, y, z, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.s, err = sdf.Box3D(sdf.V3{X: x, Y: y, Z: z}, round)
	Ck(err)
	return
}

func Cylinder(height, radius, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.s, err = sdf.Cylinder3D(height, radius, round)
	Ck(err)
	return
}

func Cone(height, r0, r1, round float64) (out Shape) {
	out = Shape{}
	var err error
	out.s, err = sdf.Cone3D(height, r0, r1, round)
	Ck(err)
	return
}
