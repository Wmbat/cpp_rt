#include "ray.hpp"

#include <cstddef>
#include <iostream>

double hit_sphere(point const& center, double radius, ray const& r)
{
   vec dist = r.origin() - center;
   auto a = glm::dot(r.direction(), r.direction());
   auto b = 2.0 * dot(dist, r.direction());
   auto c = dot(dist, dist) - radius * radius;
   auto discriminant = b * b - 4 * a * c;

   if (discriminant < 0)
   {
      return -1.0;
   }
   else
   {
      return (-b - sqrt(discriminant)) / (2.0 * a);
   }
}

colour ray_colour(ray const& r)
{
   double t = hit_sphere(point(0.0, 0.0, -1.0), 0.5, r);
   if (t > 0.0)
   {
      vec unit_dir = glm::normalize(r.at(t) - vec(0.0, 0.0, -1.0));
      return 0.5 * colour(unit_dir.x + 1, unit_dir.y + 1, unit_dir.z + 1);
   }

   vec unit_dir = glm::normalize(r.direction());
   t = 0.5 * (unit_dir.y + 1.0);
   return (1.0 - t) * colour(1.0, 1.0, 1.0) + colour(0.5, 0.7, 1.0) * t;
}

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   constexpr double aspect_ratio = 16.0 / 9.0;

   constexpr int image_width = 1080;
   constexpr int image_height = static_cast<int>(image_width / aspect_ratio);

   std::cout << "P3\n" << image_width << ' ' << image_height << "\n255\n";

   point origin(0.0, 0.0, 0.0);
   vec horizontal(4.0, 0.0, 0.0);
   vec vertical(0.0, 2.25, 0.0);

   point lower_left_corner = origin - horizontal * 0.5 - vertical * 0.5 - vec(0, 0, 1);

   for (int j = image_height - 1; j >= 0; --j)
   {
      for (int i = 0; i < image_width; ++i)
      {
         auto u = static_cast<double>(i) / (image_width - 1);
         auto v = static_cast<double>(j) / (image_height - 1);

         ray r(origin, lower_left_corner + u * horizontal + v * vertical);
         colour pixel_colour = ray_colour(r);

         std::cout << static_cast<int>(255.99 * pixel_colour.r) << ' ' << static_cast<int>(255.99 * pixel_colour.g)
                   << ' ' << static_cast<int>(255.99 * pixel_colour.b) << '\n';
      }
   }

   return 0;
}
