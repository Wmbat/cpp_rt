#pragma once

#include "pixel.hpp"

#include <vector>

class image
{
public:
   image() = default;
   image(int width, int height);

   void add_samples(int x, int y, pixel pxl);
   void add_samples(int x, int y, vec const& colour, size_t sample_count);

   image& operator+=(image const& rhs);

   constexpr int width() const noexcept { return w; }
   constexpr int height() const noexcept { return h; }

   void write() const;

private:
   int w;
   int h;

   std::vector<pixel> data;
};
