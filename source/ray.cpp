#include "ray.hpp"

ray::ray(point const& origin, vec const& direction) noexcept : orig(origin), dir(direction) {}

point ray::origin() const noexcept
{
   return orig;
}
vec ray::direction() const noexcept
{
   return dir;
}

vec ray::at(double time) const noexcept
{
   return orig + dir * time;
}
