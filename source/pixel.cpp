#include "pixel.hpp"

void pixel::add_sample(pixel const& pixel) noexcept
{
   colour += pixel.colour;
   samples_count += pixel.samples_count;
}

void pixel::add_sample(vec const& colour_in, std::size_t sample_count_in) noexcept
{
   colour += colour_in;
   samples_count += sample_count_in;
}

vec pixel::compute_colour() const noexcept
{
   if (samples_count == 0)
   {
      return colour;
   }
   else
   {
      return colour * (1.0 / samples_count);
   }
}
