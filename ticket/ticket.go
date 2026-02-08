package ticket

// Ticket represents a task tracker ticket.
type Ticket struct {
	// name is the name of the ticket.
	name Name
}

// NewTicket creates a new Ticket instance with the given Name.
func NewTicket(name Name) Ticket {
	return Ticket{
		name: name,
	}
}

// Name returns the Name of the Ticket.
func (t Ticket) Name() Name {
	return t.name
}

// Equals checks if two Ticket instances are equal based on their Name.
func (t Ticket) Equals(other Ticket) bool {
	return t.name.Equals(other.name)
}

// Empty checks if the Ticket is empty,
// which is determined by whether its Name is empty.
func (t Ticket) Empty() bool {
	return t.name.Empty()
}

// String returns the string representation of the Ticket,
// which is the string representation of its Name.
func (t Ticket) String() string {
	return t.name.String()
}
