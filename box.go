package golang_united_school_homework

import (
	"fmt"
)

var errOutOfRange = fmt.Errorf("index went out of the range")

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapes:         make([]Shape, 0, shapesCapacity),
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	if len(b.shapes) >= b.shapesCapacity {
		return fmt.Errorf("out of capacity")
	}

	b.shapes = append(b.shapes, shape)

	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errOutOfRange
	}

	return b.shapes[i], nil
}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errOutOfRange
	}

	result := b.shapes[i]
	b.shapes = append(b.shapes[:i], b.shapes[i+1:]...)

	return result, nil
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i >= len(b.shapes) {
		return nil, errOutOfRange
	}

	result := b.shapes[i]
	b.shapes[i] = shape

	return result, nil
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	result := 0.0

	for i := 0; i < len(b.shapes); i++ {
		result += b.shapes[i].CalcPerimeter()
	}

	return result
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	result := 0.0

	for i := 0; i < len(b.shapes); i++ {
		result += b.shapes[i].CalcArea()
	}

	return result
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	newShapes := make([]Shape, 0, b.shapesCapacity)

	for i := 0; i < len(b.shapes); i++ {
		if _, ok := b.shapes[i].(*Circle); !ok {
			newShapes = append(newShapes, b.shapes[i])
		}
	}

	if len(newShapes) == len(b.shapes) {
		return fmt.Errorf("circles are not exist in the list")
	}

	b.shapes = newShapes
	return nil
}
