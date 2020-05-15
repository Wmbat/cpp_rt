#pragma once

#include "hit.hpp"
#include "material.hpp"
#include "math/details.hpp"

class metal : public material
{
   using super = material;

public:
   metal(colour const& colour_in, double roughness_in) noexcept;

   std::optional<super::scatter_data> scatter(ray const& ray_in, hit const& hit_in) const override;

private:
   double roughness;
};
