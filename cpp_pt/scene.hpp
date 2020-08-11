#pragma once

#include <cpp_pt/details/camera.hpp>
#include <cpp_pt/details/image.hpp>
#include <cpp_pt/details/settings.hpp>
#include <cpp_pt/hit_record.hpp>
#include <cpp_pt/materials/material.hpp>
#include <cpp_pt/ray.hpp>
#include <cpp_pt/renderables/sphere.hpp>
#include <cpp_pt/renderables/triangle.hpp>

#include <memory>
#include <optional>
#include <random>
#include <vector>

class scene
{
public:
   [[nodiscard]] auto render(const camera& cam, const render_settings& settings) -> image;

   void add_sphere(const sphere& sphere_in, std::unique_ptr<material> p_mat);
   void add_sphere(sphere&& sphere_in, std::unique_ptr<material> p_mat);
   void add_triangle(const triangle& triangle_in, std::unique_ptr<material> p_mat);
   void add_triangle(triangle&& triangle_in, std::unique_ptr<material> p_mat);
   void add_triangle(const vec& v0, const vec& v1, const vec& v2, std::unique_ptr<material> p_mat);

   void set_environment_colour(const colour& environment_in) noexcept;

private:
   [[nodiscard]] auto radiance(ray const& r, size_t u_samples, size_t v_samples, size_t depth) const
      -> colour;

   [[nodiscard]] auto intersect(const ray& r) const -> std::optional<hit_record>;
   [[nodiscard]] auto triangle_intersect(const ray& ray_in, double nearer_than) const
      -> std::optional<hit_record>;
   [[nodiscard]] auto sphere_intersect(const ray& ray_in, double nearer_than) const
      -> std::optional<hit_record>;

private:
   std::vector<triangle> triangles;
   std::vector<std::unique_ptr<material>> triangle_mats;

   std::vector<sphere> spheres;
   std::vector<std::unique_ptr<material>> sphere_mats;

   colour environment;
};
