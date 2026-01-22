package commit_test

import (
	"testing"

	"github.com/gitamix/types/commit"
	"github.com/stretchr/testify/assert"
)

func TestDescription_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		d    commit.Description
		want bool
	}{
		{
			name: "empty",
			d:    commit.NewDescription(""),
			want: true,
		},
		{
			name: "spaces",
			d:    commit.NewDescription("   "),
			want: true,
		},
		{
			name: "feat",
			d:    commit.NewDescription("feat"),
			want: false,
		},
		{
			name: "new line",
			d:    commit.NewDescription("\n"),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.d.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestDescription_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		d    commit.Description
		want string
	}{
		{
			name: "empty",
			d:    commit.NewDescription(""),
			want: "",
		},
		{
			name: "spaces",
			d:    commit.NewDescription("   "),
			want: "   ",
		},
		{
			name: "feat",
			d:    commit.NewDescription("feat"),
			want: "feat",
		},
		{
			name: "new line",
			d:    commit.NewDescription("\n"),
			want: "\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.d.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseDescription(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want commit.Description
	}{
		{
			name: "fully correct subject",
			args: args{
				s: "feat(ui): add new button",
			},
			want: commit.NewDescription("add new button"),
		},
		{
			name: "description only",
			args: args{
				s: "add new button",
			},
			want: commit.NewDescription("add new button"),
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: commit.NewDescription(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := commit.ParseDescription(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
