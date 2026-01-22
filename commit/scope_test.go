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
			name: "ok",
			s:    commit.NewScope("ui"),
			want: "ui",
		},
		{
			name: "empty",
			s:    commit.NewScope(""),
			want: "",
		},
		{
			name: "spaces",
			s:    commit.NewScope("   "),
			want: "   ",
		},
		{
			name: "newline",
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

func TestParseScope(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want commit.Scope
	}{
		{
			name: "fully correct commit subject",
			args: args{
				s: "feat(ui): add new button",
			},
			want: commit.NewScope("ui"),
		},
		{
			name: "just a string with scopes",
			args: args{
				s: "(ui)",
			},
			want: commit.NewScope("ui"),
		},
		{
			name: "just a string with scope incorrectly closed",
			args: args{
				s: "(ui",
			},
			want: commit.NewScope(""),
		},
		{
			name: "just a string with scope incorrectly opened",
			args: args{
				s: "ui)",
			},
			want: commit.NewScope(""),
		},
		{
			name: "just a string without scopes",
			args: args{
				s: "ui",
			},
			want: commit.NewScope(""),
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: commit.NewScope(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := commit.ParseScope(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
