package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var experimentsAmount int
	_, err := fmt.Fscan(reader, &experimentsAmount)
	if err != nil {
		panic(err)
	}

	for i := 0; i < experimentsAmount; i++ {
		var cars int
		var capacity int
		var cargo int
		_, err = fmt.Fscan(reader, &cars, &capacity, &cargo)
		if err != nil {
			panic(err)
		}

		slc := make([]int, cargo)
		for j := 0; j < cargo; j++ {
			_, err = fmt.Fscan(reader, slc[j])
			if err != nil {
				panic(err)
			}

			slc[j] = 1 << slc[j]
		}

	}
}
