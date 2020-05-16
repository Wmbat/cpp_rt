#pragma once

#include "materials/material.hpp"

class diffuse : public material
{
public:
   diffuse(const colour& emission_in, const colour& diffuse_in) noexcept;

   scatter_data scatter(const ray& ray_in, const hit& hit_in, double u, double v) const override;

private:
   colour emission_colour{};
   colour diffuse_colour{};
};
