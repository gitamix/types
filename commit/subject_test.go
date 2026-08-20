package commit_test

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
	"github.com/gitamix/types/ticket"
)

func TestParseSubject(t *testing.T) {
	t.Parallel()
	t.Run("fully correct message from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject(
			"feat(ui): add new button\n\n" +
				"This commit adds a new button to the UI.",
		)
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("fully correct message from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte(
			"feat(ui): add new button\n\n" +
				"This commit adds a new button to the UI.",
		))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("message with inline body from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("feat(ui): add new button. This commit adds a new button to the UI.")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button. This commit adds a new button to the UI."),
		)
		assert.Equal(t, want, got)
	})
	t.Run("message with inline body from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("feat(ui): add new button. This commit adds a new button to the UI."))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button. This commit adds a new button to the UI."),
		)
		assert.Equal(t, want, got)
	})
	t.Run("fully correct message from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte(
			"feat(ui): add new button\n\n" +
				"This commit adds a new button to the UI.",
		))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("correct subject with spaces from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("   feat  (ui) :    add new button  ")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("correct subject with spaces from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("   feat  (ui) :    add new button  "))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("subject with type and description only from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("feat: add new button")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("subject with type and description only from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("feat: add new button"))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("description only from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("add new button")
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope(""),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("description only from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("add new button"))
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope(""),
			commit.NewDescription("add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with type only from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("feat: ")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with type only from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("feat: "))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with scope only from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("(ui): ")
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope("ui"),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with scope only from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("(ui): "))
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope("ui"),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with feat and scope only from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("feat(ui): ")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with feat and scope only from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("feat(ui): "))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with invalid scope format from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("feat: ui: add new button")
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription("ui: add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("wrong subject with invalid scope format from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte("feat: ui: add new button"))
		want := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope(""),
			commit.NewDescription("ui: add new button"),
		)
		assert.Equal(t, want, got)
	})
	t.Run("empty subject from string", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject("")
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope(""),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
	t.Run("empty subject from bytes", func(t *testing.T) {
		t.Parallel()
		got := commit.ParseSubject([]byte(""))
		want := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope(""),
			commit.NewDescription(""),
		)
		assert.Equal(t, want, got)
	})
}

