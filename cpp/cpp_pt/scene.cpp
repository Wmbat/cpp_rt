#include <cpp_pt/details/image.hpp>
#include <cpp_pt/details/pixel.hpp>
#include <cpp_pt/details/progress_keeper.hpp>
#include <cpp_pt/hit_record.hpp>
#include <cpp_pt/math/details.hpp>
#include <cpp_pt/math/vec.hpp>
#include <cpp_pt/ray.hpp>
#include <cpp_pt/scene.hpp>

#include <cmath>
#include <cstddef>
#include <future>
#include <iostream>
#include <memory>
#include <optional>
#include <random>
#include <thread>

static constinit double limit = std::numeric_limits<double>::infinity();

auto scene::render(const camera& cam, const render_settings& settings) -> image
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
               auto rad =
                  radiance(ray, settings.u_samples, settings.v_samples, settings.bounce_depth);
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
   spheres.push_back(sphere_in);
   sphere_mats.push_back(std::move(material));
}

void scene::add_triangle(const triangle& triangle_in, std::unique_ptr<material> material)
{
   triangles.push_back(triangle_in);
   triangle_mats.push_back(std::move(material));
}

void scene::add_triangle(triangle&& triangle_in, std::unique_ptr<material> material)
{
   triangles.push_back(triangle_in);
   triangle_mats.push_back(std::move(material));
}

void scene::add_triangle(
   const vec& v0, const vec& v1, const vec& v2, std::unique_ptr<material> material)
{
   triangles.emplace_back(v0, v1, v2);
   triangle_mats.push_back(std::move(material));
}

void scene::set_environment_colour(const colour& environment_in) noexcept
{
   environment = environment_in;
}

auto scene::radiance(
   const ray& ray_in, size_t u_sample_count, size_t v_sample_count, size_t depth) const -> colour
{
   if (depth <= 0)
   {
      return {0.0, 0.0, 0.0};
   }

   const auto hit_record = intersect(ray_in);
   if (!hit_record)
   {
      return environment;
   }

   const auto& hit = hit_record->hit_data;
   const auto& mat = hit_record->p_mat;

   colour result{};
   for (size_t u_sample = 0; u_sample < u_sample_count; ++u_sample)
   {
      for (size_t v_sample = 0; v_sample < v_sample_count; ++v_sample)
      {
         const double u = (double(u_sample) + random_double()) / double(u_sample_count);
         const double v = (double(v_sample) + random_double()) / double(v_sample_count);

         const auto [emission, diffuse, ray] = mat->scatter(ray_in, hit, u, v);

         result += emission + diffuse * radiance(ray, u_sample_count, v_sample_count, depth - 1);
      }
   }

   return result / static_cast<double>(u_sample_count * v_sample_count);
}

auto scene::intersect(const ray& ray_in) const -> std::optional<hit_record>
{
   const auto sphere_record = sphere_intersect(ray_in, limit);
   const auto triangle_record =
      triangle_intersect(ray_in, sphere_record ? sphere_record->hit_data.distance : limit);

   return triangle_record ? triangle_record : sphere_record;
}

auto scene::triangle_intersect(const ray& ray_in, double nearer_than) const
   -> std::optional<hit_record>
{
   double current_nearest = nearer_than;

   struct nearest
   {
      size_t index;
      double det;
      double u;
      double v;
   };

   std::optional<nearest> nearest;

   constexpr double epsilon = 0.001;
   for (size_t i = 0; i < triangles.size(); ++i)
   {
      const triangle& tri = triangles[i];
      const vec p_vec = cross(ray_in.direction(), tri.v());
      const double determinant = dot(tri.u(), p_vec);

      if (std::fabs(determinant) < epsilon)
      {
         continue;
      }

      const double inverse_determinant = 1.0 / determinant;
      const vec t_vec = ray_in.origin() - tri.vertex<0>();
      const vec q_vec = cross(t_vec, tri.u());

      const double u = dot(t_vec, p_vec) * inverse_determinant;
      const double v = dot(ray_in.direction(), q_vec) * inverse_determinant;

      if ((u < 0.0) | (u > 1.0) | (v < 0.0) | (u + v > 1))
      {
         continue;
      }

      const double t = dot(tri.v(), q_vec) * inverse_determinant;
      if (t > epsilon && t < current_nearest)
      {
         current_nearest = t;
         nearest = {.index = i, .det = determinant, .u = u, .v = v};
      }
   }

   if (!nearest)
   {
      return {};
   }

   const triangle& tri = triangles[nearest->index];
   const bool front_face = nearest->det > epsilon;

   // clang-format off
   return hit_record
      {
         .hit_data = 
         {
            .position = ray_in.position_along(current_nearest), 
            .normal = front_face ? tri.normal() : -tri.normal(), 
            .distance = current_nearest, 
            .front_face = front_face
         },
         .p_mat = triangle_mats[nearest->index].get()
      };
   // clang-format on
}

auto scene::sphere_intersect(const ray& ray_in, double nearer_than) const
   -> std::optional<hit_record>
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
   {
      return {};
   }

   const auto hit_position = ray_in.position_along(current_nearest);
   const auto normal =
      (hit_position - spheres[*nearest_index].center) / spheres[*nearest_index].radius;
   const bool front_face = dot(normal, ray_in.direction()) < 0;

   // clang-format off
   return hit_record
      {
         .hit_data = 
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
