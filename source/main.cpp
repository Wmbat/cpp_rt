#include "materials/dielectric.hpp"
#include "materials/diffuse.hpp"
#include "materials/metallic.hpp"
#include "math/details.hpp"
#include "scene.hpp"
#include "settings.hpp"

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   render_settings settings{};
   camera cam{};
   scene scene{};

   // clang-format off
   scene.add_sphere(
         sphere { .center = {0.0, 0, -1}, .radius = 0.5 },
         std::make_unique<diffuse>(colour{}, colour{0.1, 0.2, 0.5})
   );
   scene.add_sphere(
         sphere { .center = {0.0, -100.5, -1}, .radius = 100 },
         std::make_unique<diffuse>(colour{}, colour{0.8, 0.8, 0.0})
   );
   scene.add_sphere(
         sphere { .center = {1.0, 0.0, -1.0}, .radius = 0.5 },
         std::make_unique<metallic>(colour{}, colour{0.8, 0.8, 0.8}, 0.0)
   );
   scene.add_sphere(
         sphere { .center = {-1.0, 0.0, -1.0}, .radius = 0.5 },
         std::make_unique<dielectric>(colour{1.0, 1.0, 1.0}, 1.5)
   );
   // clang-format on

   auto const img = scene.render(cam, settings);

   img.write();

   return 0;
}
