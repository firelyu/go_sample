package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (ens ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("can't sqrt negative number : %f\n", float64(ens))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		var err ErrNegativeSqrt = ErrNegativeSqrt(x)
		return 0, err
	}

	zn := x
	zn1 := zn - (zn * zn - x) / (zn * 2)
	for zn - zn1 > float64(1 >> 32) {
		zn = zn1
		zn1 = zn - (zn * zn - x) / (zn * 2)
	}

	return zn1, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}