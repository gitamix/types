package commit_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gitamix/types/commit"
)

func TestCommit_String(t *testing.T) {
	t.Parallel()
	type fields struct {
		msg  commit.Message
		hash commit.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "full commit",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: "abcdef1",
		},
		{
			name: "empty hash",
			fields: fields{
				hash: commit.NewHash(""),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: "",
		},
		{
			name: "empty msg",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType(""),
						commit.NewScope(""),
						commit.NewDescription(""),
					),
					commit.NewBody([]byte{}),
				),
			},
			want: "abcdef1",
		},
		{
			name: "hash lt 7 chars",
			fields: fields{
				hash: commit.NewHash("abc"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: "abc",
		},
		{
			name: "default fields",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := commit.NewCommit(
				tt.fields.hash,
				tt.fields.msg,
			)
			got := c.String()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCommit_Message(t *testing.T) {
	t.Parallel()
	type fields struct {
		msg  commit.Message
		hash commit.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Message
	}{
		{
			name: "full commit",
			fields: fields{
				hash: commit.NewHash("abcdef1"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType("feat"),
					commit.NewScope("ui"),
					commit.NewDescription("add new button"),
				),
				commit.NewBody([]byte("Added new button and covered with tests")),
			),
		},
		{
			name: "empty msg",
			fields: fields{
				hash: commit.NewHash("abcdef1"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType(""),
						commit.NewScope(""),
						commit.NewDescription(""),
					),
					commit.NewBody([]byte("")),
				),
			},
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody([]byte("")),
			),
		},
		{
			name: "default fields",
			want: commit.NewMessage(
				commit.NewSubject(
					commit.NewType(""),
					commit.NewScope(""),
					commit.NewDescription(""),
				),
				commit.NewBody(nil),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := commit.NewCommit(
				tt.fields.hash,
				tt.fields.msg,
			)
			got := c.Message()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCommit_Hash(t *testing.T) {
	t.Parallel()
	type fields struct {
		msg  commit.Message
		hash commit.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Hash
	}{
		{
			name: "full commit",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewHash("abcdef1234567890"),
		},
		{
			name: "empty hash",
			fields: fields{
				hash: commit.NewHash(""),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewHash(""),
		},
		{
			name: "empty msg",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType(""),
						commit.NewScope(""),
						commit.NewDescription(""),
					),
					commit.NewBody([]byte{}),
				),
			},
			want: commit.NewHash("abcdef1234567890"),
		},
		{
			name: "default fields",
			want: commit.NewHash(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := commit.NewCommit(
				tt.fields.hash,
				tt.fields.msg,
			)
			got := c.Hash()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCommit_Type(t *testing.T) {
	t.Parallel()
	type fields struct {
		msg  commit.Message
		hash commit.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Type
	}{
		{
			name: "full commit",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewType("feat"),
		},
		{
			name: "empty hash",
			fields: fields{
				hash: commit.NewHash(""),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewType("feat"),
		},
		{
			name: "empty msg",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType(""),
						commit.NewScope(""),
						commit.NewDescription(""),
					),
					commit.NewBody([]byte{}),
				),
			},
			want: commit.NewType(""),
		},
		{
			name: "default fields",
			want: commit.NewType(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := commit.NewCommit(
				tt.fields.hash,
				tt.fields.msg,
			)
			got := c.Type()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCommit_Scope(t *testing.T) {
	t.Parallel()
	type fields struct {
		msg  commit.Message
		hash commit.Hash
	}
	tests := []struct {
		name   string
		fields fields
		want   commit.Scope
	}{
		{
			name: "full commit",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewScope("ui"),
		},
		{
			name: "empty hash",
			fields: fields{
				hash: commit.NewHash(""),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType("feat"),
						commit.NewScope("ui"),
						commit.NewDescription("add new button"),
					),
					commit.NewBody([]byte("Added new button and covered with tests")),
				),
			},
			want: commit.NewScope("ui"),
		},
		{
			name: "empty msg",
			fields: fields{
				hash: commit.NewHash("abcdef1234567890"),
				msg: commit.NewMessage(
					commit.NewSubject(
						commit.NewType(""),
						commit.NewScope(""),
						commit.NewDescription(""),
					),
					commit.NewBody([]byte{}),
				),
			},
			want: commit.NewScope(""),
		},
		{
			name: "default fields",
			want: commit.NewScope(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := commit.NewCommit(
				tt.fields.hash,
				tt.fields.msg,
			)
			got := c.Scope()
			assert.Equal(t, tt.want, got)
		})
	}
}
