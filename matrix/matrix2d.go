package matrix

import (
	"errors"
	"fmt"
	"strings"

	"github.com/iuga/gomath/vector"
	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

// Shape declares the shape of the matrix as Width/Rows and Height/Columns
type Shape struct {
	// The Width or Columns of the Matrix
	Width int `json:"width"`
	// The Height or Rows of the Matrix
	Height int `json:"height"`
}

// NewShape returns a new Matrix shape
func NewShape(width int, height int) *Shape {
	return &Shape{
		Width:  width,
		Height: height,
	}
}

// Position declares a coordinate in the array as row(x) and column(y)
type Position struct {
	// The row to locate
	Row int `json:"row"`
	// The column to locate
	Column int `json:"column"`
}

// NewPosition returns a new Matrix position to be used in lookups and slices
func NewPosition(row int, column int) *Position {
	return &Position{
		Row:    row,
		Column: column,
	}
}

type Matrix2DOpts[T Number] func(*Matrix2D[T]) error

type Matrix2D[T Number] struct {
	// Values of the matrix
	Values [][]T
	// Shape of the matrix as (x=width and y=height)
	Shape *Shape
}

// New2D constructs a new Matrix2D from the given shape (x=width and y=height).
func New2D[T Number](shape *Shape, opts ...Matrix2DOpts[T]) (*Matrix2D[T], error) {
	rows := make([][]T, 0, shape.Height)
	for i := 0; i < shape.Height; i++ {
		col := make([]T, shape.Width)
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
			return errors.New("WithData(...) should have a value to initialize the matrix")
		}
		if len(data) == 0 {
			return fmt.Errorf("WithData(...) empty data not alowed. Use WithConstant(...) instead.")
		}
		if len(data) != m.GetShape().Height {
			return fmt.Errorf("WithData(...) 'height/rows' shape missmatch -> data: %v matrix: %v", len(data), m.GetShape())
		}
		if len(data[0]) != m.GetShape().Width {
			return fmt.Errorf("WithData(...) 'width/cols' shape missmatch -> data: %v matrix: %v", len(data), m.GetShape())
		}
		for yRow, row := range data {
			for xCol, v := range row {
				if err := m.Set(&Position{yRow, xCol}, v); err != nil {
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
		for y := 0; y < m.GetShape().Height; y++ {
			for x := 0; x < m.GetShape().Width; x++ {
				if err := m.Set(NewPosition(y, x), k); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return f
}

// Set a value in the given position as (y, x) == (row, column)
func (m Matrix2D[T]) Set(position *Position, value T) error {
	if m.isPositionOutOfBounds(position) {
		return fmt.Errorf("out of bound position %v with shape %v", position, m.Shape)
	}
	m.Values[position.Row][position.Column] = value
	return nil
}

// At returns the value of the matrix AT that position
func (m Matrix2D[T]) At(position *Position) (T, error) {
	if m.isPositionOutOfBounds(position) {
		return 0, fmt.Errorf("out of bound position %v with shape %v", position, m.Shape)
	}
	return m.Values[position.Row][position.Column], nil
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
	shape := &Shape{Width: (x.Y + 1 - x.X), Height: (x.Y + 1 - x.X)}
	if shape.Width < 0 || shape.Height < 0 || shape.Width > m.GetShape().Width || shape.Height > m.GetShape().Height {
		return nil, fmt.Errorf("slice %v out of bounds on matrix of shape %v", shape, m.Shape)
	}
	rows := m.Values[y.X : y.Y+1]
	vrows := make([][]T, len(rows))
	for row := range rows {
		vrows[row] = rows[row][x.X : x.Y+1]
	}
	return &Matrix2D[T]{
		Values: vrows,
		Shape:  shape,
	}, nil
}

// Update the matrix data with another matrix starting in position x,y
func (m *Matrix2D[T]) Update(position *Position, matrix *Matrix2D[T]) error {
	// If the position is out of bounds, fail
	if m.isPositionOutOfBounds(position) {
		return fmt.Errorf("update starting position %v is out of bounds %v", position, m.Shape)
	}
	// If section is out of bounds, fail
	if position.Column+matrix.GetShape().Width > m.GetShape().Width || position.Row+matrix.GetShape().Height > m.GetShape().Height {
		return fmt.Errorf("slide to update is out of bounds on: %v base: %v update: %v", position, m.Shape, matrix.Shape)
	}
	for y := 0; y < matrix.GetShape().Height; y++ {
		for x := 0; x < matrix.GetShape().Width; x++ {
			v, err := matrix.At(NewPosition(y, x))
			if err != nil {
				return err
			}
			m.Set(NewPosition(y+position.Row, x+position.Column), v)
		}
	}
	return nil
}

// GetShape returns a vector representing the dimensionality of the Matrix2D as (rows, columns).
func (m Matrix2D[T]) GetShape() *Shape {
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
	b.WriteString(fmt.Sprintf("shape: (h:%d,w:%d)\n", m.Shape.Height, m.Shape.Width))
	return b.String()
}

func (m Matrix2D[T]) isPositionOutOfBounds(p *Position) bool {
	return p.Row >= m.GetShape().Height || p.Column >= m.GetShape().Width || p.Row < 0 || p.Column < 0
}
