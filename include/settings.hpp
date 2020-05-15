#pragma once

#include <cstddef>

struct render_settings
{
   int window_width{1920};
   int window_height{static_cast<int>(window_width / (16.0 / 9.0))};
   size_t sample_count{64};
   size_t bounce_depth{4};
};
