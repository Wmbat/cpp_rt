#include "math/vec.hpp"

#include <cmath>

double vec::length() const noexcept
{
   return std::sqrt(length_squared());
}
