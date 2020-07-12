#pragma once

#include "materials/material.hpp"

class dielectric : public material
{
public:
   dielectric(const colour &diffuse_in, double refractive_index_in);

   [[nodiscard]] auto scatter(const ray &ray_in, const hit &hit_in, double u, double v) const
      -> scatter_data override;

private:
   const colour diffuse;
   const double refractive_index;
};
