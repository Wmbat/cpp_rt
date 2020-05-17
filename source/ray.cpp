#include "ray.hpp"

ray::ray(const vec& origin, const vec& direction) noexcept : ori(origin), dir(direction) {}

const vec& ray::origin() const noexcept
{
   return ori;
}
const vec& ray::direction() const noexcept
{
   return dir;
}

vec ray::position_along(double t) const noexcept
{
   return ori + dir * t;
}
