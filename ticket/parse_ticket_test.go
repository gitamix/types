package ticket_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/ticket"
)

func TestParseTicket(t *testing.T) {
	t.Parallel()
	type args struct {
		s  string
		re *regexp.Regexp
	}
	tests := []struct {
		name      string
		args      args
		want      ticket.Ticket
		wantPanic bool
	}{
		{
			name: "ticket in the beginning",
			args: args{
				s:  "TASK-1234 add new feature",
				re: regexp.MustCompile(`^((?:TASK|PROJ|BUG)-\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("TASK-1234"),
			),
			wantPanic: false,
		},
		{
			name: "ticket in the middle",
			args: args{
				s:  "Fix issue PROJ-5678 in the code",
				re: regexp.MustCompile(`((?:TASK|PROJ|BUG)-\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("PROJ-5678"),
			),
			wantPanic: false,
		},
		{
			name: "ticket at the end",
			args: args{
				s:  "Resolve BUG-9012",
				re: regexp.MustCompile(`((?:TASK|PROJ|BUG)-\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("BUG-9012"),
			),
			wantPanic: false,
		},
		{
			name: "github issue",
			args: args{
				s:  "Fix issue #42 in the code",
				re: regexp.MustCompile(`#(\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("42"),
			),
		},
		{
			name: "branch name with task id",
			args: args{
				s:  "feature/TASK-1234",
				re: regexp.MustCompile(`^(?:feature|bugfix|hotfix)/([A-Z]+-\d+)`),
			},
			want: ticket.NewTicket(
				ticket.NewName("TASK-1234"),
			),
		},
		{
			name: "empty string",
			args: args{
				s:  "",
				re: regexp.MustCompile(`((?:TASK|PROJ|BUG)-\d+)`),
			},
			want: ticket.Ticket{},
		},
		{
			name: "nil regex",
			args: args{
				s:  "Fix issue TASK-1234",
				re: nil,
			},
			wantPanic: true,
		},
		{
			name:      "default args",
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
					_ = ticket.ParseTicket(
						tt.args.s,
						tt.args.re,
					)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				ticket.ParseTicket(
					tt.args.s,
					tt.args.re,
				),
			)
		})
	}
}
