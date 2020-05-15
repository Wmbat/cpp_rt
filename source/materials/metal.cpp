#include "materials/metal.hpp"

metal::metal(colour const &colour_in, double roughness_in) noexcept : super(colour_in), roughness(roughness_in) {}

auto metal::scatter(const ray &ray_in, const hit &hit_in) const -> std::optional<scatter_data>
{
   vec reflected = reflect(ray_in.direction() / ray_in.direction().length(), hit_in.normal);
   ray scattered = ray(hit_in.position, reflected + roughness * random_in_unit_sphere());

   if (dot(reflected, hit_in.normal) > 0)
   {
      return std::make_pair(super::emission, scattered);
   }
   else
   {
      return {};
   }
}
