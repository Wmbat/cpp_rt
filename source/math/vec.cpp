#include "math/vec.hpp"

#include <cmath>

auto vec::length() const noexcept -> double
{
   return std::sqrt(length_squared());
}
