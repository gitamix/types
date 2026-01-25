package branch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/branch"
)

func TestName_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    branch.Name
		want bool
	}{
		{
			name: "created from issue",
			n:    branch.NewName("123-feature-branch"),
			want: false,
		},
		{
			name: "feature with ticket number",
			n:    branch.NewName("feature/456-add-new-endpoint"),
			want: false,
		},
		{
			name: "bugfix branch",
			n:    branch.NewName("bugfix/789-fix-login-issue"),
			want: false,
		},
		{
			name: "release branch",
			n:    branch.NewName("release/v1.2.0"),
			want: false,
		},
		{
			name: "hotfix branch",
			n:    branch.NewName("hotfix/v1.2.1"),
			want: false,
		},
		{
			name: "invalid with spaces only",
			n:    branch.NewName("   "),
			want: true,
		},
		{
			name: "with numbers only",
			n:    branch.NewName("123"),
			want: false,
		},
		{
			name: "default value",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.n.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestName_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    branch.Name
		want string
	}{
		{
			name: "created from issue",
			n:    branch.NewName("123-feature-branch"),
			want: "123-feature-branch",
		},
		{
			name: "feature with ticket number",
			n:    branch.NewName("feature/456-add-new-endpoint"),
			want: "feature/456-add-new-endpoint",
		},
		{
			name: "bugfix branch",
			n:    branch.NewName("bugfix/789-fix-login-issue"),
			want: "bugfix/789-fix-login-issue",
		},
		{
			name: "release branch",
			n:    branch.NewName("release/v1.2.0"),
			want: "release/v1.2.0",
		},
		{
			name: "hotfix branch",
			n:    branch.NewName("hotfix/v1.2.1"),
			want: "hotfix/v1.2.1",
		},
		{
			name: "invalid with spaces only",
			n:    branch.NewName("   "),
			want: "   ",
		},
		{
			name: "with numbers only",
			n:    branch.NewName("123"),
			want: "123",
		},
		{
			name: "correct but with spaces",
			n:    branch.NewName("  feature/branch  "),
			want: "  feature/branch  ",
		},
		{
			name: "default value",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.n.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
