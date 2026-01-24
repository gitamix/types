package commit

import "bytes"

// Message represents a git commit message.
type Message struct {
	// subject is the subject line of the commit message.
	subject Subject
	// body is the body of the commit message.
	body Body
}

// NewMessage creates a new Message instance.
//
//   - subject: the subject of the commit message.
//   - body: the body of the commit message.
func NewMessage(
	subject Subject,
	body Body,
) Message {
	return Message{
		subject: subject,
		body:    body,
	}
}

// ParseMessage parses a commit message from a byte slice or string.
//
// It splits the input into subject and body at the first newline character.
// If there is no newline, the entire input is treated as the subject, and the body is empty.
func ParseMessage[T []byte | string](v T) Message {
	var bb, subjbb, bodybb []byte
	var subji, bodyi int
	switch val := any(v).(type) {
	case string:
		bb = []byte(val)
	case []byte:
		bb = val
	}
	subji = max(0, bytes.IndexByte(bb, '\n'))
	if subji > 0 {
		subjbb = bb[:subji]
		bodyi = subji + 1
	} else if subji == 0 {
		subjbb = bb
	}
	if bodyi > 0 {
		lni := bytes.IndexByte(bb[bodyi:], '\n')
		if lni != -1 {
			bodyi += lni + 1
		}
		bodybb = bb[bodyi:]
	}
	return NewMessage(
		ParseSubject(subjbb),
		NewBody(bodybb),
	)
}

// String returns the string representation of the Message.
//
// If the body is empty, it returns only the subject.
// Otherwise, it concatenates the subject and body with two newlines in between.
func (m Message) String() string {
	if m.body.Empty() {
		return m.subject.String()
	}
	return m.subject.String() + "\n\n" + m.body.String()
}

// Subject returns the subject of the commit message.
func (m Message) Subject() Subject {
	return m.subject
}

// Body returns the body of the commit message.
func (m Message) Body() Body {
	return m.body
}
