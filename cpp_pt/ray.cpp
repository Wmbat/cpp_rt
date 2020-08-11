#include <cpp_pt/ray.hpp>

ray::ray(const vec& origin, const vec& direction) noexcept : ori(origin), dir(direction) {}

auto ray::origin() const noexcept -> const vec&
{
   return ori;
}
auto ray::direction() const noexcept -> const vec&
{
   return dir;
}

auto ray::position_along(double t) const noexcept -> vec
{
   return ori + dir * t;
}
