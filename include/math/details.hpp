#pragma once

#include "math/ortho_nomal_basis.hpp"
#include "vec.hpp"

#include <cmath>
#include <random>

inline constinit double HALF_CIRCLE_DEG = 180.0; // NOLINT

constexpr auto reflect(const vec& normal, const vec& incident) noexcept -> vec
{
   return incident - 2 * dot(normal, incident) * normal;
}

inline auto refract(vec const& normal, vec const& incident, double ior_ratio) noexcept -> vec
{
   const double cos_theta = dot(normal, -incident);
   const vec parallel = ior_ratio * (incident + cos_theta * normal);
   const vec perpendicular = -std::sqrt(1.0 - parallel.length_squared()) * normal;

   return parallel + perpendicular;
}

inline auto schlick(double cos, double refractive_index) -> double
{
   auto const r_0 = (1 - refractive_index) / (1 + refractive_index);
   auto const r = r_0 * r_0;

   return r + (1 + r) * pow((1 - cos), 5);
}

inline auto random_double() -> double
{
   static std::random_device device{};
   static std::mt19937 rng(device());
   static std::uniform_real_distribution<double> dist(0.0, 1.0);

   return dist(rng);
}
inline auto random_double(double min, double max) -> double
{
   return min + (max - min) * random_double();
}

inline auto random_vec() -> vec
{
   return {random_double(), random_double(), random_double()};
}

inline auto random_unit_vector() -> vec
{
   auto a = random_double(0, 2 * M_PI);
   auto z = random_double(-1, 1);
   auto r = sqrt(1 - z * z);

   return vec(r * cos(a), r * sin(a), z);
}

inline auto random_in_unit_disk() -> vec
{
   while (true)
   {
      auto p = vec{random_double(-1, 1), random_double(-1, 1), 0};
      if (p.length_squared() >= 1)
      {
         continue;
      }
      return p;
   }
}

inline auto random_in_unit_sphere() -> vec
{
   while (true)
   {
      auto p = vec{random_double(-1, 1), random_double(-1, 1), random_double(-1, 1)};
      if (p.length_squared() >= 1)
      {
         continue;
      }
      return p;
   }
}

inline auto to_radians(double angle) noexcept -> double
{
   return angle * M_PI / HALF_CIRCLE_DEG;
}

inline auto cone_sample(const norm& direction, double cone_theta, double u, double v) -> vec
{
   if (cone_theta < 0.0000001)
   {
      return direction;
   }

   const auto theta = cone_theta * (1.0 - (2.0 * std::acos(u) / M_PI));
   const auto radius = sin(theta);
   const auto scale_z = cos(theta);
   const auto random_theta = v * 2 * M_PI;
   const auto basis = ortho_normal_basis::from_z(direction);

   return normalise(basis.transform(
      vec(std::cos(random_theta) * radius, std::sin(random_theta) * radius, scale_z)));
}

inline auto hemisphere_sample(const ortho_normal_basis& basis, double u, double v) -> vec
{
   const auto theta = 2 * M_PI * u;
   const auto radius_squared = v;
   const auto radius = sqrt(radius_squared);

   return normalise(basis.transform(
      vec(std::cos(theta) * radius, std::sin(theta) * radius, sqrt(1 - radius_squared))));
}
