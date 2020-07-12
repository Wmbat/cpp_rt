#pragma once

#include <array>

class vec final
{
public:
   constexpr vec() noexcept = default;
   constexpr vec(double x, double y, double z) noexcept
   {
      data[0] = x;
      data[1] = y;
      data[2] = z;
   }

   constexpr auto operator+(vec const& rhs) const noexcept -> vec
   {
      return vec(x() + rhs.x(), y() + rhs.y(), z() + rhs.z());
   }
   constexpr auto operator-(vec const& rhs) const noexcept -> vec
   {
      return vec(x() - rhs.x(), y() - rhs.y(), z() - rhs.z());
   }
   constexpr auto operator*(vec const& rhs) const noexcept -> vec
   {
      return vec(x() * rhs.x(), y() * rhs.y(), z() * rhs.z());
   }
   constexpr auto operator*(double scalar) const noexcept -> vec
   {
      return vec(x() * scalar, y() * scalar, z() * scalar);
   }
   constexpr auto operator/(double scalar) const noexcept -> vec
   {
      const auto reciprocal = 1 / scalar;
      return vec(x() * reciprocal, y() * reciprocal, z() * reciprocal);
   }
   constexpr auto operator-() const noexcept -> vec { return vec(-x(), -y(), -z()); }

   constexpr auto operator+=(vec const& rhs) noexcept -> vec&
   {
      data[0] += rhs.x();
      data[1] += rhs.y();
      data[2] += rhs.z();

      return *this;
   }
   constexpr auto operator-=(vec const& rhs) noexcept -> vec&
   {
      data[0] -= rhs.x();
      data[1] -= rhs.y();
      data[2] -= rhs.z();

      return *this;
   }
   constexpr auto operator*=(vec const& rhs) noexcept -> vec&
   {
      data[0] *= rhs.x();
      data[1] *= rhs.y();
      data[2] *= rhs.z();

      return *this;
   }
   constexpr auto operator*=(double scalar) noexcept -> vec&
   {
      data[0] *= scalar;
      data[1] *= scalar;
      data[2] *= scalar;

      return *this;
   }
   constexpr auto operator/=(double scalar) noexcept -> vec&
   {
      const auto reciprocal = 1 / scalar;
      data[0] *= reciprocal;
      data[1] *= reciprocal;
      data[2] *= reciprocal;

      return *this;
   }

   constexpr auto operator==(vec const& rhs) const noexcept -> bool = default;

   constexpr friend auto operator*(double lhs, vec const& rhs) -> vec
   {
      return vec(lhs * rhs.x(), lhs * rhs.y(), lhs * rhs.z());
   }
   constexpr friend auto operator/(double lhs, vec const& rhs) -> vec
   {
      return vec(lhs / rhs.x(), lhs / rhs.y(), lhs / rhs.z());
   }

   [[nodiscard]] constexpr auto length_squared() const noexcept -> double
   {
      return x() * x() + y() * y() + z() * z();
   }
   [[nodiscard]] auto length() const noexcept -> double;

   [[nodiscard]] constexpr auto x() const -> double { return data[0]; }
   [[nodiscard]] constexpr auto y() const -> double { return data[1]; }
   [[nodiscard]] constexpr auto z() const -> double { return data[2]; }

private:
   std::array<double, 3> data = {0.0, 0.0, 0.0};
};

constexpr auto dot(vec const& lhs, vec const& rhs) noexcept -> double
{
   return lhs.x() * rhs.x() + lhs.y() * rhs.y() + lhs.z() * rhs.z();
}
constexpr auto cross(vec const& lhs, vec const& rhs) noexcept -> vec
{
   auto x = lhs.y() * rhs.z() - lhs.z() * rhs.y();
   auto y = lhs.z() * rhs.x() - lhs.x() * rhs.z();
   auto z = lhs.x() * rhs.y() - lhs.y() * rhs.x();

   return vec(x, y, z);
}
inline auto normalise(vec const& value) noexcept -> vec
{
   return value / value.length();
}

using colour = vec;
using norm = vec;
using position = vec;
