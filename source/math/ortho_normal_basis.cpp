#include "math/ortho_nomal_basis.hpp"
#include "math/vec.hpp"

#include <cmath>

ortho_normal_basis::ortho_normal_basis(const norm &x_in, const norm &y_in, const norm &z_in)
{
   data[0] = x_in;
   data[1] = y_in;
   data[2] = z_in;
}

ortho_normal_basis ortho_normal_basis::from_xy(const norm &x_in, const norm &y_in)
{
   const norm z = normalise(cross(x_in, y_in));
   const norm y = cross(z, x_in);

   return ortho_normal_basis{x_in, y, z};
}

ortho_normal_basis ortho_normal_basis::from_yx(const norm &y_in, const norm &x_in)
{
   const norm z = normalise(cross(x_in, y_in));
   const norm x = cross(y_in, z);

   return ortho_normal_basis{x, y_in, z};
}

ortho_normal_basis ortho_normal_basis::from_xz(const norm &x_in, const norm &z_in)
{
   const norm y = normalise(cross(z_in, x_in));
   const norm z = cross(x_in, y);

   return ortho_normal_basis{x_in, y, z};
}

ortho_normal_basis ortho_normal_basis::from_zx(const norm &z_in, const norm &x_in)
{
   const norm y = normalise(cross(z_in, x_in));
   const norm x = cross(y, z_in);

   return ortho_normal_basis{x, y, z_in};
}

ortho_normal_basis ortho_normal_basis::from_yz(const norm &y_in, const norm &z_in)
{
   const norm x = normalise(cross(y_in, z_in));
   const norm z = cross(x, y_in);

   return ortho_normal_basis{x, y_in, z};
}

ortho_normal_basis ortho_normal_basis::from_zy(const norm &z_in, const norm &y_in)
{
   const norm x = normalise(cross(y_in, z_in));
   const norm y = cross(x, z_in);

   return ortho_normal_basis{x, y, z_in};
}

ortho_normal_basis ortho_normal_basis::from_z(const norm &z_in)
{
   const norm x = normalise(cross(
      std::fabs(dot(z_in, norm(1.0, 0.0, 0.0))) > 0.999 ? norm(0.0, 1.0, 0.0) : norm(1.0, 0.0, 0.0),
      z_in));
   const norm y = normalise(cross(z_in, x));

   return ortho_normal_basis{x, y, z_in};
}
