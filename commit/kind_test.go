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

	t.Run("const merge returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindMerge.Default())
	})

	t.Run("const revert returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindRevert.Default())
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

func TestKind_Merge(t *testing.T) {
	t.Parallel()

	t.Run("const merge returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.KindMerge.Merge())
	})

	t.Run("const default returns true", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindDefault.Merge())
	})

	t.Run("const revert returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindRevert.Merge())
	})

	t.Run("type of zero returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(0).Merge())
	})

	t.Run("default type returns false", func(t *testing.T) {
		t.Parallel()
		var k impl.Kind
		assert.False(t, k.Merge())
	})

	t.Run("type of one returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.Kind(1).Merge())
	})

	t.Run("type of two returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(2).Merge())
	})

	t.Run("type of three returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(3).Merge())
	})

	t.Run("type of max uint8 returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(255).Merge())
	})
}

func TestKind_Revert(t *testing.T) {
	t.Parallel()

	t.Run("const revert returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.KindRevert.Revert())
	})

	t.Run("const merge returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindMerge.Revert())
	})

	t.Run("const default returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.KindDefault.Revert())
	})

	t.Run("type of zero returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(0).Revert())
	})

	t.Run("default type returns false", func(t *testing.T) {
		t.Parallel()
		var k impl.Kind
		assert.False(t, k.Revert())
	})

	t.Run("type of one returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(1).Revert())
	})

	t.Run("type of two returns true", func(t *testing.T) {
		t.Parallel()
		assert.True(t, impl.Kind(2).Revert())
	})

	t.Run("type of three returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(3).Revert())
	})

	t.Run("type of max uint8 returns false", func(t *testing.T) {
		t.Parallel()
		assert.False(t, impl.Kind(255).Revert())
	})
}
