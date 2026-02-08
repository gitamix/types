package ticket_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/ticket"
)

func TestTicket_Name(t *testing.T) {
	t.Parallel()
	type fields struct {
		name ticket.Name
	}
	tests := []struct {
		name   string
		fields fields
		want   ticket.Name
	}{
		{
			name: "ticket with name",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			want: ticket.NewName("JIRA-1234"),
		},
		{
			name: "ticket with empty name",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: ticket.NewName(""),
		},
		{
			name: "ticket with simple name",
			fields: fields{
				name: ticket.NewName("Fix login issue"),
			},
			want: ticket.NewName("Fix login issue"),
		},
		{
			name: "empty ticket",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: ticket.NewName(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tkt := ticket.NewTicket(tt.fields.name)
			got := tkt.Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTicket_Equals(t *testing.T) {
	t.Parallel()
	type fields struct {
		name ticket.Name
	}
	type args struct {
		other ticket.Ticket
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "same names",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			args: args{
				other: ticket.NewTicket(ticket.NewName("JIRA-1234")),
			},
			want: true,
		},
		{
			name: "different names",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			args: args{
				other: ticket.NewTicket(ticket.NewName("JIRA-5678")),
			},
			want: false,
		},
		{
			name: "same names but one with spaces around",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			args: args{
				other: ticket.NewTicket(ticket.NewName(" JIRA-1234 ")),
			},
			want: false,
		},
		{
			name: "empty tickets",
			fields: fields{
				name: ticket.NewName(""),
			},
			args: args{
				other: ticket.NewTicket(ticket.NewName("")),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tkt := ticket.NewTicket(tt.fields.name)
			got := tkt.Equals(tt.args.other)
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestTicket_Empty(t *testing.T) {
	t.Parallel()
	type fields struct {
		name ticket.Name
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "same names",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			want: false,
		},
		{
			name: "different names",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			want: false,
		},
		{
			name: "same names but one with spaces around",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			want: false,
		},
		{
			name: "name with spaces only",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: true,
		},
		{
			name: "empty tickets",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tkt := ticket.NewTicket(tt.fields.name)
			got := tkt.Empty()
			if tt.want {
				assert.True(t, got)
			} else {
				assert.False(t, got)
			}
		})
	}
}

func TestTicket_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		name ticket.Name
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "jira ticket name",
			fields: fields{
				name: ticket.NewName("JIRA-1234"),
			},
			want: "JIRA-1234",
		},
		{
			name: "empty ticket name",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: "",
		},
		{
			name: "simple ticket name",
			fields: fields{
				name: ticket.NewName("Fix login issue"),
			},
			want: "Fix login issue",
		},
		{
			name: "with spaces only",
			fields: fields{
				name: ticket.NewName("   "),
			},
			want: "   ",
		},
		{
			name: "with spaces on both side of string",
			fields: fields{
				name: ticket.NewName("  JIRA-1234  "),
			},
			want: "  JIRA-1234  ",
		},
		{
			name: "empty ticket name",
			fields: fields{
				name: ticket.NewName(""),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			tkt := ticket.NewTicket(tt.fields.name)
			got := tkt.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
