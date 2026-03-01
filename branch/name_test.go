package branch_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/branch"
	"github.com/gitamix/types/ticket"
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

func TestName_Ticket(t *testing.T) {
	t.Parallel()
	type args struct {
		re *regexp.Regexp
	}
	tests := []struct {
		name      string
		n         branch.Name
		args      args
		want      ticket.Ticket
		wantPanic bool
	}{
		{
			name: "branch name with task id",
			n:    branch.NewName("feature/TASK-1234"),
			args: args{
				re: regexp.MustCompile(`^(?:feature|bugfix|hotfix)/([A-Z]+-\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("TASK-1234"),
			),
			wantPanic: false,
		},
		{
			name: "branch name without ticket",
			n:    branch.NewName("feature/add-new-endpoint"),
			args: args{
				re: regexp.MustCompile(`^(?:feature|bugfix|hotfix)/([A-Z]+-\d+)`),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "empty branch name",
			n:    branch.NewName(""),
			args: args{
				re: regexp.MustCompile(`^(?:feature|bugfix|hotfix)/([A-Z]+-\d+)`),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "nil regex",
			n:    branch.NewName("feature/TASK-1234"),
			args: args{
				re: nil,
			},
			want:      ticket.Ticket{},
			wantPanic: true,
		},
		{
			name:      "default args",
			n:         branch.NewName(""),
			args:      args{},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.wantPanic {
				assert.Panics(t, func() {
					_ = tt.n.Ticket(tt.args.re)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				tt.n.Ticket(tt.args.re),
			)
		})
	}
}
