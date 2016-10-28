package main

import (
	"fmt"
	"math"
)

func findCompositeShape(a, b float64) float64 {
	leftSide := a*a*b*b - a*a - b*b

	leftSide = leftSide / (-2 * a * a * b * b)

	return math.Acos(leftSide) / 2
}

func findShape(a float64) float64 {
	return math.Acos(.5 / (a * a))
}

func findSigma(a, b float64) float64 {
	shapeA := findShape(a)
	shapeB := findShape(b)

	return 2 * math.Acos(math.Min(shapeA, shapeB)/math.Max(shapeA, shapeB))
}

func findLargerPrime(a, b float64) float64 {
	sigma := findSigma(a, b)

	value := 2 * sigma / math.Abs(findShape(b)-findShape(a))

	if (int64(math.Floor(value)) % 2) == 0 {
		return math.Ceil(value)
	} else {
		return math.Floor(value)
	}
}

func findDeltaPrime(a, b float64) float64 {
	sigma := findSigma(a, b)

	fmt.Println("Sigma: ", sigma)

	delta := sigma / math.Abs(findShape(b)-findShape(a))

	return (a * b) / delta
}

func findDeltaShape(a, b float64) float64 {
	theta := math.Abs(findShape(b) - findShape(a))

	return theta + findSigma(a, b)
}

const ShapeOfOne float64 = 0.006817

func findPrimordial(a, b float64) float64 {
	zulu := (findShape(a) - findShape(b)) / findDeltaPrime(a, b)

	//fmt.Println("DeltaShape: ", findDeltaShape(a,b) )
	//fmt.Println("CompositeShape: ", findCompositeShape(a,b))
	fmt.Println("DeltaPrime: ", findDeltaPrime(a, b))
	//fmt.Println("Primordial: ", ShapeOfOne / zulu)

	return ShapeOfOne / math.Abs(zulu)
}

func main() {
	fmt.Println("Prime Number Theorem and Primordials")
	fmt.Println("Author : Chris Pergrossi [Aeon]")
	fmt.Println("  Date : October 27, 2016")

	// grab two prime numbers
	fmt.Println("Enter Two Prime Numbers")
	var a, b float64

	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		fmt.Println("Invalid Value")
		return
	}
	_, err = fmt.Scanf("%f", &b)
	if err != nil {
		fmt.Println("Invalid Value")
		return
	}

	// find the next highest prime number
	largerPrime := findLargerPrime(a, b)

	// find the next primordial
	primordial := findPrimordial(a, b)

	// print them to screen
	fmt.Println("Next Largest Prime: ", largerPrime)
	fmt.Println("    Primordial    : ", primordial)
}
