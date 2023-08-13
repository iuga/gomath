package vector

import (
	"golang.org/x/exp/constraints"
)

var (
	// Integer Zero vector, a vector with all components set to 0.
	ZERO = New2D[int](0, 0)
	// Integer One vector, a vector with all components set to 1.
	ONE = New2D[int](1, 1)
	// Left unit vector. Represents the direction of left.
	LEFT = New2D[int](-1, 0)
	// Right unit vector. Represents the direction of right.
	RIGHT = New2D[int](1, 0)
	// Up unit vector. Y is down in 2D, so this vector points -Y.
	UP = New2D[int](0, -1)
	// Down unit vector. Y is down in 2D, so this vector points +Y.
	DOWN = New2D[int](0, 1)
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Vector2D is a 2D vector using numeric coordinates as generics.
// A 2-element structure that can be used to represent 2D coordinates or any other pair of numeric values.
type Vector2D[T Number] struct {
	// The vector's X component
	X T `json:"x"`
	// The vector's Y component
	Y T `json:"y"`
}

// New2D constructs a new Vector2 from the given x and y.
func New2D[T Number](x, y T) *Vector2D[T] {
	return &Vector2D[T]{X: x, Y: y}
}

// Returns the squared length (squared magnitude) of this vector.
func (v *Vector2D[T]) LengthSquared() T {
	return v.X*v.X + v.Y*v.Y
}

// Returns the dot product of this vector and with. This can be used to compare the angle between two vectors.
// For example, this can be used to determine whether an enemy is facing the player.
// The dot product will be 0 for a straight angle (90 degrees), greater than 0 for angles narrower than 90 degrees
// and lower than 0 for angles wider than 90 degrees. When using unit (normalized) vectors, the result will always
// be between -1.0 (180 degree angle) when the vectors are facing opposite directions, and 1.0 (0 degree angle) when the
// vectors are aligned.
// Note: a.dot(b) is equivalent to b.dot(a).
func (v *Vector2D[T]) Dot(w *Vector2D[T]) T {
	return v.X*w.X + v.Y*w.Y
}

// Returns the 2D analog of the cross product for this vector and with.
// This is the signed area of the parallelogram formed by the two vectors. If the second vector is clockwise from the
// / first vector, then the cross product is the positive area. If counter-clockwise, the cross product is the negative area.
// Note: Cross product is not defined in 2D mathematically. This method embeds the 2D vectors in the XY plane of 3D space
// and uses their cross product's Z component as the analog.
func (v *Vector2D[T]) Cross(w *Vector2D[T]) T {
	return v.X*w.Y + v.Y*w.X
}

// Returns the squared distance between this vector and w.
func (v *Vector2D[T]) DistanceSquaredTo(w *Vector2D[T]) T {
	return (v.X-w.X)*(v.X-w.X) + (v.Y-w.Y)*(v.Y-w.Y)
}
