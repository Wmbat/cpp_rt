#pragma once

#include <cpp_pt/details/pixel.hpp>

#include <cassert>
#include <fstream>
#include <vector>

class image
{
public:
   image() = default;
   image(int width, int height) : m_width{width}, m_height{height}
   {
      m_pixels.resize(m_width * m_height);
   }

   void add_samples(int x, int y, const pixel& pxl)
   {
      assert(x >= 0 && x < m_width);  // NOLINT
      assert(y >= 0 && y < m_height); // NOLINT

      m_pixels[x + y * width()].add_sample(pxl);
   }
   void add_samples(int x, int y, const vec& colour, size_t sample_count)
   {
      assert(x >= 0 && x < width()); // NOLINT
      assert(y >= 0);                // NOLINT
      assert(y < height());          // NOLINT

      m_pixels[x + y * width()].add_sample(colour, sample_count);
   }

   auto operator+=(const image& rhs) -> image&
   {
      assert(m_pixels.size() == rhs.m_pixels.size());

      for (size_t i = 0; i < m_pixels.size(); ++i)
      {
         m_pixels[i].add_sample(rhs.m_pixels[i]);
      }

      return *this;
   }

   [[nodiscard]] constexpr auto width() const noexcept -> int { return m_width; }
   [[nodiscard]] constexpr auto height() const noexcept -> int { return m_height; }

   void write() const
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

private:
   int m_width{0};
   int m_height{0};

   std::vector<pixel> m_pixels{};
};
