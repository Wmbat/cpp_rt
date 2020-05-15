#include "image.hpp"

#include <cassert>
#include <iostream>

image::image(int width, int height) : w(width), h(height)
{
   data.resize(w * h);
}

void image::add_samples(int x, int y, pixel pxl) noexcept
{
   assert(x >= 0 && x < w);
   assert(y >= 0 && y < h);

   data[x + y * width()].add_sample(pxl);
}

void image::add_samples(int x, int y, vec const& colour, size_t sample_count) noexcept
{
   assert(x >= 0 && x < w);
   assert(y >= 0);
   assert(y < h);

   data[x + y * width()].add_sample(colour, sample_count);
}

image& image::operator+=(image const& rhs)
{
   assert(data.size() == rhs.data.size());

   for (size_t i = 0; i < data.size(); ++i)
   {
      data[i].add_sample(rhs.data[i]);
   }

   return *this;
}

void image::write() const
{
   std::cout << "P3\n" << width() << " " << height() << "\n255\n";
   for (size_t i = 0; i < data.size(); ++i)
   {
      auto const colour = data[i].compute_colour();
      std::cout << static_cast<int>(255.99 * colour.x) << ' ' << static_cast<int>(255.99 * colour.y) << ' '
                << static_cast<int>(255.99 * colour.z) << '\n';
   }
}
