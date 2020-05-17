#pragma once

#include "math/vec.hpp"

class ray
{
public:
   ray() = default;
   ray(const vec& origin, const vec& direction) noexcept;

   const vec& origin() const noexcept;
   const vec& direction() const noexcept;

   [[nodiscard]] vec position_along(double t) const noexcept;

private:
   vec ori;
   vec dir;
};
