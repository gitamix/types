package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestParseSubject(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want commit.Subject
	}{
		{
			name: "fully correct subject",
			args: args{
				s: "feat(ui): add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "correct subject with spaces",
			args: args{
				s: "   feat  (ui) :    add new button  ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "subject with type and description only",
			args: args{
				s: "feat: add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "description only",
			args: args{
				s: "add new button",
			},
			want: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope(""),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "wrong subject with type only",
			args: args{
				s: "feat: ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with scope only",
			args: args{
				s: "(ui): ",
			},
			want: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope("ui"),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with feat and scope only",
			args: args{
				s: "feat(ui): ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with invalid scope format",
			args: args{
				s: "feat: ui: add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription("ui: add new button"),
			),
		},
		{
			name: "empty subject",
			args: args{
				s: "",
			},
			want: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope(""),
				commit.NewDescription(""),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := commit.ParseSubject(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
