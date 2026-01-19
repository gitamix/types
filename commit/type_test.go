package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestType_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		tr   commit.Type
		want bool
	}{
		{
			name: "empty",
			tr:   commit.NewType(""),
			want: true,
		},
		{
			name: "spaces",
			tr:   commit.NewType("   "),
			want: true,
		},
		{
			name: "feat",
			tr:   commit.NewType("feat"),
			want: false,
		},
		{
			name: "new line",
			tr:   commit.NewType("\n"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.tr.Empty()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestType_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		tr   commit.Type
		want string
	}{
		{
			name: "empty",
			tr:   commit.NewType(""),
			want: "",
		},
		{
			name: "spaces",
			tr:   commit.NewType("   "),
			want: "   ",
		},
		{
			name: "feat",
			tr:   commit.NewType("feat"),
			want: "feat",
		},
		{
			name: "new line",
			tr:   commit.NewType("\n"),
			want: "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.tr.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
