package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestMessage_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		subject commit.Subject
		body    commit.Body
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
