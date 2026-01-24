package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
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
