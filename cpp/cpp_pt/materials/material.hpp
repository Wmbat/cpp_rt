#pragma once

#include <cpp_pt/hit.hpp>
#include <cpp_pt/math/vec.hpp>
#include <cpp_pt/ray.hpp>

#include <optional>
#include <utility>

class material
{
public:
   struct scatter_data
   {
      vec emission{};
      vec diffuse{};

      ray scattered_ray{};
   };

public:
   material() = default;
   material(const material&) = default;
   material(material&&) = default;
   virtual ~material() = default;

   auto operator=(const material&) -> material& = default;
   auto operator=(material&&) -> material& = default;

   [[nodiscard]] virtual auto scatter(
      const ray& ray_in, const hit& hit_in, double u, double v) const -> scatter_data = 0;
};
