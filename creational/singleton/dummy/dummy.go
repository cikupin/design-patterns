package dummy

// IDummy defines interfce for dummy
type IDummy interface {
	SetDescription(desc string)
	GetDescription() string
}

// Dummy defines dummy struct
type Dummy struct {
	description string
}

// SetDescription will set dummy's description
func (d *Dummy) SetDescription(desc string) {
	d.description = desc
}

// GetDescription will retrieve dummy's description
func (d *Dummy) GetDescription() string {
	return d.description
}
