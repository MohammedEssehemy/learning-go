// package main

// import (
// 	"fmt"
// 	"time"
// )

// type MyError struct {
// 	when time.Time
// 	what string
// }

// func (err *MyError) Error() string {
// 	return fmt.Sprintf("at %v, %s", err.when, err.what)
// }

// func run() error {
// 	return &MyError{
// 		time.Now(),
// 		"it didn't work",
// 	}
// }

// func main() {
// 	if err := run(); err != nil {
// 		fmt.Println(err)
// 	}
// }

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", err)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
