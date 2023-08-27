package vector

import (
	"math"

	"golang.org/x/exp/constraints"
)

var (
	// EPSILON define small comparisons
	EPSILON = float64(0.00001)
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

// Returns the distance between this vector and w.
func (v *Vector2D[T]) DistanceTo(w *Vector2D[T]) T {
	return T(math.Sqrt(float64(v.DistanceSquaredTo(w))))
}

// Returns the squared distance between this vector and w.
func (v *Vector2D[T]) DistanceSquaredTo(w *Vector2D[T]) T {
	return (v.X-w.X)*(v.X-w.X) + (v.Y-w.Y)*(v.Y-w.Y)
}

// Returns the result of the linear interpolation between this vector and to by amount weight.
// weight is on the range of 0.0 to 1.0, representing the amount of interpolation.
// interpolation = A * (1 - t) + B * t =  A + (B - A) * t
func (v *Vector2D[T]) LERP(to *Vector2D[T], weight float64) *Vector2D[T] {
	// return p_from + (p_to - p_from) * p_weight;
	return New2D[T](
		v.X+(to.X-v.X)*T(weight),
		v.Y+(to.Y-v.Y)*T(weight),
	)
}

// Add one vector to another
func (v *Vector2D[T]) Add(to *Vector2D[T]) *Vector2D[T] {
	return New2D[T](v.X+to.X, v.Y+to.Y)
}

// Subtract one vector from another
func (v *Vector2D[T]) Subtract(u *Vector2D[T]) *Vector2D[T] {
	return New2D[T](v.X-u.X, v.Y-u.Y)
}

// Lenght returns the length (magnitude) of this vector.
func (v *Vector2D[T]) Length() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

// Returns the result of scaling the vector to unit length. Equivalent to v / v.length().
// Note: This function may return incorrect values if the input vector length is near zero.
func (v *Vector2D[T]) Normalized() *Vector2D[T] {
	// real_t l = x * x + y * y;
	l := v.X*v.X + v.Y*v.Y
	if l != 0 {
		l = T(math.Sqrt(float64(l)))
		return New2D[T](v.X/l, v.Y/l)
	}
	return New2D[T](v.X, v.Y)
}

// MoveToward returns a new vector moved toward to by the fixed delta amount. Will not go past the final value.
func (v *Vector2D[T]) MoveToward(to *Vector2D[T], delta T) *Vector2D[T] {
	vd := to.Subtract(v)
	l := vd.Length()
	if l <= delta || l < T(EPSILON) {
		return New2D[T](to.X, to.Y)
	}
	return New2D[T](v.X+vd.X/l*delta, v.Y+vd.Y/l*delta)
}

// DirectionTo returns the normalized vector pointing from this vector to to.
// This is equivalent to using (b - a).normalized().
func (v *Vector2D[T]) DirectionTo(to *Vector2D[T]) *Vector2D[T] {
	return to.Subtract(v).Normalized()
}

// Clone returns a deep-copy of the current vector.
func (v *Vector2D[T]) Clone() *Vector2D[T] {
	return New2D[T](v.X, v.Y)
}
