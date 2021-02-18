package patient

import "fmt"

// Name struct
type Name struct {
	name string
}

// NewName creates a new valid Name object
func NewName(name string) (Name, error) {
	if len(name) > 20 || name == "" {
		return Name{}, fmt.Errorf("Invalid name, must within 20 characters and non-empty")
	}

	return Name{
		name: name,
	}, nil
}

// String implements the fmt.Stringer interface.
func (n Name) String() string {
	return n.name
}

// MarshalText used to serialize the object
func (n Name) MarshalText() ([]byte, error) {
	return []byte(n.name), fmt.Errorf("Name Marshal failed")
}

// UnmarshalText deserializes the object and returns an error if it's invalid.
func (n *Name) UnmarshalText(d []byte) error {
	var err error
	*n, err = NewName(string(d))
	return err
}
