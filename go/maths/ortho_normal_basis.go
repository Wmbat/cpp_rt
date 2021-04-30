package maths

import "math"

type OrthoNormalBasis struct {
   X Vec3
   Y Vec3
   Z Vec3
}

func (lhs *OrthoNormalBasis) Transform(rhs *Vec3) Vec3 {
   x := MultScalar(&lhs.X, rhs.X)
   y := MultScalar(&lhs.Y, rhs.Y)
   z := MultScalar(&lhs.Z, rhs.Z)

   xy := Add(&x, &y);

   return Add(&xy, &z)
}

func OrthoNormalBasisFromXY(x *Vec3, y *Vec3) OrthoNormalBasis {
   crossXY := Cross(x, y)
   normZ := Normalise(&crossXY)
   normY := Cross(&normZ, x)
   
   return OrthoNormalBasis{X: *x, Y: normY, Z: normZ}
}

func OrthoNormalBasisFromYX(y *Vec3, x *Vec3) OrthoNormalBasis {
    crossXY := Cross(x, y)
    normZ := Normalise(&crossXY)
    normX := Cross(y, &normZ)

    return OrthoNormalBasis{X: normX, Y: *y, Z: normZ}
}

func OrthoNormalBasisFromXZ(x *Vec3, z *Vec3) OrthoNormalBasis {
    crossZX := Cross(z, x)
    normY := Normalise(&crossZX)
    normZ := Cross(x, &normY)

    return OrthoNormalBasis{X: *x, Y: normY, Z: normZ}
}

func OrthoNormalBasisFromZX(z *Vec3, x *Vec3) OrthoNormalBasis {
    crossZX := Cross(z, x)
    normY := Normalise(&crossZX)
    normX := Cross(&normY, z)

    return OrthoNormalBasis{X: normX, Y: normY, Z: *z}
}

func OrthoNormalBasisFromYZ(y *Vec3, z *Vec3) OrthoNormalBasis {
    crossYZ := Cross(y, z) 
    normX := Normalise(&crossYZ)
    normZ := Cross(&normX, y)

    return OrthoNormalBasis{X: normX, Y: *y, Z: normZ}
}

func OrthoNormalBasisFromZY(z *Vec3, y *Vec3) OrthoNormalBasis {
    crossYZ := Cross(y, z)
    normX := Normalise(&crossYZ)
    normY := Cross(&normX, z)

    return OrthoNormalBasis{X: normX, Y: normY, Z: *z}
}

func OrthoNormalBasisFromZ(z* Vec3) OrthoNormalBasis {
    xAxis := Vec3{X: 1.0}

    if math.Abs(Dot(z, &xAxis)) >= 0.999 {
        basis := Vec3{Y: 1.0}
        crossBZ := Cross(&basis, z)
        normX := Normalise(&crossBZ)
        
        crossZX := Cross(z, &normX)
        normY := Normalise(&crossZX)

        return OrthoNormalBasis{X: normX, Y: normY, Z: *z}

    } else {
        basis := xAxis
        crossBZ := Cross(&basis, z)
        normX := Normalise(&crossBZ)
        
        crossZX := Cross(z, &normX)
        normY := Normalise(&crossZX) 

        return OrthoNormalBasis{X: normX, Y: normY, Z: *z}
    }
}
