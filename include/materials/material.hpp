#pragma once

#include "hit.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <optional>
#include <utility>

struct material_info
{
   bool operator==(material_info const& rhs) const = default;

   vec emission{};
   vec diffuse{};

   double refraction_index{1.0};
   double reflectivity{-1.0};
   double reflection_angle{0.0};
};
