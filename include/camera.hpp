#pragma once

#include "ray.hpp"

class camera
{
public:
   camera() : lower_left_corner(origin - horizontal / 2.0 - vertical / 2.0 - vec(0.0, 0.0, 1.0)) {}

   ray shoot_ray(double u, double v) const { return ray(origin, lower_left_corner + u * horizontal + v * vertical); }

private:
   vec origin{0.0, 0.0, 0.0};
   vec horizontal{4.0, 0.0, 0.0};
   vec vertical{0.0, 2.25, 0.0};
   vec lower_left_corner;
};
