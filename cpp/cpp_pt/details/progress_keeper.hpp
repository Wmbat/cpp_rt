#pragma once

#include <cstddef>

class progress_keeper
{
public:
   explicit progress_keeper(size_t sample_count_in) noexcept;

   void update(size_t samples_done) noexcept;

private:
   size_t sample_count{};
   const double minimum_progress{5.0};
   double last_progress{};
};
