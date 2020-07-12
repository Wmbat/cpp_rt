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
      double vertical_fov{0.0};
      double aspect_ratio{0.0};
      double aperture{0.0};
      double focus_distance{0.0};
   };

public:
   camera(const create_info& info) noexcept;

   [[nodiscard]] auto shoot_ray(double u, double v) const noexcept -> ray;

private:
   const vec origin;
   vec horizontal;
   vec vertical;
   vec lower_left_corner;
   ortho_normal_basis axis;
   double lens_radius;
};
