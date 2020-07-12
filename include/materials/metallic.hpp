#pragma once

#include "materials/material.hpp"

class metallic : public material
{
public:
   metallic(const colour& emission_in, const colour& diffuse_in, double roughness_in) noexcept;

   [[nodiscard]] auto scatter(const ray& ray_in, const hit& hit_in, double u, double v) const
      -> scatter_data override;

private:
   const colour emission;
   const colour diffuse;
   const double roughness;
};
