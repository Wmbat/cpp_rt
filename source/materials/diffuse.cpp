#include "materials/diffuse.hpp"
#include "math/details.hpp"
#include "math/ortho_nomal_basis.hpp"

diffuse::diffuse(const colour& emission_in, const colour& diffuse_in) noexcept :
   emission_colour(emission_in), diffuse_colour(diffuse_in)
{}

auto diffuse::scatter(const ray& ray_in, const hit& hit_in, double u, double v) const -> scatter_data
{
   const auto basis = ortho_normal_basis::from_z(hit_in.normal);
   const vec scatter_dir = hemisphere_sample(basis, u, v);

   // clang-format off
   return scatter_data {
      .emission = emission_colour, 
      .diffuse = diffuse_colour, 
      .scattered_ray = ray(hit_in.position, scatter_dir)
   };
   // clang-format on
}
