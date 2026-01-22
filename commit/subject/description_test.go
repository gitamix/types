package subject_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit/subject"
)

func TestDescription_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		d    subject.Description
		want bool
	}{
		{
			name: "empty",
			d:    subject.NewDescription(""),
			want: true,
		},
		{
			name: "spaces",
			d:    subject.NewDescription("   "),
			want: true,
		},
		{
			name: "feat",
			d:    subject.NewDescription("feat"),
			want: false,
		},
		{
			name: "new line",
			d:    subject.NewDescription("\n"),
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
		d    subject.Description
		want string
	}{
		{
			name: "empty",
			d:    subject.NewDescription(""),
			want: "",
		},
		{
			name: "spaces",
			d:    subject.NewDescription("   "),
			want: "   ",
		},
		{
			name: "feat",
			d:    subject.NewDescription("feat"),
			want: "feat",
		},
		{
			name: "new line",
			d:    subject.NewDescription("\n"),
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
		want subject.Description
	}{
		{
			name: "fully correct subject",
			args: args{
				s: "feat(ui): add new button",
			},
			want: subject.NewDescription("add new button"),
		},
		{
			name: "description only",
			args: args{
				s: "add new button",
			},
			want: subject.NewDescription("add new button"),
		},
		{
			name: "empty string",
			args: args{
				s: "",
			},
			want: subject.NewDescription(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := subject.ParseDescription(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
