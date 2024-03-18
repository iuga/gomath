package matrix

import (
	"testing"

	"github.com/iuga/gomath/vector"
	"github.com/stretchr/testify/require"
)

func TestMatrix2DSlice(t *testing.T) {

	m := New2D[int](vector.New2D[int](4, 4))
	m.Set(vector.New2D[int](0, 0), 1)
	m.Set(vector.New2D[int](1, 1), 1)
	m.Set(vector.New2D[int](2, 2), 1)
	m.Set(vector.New2D[int](3, 3), 1)
	t.Log(m)

	s, _ := m.Slice(vector.New2D[int](1, 3), vector.New2D[int](0, 1))
	t.Log(s)
	t.Fail()
}

func TestMatrix2D(t *testing.T) {
	m := New2D[int](vector.New2D[int](5, 6))

	m.Set(vector.New2D[int](0, 0), 1)
	m.Set(vector.New2D[int](1, 1), 1)
	m.Set(vector.New2D[int](2, 2), 1)
	m.Set(vector.New2D[int](3, 3), 1)
	m.Set(vector.New2D[int](4, 4), 1)
	t.Log(m)
	t.Fail()

	s, err := m.Slice(vector.New2D[int](1, 4), vector.New2D[int](1, 2))
	require.NoError(t, err)
	require.NotNil(t, s)

	s, _ = m.Slice(vector.New2D[int](1, 3), vector.New2D[int](1, 3))
	require.NoError(t, err)
	require.NotNil(t, s)

	s, err = m.Slice(vector.New2D[int](-1, -13), vector.New2D[int](1, 3))
	require.Error(t, err)
	require.Nil(t, s)

	s, err = m.Slice(vector.New2D[int](0, 5), vector.New2D[int](0, 6))
	require.Error(t, err)
	require.Nil(t, s)
}
