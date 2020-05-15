#pragma once

#include "camera.hpp"
#include "hit_record.hpp"
#include "image.hpp"
#include "materials/material.hpp"
#include "ray.hpp"
#include "settings.hpp"
#include "sphere.hpp"

#include <memory>
#include <optional>
#include <random>
#include <vector>

class scene
{
public:
   [[nodiscard]] image render(camera const& cam, render_settings const& settings);

   void add_sphere(sphere const& sphere_in, material_info const& mat);
   void add_sphere(sphere&& sphere_in, material_info&& mat);

private:
   [[nodiscard]] colour radiance(ray const& r, size_t u_samples, size_t v_samples, size_t depth) const;

   [[nodiscard]] std::optional<hit_record> intersect(ray const& r) const;
   [[nodiscard]] std::optional<hit_record> sphere_intersect(ray const& ray_in, double nearer_than) const;

private:
   std::vector<sphere> spheres;
   std::vector<material_info> sphere_mats;
};
