#include "materials/lambertian.hpp"
#include "materials/metal.hpp"
#include "scene.hpp"
#include "settings.hpp"
#include <memory>

int main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[])
{
   camera cam{};
   render_settings settings{};
   scene scene{};

   scene.add_sphere(sphere{.center = {0.0, 0, -1}, .radius = 0.5}, std::make_shared<lambertian>(colour(0.7, 0.3, 0.3)));
   scene.add_sphere(
      sphere{.center = {0.0, -100.5, -1}, .radius = 100}, std::make_shared<lambertian>(colour(0.8, 0.8, 0.0)));
   scene.add_sphere(sphere{.center = {1.0, 0, -1}, .radius = 0.5}, std::make_shared<metal>(colour(0.3, 0.3, 0.3), 5.0));
   scene.add_sphere(sphere{.center = {-1.0, 0, -1}, .radius = 0.5}, std::make_shared<metal>(colour(0.3, 0.3, 0.3), 0.0));

   auto const img = scene.render(cam, settings);

   img.write();

   return 0;
}
