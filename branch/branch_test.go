package branch_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/branch"
)

func TestBranch_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		name branch.Name
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "created from issue",
			fields: fields{
				name: branch.NewName("123-feature-branch"),
			},
			want: "123-feature-branch",
		},
		{
			name: "feature with ticket number",
			fields: fields{
				name: branch.NewName("feature/456-add-new-endpoint"),
			},
			want: "feature/456-add-new-endpoint",
		},
		{
			name: "bugfix branch",
			fields: fields{
				name: branch.NewName("bugfix/789-fix-login-issue"),
			},
			want: "bugfix/789-fix-login-issue",
		},
		{
			name: "release branch",
			fields: fields{
				name: branch.NewName("release/v1.2.0"),
			},
			want: "release/v1.2.0",
		},
		{
			name: "hotfix branch",
			fields: fields{
				name: branch.NewName("hotfix/v1.2.1"),
			},
			want: "hotfix/v1.2.1",
		},
		{
			name: "invalid with spaces only",
			fields: fields{
				name: branch.NewName("   "),
			},
			want: "   ",
		},
		{
			name: "with numbers only",
			fields: fields{
				name: branch.NewName("123"),
			},
			want: "123",
		},
		{
			name: "correct but with spaces",
			fields: fields{
				name: branch.NewName("  feature/branch  "),
			},
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
			b := branch.NewBranch(tt.fields.name)
			got := b.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
