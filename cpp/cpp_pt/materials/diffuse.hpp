#pragma once

#include <cpp_pt/materials/material.hpp>

class diffuse : public material
{
public:
   diffuse(const colour& emission_in, const colour& diffuse_in) noexcept;

   [[nodiscard]] auto scatter(const ray& ray_in, const hit& hit_in, double u, double v) const
      -> scatter_data override;

private:
   const colour emission_colour{};
   const colour diffuse_colour{};
};
