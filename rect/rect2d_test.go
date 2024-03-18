package rect_test

import (
	"testing"

	"github.com/iuga/gomath/rect"
	"github.com/iuga/gomath/vector"
	"github.com/stretchr/testify/require"
)

func TestRect(t *testing.T) {
	r := rect.New2D[int](vector.New2D[int](1, 1), vector.New2D[int](2, 3))
	require.Equal(t, 6, r.GetArea())
	require.Equal(t, true, r.HasPoint(vector.New2D[int](2, 2)))
	require.Equal(t, false, r.HasPoint(vector.New2D[int](6, 2)))
	c := r.GetCenter()
	require.Equal(t, 2, c.X)
	require.Equal(t, 2, c.Y)
}
