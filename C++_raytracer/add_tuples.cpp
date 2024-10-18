#include <iostream>
#include <cmath>

using namespace std;

class Tuple {
public:
    double x, y, z, w;

    Tuple(double x, double y, double z, double w) : x(x), y(y), z(z), w(w) {}

    bool isPoint() const {
        return w == 1.0;
    }

    bool isVector() const {
        return w == 0.0;
    }

    Tuple operator+(const Tuple& other) const {
        return Tuple (
            this-> x + other.x,
            this-> y + other.y,
            this-> z + other.z,
            this-> w + other.w
        );
    }

    Tuple operator-(const Tuple& other) const {
        return Tuple (
            this-> x - other.x,
            this-> y - other.y,
            z - other.z,
            w - other.w
        );
    }

    Tuple operator*(double other) const {
        return Tuple (
            x * other,
            y * other,
            z * other,
            w * other
        );
    }

    Tuple operator/(double other) const {
        return Tuple (
            x / other,
            y / other,
            z / other,
            w / other
        );
    }

    Tuple Neg() const {
        return Tuple (
            -1 * x,
            -1 * y,
            -1 * z,
            -1 * w
        );
    }

    double magnitude() const {
        return sqrt(x*x + y*y + z*z);
    }

    Tuple normalizing() {
        double mag = magnitude();
        return (*this) / mag;
    }
};

Tuple Point(double x, double y, double z) {
    return Tuple(x, y, z, 1.0);
}

Tuple Vector(double x, double y, double z) {
    return Tuple(x, y, z, 0.0);
}

double Dot(Tuple a, Tuple b) {
    return (a.x * b.x + a.y * b.y + a.z * b.z + a.w * b.w);
}

Tuple cross(Tuple a, Tuple b) {
    return Vector (
        a.y * b.z - a.z * b.y,
        a.z * b.x - a.x * b.z,
        a.x * b.y - a.y * b.x
    );
}

int main(){
    Tuple a = Vector(1, 2, 3);
    Tuple b = Vector(2, 3, 4);
    Tuple c = cross(a, b);
    cout << c.x << c.y << c.z << endl;
}