#pragma once

#include "math/vec.hpp"

#include <cstddef>

class pixel
{
public:
   void add_sample(pixel const& pixel) noexcept;
   void add_sample(vec const& colour, std::size_t num) noexcept;

   [[nodiscard]] vec compute_colour() const noexcept;

private:
   vec colour{0.0, 0.0, 0.0};
   size_t samples_count{0};
};
