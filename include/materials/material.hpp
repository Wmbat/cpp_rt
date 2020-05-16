#pragma once

#include "hit.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <optional>
#include <utility>

class material
{
public:
   struct scatter_data
   {
      vec emission{};
      vec diffuse{};

      ray scattered_ray{};
   };

public:
   virtual ~material() = default;

   virtual scatter_data scatter(const ray& ray_in, const hit& hit_in, double u, double v) const = 0;
};
