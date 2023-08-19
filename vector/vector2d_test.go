package vector_test

import (
	"testing"

	"github.com/iuga/gomath/vector"
	"github.com/stretchr/testify/require"
)

func TestMoveToward(t *testing.T) {
	v := vector.New2D[float32](10.0, 10.0)
	v = v.MoveToward(vector.New2D[float32](0.0, 0.0), 0.5)
	require.Less(t, float32(v.X), float32(10.0))
	require.Less(t, float32(v.Y), float32(10.0))
	v = v.MoveToward(vector.New2D[float32](0.0, 0.0), 0.5)
	require.Less(t, float32(v.X), float32(9.5))
	require.Less(t, float32(v.Y), float32(9.5))
	v = v.MoveToward(vector.New2D[float32](0.0, 0.0), 0.5)
	require.Less(t, float32(v.X), float32(9.0))
	require.Less(t, float32(v.Y), float32(9.0))
}
