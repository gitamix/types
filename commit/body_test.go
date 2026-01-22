package commit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBody_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		b    Body
		want bool
	}{
		{
			name: "empty body",
			b:    NewBody([]byte("")),
			want: true,
		},
		{
			name: "body with only spaces",
			b:    NewBody([]byte("     ")),
			want: true,
		},
		{
			name: "body with only newlines",
			b:    NewBody([]byte("\n\n\n")),
			want: true,
		},
		{
			name: "body with spaces and newlines",
			b:    NewBody([]byte("  \n  \n")),
			want: true,
		},
		{
			name: "non-empty body",
			b:    NewBody([]byte("This is a commit body.")),
		},
		{
			name: "body with leading and trailing spaces",
			b:    NewBody([]byte("   This is a commit body.   ")),
		},
		{
			name: "body with multiple lines of text",
			b:    NewBody([]byte("Line 1\nLine 2\nLine 3")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestBody_String(t *testing.T) {
	tests := []struct {
		name string
		b    Body
		want string
	}{
		{
			name: "body with text",
			b:    NewBody([]byte("This is a commit body.")),
			want: "This is a commit body.",
		},
		{
			name: "empty body",
			b:    NewBody([]byte("")),
			want: "",
		},
		{
			name: "body with multiple lines",
			b:    NewBody([]byte("Line 1\nLine 2\nLine 3")),
			want: "Line 1\nLine 2\nLine 3",
		},
		{
			name: "body with leading and trailing spaces",
			b:    NewBody([]byte("   This is a commit body.   ")),
			want: "   This is a commit body.   ",
		},
		{
			name: "body with special characters",
			b:    NewBody([]byte("Special chars: !@#$%^&*()_+")),
			want: "Special chars: !@#$%^&*()_+",
		},
		{
			name: "body with unicode characters",
			b:    NewBody([]byte("Unicode test: 測試")),
			want: "Unicode test: 測試",
		},
		{
			name: "body with tabs and newlines",
			b:    NewBody([]byte("Line 1\twith tab\nLine 2")),
			want: "Line 1\twith tab\nLine 2",
		},
		{
			name: "body with only whitespace",
			b:    NewBody([]byte("     \n   \t  ")),
			want: "     \n   \t  ",
		},
		{
			name: "body with mixed content",
			b:    NewBody([]byte("  Line 1  \n\nLine 2\t\n  ")),
			want: "  Line 1  \n\nLine 2\t\n  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.b.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