func TestSubject_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		typ         commit.Type
		scope       commit.Scope
		description commit.Description
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "fully correct subject",
			fields: fields{
				typ:         commit.NewType("feat"),
				scope:       commit.NewScope("ui"),
				description: commit.NewDescription("add new button"),
			},
			want: "feat(ui): add new button",
		},
		{
			name: "correct subject with type and description only",
			fields: fields{
				typ:         commit.NewType("feat"),
				scope:       commit.NewScope(""),
				description: commit.NewDescription("add new button"),
			},
			want: "feat: add new button",
		},
		{
			name: "description only",
			fields: fields{
				typ:         commit.NewType(""),
				scope:       commit.NewScope(""),
				description: commit.NewDescription("add new button"),
			},
			want: "add new button",
		},
		{
			name: "empty subject",
			fields: fields{
				typ:         commit.NewType(""),
				scope:       commit.NewScope(""),
				description: commit.NewDescription(""),
			},
			want: "",
		},
		{
			name: "subject with spaces",
			fields: fields{
				typ:         commit.NewType("  feat  "),
				scope:       commit.NewScope("  ui  "),
				description: commit.NewDescription("   add new button   "),
			},
			want: "  feat  (  ui  ):    add new button   ",
		},
		{
			name: "subject with empty scope",
			fields: fields{
				typ:         commit.NewType("fix"),
				scope:       commit.NewScope(""),
				description: commit.NewDescription("resolve button issue"),
			},
			want: "fix: resolve button issue",
		},
		{
			name: "subject with empty type",
			fields: fields{
				typ:         commit.NewType(""),
				scope:       commit.NewScope("core"),
				description: commit.NewDescription("update core module"),
			},
			want: "update core module",
		},
		{
			name: "subject with empty description",
			fields: fields{
				typ:         commit.NewType("chore"),
				scope:       commit.NewScope("deps"),
				description: commit.NewDescription(""),
			},
			want: "chore(deps): ",
		},
		{
			name: "subject with only type",
			fields: fields{
				typ:         commit.NewType("refactor"),
				scope:       commit.NewScope(""),
				description: commit.NewDescription(""),
			},
			want: "refactor: ",
		},
		{
			name: "subject with only scope",
			fields: fields{
				typ:         commit.NewType(""),
				scope:       commit.NewScope("api"),
				description: commit.NewDescription(""),
			},
			want: "",
		},
		{
			name: "subject with spaces only",
			fields: fields{
				typ:         commit.NewType("   "),
				scope:       commit.NewScope("   "),
				description: commit.NewDescription("   "),
			},
			want: "   ",
		},
		{
			name: "subject with new lines",
			fields: fields{
				typ:         commit.NewType("\n"),
				scope:       commit.NewScope("\n"),
				description: commit.NewDescription("\n"),
			},
			want: "\n",
		},
		{
			name: "subject with special characters",
			fields: fields{
				typ:         commit.NewType("feat!"),
				scope:       commit.NewScope("ui$"),
				description: commit.NewDescription("add new button@"),
			},
			want: "feat!(ui$): add new button@",
		},
		{
			name: "subject with unicode characters",
			fields: fields{
				typ:         commit.NewType("修正"),
				scope:       commit.NewScope("ユーザーインターフェース"),
				description: commit.NewDescription("新しいボタンを追加"),
			},
			want: "修正(ユーザーインターフェース): 新しいボタンを追加",
		},
		{
			name: "subject with mixed content",
			fields: fields{
				typ:         commit.NewType("feat123"),
				scope:       commit.NewScope("module_xyz"),
				description: commit.NewDescription("implement feature #42"),
			},
			want: "feat123(module_xyz): implement feature #42",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			s := commit.NewSubject(
				tt.fields.typ,
				tt.fields.scope,
				tt.fields.description,
			)
			got := s.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSubject_Type(t *testing.T) {
	t.Parallel()

	t.Run("full subject", func(t *testing.T) {
		t.Parallel()
		want := commit.NewType("feat")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					want,
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				).
				Type(),
		)
	})

	t.Run("type and description only", func(t *testing.T) {
		t.Parallel()
		want := commit.NewType("fix")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					want,
					commit.NewScope(""),
					commit.NewDescription("resolve issue"),
				).
				Type(),
		)
	})

	t.Run("empty type", func(t *testing.T) {
		t.Parallel()
		want := commit.NewType("")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					want,
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				).
				Type(),
		)
	})

	t.Run("empty subject", func(t *testing.T) {
		t.Parallel()
		want := commit.NewType("")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					want,
					commit.NewScope(""),
					commit.NewDescription(""),
				).
				Type(),
		)
	})

	t.Run("type with spaces", func(t *testing.T) {
		t.Parallel()
		want := commit.NewType("  feat  ")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					want,
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				).
				Type(),
		)
	})
}

func TestSubject_Scope(t *testing.T) {
	t.Parallel()

	t.Run("full subject", func(t *testing.T) {
		t.Parallel()
		want := commit.NewScope("ui")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType("feat"),
					want,
					commit.NewDescription("add new button"),
				).
				Scope(),
		)
	})

	t.Run("scope only", func(t *testing.T) {
		t.Parallel()
		want := commit.NewScope("core")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType(""),
					want,
					commit.NewDescription(""),
				).
				Scope(),
		)
	})

	t.Run("empty scope", func(t *testing.T) {
		t.Parallel()
		want := commit.NewScope("")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType("feat"),
					want,
					commit.NewDescription("add new button"),
				).
				Scope(),
		)
	})

	t.Run("empty subject", func(t *testing.T) {
		t.Parallel()
		want := commit.NewScope("")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType(""),
					want,
					commit.NewDescription(""),
				).
				Scope(),
		)
	})

	t.Run("scope with spaces", func(t *testing.T) {
		t.Parallel()
		want := commit.NewScope("  ui  ")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType("feat"),
					want,
					commit.NewDescription("add new button"),
				).
				Scope(),
		)
	})
}

