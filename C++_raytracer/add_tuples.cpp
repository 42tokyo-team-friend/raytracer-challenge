#include <iostream>

using namespace std;

class Tuple {
public:
    int x, y, z, w;

    Tuple(int x, int y, int z, int w) : x(x), y(y), z(z), w(w) {}

    Tuple operator+(const Tuple& other) const {
        return Tuple (
            this-> x + other.x,
            this-> y + other.y,
            this-> z + other.z,
            this-> w + other.w,
        );
    }

};

int main(){
    Tuple a1(3, -2, 5, 1);
    Tuple a2(-2, 3, 1, 0);
    Tuple a3 = a1 + a2;

}