#pragma once

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

   constexpr vec operator+(vec const& rhs) const noexcept { return vec(x() + rhs.x(), y() + rhs.y(), z() + rhs.z()); }
   constexpr vec operator-(vec const& rhs) const noexcept { return vec(x() - rhs.x(), y() - rhs.y(), z() - rhs.z()); }
   constexpr vec operator*(vec const& rhs) const noexcept { return vec(x() * rhs.x(), y() * rhs.y(), z() * rhs.z()); }
   constexpr vec operator*(double scalar) const noexcept { return vec(x() * scalar, y() * scalar, z() * scalar); }
   constexpr vec operator/(double scalar) const noexcept
   {
      const auto reciprocal = 1 / scalar;
      return vec(x() * reciprocal, y() * reciprocal, z() * reciprocal);
   }
   constexpr vec operator-() const noexcept { return vec(-x(), -y(), -z()); }

   constexpr vec& operator+=(vec const& rhs) noexcept
   {
      data[0] += rhs.x();
      data[1] += rhs.y();
      data[2] += rhs.z();

      return *this;
   }
   constexpr vec& operator-=(vec const& rhs) noexcept
   {
      data[0] -= rhs.x();
      data[1] -= rhs.y();
      data[2] -= rhs.z();

      return *this;
   }
   constexpr vec& operator*=(vec const& rhs) noexcept
   {
      data[0] *= rhs.x();
      data[1] *= rhs.y();
      data[2] *= rhs.z();

      return *this;
   }
   constexpr vec& operator*=(double scalar) noexcept
   {
      data[0] *= scalar;
      data[1] *= scalar;
      data[2] *= scalar;

      return *this;
   }
   constexpr vec& operator/=(double scalar) noexcept
   {
      const auto reciprocal = 1 / scalar;
      data[0] *= reciprocal;
      data[1] *= reciprocal;
      data[2] *= reciprocal;

      return *this;
   }

   constexpr bool operator==(vec const& rhs) const noexcept = default;

   friend vec operator*(double lhs, vec const& rhs) { return vec(lhs * rhs.x(), lhs * rhs.y(), lhs * rhs.z()); }
   friend vec operator/(double lhs, vec const& rhs) { return vec(lhs / rhs.x(), lhs / rhs.y(), lhs / rhs.z()); }

   constexpr double length_squared() const noexcept { return x() * x() + y() * y() + z() * z(); }
   double length() const noexcept;

   constexpr double x() const { return data[0]; }
   constexpr double y() const { return data[1]; }
   constexpr double z() const { return data[2]; }

private:
   double data[3];
};

constexpr double dot(vec const& lhs, vec const& rhs) noexcept
{
   return lhs.x() * rhs.x() + lhs.y() * rhs.y() + lhs.z() * rhs.z();
}
constexpr vec cross(vec const& lhs, vec const& rhs) noexcept
{
   auto x = lhs.y() * rhs.z() - lhs.z() * rhs.y();
   auto y = lhs.z() * rhs.x() - lhs.x() * rhs.z();
   auto z = rhs.x() * rhs.y() - lhs.y() * rhs.x();

   return vec(x, y, z);
}
constexpr vec normalise(vec const& vec) noexcept
{
   return vec / vec.length();
}

using colour = vec;
using norm = vec;
using position = vec;
