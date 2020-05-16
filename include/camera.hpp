#pragma once

#include "math/details.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <cmath>

class camera
{
public:
   camera() noexcept : origin(0.0, 0.0, 0.0)
   {
      lower_left_corner = vec{-2, -1, -1.0};
      horizontal = vec{4.0, 0.0, 0.0};
      vertical = vec{0.0, 2.0, 0.0};
   }

   ray shoot_ray(double u, double v) const noexcept
   {
      return ray(origin, lower_left_corner + u * horizontal + v * vertical);
   }

private:
   vec origin;
   vec horizontal;
   vec vertical;
   vec lower_left_corner;
};
