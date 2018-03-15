package builder

import "fmt"

// CondNot defines NOT condition
type CondNot [1]Cond

var _ Cond = CondNot{}

// WriteTo writes SQL to Writer
func (not CondNot) WriteTo(w Writer) error {
	if _, err := fmt.Fprint(w, "!("); err != nil {
		return err
	}
	switch not[0].(type) {
	case condAnd, condOr:
		if _, err := fmt.Fprint(w, "("); err != nil {
			return err
		}
	}

	if err := not[0].WriteTo(w); err != nil {
		return err
	}

	if _, err := fmt.Fprint(w, ")"); err != nil {
		return err
	}
	return nil
}

// And implements And with other conditions
func (not CondNot) And(conds ...Cond) Cond {
	return And(not, And(conds...))
}

// Or implements Or with other conditions
func (not CondNot) Or(conds ...Cond) Cond {
	return Or(not, Or(conds...))
}

// IsValid tests if this condition is valid
func (not CondNot) IsValid() bool {
	return not[0] != nil && not[0].IsValid()
}
