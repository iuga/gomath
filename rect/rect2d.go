package rect

import (
	"math/rand"

	"github.com/iuga/gomath/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// A 2D axis-aligned bounding box using floating-point coordinates.
// Represents an axis-aligned rectangle in a 2D space. It is defined by its position and size, which are Vector2D.
type Rect2D[T Number] struct {
	Position *vector.Vector2D[T] `json:"position"`
	Size     *vector.Vector2D[T] `json:"size"`
}

// New2D constructs a new VRect2D from the given position and size.
// position is the rectangle start x and y
// size is the rectangle width and height
func New2D[T Number](position, size *vector.Vector2D[T]) *Rect2D[T] {
	return &Rect2D[T]{Position: position, Size: size}
}

// HasPoint returns true if the rectangle contains the given point. By convention, points on the right and bottom edges are not included.
func (r *Rect2D[T]) HasPoint(point *vector.Vector2D[T]) bool {
	if r.Size.X < 0 || r.Size.Y < 0 {
		return false
	}
	if point.X < r.Position.X {
		return false
	}
	if point.Y < r.Position.Y {
		return false
	}
	if point.X >= (r.Position.X + r.Size.X) {
		return false
	}
	if point.Y >= (r.Position.Y + r.Size.Y) {
		return false
	}
	return true
}

// GetCenter returns the center point of the rectangle. This is the same as position + (size / 2.0).
func (r *Rect2D[T]) GetCenter() *vector.Vector2D[T] {
	return r.Position.Add(vector.New2D[T](r.Size.X/2, r.Size.Y/2))
}

// GetArea returns the rectangle's area. This is equivalent to size.x * size.y
func (r *Rect2D[T]) GetArea() T {
	return r.Size.X * r.Size.Y
}

// GetPosition returns a vector with the starting point of the rectangle as x and y
func (r *Rect2D[T]) GetPosition() *vector.Vector2D[T] {
	return r.Position
}

// GetSize returns a vector with the current size of the Rectangle as width and height
func (r *Rect2D[T]) GetSize() *vector.Vector2D[T] {
	return r.Size
}

// Abs returns a Rect2D equivalent to this rectangle, with its width and height modified to be non-negative values, and
// with its position being the top-left corner of the rectangle.
// E.g: Rect2D(25, 25, -100, -50) ~> Rect2D(-75, -25, 100, 50)
func (r *Rect2D[T]) Abs() *Rect2D[T] {
	return New2D[T](
		vector.New2D[T](
			r.Position.X+r.min(r.Size.X, 0),
			r.Position.Y+r.min(r.Size.Y, 0),
		),
		r.Size.Abs(),
	)
}

// GetRandomPoint generates a random coordinates vector within the rectangle's bounds
func (r *Rect2D[T]) GetRandomPoint() *vector.Vector2D[T] {
	return vector.New2D[T](
		T(rand.Float64())*(r.Size.X)+r.Position.X,
		T(rand.Float64())*(r.Size.Y)+r.Position.Y,
	)
}

func (r *Rect2D[T]) min(a, b T) T {
	if a < b {
		return a
	}
	return b
}
