#pragma once

#include "math/details.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <cmath>

class camera
{
public:
   struct create_info
   {
      vec eye;
      vec look_at;
      vec up;
      double vertical_fov;
      double aspect_ratio;
      double aperture;
      double focus_distance;
   };

public:
   camera(const create_info& info) noexcept;

   [[nodiscard]] ray shoot_ray(double u, double v) const noexcept;

private:
   const vec origin;
   vec horizontal;
   vec vertical;
   vec lower_left_corner;
   ortho_normal_basis axis;
   double lens_radius;
};
