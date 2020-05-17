#pragma once

#include "details/camera.hpp"
#include "details/image.hpp"
#include "details/settings.hpp"
#include "hit_record.hpp"
#include "materials/material.hpp"
#include "ray.hpp"
#include "renderables/sphere.hpp"

#include <memory>
#include <optional>
#include <random>
#include <vector>

class scene
{
public:
   [[nodiscard]] image render(camera const& cam, render_settings const& settings);

   void add_sphere(sphere const& sphere_in, std::unique_ptr<material> p_mat);
   void add_sphere(sphere&& sphere_in, std::unique_ptr<material> p_mat);

   void set_environment_colour(const colour& environment_in) noexcept;

private:
   [[nodiscard]] colour radiance(
      ray const& r, size_t u_samples, size_t v_samples, size_t depth) const;

   [[nodiscard]] std::optional<hit_record> intersect(ray const& r) const;
   [[nodiscard]] std::optional<hit_record> sphere_intersect(
      ray const& ray_in, double nearer_than) const;

private:
   std::vector<sphere> spheres;
   std::vector<std::unique_ptr<material>> sphere_mats;

   colour environment;
};
