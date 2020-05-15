#pragma once

#include "vec.hpp"

#include <cmath>
#include <random>

inline vec reflect(vec const& v, vec const& n)
{
   return v - 2 * dot(v, n) * n;
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
