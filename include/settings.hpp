#pragma once

#include <cstddef>

struct render_settings
{
   int window_width{1920};
   int window_height{static_cast<int>(window_width / (16.0 / 9.0))};
   size_t u_samples{1};
   size_t v_samples{1};
   size_t sample_count{256};
   size_t bounce_depth{8};
};
