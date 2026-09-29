package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	impl "github.com/gitamix/types/commit"
)

func TestKind_Default(t *testing.T) {
	t.Parallel()

	t.Run("const default returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.KindDefault.Default())
	})

	t.Run("type of zero returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.Kind(0).Default())
	})

	t.Run("default type returns true", func(t *testing.T) {
		t.Parallel()
		var k impl.Kind
		assert.True(t, k.Default())
	})

	t.Run("type of one returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(1).Default())
	})

	t.Run("type of two returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(2).Default())
	})

	t.Run("type of three returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(3).Default())
	})

	t.Run("type of max uint8 returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(255).Default())
	})
}
