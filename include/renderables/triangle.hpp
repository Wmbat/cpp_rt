#pragma once

#include "math/vec.hpp"

#include <cstddef>

class triangle
{
public:
   triangle(const vec& v0, const vec& v1, const vec& v2) noexcept;

   // clang-format off
   template <size_t index> requires(index < 3) 
   [[nodiscard]] constexpr const vec& vertex() const noexcept
   {
      return data[index];
   }
   // clang-format on

   [[nodiscard]] constexpr const vec& vertex(size_t index) const noexcept { return data[index]; }

   [[nodiscard]] constexpr const vec u() const noexcept { return vertex<1>() - vertex<0>(); }
   [[nodiscard]] constexpr const vec v() const noexcept { return vertex<2>() - vertex<0>(); }

   [[nodiscard]] const norm normal() const;

private:
   vec data[3];
};
