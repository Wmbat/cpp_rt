package tracer 

import "go_pt/maths"

type Ray struct {
   Origin maths.Vec3
   Direction maths.Vec3
}

func (ray *Ray) PositionAlong(t float64) maths.Vec3 {
   timeVector := maths.MultScalar(&ray.Direction, t)

   return maths.Add(&ray.Direction, &timeVector);
}
