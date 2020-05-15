#pragma once

#include "hit.hpp"
#include "material.hpp"
#include "math/details.hpp"

class lambertian : public material
{
   using super = material;

public:
   lambertian(colour const& colour_in) noexcept;

   std::optional<super::scatter_data> scatter(ray const& ray_in, hit const& hit_in) const override;
};
