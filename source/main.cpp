#include "materials/dielectric.hpp"
#include "materials/diffuse.hpp"
#include "materials/metallic.hpp"
#include "math/details.hpp"
#include "scene.hpp"
#include "settings.hpp"

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   render_settings settings{};

   vec eye{13, 2, 3.0};
   vec look_at{0.0, 0.0, 0.0};

   // clang-format off
   camera::create_info info
   {
      .eye = eye,
      .look_at = look_at,
      .up = {0.0, 1.0, 0.0},
      .vertical_fov = 20,
      .aspect_ratio = (double)settings.window_width / settings.window_height,
      .aperture = 0.1,
      .focus_distance = 10
   };
   // clang-format on

   camera cam{info};
   scene scene{};

   for (int a = -11; a < 11; ++a)
   {
      for (int b = -11; b < 11; ++b)
      {
         const double random_value = random_double();
         vec center{a + 0.9 * random_double(), 0.2, b + 0.9 * random_double()};
         if ((center - vec{4, 0.2, 0}).length() > 0.9)
         {
            if (random_value < 0.7)
            {
               scene.add_sphere(sphere{.center = center, .radius = 0.2},
                  std::make_unique<diffuse>(colour{}, random_vec() * random_vec()));
            }
            else if (random_value < 0.9)
            {
               scene.add_sphere(sphere{.center = center, .radius = 0.2},
                  std::make_unique<metallic>(colour{}, random_vec() * random_vec(), random_double(0.0, 0.5)));
            }
            else
            {
               scene.add_sphere(sphere{.center = {-1.0, 0.0, -1.0}, .radius = 0.5},
                  std::make_unique<dielectric>(colour{1.0, 1.0, 1.0}, 1.5));
            }
         }
      }
   }

   // clang-format off
   scene.add_sphere(
         sphere{ .center = {0.0, -1000.0, 0.0}, .radius = 1000 },
         std::make_unique<diffuse>(colour{}, colour{0.5, 0.5, 0.5})
   );

   scene.add_sphere(
         sphere { .center = {0.0, 1.0, 0.0}, .radius = 1 },
         std::make_unique<diffuse>(colour{}, colour{0.4, 0.2, 0.1})
   );
   scene.add_sphere(
         sphere { .center = {-4.0, 1.0, 0.0}, .radius = 1 },
         std::make_unique<metallic>(colour{}, colour{0.7, 0.6, 0.5}, 0.0)
   );
   scene.add_sphere(
         sphere { .center = {4.0, 1.0, 0.0}, .radius = 1 },
         std::make_unique<dielectric>(colour{1.0, 1.0, 1.0}, 1.5)
   );
   // clang-format on

   auto const img = scene.render(cam, settings);

   img.write();

   return 0;
}
