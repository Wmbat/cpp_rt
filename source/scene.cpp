#include "scene.hpp"
#include "details/image.hpp"
#include "details/pixel.hpp"
#include "details/progress_keeper.hpp"
#include "hit_record.hpp"
#include "math/details.hpp"
#include "math/vec.hpp"
#include "ray.hpp"

#include <algorithm>
#include <cmath>
#include <cstddef>
#include <future>
#include <iostream>
#include <memory>
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
               auto rad = radiance(ray, settings.u_samples, settings.v_samples, settings.bounce_depth);
               img.add_samples(x, new_y, rad, 1);
            }
         };

         return img;
      });
   };

   image final_img{settings.window_width, settings.window_height};

   std::vector<std::future<image>> futures;
   futures.reserve(settings.sample_count);

   size_t num_done = 0;
   progress_keeper progress_keeper{settings.sample_count};
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
         progress_keeper.update(num_done);
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

void scene::add_sphere(const sphere& sphere_in, std::unique_ptr<material> material)
{
   spheres.push_back(sphere_in);
   sphere_mats.push_back(std::move(material));
}

void scene::add_sphere(sphere&& sphere_in, std::unique_ptr<material> material)
{
   spheres.push_back(std::move(sphere_in));
   sphere_mats.push_back(std::move(material));
}

colour scene::radiance(const ray& ray_in, size_t u_sample_count, size_t v_sample_count, size_t depth) const
{
   if (depth <= 0)
   {
      return {0.0, 0.0, 0.0};
   }

   const auto hit_record = intersect(ray_in);
   if (!hit_record)
   {
      vec unit_direction = normalise(ray_in.direction());
      auto t = 0.5 * (unit_direction.y + 1.0);
      return (1.0 - t) * colour(1.0, 1.0, 1.0) + t * colour(0.5, 0.7, 1.0);
   }

   const auto& hit = hit_record->hit;
   const auto& mat = hit_record->p_mat;

   colour result{};

   for (size_t u_sample = 0; u_sample < u_sample_count; ++u_sample)
   {
      for (size_t v_sample = 0; v_sample < v_sample_count; ++v_sample)
      {
         const double u = (u_sample + random_double()) / static_cast<double>(u_sample_count);
         const double v = (v_sample + random_double()) / static_cast<double>(v_sample_count);

         const auto [emission, diffuse, ray] = mat->scatter(ray_in, hit, u, v);

         result += emission + diffuse * radiance(ray, u_sample_count, v_sample_count, depth - 1);
      }
   }

   return result / (u_sample_count * v_sample_count);
}

std::optional<hit_record> scene::intersect(ray const& ray_in) const
{
   return sphere_intersect(ray_in, std::numeric_limits<double>::max());
}

std::optional<hit_record> scene::sphere_intersect(ray const& ray_in, double nearer_than) const
{
   double current_nearest = nearer_than;
   std::optional<size_t> nearest_index;

   constexpr double epsilon = 0.001;
   for (size_t i = 0; i < spheres.size(); ++i)
   {
      const auto oc = ray_in.origin() - spheres[i].center;
      const auto a = ray_in.direction().length_squared();
      const auto b = dot(oc, ray_in.direction());
      const auto c = oc.length_squared() - spheres[i].radius * spheres[i].radius;
      const auto discriminant = b * b - a * c;

      if (discriminant > 0)
      {
         const auto determinant = std::sqrt(discriminant);
         const auto t_min = (-b - determinant) / a;
         const auto t_max = (-b + determinant) / a;

         if (t_min < current_nearest && t_min > epsilon)
         {
            nearest_index = i;
            current_nearest = t_min;
         }

         if (t_max < current_nearest && t_max > epsilon)
         {
            nearest_index = i;
            current_nearest = t_max;
         }
      }
   }

   if (!nearest_index)
      return {};

   const auto hit_position = ray_in.position_along(current_nearest);
   const auto normal = (hit_position - spheres[*nearest_index].center) / spheres[*nearest_index].radius;
   const bool front_face = dot(normal, ray_in.direction()) < 0;

   // clang-format off
   return hit_record
      {
         .hit = 
         {
            .position = hit_position, 
            .normal = front_face ? normal : -normal, 
            .distance = current_nearest, 
            .front_face = front_face
         },
         .p_mat = sphere_mats[*nearest_index].get()
      };
   // clang-format on
}
