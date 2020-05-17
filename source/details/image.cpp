#include "details/image.hpp"

#include <cassert>
#include <fstream>

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
   std::ofstream out("image.ppm");

   out << "P3\n" << width() << " " << height() << "\n255\n";
   for (size_t i = 0; i < data.size(); ++i)
   {
      auto const colour = data[i].compute_colour();
      out << static_cast<int>(256 * std::clamp(colour.x, 0.0, 0.999)) << ' '
          << static_cast<int>(256 * std::clamp(colour.y, 0.0, 0.999)) << ' '
          << static_cast<int>(256 * std::clamp(colour.z, 0.0, 0.999)) << '\n';
   }

   out << std::endl;
   out.close();
}
