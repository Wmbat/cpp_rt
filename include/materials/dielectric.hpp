#pragma once

#include "materials/material.hpp"

class dielectric : public material
{
public:
   dielectric(const colour &diffuse_in, double refractive_index_in);

   scatter_data scatter(const ray &ray_in, const hit &hit_in, double u, double v) const override;

private:
   colour diffuse;
   double refractive_index;
};
