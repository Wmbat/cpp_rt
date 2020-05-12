#pragma once

#include <glm/glm.hpp>

struct hit
{
   glm::dvec3 position{0.0, 0.0, 0.0};
   glm::dvec3 normal{0.0, 0.0, 0.0};
   double distance{0.0};
};
