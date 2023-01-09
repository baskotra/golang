package main

import (
	"errors"
	"fmt"
)

type areaError struct {
	err    string
	length float64
	width  float64
}

func (a *areaError) Error() string {
	return a.err
}

func (a *areaError) lengthNegative() bool {
	return a.length < 0
}

func (a *areaError) widthNegative() bool {
	return a.width < 0
}

func calculateArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "Length is less than zero."
	}
	if width < 0 {
		if err == "" {
			err = "Width is less than zero."
		} else {
			err += "Width is less than zero."
		}
	}
	if err != "" {
		return 0, &areaError{
			err:    err,
			length: length,
			width:  width,
		}
	}
	return length * width, nil
}

func main() {
	answer, err := calculateArea(45, -62)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			if areaError.lengthNegative() {
				fmt.Println("Length cannot be negative.")
				return
			}
			if areaError.widthNegative() {
				fmt.Println("Width cannot be negative")
				return
			}
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Area of rectangle is:", answer)

}
