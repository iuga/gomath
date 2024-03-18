package matrix

import (
	"fmt"
	"strings"

	"github.com/iuga/gomath/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Matrix2DOpts[T Number] func(*Matrix2D[T])

type Matrix2D[T Number] struct {
	values [][]T
	shape  *vector.Vector2D[int]
}

// New2D constructs a new Matrix2D from the given shape (x=rows, y=columns).
func New2D[T Number](shape *vector.Vector2D[int], opts ...Matrix2DOpts[T]) *Matrix2D[T] {
	rows := make([][]T, 0, shape.X)
	for i := 0; i < shape.X; i++ {
		col := make([]T, shape.Y)
		rows = append(rows, col)
	}
	m := &Matrix2D[T]{
		values: rows,
		shape:  shape,
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// WithData will use the slice of slice values to initialize the matrix
func WithData[T Number](data [][]T) Matrix2DOpts[T] {
	f := func(m *Matrix2D[T]) {
		if data == nil {
			return
		}
		if len(data) != m.GetShape().Y {
			return
		}
		for y, row := range data {
			for x, v := range row {
				m.Set(vector.New2D[int](x, y), v)
			}
		}
	}
	return f
}

// WithConstant initialized the matrix with a constant value. E.g: 1
func WithConstant[T Number](k T) Matrix2DOpts[T] {
	f := func(m *Matrix2D[T]) {
		if k == 0 {
			return // Zero Value
		}
		rows := m.GetShape().X
		cols := m.GetShape().Y
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				m.Set(vector.New2D[int](x, y), k)
			}
		}
	}
	return f
}

// Set a value in the given position as (x, y) == (column, row)
func (m Matrix2D[T]) Set(position *vector.Vector2D[int], value T) error {
	if position.X >= m.shape.Y || position.Y >= m.shape.X || position.X < 0 || position.Y < 0 {
		return fmt.Errorf("out of bound position %v with shape %v", position, m.shape)
	}
	m.values[position.Y][position.X] = value
	return nil
}

// At returns the value of the matrix AT that position
func (m Matrix2D[T]) At(position *vector.Vector2D[int]) (T, error) {
	if position.X >= m.shape.Y || position.Y >= m.shape.X || position.X < 0 || position.Y < 0 {
		return 0, fmt.Errorf("out of bound position %v with shape %v", position, m.shape)
	}
	return m.values[position.Y][position.X], nil
}

// Slice returns a subset of the matrix as a slice of slices taking two coordinates.
// m := New2D[int](vector.New2D[int](4, 4))
// m.Set(vector.New2D[int](0, 0), 1)
// m.Set(vector.New2D[int](1, 1), 1)
// m.Set(vector.New2D[int](2, 2), 1)
// m.Set(vector.New2D[int](3, 3), 1)
//
// [1 0 0 0]
// [0 1 0 0]
// [0 0 1 0]
// [0 0 0 1]
//
// In X from 1 to 3 (3 columns) and in Y from 0 to 1 (2 rows)
// m.Slice(vector.New2D[int](1, 3), vector.New2D[int](0, 1))
//
// [0 0 0]
// [1 0 0]
func (m Matrix2D[T]) Slice(x *vector.Vector2D[int], y *vector.Vector2D[int]) (*Matrix2D[T], error) {
	// This selects a half-open range which includes the first element, but excludes the last one.
	sShape := vector.New2D[int]((x.Y + 1 - x.X), (y.Y + 1 - y.X))
	if sShape.X < 0 || sShape.Y < 0 || sShape.X > m.shape.X || sShape.Y > m.shape.Y {
		return nil, fmt.Errorf("slice %v out of bounds on matrix of shape %v", sShape, m.shape)
	}
	rows := m.values[y.X : y.Y+1]
	vrows := make([][]T, len(rows))
	for row := range rows {
		vrows[row] = rows[row][x.X : x.Y+1]
	}
	return &Matrix2D[T]{
		values: vrows,
		shape:  sShape,
	}, nil
}

// Update the matrix data with another matrix starting in position x,y
func (m *Matrix2D[T]) Update(position *vector.Vector2D[int], matrix *Matrix2D[T]) error {
	// If the position is out of bounds, fail
	if position.X >= m.shape.X || position.Y >= m.shape.Y {
		return fmt.Errorf("update starting position %v is out of bounds %v", position, m.shape)
	}
	// If section is out of bounds, fail
	if position.X+matrix.shape.X >= m.shape.X || position.Y+matrix.shape.Y >= m.shape.Y {
		return fmt.Errorf("slide to update is out of bounds %v", m.shape)
	}
	rows := matrix.GetShape().X
	cols := matrix.GetShape().Y
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			v, err := matrix.At(vector.New2D[int](x, y))
			if err != nil {
				return err
			}
			m.Set(vector.New2D[int](x+position.X, y+position.Y), v)
		}
	}
	return nil
}

// GetShape returns a vector representing the dimensionality of the Matrix2D as (rows, columns).
func (m Matrix2D[T]) GetShape() *vector.Vector2D[int] {
	return m.shape
}

// GetValues return the invernal values as a slice of slices
func (m Matrix2D[T]) GetValues() [][]T {
	return m.values
}

// String returns a human-readable representation of the Matrix2D
func (m Matrix2D[T]) String() string {
	var b strings.Builder
	b.WriteString("\n")
	for _, row := range m.values {
		for _, col := range row {
			b.WriteString(fmt.Sprintf(" %04v ", col))
		}
		b.WriteString("\n")
	}
	b.WriteString(fmt.Sprintf("shape: (%d,%d)\n", m.shape.X, m.shape.Y))
	return b.String()
}
