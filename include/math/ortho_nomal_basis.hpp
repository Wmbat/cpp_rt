#pragma once

#include "vec.hpp"

class ortho_normal_basis
{
public:
   ortho_normal_basis() = default;
   ortho_normal_basis(const norm& x_in, const norm& y_in, const norm& z_in);

   constexpr const norm& x() const noexcept { return data[0]; }
   constexpr const norm& y() const noexcept { return data[1]; }
   constexpr const norm& z() const noexcept { return data[2]; }

   constexpr vec transform(const vec& pos) const noexcept
   {
      return x() * pos.x() + y() * pos.y() + z() * pos.z();
   }

   static ortho_normal_basis from_xy(const norm& x_in, const norm& y_in);
   static ortho_normal_basis from_yx(const norm& y_in, const norm& x_in);
   static ortho_normal_basis from_xz(const norm& x_in, const norm& z_in);
   static ortho_normal_basis from_zx(const norm& z_in, const norm& x_in);
   static ortho_normal_basis from_yz(const norm& y_in, const norm& z_in);
   static ortho_normal_basis from_zy(const norm& z_in, const norm& y_in);
   static ortho_normal_basis from_z(const norm& z_in);

private:
   norm data[3] = {norm(0.0, 0.0, 0.0), norm(0.0, 0.0, 0.0), norm(0.0, 0.0, 0.0)};
};
