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

   image& operator+=(const image& rhs);

   constexpr int width() const noexcept { return w; }
   constexpr int height() const noexcept { return h; }

   void write() const;

private:
   int w;
   int h;

   std::vector<pixel> pixels;
};
