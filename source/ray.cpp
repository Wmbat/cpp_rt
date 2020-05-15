#include "ray.hpp"

ray::ray(vec const& origin, vec const& direction) noexcept : ori(origin), dir(direction) {}

vec const& ray::origin() const noexcept
{
   return ori;
}
vec const& ray::direction() const noexcept
{
   return dir;
}

vec ray::position_along(double t) const noexcept
{
   return ori + dir * t;
}
