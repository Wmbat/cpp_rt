#include <cpp_pt/details/pixel.hpp>
#include <cpp_pt/details/settings.hpp>
#include <cpp_pt/materials/dielectric.hpp>
#include <cpp_pt/materials/diffuse.hpp>
#include <cpp_pt/materials/diffuse_light.hpp>
#include <cpp_pt/materials/metallic.hpp>
#include <cpp_pt/math/details.hpp>
#include <cpp_pt/math/vec.hpp>
#include <cpp_pt/camera.hpp>

#include <iostream>
#include <memory>

import scene;
import renderables.sphere;
import renderables.triangle;

void random_sphere_scene(const render_settings& settings);
void cornell_box_scene(const render_settings& settings);

auto main([[maybe_unused]] int argc, [[maybe_unused]] char* argv[]) -> int
{
   render_settings settings{};
   settings.window_height = 1080u;
   settings.u_samples = 1u;
   settings.v_samples = 1u;
   settings.bounce_depth = 3;
   settings.sample_count = 6u;

   random_sphere_scene(settings);

   // settings.window_width = 556;
   // settings.window_height = 556;
   // cornell_box_scene(settings);

   return 0;
}

void random_sphere_scene(const render_settings& settings)
{
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

   scene.set_environment_colour({135 / 256.0, 206 / 256.0, 235 / 256.0});

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
                  std::make_unique<metallic>(
                     colour{}, random_vec() * random_vec(), random_double(0.0, 0.5)));
            }
            else
            {
               scene.add_sphere(sphere{.center = {-1.0, 0.0, -1.0}, .radius = 0.5},
                  std::make_unique<dielectric>(colour{1.0, 1.0, 1.0}, 1.5));
            }
         }
      }
   }

   scene.add_sphere(sphere{.center{-1000.0, 1000.0, 100.0}, .radius = 100},
      std::make_unique<diffuse_light>(colour{20.0, 20.0, 20.0}));
   scene.add_sphere(sphere{.center = {0.0, -1000.0, 0.0}, .radius = 1000},
      std::make_unique<diffuse>(colour{}, colour{0.5, 0.5, 0.5}));
   scene.add_sphere(sphere{.center = {0.0, 1.0, 0.0}, .radius = 1},
      std::make_unique<diffuse>(colour{}, colour{0.4, 0.2, 0.1}));
   scene.add_sphere(sphere{.center = {-4.0, 1.0, 0.0}, .radius = 1},
      std::make_unique<metallic>(colour{}, colour{0.7, 0.6, 0.5}, 0.0));
   scene.add_sphere(sphere{.center = {4.0, 1.0, 0.0}, .radius = 1},
      std::make_unique<dielectric>(colour{1.0, 1.0, 1.0}, 1.5));

   auto const img = scene.render(cam, settings);

   img.write();
}

void cornell_box_scene(const render_settings& settings)
{
   vec eye{278, 278, -800};
   vec look_at{278, 278, 0};

   // clang-format off
   camera::create_info info
   {
      .eye = eye,
      .look_at = look_at,
      .up = {0.0, 1.0, 0.0},
      .vertical_fov = 40,
      .aspect_ratio = (double)settings.window_width / settings.window_height,
      .aperture = 0.0,
      .focus_distance = 10
   };
   // clang-format on

   camera cam{info};

   scene scene{};
   // scene.set_environment_colour({135 / 256.0, 206 / 256.0, 235 / 256.0});

   // clang-format on

   // back rectangle
   scene.add_triangle(vec(0, 0, 555), vec(0, 555, 555), vec(555, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   scene.add_triangle(vec(555, 555, 555), vec(555, 0, 555), vec(0, 555, 555),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   // bottom rectangle
   scene.add_triangle(vec(0, 0, 555), vec(0, 0, 0), vec(555, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   scene.add_triangle(vec(555, 0, 0), vec(555, 0, 555), vec(0, 0, 0),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   // top rectangle
   scene.add_triangle(vec(0, 555, 555), vec(555, 555, 555), vec(0, 555, 0),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   scene.add_triangle(vec(555, 555, 0), vec(555, 555, 555), vec(0, 555, 0),
      std::make_unique<diffuse>(colour{}, colour{.73, .73, .73}));

   // left rectangle
   scene.add_triangle(vec(0, 0, 0), vec(0, 555, 0), vec(0, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.12, .45, .15}));

   scene.add_triangle(vec(0, 555, 555), vec(0, 555, 0), vec(0, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.12, .45, .15}));

   // right rectangle
   scene.add_triangle(vec(555, 0, 0), vec(555, 555, 0), vec(555, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.65, .05, .05}));

   scene.add_triangle(vec(555, 555, 555), vec(555, 555, 0), vec(555, 0, 555),
      std::make_unique<diffuse>(colour{}, colour{.65, .05, .05}));

   // light rectangle
   scene.add_triangle(vec(213, 554, 332), vec(343, 554, 332), vec(213, 555, 227),
      std::make_unique<diffuse_light>(colour{25, 25, 25}));

   scene.add_triangle(vec(343, 554, 227), vec(343, 554, 332), vec(213, 554, 227),
      std::make_unique<diffuse_light>(colour{25, 25, 25}));

   auto const img = scene.render(cam, settings);
   img.write();
};
