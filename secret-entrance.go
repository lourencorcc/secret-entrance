package main

import (
	"fmt"
	"os"
)

type Rotation struct {
	direction string
	dist      int
}

type CircularSlice[T any] struct {
	data []T
	i    int
}

func NewCircularSlice[T any](data []T) *CircularSlice[T] {
	return &CircularSlice[T]{data: data, i: 50}
}

func (c *CircularSlice[T]) Next(howMany int) T {
	c.i = (c.i + howMany) % len(c.data)
	v := c.data[c.i]
	return v
}

func (c *CircularSlice[T]) Previous(howMany int) T {
	c.i = (c.i - howMany) % len(c.data)
	if c.i < 0 {
		c.i += len(c.data)
		// c.i *= -1
	}
	v := c.data[c.i]
	return v
}

func loadRotationsFile() []Rotation {
	// read into byte array
	// 1st byte: direction
	// read bytes until byte is \n: dist
	// EOF finish
	var fileContent []byte

	fileContent, err := os.ReadFile("./instructions.txt")

	if err != nil {
		fmt.Println("Deu merda a ler o ficheiro")
	}

	var rotations []Rotation

	for i := 0; i < len(fileContent); {
		var rotation Rotation
		rotation.direction = string(fileContent[i]) // converts to character in string
		i++
		for fileContent[i] != '\n' {

			rotation.dist = rotation.dist*10 + int(fileContent[i]-'0') // or  strconv.ParseUint(fileContent[i], 10, 64) but this returns err so its annoying to work with
			i++
		}
		rotations = append(rotations, rotation)
		// fmt.Println(rotations)
		i++ // consume the \n
	}
	fmt.Println(len(rotations))
	// fmt.Println(rotations)
	return rotations
}

func getPassword(dial *CircularSlice[int], rotations []Rotation) int {
	var password int
	var dialVal int
	for _, v := range rotations {
		// fmt.Println(rotations[i])
		switch v.direction {
		case "L":
			dialVal = dial.Previous(v.dist)
		case "R":
			dialVal = dial.Next(v.dist)
		}

		if dialVal == 0 {
			password += 1
		}

	}

	return password
}

func main() {
	fmt.Println("hello")
	rotations := loadRotationsFile()
	dial := NewCircularSlice([]int{
		0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
		10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		50, 51, 52, 53, 54, 55, 56, 57, 58, 59,
		60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
		70, 71, 72, 73, 74, 75, 76, 77, 78, 79,
		80, 81, 82, 83, 84, 85, 86, 87, 88, 89,
		90, 91, 92, 93, 94, 95, 96, 97, 98, 99,
	})

	result := getPassword(dial, rotations)

	fmt.Println(result)

}
