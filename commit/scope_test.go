package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestScope_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    commit.Scope
		want bool
	}{
		{
			name: "true on empty scope",
			s:    commit.NewScope(""),
			want: true,
		},
		{
			name: "true on scope with spaces",
			s:    commit.NewScope("   "),
			want: true,
		},
		{
			name: "false on non-empty scope",
			s:    commit.NewScope("ui"),
			want: false,
		},
		{
			name: "true on scope with new line only",
			s:    commit.NewScope("\n"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.s.Empty()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestScope_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		s    commit.Scope
		want string
	}{
		{
			name: "empty string on empty scope",
			s:    commit.NewScope(""),
			want: "",
		},
		{
			name: "preserves spaces",
			s:    commit.NewScope("   "),
			want: "   ",
		},
		{
			name: "returns scope value",
			s:    commit.NewScope("ui"),
			want: "ui",
		},
		{
			name: "preserves newline",
			s:    commit.NewScope("\n"),
			want: "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.s.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
