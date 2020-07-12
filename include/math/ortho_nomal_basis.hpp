#pragma once

#include "vec.hpp"

#include <array>

class ortho_normal_basis
{
public:
   ortho_normal_basis() = default;
   ortho_normal_basis(const norm& x_in, const norm& y_in, const norm& z_in);

   [[nodiscard]] constexpr auto x() const noexcept -> const norm& { return data[0]; }
   [[nodiscard]] constexpr auto y() const noexcept -> const norm& { return data[1]; }
   [[nodiscard]] constexpr auto z() const noexcept -> const norm& { return data[2]; }

   [[nodiscard]] constexpr auto transform(const vec& pos) const noexcept -> vec
   {
      return x() * pos.x() + y() * pos.y() + z() * pos.z();
   }

   static auto from_xy(const norm& x_in, const norm& y_in) -> ortho_normal_basis;
   static auto from_yx(const norm& y_in, const norm& x_in) -> ortho_normal_basis;
   static auto from_xz(const norm& x_in, const norm& z_in) -> ortho_normal_basis;
   static auto from_zx(const norm& z_in, const norm& x_in) -> ortho_normal_basis;
   static auto from_yz(const norm& y_in, const norm& z_in) -> ortho_normal_basis;
   static auto from_zy(const norm& z_in, const norm& y_in) -> ortho_normal_basis;
   static auto from_z(const norm& z_in) -> ortho_normal_basis;

private:
   std::array<norm, 3> data = {norm(0.0, 0.0, 0.0), norm(0.0, 0.0, 0.0), norm(0.0, 0.0, 0.0)};
};
