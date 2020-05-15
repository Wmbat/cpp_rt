#pragma once

#include "math/vec.hpp"

class ray
{
public:
   ray() = default;
   ray(vec const& origin, vec const& direction) noexcept;

   vec const& origin() const noexcept;
   vec const& direction() const noexcept;

   [[nodiscard]] vec position_along(double t) const noexcept;

private:
   vec ori;
   vec dir;
};
