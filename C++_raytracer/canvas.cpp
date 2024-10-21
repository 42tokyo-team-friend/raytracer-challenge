#include <iostream>
#include <vector>
#include <cmath>
#include <iomanip> 

#include "C++_raytracer/add_tuples.cpp"

using namespace std;

class Canvas;

class Canvas {
private:
    double width, height;
    vector<Tuple> pixels;

public:
    Canvas(double width, double height) : 
    width(width), height(height), pixels(width * height, Color(0, 0, 0)){}

    double getWidth() {
        return width;
    }

    double getHeight() {
        return height;
    }
};

int main(){
    Canvas c = Canvas(10, 20);
    cout << c.width << endl;
}