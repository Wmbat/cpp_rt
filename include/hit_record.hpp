#pragma once

#include "hit.hpp"
#include "materials/material.hpp"

#include <memory>

struct hit_record
{
   hit hit;
   material_info mat;
};
