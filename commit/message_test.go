package commit_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
	"github.com/gitamix/types/ticket"
)

func TestMessage_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		subject commit.Subject
		body    commit.Body
		raw     []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full message",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				body: commit.NewBody([]byte("This commit adds a new button to the UI.")),
				raw: []byte(
					"feat(ui): add new button\n\nThis commit adds a new button to the UI.",
				),
			},
			want: "feat(ui): add new button\n\nThis commit adds a new button to the UI.",
		},
		{
			name: "subject only",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				body: commit.NewBody([]byte("")),
			},
			want: "fix(backend): resolve issue with API",
		},
		{
			name: "empty body with spaces",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("docs"),
					commit.NewScope(""),
					commit.NewDescription("update README"),
				),
				body: commit.NewBody([]byte("   ")),
			},
			want: "docs: update README",
		},
		{
			name: "empty subject and body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				body: commit.NewBody([]byte("")),
			},
			want: "",
		},
		{
			name: "body with new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("chore"),
					commit.NewScope("ci"),
					commit.NewDescription("update pipeline config"),
				),
				body: commit.NewBody([]byte("Updated the CI pipeline to include new tests.\nPlease review the changes.")),
			},
			want: "chore(ci): update pipeline config\n\n" +
				"Updated the CI pipeline to include new tests.\n" +
				"Please review the changes.",
		},
		{
			name: "subject with spaces only and body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("   "),
					commit.NewScope("   "),
					commit.NewDescription("   "),
				),
				body: commit.NewBody([]byte("Body content here.")),
			},
			want: "   \n\nBody content here.",
		},
		{
			name: "subject with new lines and empty body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("\n"),
					commit.NewScope("\n"),
					commit.NewDescription("\n"),
				),
				body: commit.NewBody([]byte("")),
			},
			want: "\n",
		},
		{
			name: "body with only new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("style"),
					commit.NewScope("formatting"),
					commit.NewDescription("improve code formatting"),
				),
				body: commit.NewBody([]byte("\n\n\n")),
			},
			want: "style(formatting): improve code formatting",
		},
		{
			name: "both subject and body with new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("\n"),
					commit.NewScope("\n"),
					commit.NewDescription("\n"),
				),
				body: commit.NewBody([]byte("\n\n")),
			},
			want: "\n",
		},
		{
			name: "complex body content",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("api"),
					commit.NewDescription("add new endpoint"),
				),
				body: commit.NewBody([]byte("This commit introduces a new API endpoint.\n\n" +
					"Changes include:\n" +
					"- Added GET /v2/resource\n" +
					"- Updated documentation\n\n" +
					"Please refer to the API docs for more details.")),
			},
			want: "feat(api): add new endpoint\n\n" +
				"This commit introduces a new API endpoint.\n\n" +
				"Changes include:\n" +
				"- Added GET /v2/resource\n" +
				"- Updated documentation\n\n" +
				"Please refer to the API docs for more details.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := commit.NewMessage(
				tt.fields.subject,
				tt.fields.body,
				commit.WithRaw(tt.fields.raw),
			)
			got := m.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMessage_Subject(t *testing.T) {
	t.Parallel()
	type fields struct {
		subject commit.Subject
		body    commit.Body
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Subject
	}{
		{
			name: "ok subject",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				body: commit.NewBody([]byte("This commit adds a new button to the UI.")),
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "subject with empty body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				body: commit.NewBody([]byte("")),
			},
			want: commit.NewSubject(
				commit.NewType("fix"),
				commit.NewScope("backend"),
				commit.NewDescription("resolve issue with API"),
			),
		},
		{
			name: "subject with spaces only",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("   "),
					commit.NewScope(""),
					commit.NewDescription("   "),
				),
				body: commit.NewBody([]byte("   ")),
			},
			want: commit.NewSubject(
				commit.NewType("   "),
				commit.NewScope(""),
				commit.NewDescription("   "),
			),
		},
		{
			name: "empty subject and body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				body: commit.NewBody([]byte("")),
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
			m := commit.NewMessage(
				tt.fields.subject,
				tt.fields.body,
			)
			got := m.Subject()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMessage_Body(t *testing.T) {
	t.Parallel()
	type fields struct {
		subject commit.Subject
		body    commit.Body
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Body
	}{
		{
			name: "ok message",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				body: commit.NewBody([]byte("This commit adds a new button to the UI.")),
			},
			want: commit.NewBody([]byte("This commit adds a new button to the UI.")),
		},
		{
			name: "empty body",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				body: commit.NewBody([]byte("")),
			},
			want: commit.NewBody([]byte("")),
		},
		{
			name: "body with spaces only",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("docs"),
					commit.NewScope(""),
					commit.NewDescription("update README"),
				),
				body: commit.NewBody([]byte("   ")),
			},
			want: commit.NewBody([]byte("   ")),
		},
		{
			name: "body with new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("chore"),
					commit.NewScope("ci"),
					commit.NewDescription("update pipeline config"),
				),
				body: commit.NewBody([]byte("Updated the CI pipeline to include new tests.\nPlease review the changes.")),
			},
			want: commit.NewBody([]byte("Updated the CI pipeline to include new tests.\nPlease review the changes.")),
		},
		{
			name: "body with only new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("style"),
					commit.NewScope("formatting"),
					commit.NewDescription("improve code formatting"),
				),
				body: commit.NewBody([]byte("\n\n\n")),
			},
			want: commit.NewBody([]byte("\n\n\n")),
		},
		{
			name: "both subject and body with new lines",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("\n"),
					commit.NewScope("\n"),
					commit.NewDescription("\n"),
				),
				body: commit.NewBody([]byte("\n\n")),
			},
			want: commit.NewBody([]byte("\n\n")),
		},
		{
			name: "complex body content",
			fields: fields{
				subject: commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("api"),
					commit.NewDescription("add new endpoint"),
				),
				body: commit.NewBody([]byte("This commit introduces a new API endpoint.\n\n" +
					"Changes include:\n" +
					"- Added GET /v2/resource\n" +
					"- Updated documentation\n\n" +
					"Please refer to the API docs for more details.")),
			},
			want: commit.NewBody([]byte("This commit introduces a new API endpoint.\n\n" +
				"Changes include:\n" +
				"- Added GET /v2/resource\n" +
				"- Updated documentation\n\n" +
				"Please refer to the API docs for more details.")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			m := commit.NewMessage(
				tt.fields.subject,
				tt.fields.body,
			)
			got := m.Body()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseMessage(t *testing.T) {
	t.Parallel()
	type args struct {
		bb []byte
	}
	tests := []struct {
		name string
		args args
		want commit.Message
	}{
		{
			name: "fully correct message",
			args: args{
				bb: []byte(
					"feat(ui): add new button\n\n" +
						"This commit adds a new button to the UI.",
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte("This commit adds a new button to the UI.")),
				commit.WithRaw([]byte(
					"feat(ui): add new button\n\n"+
						"This commit adds a new button to the UI.",
				)),
			),
		},
		{
			name: "fully correct message with extra new lines",
			args: args{
				bb: []byte(
					"feat(ui): add new button\n\n" +
						"This commit adds a new button to the UI.\n\n" +
						"Thank you!",
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte("This commit adds a new button to the UI.\n\nThank you!")),
				commit.WithRaw([]byte(
					"feat(ui): add new button\n\n"+
						"This commit adds a new button to the UI.\n\n"+
						"Thank you!",
				)),
			),
		},
		{
			name: "subject and body separated by single newline",
			args: args{
				bb: []byte(
					"feat(ui): add new button\n" +
						"This commit adds a new button to the UI.",
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte("This commit adds a new button to the UI.")),
				commit.WithRaw([]byte(
					"feat(ui): add new button\n"+
						"This commit adds a new button to the UI.",
				)),
			),
		},
		{
			name: "subject and multi-line body separated by single newline",
			args: args{
				bb: []byte(
					"feat(ui): add new button\n" +
						"This commit adds a new button to the UI.\n" +
						"Fixes #123",
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte(
					"This commit adds a new button to the UI.\n"+
						"Fixes #123",
				)),
				commit.WithRaw([]byte(
					"feat(ui): add new button\n"+
						"This commit adds a new button to the UI.\n"+
						"Fixes #123",
				)),
			),
		},
		{
			name: "message with inline body",
			args: args{
				bb: []byte("feat(ui): add new button. This commit adds a new button to the UI."),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button. This commit adds a new button to the UI."),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("feat(ui): add new button. This commit adds a new button to the UI."),
				),
			),
		},
		{
			name: "correct subject with spaces",
			args: args{
				bb: []byte("   feat  (ui) :    add new button  "),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("   feat  (ui) :    add new button  "),
				),
			),
		},
		{
			name: "subject with type and description only",
			args: args{
				bb: []byte("feat: add new button"),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope(""),
					commit.NewDescription("add new button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("feat: add new button"),
				),
			),
		},
		{
			name: "subject with scope and description only",
			args: args{
				bb: []byte("(ui): add new button"),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("(ui): add new button"),
				),
			),
		},
		{
			name: "description only",
			args: args{
				bb: []byte("add new button"),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription("add new button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("add new button"),
				),
			),
		},
		{
			name: "description with body and extra new lines",
			args: args{
				bb: []byte(
					"add new button" + "\n\n" +
						"This commit adds a new button to the UI.\n" +
						"Fixes #123\n\n" +
						"Please review.",
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte(
					"This commit adds a new button to the UI.\n"+
						"Fixes #123\n\n"+
						"Please review.",
				)),
				commit.WithRaw(
					[]byte(
						"add new button"+"\n\n"+
							"This commit adds a new button to the UI.\n"+
							"Fixes #123\n\n"+
							"Please review.",
					),
				),
			),
		},
		{
			name: "wrong subject with type only",
			args: args{
				bb: []byte("feat: "),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("feat: "),
				),
			),
		},
		{
			name: "wrong subject with scope only",
			args: args{
				bb: []byte("(ui): "),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope("ui"),
					commit.NewDescription(""),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("(ui): "),
				),
			),
		},
		{
			name: "wrong subject with feat and scope only",
			args: args{
				bb: []byte("feat(ui): "),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription(""),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("feat(ui): "),
				),
			),
		},
		{
			name: "wrong subject with invalid scope format",
			args: args{
				bb: []byte("feat: ui: add new button"),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope(""),
					commit.NewDescription("ui: add new button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("feat: ui: add new button"),
				),
			),
		},
		{
			name: "empty message",
			args: args{
				bb: []byte(""),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte(""),
				),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" from bytes", func(t *testing.T) {
			t.Parallel()
			got := commit.ParseMessage(tt.args.bb)
			assert.Equal(t, tt.want, got)
		})
		t.Run(tt.name+" from string", func(t *testing.T) {
			t.Parallel()
			got := commit.ParseMessage(string(tt.args.bb))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMessage_Ticket(t *testing.T) {
	t.Parallel()
	type args struct {
		re *regexp.Regexp
	}
	tests := []struct {
		name      string
		m         commit.Message
		args      args
		want      ticket.Ticket
		wantPanic bool
	}{
		{
			name: "ticket from raw msg",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				commit.NewBody([]byte("Closes BUG-4331")),
				commit.WithRaw(
					[]byte(
						"TASK-1234 fix(backend): resolve issue with API\n\nCloses BUG-4331",
					),
				),
			),
			args: args{
				re: regexp.MustCompile("((TASK|PROJ|BUG)-[0-9]+)"),
			},
			want: ticket.NewTicket(
				ticket.NewName("TASK-1234"),
			),
			wantPanic: false,
		},
		{
			name: "ticket from single-line raw msg with no newline",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("ui"),
					commit.NewDescription("add button"),
				),
				commit.NewBody(nil),
				commit.WithRaw(
					[]byte("TASK-1234 fix(ui): add button"),
				),
			),
			args: args{
				re: regexp.MustCompile("((TASK|PROJ|BUG)-[0-9]+)"),
			},
			want: ticket.NewTicket(
				ticket.NewName("TASK-1234"),
			),
			wantPanic: false,
		},
		{
			name: "no ticket in subject but in body",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				commit.NewBody([]byte("Closes BUG-4331")),
			),
			args: args{
				re: regexp.MustCompile("((TASK|PROJ|BUG)-[0-9]+)"),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "no ticket in subject and body",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				commit.NewBody([]byte("No tickets mentioned.")),
			),
			args: args{
				re: regexp.MustCompile("((TASK|PROJ|BUG)-[0-9]+)"),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "empty subject and body",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody([]byte("")),
			),
			args: args{
				re: regexp.MustCompile("((TASK|PROJ|BUG)-[0-9]+)"),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "nil regex",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("fix"),
					commit.NewScope("backend"),
					commit.NewDescription("resolve issue with API"),
				),
				commit.NewBody([]byte("Closes BUG-4331")),
				commit.WithRaw(
					[]byte(
						"TASK-1234 fix(backend): resolve issue with API\n\nCloses BUG-4331",
					),
				),
			),
			args: args{
				re: nil,
			},
			want:      ticket.Ticket{},
			wantPanic: true,
		},
		{
			name: "default args",
			m: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody([]byte("")),
			),
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
					_ = tt.m.Ticket(tt.args.re)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				tt.m.Ticket(tt.args.re),
			)
		})
	}
}
