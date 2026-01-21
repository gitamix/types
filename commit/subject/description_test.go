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
