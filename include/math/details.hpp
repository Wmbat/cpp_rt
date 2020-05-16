#pragma once

#include "math/ortho_nomal_basis.hpp"
#include "vec.hpp"

#include <cmath>
#include <random>

constexpr vec reflect(const vec& normal, const vec& incident) noexcept
{
   return incident - 2 * dot(normal, incident) * normal;
}

inline double reflectance(const norm& normal, const norm& incident, double ior_from, double ior_to) noexcept
{
   const auto ior_ratio = ior_from / ior_to;
   const auto cos_theta_i = -dot(normal, incident);
   const auto sin_theta_t_squared = ior_ratio * ior_ratio * (1 - cos_theta_i * cos_theta_i);
   if (sin_theta_t_squared > 1)
   {
      return 1.0;
   }

   const auto cos_theta_t = sqrt(1 - sin_theta_t_squared);
   auto ray_perp = (ior_from * cos_theta_i - ior_to * cos_theta_t) / (ior_from * cos_theta_i + ior_to * cos_theta_t);
   auto ray_para = (ior_to * cos_theta_i - ior_from * cos_theta_t) / (ior_to * cos_theta_i + ior_from * cos_theta_t);

   return (ray_perp * ray_perp + ray_para * ray_para) / 2;
}

inline vec refract(vec const& normal, vec const& incident, double ior_ratio) noexcept
{
   const double cos_theta = dot(normal, -incident);
   const vec parallel = ior_ratio * (incident + cos_theta * normal);
   const vec perpendicular = -std::sqrt(1.0 - parallel.length_squared()) * normal;

   return parallel + perpendicular;
}

inline double schlick(double cos, double refractive_index)
{
   auto const r_0 = (1 - refractive_index) / (1 + refractive_index);
   auto const r = r_0 * r_0;

   return r + (1 + r) * pow((1 - cos), 5);
}

inline double random_double()
{
   static std::random_device device{};
   static std::mt19937 rng(device());
   static std::uniform_real_distribution<double> dist(0.0, 1.0);

   return dist(rng);
}
inline double random_double(double min, double max)
{
   return min + (max - min) * random_double();
}

inline vec random_unit_vector()
{
   auto a = random_double(0, 2 * M_PI);
   auto z = random_double(-1, 1);
   auto r = sqrt(1 - z * z);

   return vec(r * cos(a), r * sin(a), z);
}

inline vec random_in_unit_sphere()
{
   while (true)
   {
      auto p = vec{random_double(), random_double(), random_double()};
      if (p.length_squared() >= 1)
      {
         continue;
      }
      return p;
   }
}

inline double to_radians(double angle) noexcept
{
   return angle * M_PI / 180;
}

inline vec cone_sample(const norm& direction, double cone_theta, double u, double v)
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

   return normalise(basis.transform(vec(std::cos(random_theta) * radius, std::sin(random_theta) * radius, scale_z)));
}

inline vec hemisphere_sample(const ortho_normal_basis& basis, double u, double v)
{
   const auto theta = 2 * M_PI * u;
   const auto radius_squared = v;
   const auto radius = sqrt(radius_squared);

   return normalise(basis.transform(vec(std::cos(theta) * radius, std::sin(theta) * radius, sqrt(1 - radius_squared))));
}
