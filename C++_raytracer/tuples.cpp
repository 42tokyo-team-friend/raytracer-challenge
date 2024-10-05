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
};

int main(){
    Tuple point(1.0, 2.0, 3.0, 1.0);
    Tuple vector(1.0, 2.0, 3.0, 0.0);

    cout << point.isPoint() << endl;
    cout << point.isVector() << endl;
    cout << vector.isPoint() << endl;
    cout << vector.isVector() << endl;
}
