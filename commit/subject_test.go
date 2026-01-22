package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestParseSubject(t *testing.T) {
	t.Parallel()
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want commit.Subject
	}{
		{
			name: "fully correct subject",
			args: args{
				s: "feat(ui): add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "correct subject with spaces",
			args: args{
				s: "   feat  (ui) :    add new button  ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "subject with type and description only",
			args: args{
				s: "feat: add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "description only",
			args: args{
				s: "add new button",
			},
			want: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope(""),
				commit.NewDescription("add new button"),
			),
		},
		{
			name: "wrong subject with type only",
			args: args{
				s: "feat: ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with scope only",
			args: args{
				s: "(ui): ",
			},
			want: commit.NewSubject(
				commit.NewType(""),
				commit.NewScope("ui"),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with feat and scope only",
			args: args{
				s: "feat(ui): ",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope("ui"),
				commit.NewDescription(""),
			),
		},
		{
			name: "wrong subject with invalid scope format",
			args: args{
				s: "feat: ui: add new button",
			},
			want: commit.NewSubject(
				commit.NewType("feat"),
				commit.NewScope(""),
				commit.NewDescription("ui: add new button"),
			),
		},
		{
			name: "empty subject",
			args: args{
				s: "",
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
			got := commit.ParseSubject(tt.args.s)
			assert.Equal(t, tt.want, got)
		})
	}
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
