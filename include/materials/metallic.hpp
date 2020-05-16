#pragma once

#include "materials/material.hpp"

class metallic : public material
{
public:
   metallic(const colour& emission_in, const colour& diffuse_in, double roughness_in) noexcept;

   scatter_data scatter(const ray& ray_in, const hit& hit_in, double u, double v) const override;

private:
   colour emission;
   colour diffuse;
   double roughness;
};
