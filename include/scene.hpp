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

   void add_sphere(sphere const& sphere_in, std::shared_ptr<material> material);
   void add_sphere(sphere&& sphere_in, std::shared_ptr<material> material);

private:
   [[nodiscard]] colour radiance(ray const& r, size_t depth);

   [[nodiscard]] std::optional<hit_record> intersect(ray const& r);
   [[nodiscard]] std::optional<hit_record> sphere_intersect(ray const& ray_in, double nearer_than);

private:
   std::vector<sphere> spheres;
   std::vector<std::shared_ptr<material>> sphere_materials;
};
