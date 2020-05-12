#pragma once

#include <glm/glm.hpp>

using point = glm::dvec3;
using vec = point;
using colour = point;

class ray
{
public:
   ray() = default;
   ray(point const& origin, vec const& direction) noexcept;

   point origin() const noexcept;
   vec direction() const noexcept;

   vec at(double time) const noexcept;

private:
   point orig;
   vec dir;
};
