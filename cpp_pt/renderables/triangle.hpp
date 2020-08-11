#pragma once

#include <cpp_pt/math/vec.hpp>

#include <array>
#include <cstddef>

class triangle
{
public:
   triangle(const vec& v0, const vec& v1, const vec& v2) noexcept;

   // clang-format off
   template <size_t index> requires(index < 3) 
   [[nodiscard]] constexpr auto vertex() const noexcept -> const vec&
   {
      return data[index];
   }
   // clang-format on

   [[nodiscard]] constexpr auto vertex(size_t index) const noexcept -> const vec&
   {
      return data.at(index);
   }

   [[nodiscard]] constexpr auto u() const noexcept -> const vec
   {
      return vertex<1>() - vertex<0>();
   }
   [[nodiscard]] constexpr auto v() const noexcept -> const vec
   {
      return vertex<2>() - vertex<0>();
   }

   [[nodiscard]] auto normal() const -> const norm;

private:
   std::array<vec, 3> data;
};
