package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	tmp := make([]bytes.Buffer, numRows)

	var currentRow int

	var goDown bool
	for _, value := range s {
		_, err := tmp[currentRow].WriteRune(value)
		if err != nil {
			log.Println(err)
		}

		if currentRow == 0 || currentRow == numRows-1 {
			goDown = !goDown
		}

		if goDown {
			currentRow++
		} else {
			currentRow--
		}
	}

	for i := 1; i < numRows; i++ {
		_, err := io.Copy(&tmp[0], &tmp[i])
		if err != nil {
			log.Println(err)
		}
	}

	return tmp[0].String()
}
