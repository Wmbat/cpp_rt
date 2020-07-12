#pragma once

#include "materials/material.hpp"

class diffuse_light : public material
{
public:
   diffuse_light(const colour &emission_in) noexcept;

   [[nodiscard]] auto scatter(const ray &ray_in, const hit &hit_in, double u, double v) const
      -> scatter_data override;

private:
   const colour emission{};
};
