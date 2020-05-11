#include <glm/vec3.hpp>

#include <cstddef>
#include <iostream>

using point = glm::dvec3;
using colour = point;

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   constexpr int image_width = 1080;
   constexpr int image_height = 720;

   std::cout << "P3\n" << image_width << ' ' << image_height << "\n255\n";

   for (int j = image_height - 1; j >= 0; --j)
   {
      for (int i = 0; i < image_width; ++i)
      {
         colour c(static_cast<double>(i) / (image_width - 1), static_cast<double>(j) / (image_height - 1), 0.25);

         std::cout << static_cast<int>(255.99 * c.r) << ' ' << static_cast<int>(255.99 * c.g) << ' '
                   << static_cast<int>(255.99 * c.b) << '\n';
      }
   }

   return 0;
}
