package main

import (
	"fmt"
	"log"
	"vaawebdev/sandbox/fib"
	"vaawebdev/sandbox/parallelize"
)

func main() {
	i := 999999

	parallelize.Parallelize(
		func() {
			res, err := fib.FibRec(i)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(fmt.Sprintf("Recursion %d!", res))
		},
		func() {
			res, err := fib.FibRec(i)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(fmt.Sprintf("Loop %d!", res))
		},
	)
}
