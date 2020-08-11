#pragma once

#include <cpp_pt/hit.hpp>
#include <cpp_pt/materials/material.hpp>

#include <memory>

struct hit_record
{
   hit hit_data{};
   material* p_mat{nullptr};
};
