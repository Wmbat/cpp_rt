#include "materials/metallic.hpp"
#include "math/details.hpp"

#include <algorithm>
#include <cmath>

metallic::metallic(
   const colour& emission_in, const colour& diffuse_in, double roughness_in) noexcept :
   emission(emission_in),
   diffuse(diffuse_in), roughness(std::clamp(roughness_in, 0.0, 1.0))
{}

auto metallic::scatter(const ray& ray_in, const hit& hit_in, [[maybe_unused]] double u,
   [[maybe_unused]] double v) const -> scatter_data
{
   const norm reflected_dir = reflect(hit_in.normal, normalise(ray_in.direction()));
   const ray scattered = ray(hit_in.position, reflected_dir + roughness * random_in_unit_sphere());

   // clang-format off
   return scatter_data {
      .emission = emission, 
      .diffuse = dot(scattered.direction(), hit_in.normal) > 0 ? diffuse : colour{}, 
      .scattered_ray = scattered
   };
   // clang-format on
}
