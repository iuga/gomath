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

type Matrix2DOpts[T Number] func(*Matrix2D[T]) error

type Matrix2D[T Number] struct {
	Values [][]T
	Shape  *vector.Vector2D[int]
}

// New2D constructs a new Matrix2D from the given shape (x=rows, y=columns).
func New2D[T Number](shape *vector.Vector2D[int], opts ...Matrix2DOpts[T]) (*Matrix2D[T], error) {
	rows := make([][]T, 0, shape.X)
	for i := 0; i < shape.X; i++ {
		col := make([]T, shape.Y)
		rows = append(rows, col)
	}
	m := &Matrix2D[T]{
		Values: rows,
		Shape:  shape,
	}
	for _, opt := range opts {
		if err := opt(m); err != nil {
			return nil, err
		}
	}
	return m, nil
}

// WithData will use the slice of slice values to initialize the matrix
func WithData[T Number](data [][]T) Matrix2DOpts[T] {
	f := func(m *Matrix2D[T]) error {
		if data == nil {
			return nil
		}
		if len(data) == 0 {
			return fmt.Errorf("WithData(...) empty data not alowed. Use WithConstant(...) instead.")
		}
		if len(data) != m.GetShape().Y {
			return fmt.Errorf("WithData(...) 'y' shape missmatch -> data: %v matrix: %v", len(data), m.GetShape().Y)
		}
		if len(data[0]) != m.GetShape().X {
			return fmt.Errorf("WithData(...) 'x' shape missmatch -> data: %v matrix: %v", len(data), m.GetShape().Y)
		}
		for yRow, row := range data {
			for xCol, v := range row {
				if err := m.Set(vector.New2D[int](xCol, yRow), v); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return f
}

// WithConstant initialized the matrix with a constant value. E.g: 1
func WithConstant[T Number](k T) Matrix2DOpts[T] {
	f := func(m *Matrix2D[T]) error {
		if k == 0 {
			return nil // Zero Value, nothing to do...
		}
		rows := m.GetShape().X
		cols := m.GetShape().Y
		for y := 0; y < rows; y++ {
			for x := 0; x < cols; x++ {
				if err := m.Set(vector.New2D[int](x, y), k); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return f
}

// Set a value in the given position as (x, y) == (column, row)
func (m Matrix2D[T]) Set(position *vector.Vector2D[int], value T) error {
	if position.X >= m.Shape.Y || position.Y >= m.Shape.X || position.X < 0 || position.Y < 0 {
		return fmt.Errorf("out of bound position %v with shape %v", position, m.Shape)
	}
	m.Values[position.Y][position.X] = value
	return nil
}

// At returns the value of the matrix AT that position
func (m Matrix2D[T]) At(position *vector.Vector2D[int]) (T, error) {
	if position.X >= m.Shape.Y || position.Y >= m.Shape.X || position.X < 0 || position.Y < 0 {
		return 0, fmt.Errorf("out of bound position %v with shape %v", position, m.Shape)
	}
	return m.Values[position.Y][position.X], nil
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
	if sShape.X < 0 || sShape.Y < 0 || sShape.X > m.Shape.X || sShape.Y > m.Shape.Y {
		return nil, fmt.Errorf("slice %v out of bounds on matrix of shape %v", sShape, m.Shape)
	}
	rows := m.Values[y.X : y.Y+1]
	vrows := make([][]T, len(rows))
	for row := range rows {
		vrows[row] = rows[row][x.X : x.Y+1]
	}
	return &Matrix2D[T]{
		Values: vrows,
		Shape:  sShape,
	}, nil
}

// Update the matrix data with another matrix starting in position x,y
func (m *Matrix2D[T]) Update(position *vector.Vector2D[int], matrix *Matrix2D[T]) error {
	// If the position is out of bounds, fail
	if position.X >= m.Shape.X || position.Y >= m.Shape.Y {
		return fmt.Errorf("update starting position %v is out of bounds %v", position, m.Shape)
	}
	// If section is out of bounds, fail
	if position.X+matrix.Shape.X > m.Shape.X || position.Y+matrix.Shape.Y > m.Shape.Y {
		return fmt.Errorf("slide to update is out of bounds on: %v base: %v update: %v", position, m.Shape, matrix.Shape)
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
	return m.Shape
}

// GetValues return the invernal values as a slice of slices
func (m Matrix2D[T]) GetValues() [][]T {
	return m.Values
}

// String returns a human-readable representation of the Matrix2D
func (m Matrix2D[T]) String() string {
	var b strings.Builder
	b.WriteString("\n")
	for _, row := range m.Values {
		for _, col := range row {
			b.WriteString(fmt.Sprintf(" %04v ", col))
		}
		b.WriteString("\n")
	}
	b.WriteString(fmt.Sprintf("shape: (%d,%d)\n", m.Shape.X, m.Shape.Y))
	return b.String()
}
