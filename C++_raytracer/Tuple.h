#ifndef TUPLE_H
#define TUPLE_H

#include <iostream>
#include <cmath>
#include <iomanip> 

class Tuple {
public:
    double x, y, z, w;

    double& red;
    double& green;
    double& blue;

    Tuple(double x, double y, double z, double w);

    bool isPoint() const;
    bool isVector() const;

    Tuple operator+(const Tuple& other) const;
    Tuple operator-(const Tuple& other) const;
    Tuple operator*(double scalar) const;
    Tuple operator*(const Tuple& other) const;
    Tuple operator/(double scalar) const;

    Tuple Neg() const;
    double magnitude() const;
    Tuple normalizing() const;
};

Tuple Point(double x, double y, double z);
Tuple Vector(double x, double y, double z);
Tuple Color(double red, double green, double blue);

double Dot(const Tuple& a, const Tuple& b);
Tuple cross(const Tuple& a, const Tuple& b);

#endif
