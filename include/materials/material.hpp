#pragma once

#include "hit.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <optional>
#include <utility>

class material
{
protected:
   material(colour const& colour_in) : emission(colour_in) {}

public:
   using scatter_data = std::pair<colour, ray>;

   virtual std::optional<scatter_data> scatter(ray const& ray_in, hit const& hit) const = 0;

protected:
   colour emission;
};
