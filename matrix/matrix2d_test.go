package matrix

import (
	"testing"

	"github.com/iuga/gomath/vector"
	"github.com/stretchr/testify/require"
)

func TestMatrix2DUpdate(t *testing.T) {
	m, err := New2D[int](NewShape(5, 5), WithConstant[int](0))
	require.NoError(t, err)
	require.NotNil(t, m)

	mu, err := New2D[int](NewShape(3, 2), WithConstant[int](1))
	require.NoError(t, err)
	require.NotNil(t, mu)

	err = m.Update(NewPosition(2, 1), mu)
	require.NoError(t, err)

	t.Log(m)
	require.Equal(t, [][]int{
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 0, 0, 0},
		[]int{0, 1, 1, 1, 0},
		[]int{0, 1, 1, 1, 0},
		[]int{0, 0, 0, 0, 0},
	}, m.GetValues())

	mu2, err := New2D[int](NewShape(2, 3), WithConstant[int](2))
	require.NoError(t, err)
	require.NotNil(t, mu2)

	err = m.Update(NewPosition(1, 2), mu2)
	require.NoError(t, err)
	t.Log(m)
	require.Equal(t, [][]int{
		[]int{0, 0, 0, 0, 0},
		[]int{0, 0, 2, 2, 0},
		[]int{0, 1, 2, 2, 0},
		[]int{0, 1, 2, 2, 0},
		[]int{0, 0, 0, 0, 0},
	}, m.GetValues())
}

func TestMatrix2DSlice(t *testing.T) {

	m, err := New2D[int](NewShape(4, 4))
	require.NoError(t, err)

	m.Set(NewPosition(0, 0), 1)
	m.Set(NewPosition(1, 1), 1)
	m.Set(NewPosition(2, 2), 1)
	m.Set(NewPosition(3, 3), 1)
	t.Log(m)

	s, _ := m.Slice(vector.New2D[int](1, 3), vector.New2D[int](0, 1))
	t.Log(s)
	require.Equal(t, [][]int{
		[]int{0, 0, 0},
		[]int{1, 0, 0},
	}, s.GetValues())
}

func TestMatrix2D(t *testing.T) {
	m, err := New2D[int](NewShape(5, 6))
	require.NoError(t, err)

	m.Set(NewPosition(0, 0), 1)
	m.Set(NewPosition(1, 1), 1)
	m.Set(NewPosition(2, 2), 1)
	m.Set(NewPosition(3, 3), 1)
	m.Set(NewPosition(4, 4), 1)
	t.Log(m)

	s, err := m.Slice(vector.New2D[int](1, 4), vector.New2D[int](1, 2))
	require.NoError(t, err)
	require.NotNil(t, s)
	t.Log(s)
	require.Equal(t, [][]int{
		[]int{1, 0, 0, 0},
		[]int{0, 1, 0, 0},
	}, s.GetValues())

	s, _ = m.Slice(vector.New2D[int](1, 3), vector.New2D[int](1, 3))
	require.NoError(t, err)
	require.NotNil(t, s)
	t.Log(s)
	require.Equal(t, [][]int{
		[]int{1, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 1},
	}, s.GetValues())

	s, err = m.Slice(vector.New2D[int](-1, -13), vector.New2D[int](1, 3))
	require.Error(t, err)
	require.Nil(t, s)

	s, err = m.Slice(vector.New2D[int](0, 5), vector.New2D[int](0, 6))
	require.Error(t, err)
	require.Nil(t, s)
}
