package main

import (
	"fmt"
	"math"
)

const prec = 5000

/*
func pow(x *big.Float, pow int) *big.Float {
	if pow == 0 {
		return new(big.Float).SetPrec(prec).SetFloat64(1)
	}

	a := x

	for i := 2; i < pow; i++ {
		a.Mul(a, x)
	}

	return a
}

func atan(x *big.Float, steps int) *big.Float {
	var res, v *big.Float

	res = new(big.Float).SetPrec(prec)
	v = new(big.Float).SetPrec(prec)

	for i := 0; i < steps; i++ {
		res.Add(res, v.Mul(pow(new(big.Float).SetFloat64(-1), i), pow(x, i*2+1)).Quo(v, new(big.Float).SetFloat64(float64(i*2+1))))
	}

	return res
}

func pi(steps int64) *big.Float {
	var res, two, den, n *big.Float
	var t1, t2 *big.Float

	res = new(big.Float).SetPrec(prec)
	two = new(big.Float).SetPrec(prec).SetFloat64(2.0)
	den = new(big.Float).SetPrec(prec)
	n = new(big.Float).SetPrec(prec)
	t1 = new(big.Float).SetPrec(prec)
	t2 = new(big.Float).SetPrec(prec)

	for i := int64(0); i < steps; i++ {

		n.SetInt64(i)

		res.Add(res)
	}
}

func acos(x *big.Float, steps int) *big.Float {

	return 0.5*pi(steps) - atan(x/math.Sqrt(1-x*x), steps)

}*/

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

	return (a * b) / sigma
}

const ShapeOfOne float64 = 0.006817

func findPrimordial(a, b float64) float64 {

	//fmt.Println("DeltaShape: ", findDeltaShape(a,b) )
	//fmt.Println("CompositeShape: ", findCompositeShape(a,b))
	fmt.Println("DeltaPrime: ", findDeltaPrime(a, b))
	//fmt.Println("Primordial: ", ShapeOfOne / zulu)

	return ShapeOfOne / math.Abs(findDeltaPrime(a, b))
}

func main() {
	fmt.Println("Prime Number Theorem and Primordials")
	fmt.Println("Author : Chris Pergrossi [Aeon]")
	fmt.Println("  Date : October 27, 2016")

	/*
			fA := math.Acos(math.Cos(math.Pi / 2.0))
			fB := acos(math.Cos(math.Pi/2.0), 5)
			fC := acos(math.Cos(math.Pi/2.0), 50)
			fD := acos(math.Cos(math.Pi/2.0), 5000)

		fmt.Println("A: ", fA)
		fmt.Println("B: ", fB)
		fmt.Println("C: ", fC)
		fmt.Println("D: ", fD)*/

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
