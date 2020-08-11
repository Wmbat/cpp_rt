#pragma once

#include <cpp_pt/math/vec.hpp>

struct hit
{
   vec position{0.0, 0.0, 0.0};
   vec normal{0.0, 0.0, 0.0};
   double distance{0.0};
   bool front_face{false};
};
