#pragma once

#include <cpp_pt/math/vec.hpp>

#include <cstddef>
#include <ostream>

class pixel
{
public:
   void add_sample(pixel const& pixel) noexcept;
   void add_sample(vec const& colour, std::size_t num) noexcept;

   [[nodiscard]] auto compute_colour() const noexcept -> vec;

private:
   vec colour{0.0, 0.0, 0.0};
   size_t samples_count{0};
};

inline auto operator<<(std::ostream& out, const pixel& pxl) -> std::ostream&
{
   const colour col = pxl.compute_colour();

   return out << static_cast<int>(256 * std::clamp(col.x(), 0.0, 0.999)) << ' '
              << static_cast<int>(256 * std::clamp(col.y(), 0.0, 0.999)) << ' '
              << static_cast<int>(256 * std::clamp(col.z(), 0.0, 0.999)) << '\n';
}
