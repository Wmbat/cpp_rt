#include "materials/diffuse_light.hpp"
#include "math/details.hpp"
#include "math/ortho_nomal_basis.hpp"

diffuse_light::diffuse_light(const colour &emission_in) noexcept : emission(emission_in) {}

auto diffuse_light::scatter(const ray &ray_in, const hit &hit_in, double u, double v) const
   -> scatter_data
{
   const auto basis = ortho_normal_basis::from_z(hit_in.normal);
   const vec scatter_dir = hemisphere_sample(basis, u, v);

   // clang-format off
   return scatter_data
      {
         .emission = emission,
         .diffuse = {0.0, 0.0, 0.0},
         .scattered_ray = ray(hit_in.position, scatter_dir)
      };
   // clang-format on
}
