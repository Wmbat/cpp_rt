#include <cpp_pt/details/progress_keeper.hpp>

#include <chrono>
#include <iomanip>
#include <iostream>
#include <ostream>
#include <sstream>

progress_keeper::progress_keeper(size_t sample_count_in) noexcept : sample_count(sample_count_in) {}

void progress_keeper::update(size_t samples_done) noexcept
{
   const double progress = static_cast<double>(samples_done) / sample_count * 100;
   if (progress >= last_progress + minimum_progress)
   {
      std::cout << std::fixed << std::setprecision(2) << progress << "% (" << samples_done << " / "
                << sample_count << ")\n"
                << std::flush;

      last_progress = progress;
   }
}
