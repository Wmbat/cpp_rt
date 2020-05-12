#pragma once

#include "ray.hpp"

struct sphere
{
   sphere() = default;
   sphere(point const& center, double radius);

   point center{0.0, 0.0, 0.0};
   double radius{0.0};
};
