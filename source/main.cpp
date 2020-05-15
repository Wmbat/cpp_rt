#include "math/details.hpp"
#include "scene.hpp"
#include "settings.hpp"

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   camera cam{};
   render_settings settings{};
   scene scene{};

   // clang-format off
   scene.add_sphere(
         sphere { .center = {0.0, 0, -1}, .radius = 0.5 },
         material_info { .emission = colour{}, .diffuse = colour{0.7, 0.3, 0.3} }
   );
   scene.add_sphere(
         sphere { .center = {0.0, -100.5, -1}, .radius = 100 },
         material_info { .emission = colour{}, .diffuse = colour{0.8, 0.8, 0.8} }
   );
   scene.add_sphere(
         sphere { .center = {1.0, 0.0, -1.0}, .radius = 0.5 },
         material_info { 
            .emission = colour{}, 
            .diffuse = colour{0.8, 0.6, 0.2}, 
            .reflectivity = 0.95, 
            .reflection_angle = to_radians(0) 
         }
   );
   // clang-format on

   auto const img = scene.render(cam, settings);

   img.write();

   return 0;
}
