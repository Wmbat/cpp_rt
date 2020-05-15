#include "scene.hpp"
#include "hit_record.hpp"
#include "image.hpp"
#include "math/details.hpp"
#include "math/vec.hpp"
#include "pixel.hpp"
#include "ray.hpp"

#include <algorithm>
#include <cmath>
#include <cstddef>
#include <future>
#include <iostream>
#include <optional>
#include <random>
#include <thread>

image scene::render(camera const& cam, render_settings const& settings)
{
   size_t curr_samples = 0;
   auto launch = [&] {
      return std::async(std::launch::async, [&] {
         ++curr_samples;

         image img{settings.window_width, settings.window_height};
         for (int y = settings.window_height - 1; y >= 0; --y)
         {
            for (int x = 0; x < settings.window_width; ++x)
            {
               auto const u = (x + random_double()) / (settings.window_width - 1);
               auto const v = (y + random_double()) / (settings.window_height - 1);
               auto const ray = cam.shoot_ray(u, v);

               int const new_y = std::abs(y - (settings.window_height - 1));
               img.add_samples(x, new_y, radiance(ray, 32), 1);
            }
         };

         return img;
      });
   };

   image final_img{settings.window_width, settings.window_height};

   std::vector<std::future<image>> futures;
   futures.reserve(settings.sample_count);

   size_t num_done = 0;
   do
   {
      auto const num_left = settings.sample_count - curr_samples;
      auto const num_space_cpu = std::thread::hardware_concurrency() - futures.size();
      for (size_t i = 0; i < std::min<size_t>(num_left, num_space_cpu); ++i)
      {
         futures.emplace_back(launch());
      }

      auto const finished_img = std::find_if(futures.begin(), futures.end(), [&](auto& f) {
         return f.wait_for(std::chrono::milliseconds(1)) == std::future_status::ready;
      });

      if (finished_img != futures.end())
      {
         ++num_done;
         final_img += finished_img->get();
         futures.erase(finished_img);
      }
      else
      {
         std::this_thread::sleep_for(std::chrono::milliseconds(250));
      }
   } while (num_done < settings.sample_count);

   return final_img;
}

void scene::add_sphere(const sphere& sphere_in, std::shared_ptr<material> material)
{
   spheres.push_back(sphere_in);
   sphere_materials.push_back(material);
}

void scene::add_sphere(sphere&& sphere_in, std::shared_ptr<material> material)
{
   spheres.push_back(std::move(sphere_in));
   sphere_materials.push_back(material);
}

colour scene::radiance(ray const& ray_in, size_t depth)
{
   if (depth <= 0)
   {
      return {0.0, 0.0, 0.0};
   }

   auto const hit_record = intersect(ray_in);
   if (!hit_record)
   {
      vec unit_dir = ray_in.direction() / ray_in.direction().length();
      auto t = 0.5 * (unit_dir.y + 1.0);
      return (1.0 - t) * vec(1.0, 1.0, 1.0) + t * vec(0.5, 0.7, 1.0);
   }

   auto const& hit = hit_record->hit;
   auto const& mat = hit_record->material;

   auto scatter = mat->scatter(ray_in, hit);
   if (scatter)
   {
      return scatter->first * radiance(scatter->second, depth - 1);
   }
   else
   {
      return vec{};
   }
}

std::optional<hit_record> scene::intersect(ray const& ray_in)
{
   return sphere_intersect(ray_in, std::numeric_limits<double>::max());
}

std::optional<hit_record> scene::sphere_intersect(ray const& ray_in, double nearer_than)
{
   double current_nearest = nearer_than;
   std::optional<size_t> nearest_index;

   double epsilon = 0.001;

   for (size_t i = 0; i < spheres.size(); ++i)
   {
      auto const oc = ray_in.origin() - spheres[i].center;
      auto const a = ray_in.direction().length_squared();
      auto const b = dot(oc, ray_in.direction());
      auto const c = oc.length_squared() - spheres[i].radius * spheres[i].radius;
      auto const discriminant = b * b - a * c;

      if (discriminant > 0)
      {
         auto const determinant = std::sqrt(discriminant);
         auto const t_min = (-b - determinant) / a;
         auto const t_max = (-b + determinant) / a;

         if (t_min < epsilon && t_max < epsilon)
         {
            continue;
         }

         auto const t = t_min > epsilon ? t_min : t_max;
         if (t < current_nearest)
         {
            nearest_index = i;
            current_nearest = t_min;
         }
      }
   }

   if (!nearest_index)
      return {};

   auto hit_position = ray_in.position_along(current_nearest);
   auto normal = (hit_position - spheres[*nearest_index].center) / spheres[*nearest_index].radius;
   bool inside = dot(normal, ray_in.direction()) > 0;

   if (inside)
   {
      normal = -normal;
   }

   // clang-format off
   return hit_record
      {
         .hit = 
         {
            .position = hit_position, 
            .normal = normal, 
            .distance = current_nearest, 
            .inside = inside
         },
         .material = sphere_materials[*nearest_index]
      };
   // clang-format on
}
