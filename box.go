package golang_united_school_homework

// package main

import (
	"fmt"
)

// box contains list of shapes and able to perform operations on them
type box struct {
	shapes         []Shape
	shapesCapacity int // Maximum quantity of shapes that can be inside the box.
}

// NewBox creates new instance of box
func NewBox(shapesCapacity int) *box {
	return &box{
		shapesCapacity: shapesCapacity,
	}
}

// AddShape adds shape to the box
// returns the error in case it goes out of the shapesCapacity range.
func (b *box) AddShape(shape Shape) error {
	// panic("implement me")
	if len(b.shapes) < b.shapesCapacity {
		b.shapes = append(b.shapes, shape)
	} else {
		return fmt.Errorf("capacity is already full")
	}
	// fmt.Printf("AddShape: shapes len %d \n", len(b.shapes))
	return nil
}

// GetByIndex allows getting shape by index
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) GetByIndex(i int) (Shape, error) {
	// panic("implement me")
	if i < 1 || i > len(b.shapes) {
		return nil, fmt.Errorf("index out of range ")
	} else {
		return b.shapes[i], nil
	}

}

// ExtractByIndex allows getting shape by index and removes this shape from the list.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ExtractByIndex(i int) (Shape, error) {
	// fmt.Printf("index to remove %d, total shapes len %d \n", i, len(b.shapes))
	if i < 1 || i > len(b.shapes) {
		return nil, fmt.Errorf("index out of range ")
	} else {
		temp := b.shapes[i]
		// fmt.Printf("ExtractByIndex %d \n", len(b.shapes))
		copy(b.shapes[i:], b.shapes[i+1:])
		b.shapes = b.shapes[:len(b.shapes)-1]
		// fmt.Printf("ExtractByIndex %d \n", len(b.shapes))
		return temp, nil
	}
}

// ReplaceByIndex allows replacing shape by index and returns removed shape.
// whether shape by index doesn't exist or index went out of the range, then it returns an error
func (b *box) ReplaceByIndex(i int, shape Shape) (Shape, error) {
	if i < 1 || i > len(b.shapes) {
		return nil, fmt.Errorf("index out of range ")
	} else {
		temp := b.shapes[i]
		b.shapes[i] = shape
		return temp, nil
	}
}

// SumPerimeter provides sum perimeter of all shapes in the list.
func (b *box) SumPerimeter() float64 {
	// panic("implement me")
	sum := 0.0
	for _, k := range b.shapes {
		sum += k.CalcPerimeter()
	}
	return sum
}

// SumArea provides sum area of all shapes in the list.
func (b *box) SumArea() float64 {
	// panic("implement me")
	sum := 0.0
	for _, k := range b.shapes {
		sum += k.CalcArea()
	}
	return sum
}

// RemoveAllCircles removes all circles in the list
// whether circles are not exist in the list, then returns an error
func (b *box) RemoveAllCircles() error {
	// panic("implement me")
	// return fmt.Errorf("")
	if len(b.shapes) < 1 {
		return fmt.Errorf("empty Box")
	} else {
		var tempShapes []Shape
		for _, k := range b.shapes {
			if fmt.Sprintf("%T", k) != "golang_united_school_homework.Circle" {
				tempShapes = append(tempShapes, k)
			}
		}
		if len(b.shapes) == len(tempShapes) {
			return fmt.Errorf("No circles found")
		}
		b.shapes = tempShapes
		return nil
		// fmt.Println(len(circleIndexes))
		// for ix, x := range circleIndexes {
		// 	fmt.Printf("RemoveAllCircles: %d %d \n", ix, x)
		// 	fmt.Printf("RemoveAllCircles: %d \n", len(b.shapes))
		// 	b.ExtractByIndex(x)
		// 	fmt.Printf("RemoveAllCircles: %d \n", len(b.shapes))
		// }
	}
	return nil
}

// func main() {
// 	b1 := NewBox(4)
// 	t1 := Triangle{Side: 10}
// 	r1 := Rectangle{Height: 10, Weight: 20}
// 	c1 := Circle{Radius: 2.0}
// 	b1.AddShape(r1)
// 	b1.AddShape(t1)
// 	b1.AddShape(c1)
// 	// c2 := Circle{Radius: 2.0}
// 	// b1.AddShape(c2)
// 	fmt.Println(b1.SumArea())
// 	b1.RemoveAllCircles()
// 	fmt.Println(b1.SumArea())
// }
