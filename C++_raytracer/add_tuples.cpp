#include "Tuple.h"

// Tuple.cpp
#include "Tuple.h"

// Constructor
Tuple::Tuple(double x, double y, double z, double w)
    : x(x), y(y), z(z), w(w), red(this->x), green(this->y), blue(this->z) {}

// Type checkers
bool Tuple::isPoint() const {
    return w == 1.0;
}

bool Tuple::isVector() const {
    return w == 0.0;
}

// Operator overloads
Tuple Tuple::operator+(const Tuple& other) const {
    return Tuple(
        this->x + other.x,
        this->y + other.y,
        this->z + other.z,
        this->w + other.w
    );
}

Tuple Tuple::operator-(const Tuple& other) const {
    return Tuple(
        this->x - other.x,
        this->y - other.y,
        this->z - other.z,
        this->w - other.w
    );
}

Tuple Tuple::operator*(double scalar) const {
    return Tuple(
        x * scalar,
        y * scalar,
        z * scalar,
        w * scalar
    );
}

Tuple Tuple::operator*(const Tuple& other) const {
    return Color(
        this->red * other.red,
        this->green * other.green,
        this->blue * other.blue
    );
}

Tuple Tuple::operator/(double scalar) const {
    return Tuple(
        x / scalar,
        y / scalar,
        z / scalar,
        w / scalar
    );
}

// Other member functions
Tuple Tuple::Neg() const {
    return Tuple(
        -x,
        -y,
        -z,
        -w
    );
}

double Tuple::magnitude() const {
    return std::sqrt(x * x + y * y + z * z);
}

Tuple Tuple::normalizing() const {
    double mag = magnitude();
    return (*this) / mag;
}

// Factory functions
Tuple Point(double x, double y, double z) {
    return Tuple(x, y, z, 1.0);
}

Tuple Vector(double x, double y, double z) {
    return Tuple(x, y, z, 0.0);
}

Tuple Color(double red, double green, double blue){
    return Tuple(red, green, blue, 0.0);
}

double Dot(const Tuple& a, const Tuple& b) {
    return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * b.w);
}

Tuple cross(const Tuple& a, const Tuple& b) {
    return Vector(
        a.y * b.z - a.z * b.y,
        a.z * b.x - a.x * b.z,
        a.x * b.y - a.y * b.x
    );
}

int main() {
    Tuple p1 = Point(1.0, 2.0, 3.0);
    Tuple v1 = Vector(4.0, 5.0, 6.0);
    Tuple color1 = Color(0.5, 0.4, 0.3);

    Tuple result = p1 + v1;
    std::cout << "Result: (" << result.x << ", " << result.y << ", " << result.z << ", " << result.w << ")\n";

    return 0;
}