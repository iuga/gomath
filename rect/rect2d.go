package rect

import (
	"github.com/iuga/gomath/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// A 2D axis-aligned bounding box using floating-point coordinates.
// Represents an axis-aligned rectangle in a 2D space. It is defined by its position and size, which are Vector2D.
type Rect2D[T Number] struct {
	position *vector.Vector2D[T]
	size     *vector.Vector2D[T]
}

// New2D constructs a new VRect2D from the given position and size.
// position is the rectangle start x and y
// size is the rectangle width and height
func New2D[T Number](position, size *vector.Vector2D[T]) *Rect2D[T] {
	return &Rect2D[T]{position: position, size: size}
}

// HasPoint returns true if the rectangle contains the given point. By convention, points on the right and bottom edges are not included.
func (r *Rect2D[T]) HasPoint(point *vector.Vector2D[T]) bool {
	if r.size.X < 0 || r.size.Y < 0 {
		return false
	}
	if point.X < r.position.X {
		return false
	}
	if point.Y < r.position.Y {
		return false
	}
	if point.X >= (r.position.X + r.size.X) {
		return false
	}
	if point.Y >= (r.position.Y + r.size.Y) {
		return false
	}
	return true
}

// GetCenter returns the center point of the rectangle. This is the same as position + (size / 2.0).
func (r *Rect2D[T]) GetCenter() *vector.Vector2D[T] {
	return r.position.Add(vector.New2D[T](r.size.X/2, r.size.Y/2))
}

// GetArea returns the rectangle's area. This is equivalent to size.x * size.y
func (r *Rect2D[T]) GetArea() T {
	return r.size.X * r.size.Y
}

// GetPosition returns a vector with the starting point of the rectangle as x and y
func (r *Rect2D[T]) GetPosition() *vector.Vector2D[T] {
	return r.position
}

// GetSize returns a vector with the current size of the Rectangle as width and height
func (r *Rect2D[T]) GetSize() *vector.Vector2D[T] {
	return r.size
}

// Abs returns a Rect2D equivalent to this rectangle, with its width and height modified to be non-negative values, and
// with its position being the top-left corner of the rectangle.
// E.g: Rect2D(25, 25, -100, -50) ~> Rect2D(-75, -25, 100, 50)
func (r *Rect2D[T]) Abs() *Rect2D[T] {
	return New2D[T](
		vector.New2D[T](
			r.position.X+r.min(r.size.X, 0),
			r.position.Y+r.min(r.size.Y, 0),
		),
		r.size.Abs(),
	)
}

func (r *Rect2D[T]) min(a, b T) T {
	if a < b {
		return a
	}
	return b
}
