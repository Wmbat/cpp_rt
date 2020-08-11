#include <cpp_pt/materials/dielectric.hpp>
#include <cpp_pt/math/details.hpp>

dielectric::dielectric(const colour &diffuse_in, double refractive_index_in) :
   diffuse(diffuse_in), refractive_index(refractive_index_in)
{}

auto dielectric::scatter(const ray &ray_in, const hit &hit_in, [[maybe_unused]] double u,
   [[maybe_unused]] double v) const -> scatter_data
{
   const vec unit_dir = normalise(ray_in.direction());
   const double ior_ratio = (hit_in.front_face) ? (1.0 / refractive_index) : refractive_index;
   const double cos_i = -dot(hit_in.normal, unit_dir);
   const double sin_t_squared = ior_ratio * ior_ratio * (1.0 - cos_i * cos_i);

   if (sin_t_squared > 1.0) // total internal reflection
   {
      // clang-format off
      return scatter_data {
         .emission = colour{}, 
         .diffuse = diffuse, 
         .scattered_ray = ray(hit_in.position, reflect(hit_in.normal, unit_dir))
      };
      // clang-format on
   }

   if (random_double() < schlick(cos_i, ior_ratio))
   {
      // clang-format off
      return scatter_data {
         .emission = colour{}, 
         .diffuse = diffuse, 
         .scattered_ray = ray(hit_in.position, reflect(hit_in.normal, unit_dir))
      };
      // clang-format on
   }

   const double cos_t = std::sqrt(1.0 - sin_t_squared);
   vec refracted = ior_ratio * unit_dir + (ior_ratio * cos_i - cos_t) * hit_in.normal;

   // clang-format off
   return scatter_data {
      .emission = colour{}, 
      .diffuse = diffuse, 
      .scattered_ray = ray(hit_in.position, refracted)
   };
   // clang-format on
}
