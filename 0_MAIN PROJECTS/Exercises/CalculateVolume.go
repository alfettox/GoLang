/*
Author: Giovanni De Franceschi
*/






package main

import (
	"fmt"
	"math"
)

// Define structs: Cube, Box, Sphere
type Cube struct {
	length float64
}

type Box struct {
	length, width, height float64
}

type Sphere struct {
	radius float64
}

// Define additional struct: Cuboid, Cone, Pyramid
type Cuboid struct {
	length, width, height float64
}

type Cone struct {
	radius, height float64
}

type Pyramid struct {
	baseLength, baseWidth, height float64
}

// Volume methods for shapes
func (c Cube) Volume() float64 {
	return math.Pow(c.length, 3)
}

func (b Box) Volume() float64 {
	return b.length * b.width * b.height
}

func (s Sphere) Volume() float64 {
	return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3)
}

func (c Cuboid) Volume() float64 {
	return c.length * c.width * c.height
}

func (co Cone) Volume() float64 {
	return (1.0 / 3.0) * math.Pi * math.Pow(co.radius, 2) * co.height
}

func (p Pyramid) Volume() float64 {
	return (1.0 / 3.0) * p.baseLength * p.baseWidth * p.height
}

func main() {
	// Creating instances of different shapes
	cube := Cube{length: 3.0}
	box := Box{length: 4.0, width: 5.0, height: 6.0}
	sphere := Sphere{radius: 2.0}
	cuboid := Cuboid{length: 4.0, width: 5.0, height: 6.0}
	cone := Cone{radius: 3.0, height: 8.0}
	pyramid := Pyramid{baseLength: 5.0, baseWidth: 6.0, height: 7.0}

	// Calculating volumes
	cubeVolume := cube.Volume()
	boxVolume := box.Volume()
	sphereVolume := sphere.Volume()
	cuboidVolume := cuboid.Volume()
	coneVolume := cone.Volume()
	pyramidVolume := pyramid.Volume()

	// Displaying volumes
	fmt.Printf("Cube Volume: %.2f\n", cubeVolume)
	fmt.Printf("Box Volume: %.2f\n", boxVolume)
	fmt.Printf("Sphere Volume: %.2f\n", sphereVolume)
	fmt.Printf("Cuboid Volume: %.2f\n", cuboidVolume)
	fmt.Printf("Cone Volume: %.2f\n", coneVolume)
	fmt.Printf("Pyramid Volume: %.2f\n", pyramidVolume)
}
	