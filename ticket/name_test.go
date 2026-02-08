package ticket_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/ticket"
)

func TestName_String(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    ticket.Name
		want string
	}{
		{
			name: "jira ticket name",
			n:    ticket.NewName("JIRA-1234"),
			want: "JIRA-1234",
		},
		{
			name: "empty ticket name",
			n:    ticket.NewName(""),
			want: "",
		},
		{
			name: "simple ticket name",
			n:    ticket.NewName("Fix login issue"),
			want: "Fix login issue",
		},
		{
			name: "with spaces only",
			n:    ticket.NewName("   "),
			want: "   ",
		},
		{
			name: "with spaces on both side of string",
			n:    ticket.NewName("  JIRA-1234  "),
			want: "  JIRA-1234  ",
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

func TestName_Empty(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		n    ticket.Name
		want bool
	}{
		{
			name: "jira ticket name",
			n:    ticket.NewName("JIRA-1234"),
			want: false,
		},
		{
			name: "empty ticket name",
			n:    ticket.NewName(""),
			want: true,
		},
		{
			name: "simple ticket name",
			n:    ticket.NewName("Fix login issue"),
			want: false,
		},
		{
			name: "with spaces only",
			n:    ticket.NewName("   "),
			want: false,
		},
		{
			name: "with spaces on both side of string",
			n:    ticket.NewName("  JIRA-1234  "),
			want: false,
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

func TestName_Equals(t *testing.T) {
	t.Parallel()
	type args struct {
		other ticket.Name
	}
	tests := []struct {
		name string
		n    ticket.Name
		args args
		want bool
	}{
		{
			name: "same",
			n:    ticket.NewName("JIRA-1234"),
			args: args{other: ticket.NewName("JIRA-1234")},
			want: true,
		},
		{
			name: "different",
			n:    ticket.NewName("JIRA-1234"),
			args: args{other: ticket.NewName("JIRA-5678")},
			want: false,
		},
		{
			name: "empty vs non-empty",
			n:    ticket.NewName(""),
			args: args{other: ticket.NewName("JIRA-1234")},
			want: false,
		},
		{
			name: "both empty",
			n:    ticket.NewName(""),
			args: args{other: ticket.NewName("")},
			want: true,
		},
		{
			name: "with spaces only",
			n:    ticket.NewName("   "),
			args: args{other: ticket.NewName("   ")},
			want: true,
		},
		{
			name: "with spaces on both side of string",
			n:    ticket.NewName("  JIRA-1234  "),
			args: args{other: ticket.NewName("  JIRA-1234  ")},
			want: true,
		},
		{
			name: "with spaces on both side of string - different",
			n:    ticket.NewName("  JIRA-1234  "),
			args: args{other: ticket.NewName("JIRA-1234")},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := tt.n.Equals(tt.args.other)
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}
