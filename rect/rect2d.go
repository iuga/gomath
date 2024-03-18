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