func TestSubject_Description(t *testing.T) {
	t.Parallel()

	t.Run("full subject", func(t *testing.T) {
		t.Parallel()
		s := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription("add new button"),
		)
		want := commit.NewDescription("add new button")
		assert.Equal(t, want, s.Description())
	})

	t.Run("description only", func(t *testing.T) {
		t.Parallel()
		s := commit.NewSubject(
			commit.NewType(""),
			commit.NewScope(""),
			commit.NewDescription("add new button"),
		)
		want := commit.NewDescription("add new button")
		assert.Equal(t, want, s.Description())
	})

	t.Run("empty description", func(t *testing.T) {
		t.Parallel()
		s := commit.NewSubject(
			commit.NewType("feat"),
			commit.NewScope("ui"),
			commit.NewDescription(""),
		)
		want := commit.NewDescription("")
		assert.Equal(t, want, s.Description())
	})

	t.Run("empty subject", func(t *testing.T) {
		t.Parallel()
		want := commit.NewDescription("")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					want,
				).
				Description(),
		)
	})

	t.Run("description with spaces", func(t *testing.T) {
		t.Parallel()
		want := commit.NewDescription("  add new button  ")
		assert.Equal(
			t,
			want,
			commit.
				NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					want,
				).
				Description(),
		)
	})
}

func TestSubject_Ticket(t *testing.T) {
	t.Parallel()
	type args struct {
		re *regexp.Regexp
	}
	tests := []struct {
		name      string
		s         commit.Subject
		args      args
		want      ticket.Ticket
		wantPanic bool
	}{
		{
			name: "in the beginning",
			s: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("TASK-1234 add new feature"),
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
			name: "in the beginning with scopes",
			s: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("[TASK-1234] add new feature"),
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
			name: "issue in the end",
			s: commit.NewSubject(
				commit.NewType("fix"),
				commit.NewScope("core"),
				commit.NewDescription("Fix #42"),
			),
			args: args{
				re: regexp.MustCompile("(#[0-9]+)"),
			},
			want: ticket.NewTicket(
				ticket.NewName("#42"),
			),
			wantPanic: false,
		},
		{
			name: "issue in the middle",
			s: commit.NewSubject(
				commit.NewType("fix"),
				commit.NewScope("core"),
				commit.NewDescription("Fix issue #42 in the code"),
			),
			args: args{
				re: regexp.MustCompile("(#[0-9]+)"),
			},
			want: ticket.NewTicket(
				ticket.NewName("#42"),
			),
			wantPanic: false,
		},
		{
			name: "no ticket",
			s: commit.NewSubject(
				commit.NewType("fix"),
				commit.NewScope("core"),
				commit.NewDescription("Fix issue in the code"),
			),
			args: args{
				re: regexp.MustCompile("(#[0-9]+)"),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "empty subject",
			s: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope(""),
				commit.NewDescription(""),
			),
			args: args{
				re: regexp.MustCompile("(#[0-9]+)"),
			},
			want:      ticket.Ticket{},
			wantPanic: false,
		},
		{
			name: "nil regex",
			s: commit.NewSubject(
				commit.NewType("fix"),
				commit.NewScope("core"),
				commit.NewDescription("Fix issue #42 in the code"),
			),
			args: args{
				re: nil,
			},
			want:      ticket.Ticket{},
			wantPanic: true,
		},
		{
			name: "default args",
			s: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope(""),
				commit.NewDescription(""),
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
					_ = tt.s.Ticket(tt.args.re)
				})
				return
			}
			assert.Equal(
				t,
				tt.want,
				tt.s.Ticket(tt.args.re),
			)
		})
	}
}
