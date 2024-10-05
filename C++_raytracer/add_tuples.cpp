#include <iostream>

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

    Tuple Neg() const {
        return Tuple (
            -1 * x,
            -1 * y,
            -1 * z,
            -1 * w
        );
    }
};

Tuple Point(double x, double y, double z) {
    return Tuple(x, y, z, 1.0);
}

Tuple Vector(double x, double y, double z) {
    return Tuple(x, y, z, 0.0);
}

int main(){
    Tuple a(1, -2, 3, -4);

    Tuple Neg1 = a.Neg();

    cout << Neg1.x << Neg1.y << Neg1.z << Neg1.w;
}