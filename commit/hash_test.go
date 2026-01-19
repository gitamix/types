package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gitamix/types/commit"
)

func TestHash_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		h    commit.Hash
		want string
	}{
		{
			name: "short hash",
			h:    commit.NewHash("1234567"),
			want: "1234567",
		},
		{
			name: "full hash",
			h:    commit.NewHash("1234567890abcdef"),
			want: "1234567890abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.h.String()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHash_ShortString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		h    commit.Hash
		want string
	}{
		{
			name: "on hash less than 7 chars",
			h:    commit.NewHash("123456"),
			want: "123456",
		},
		{
			name: "on hash greater than 7 chars",
			h:    commit.NewHash("1234567890"),
			want: "1234567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.h.ShortString()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestHash_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		h    commit.Hash
		want bool
	}{
		{
			name: "true on empty hash",
			h:    commit.NewHash(""),
			want: true,
		},
		{
			name: "false on non-empty hash",
			h:    commit.NewHash("1234567"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.h.Empty()
			assert.Equal(t, tt.want, got)
		})
	}
}
