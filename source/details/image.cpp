#include "details/image.hpp"

#include <bits/ranges_algo.h>
#include <cassert>
#include <fstream>

image::image(int width, int height) : w(width), h(height)
{
   pixels.resize(w * h);
}

void image::add_samples(int x, int y, const pixel& pxl)
{
   assert(x >= 0 && x < w);
   assert(y >= 0 && y < h);

   pixels[x + y * width()].add_sample(pxl);
}

void image::add_samples(int x, int y, const vec& colour, size_t sample_count)
{
   assert(x >= 0 && x < w);
   assert(y >= 0);
   assert(y < h);

   pixels[x + y * width()].add_sample(colour, sample_count);
}

image& image::operator+=(const image& rhs)
{
   assert(pixels.size() == rhs.pixels.size());

   for (size_t i = 0; i < pixels.size(); ++i)
   {
      pixels[i].add_sample(rhs.pixels[i]);
   }

   return *this;
}

void image::write() const
{
   std::ofstream out("image.ppm");

   out << "P3\n" << width() << " " << height() << "\n255\n";

   std::ranges::for_each(pixels, [&out](const auto& pxl) {
      out << pxl;
   });

   out << std::endl;
   out.close();
}
