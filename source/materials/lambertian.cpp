#include "materials/lambertian.hpp"
#include "math/details.hpp"

lambertian::lambertian(colour const &colour_in) noexcept : super(colour_in) {}

auto lambertian::scatter(const ray &ray_in, const hit &hit_in) const -> std::optional<scatter_data>
{
   vec reflected = hit_in.normal + random_unit_vector();

   return std::make_pair(super::emission, ray(hit_in.position, reflected));
}
