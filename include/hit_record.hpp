#pragma once

#include "hit.hpp"
#include "materials/material.hpp"

#include <memory>

struct hit_record
{
   hit hit;
   std::shared_ptr<material> material;
};
