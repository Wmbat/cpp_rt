#include <cpp_pt/details/image.hpp>

#include <bits/ranges_algo.h>
#include <cassert>
#include <fstream>

image::image(int width, int height) : m_width{width}, m_height{height}
{
   m_pixels.resize(m_width * m_height);
}

void image::add_samples(int x, int y, const pixel& pxl)
{
   assert(x >= 0 && x < m_width);
   assert(y >= 0 && y < m_height);

   m_pixels[x + y * width()].add_sample(pxl);
}

void image::add_samples(int x, int y, const vec& colour, size_t sample_count)
{
   assert(x >= 0 && x < width());
   assert(y >= 0);
   assert(y < height());

   m_pixels[x + y * width()].add_sample(colour, sample_count);
}

auto image::operator+=(const image& rhs) -> image&
{
   assert(m_pixels.size() == rhs.m_pixels.size());

   for (size_t i = 0; i < m_pixels.size(); ++i)
   {
      m_pixels[i].add_sample(rhs.m_pixels[i]);
   }

   return *this;
}

void image::write() const
{
   std::ofstream out("image.ppm");

   out << "P3\n" << width() << " " << height() << "\n255\n";

   for (const auto& pxl : m_pixels)
   {
      out << pxl;
   }

   out << std::endl;
   out.close();
}
