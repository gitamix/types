package commit

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
