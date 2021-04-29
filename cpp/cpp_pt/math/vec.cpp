#include <cpp_pt/math/vec.hpp>

#include <cmath>

auto vec::length() const noexcept -> double
{
   return std::sqrt(length_squared());
}
