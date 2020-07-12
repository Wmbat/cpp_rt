#pragma once

#include "math/vec.hpp"

class ray
{
public:
   ray() = default;
   ray(const vec& origin, const vec& direction) noexcept;

   [[nodiscard]] auto origin() const noexcept -> const vec&;
   [[nodiscard]] auto direction() const noexcept -> const vec&;

   [[nodiscard]] auto position_along(double t) const noexcept -> vec;

private:
   vec ori;
   vec dir;
};
