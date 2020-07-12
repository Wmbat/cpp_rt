#pragma once

#include "pixel.hpp"

#include <vector>

class image
{
public:
   image() = default;
   image(int width, int height);

   void add_samples(int x, int y, const pixel& pxl);
   void add_samples(int x, int y, const vec& colour, size_t sample_count);

   auto operator+=(const image& rhs) -> image&;

   [[nodiscard]] constexpr auto width() const noexcept -> int { return m_width; }
   [[nodiscard]] constexpr auto height() const noexcept -> int { return m_height; }

   void write() const;

private:
   int m_width{0};
   int m_height{0};

   std::vector<pixel> m_pixels;
};
