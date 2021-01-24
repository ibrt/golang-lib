package stringz

// Predicate describes a predicate.
type Predicate func(s string) bool

// Not negates a predicate.
func Not(p Predicate) Predicate {
	return func(s string) bool {
		return !p(s)
	}
}

// And returns a predicate that is true if all given predicates are true.
func And(ps ...Predicate) Predicate {
	return func(s string) bool {
		for _, p := range ps {
			if !p(s) {
				return false
			}
		}
		return true
	}
}

// Or returns a predicate that is true if any of the given predicates are true.
func Or(ps ...Predicate) Predicate {
	return func(s string) bool {
		for _, p := range ps {
			if p(s) {
				return true
			}
		}
		return false
	}
}

// Equal returns a predicate matching "eq".
func Equal(eq string) Predicate {
	return func(s string) bool {
		return eq == s
	}
}

// Filter returns a shallow of the slice containing only the elements for which the predicate is true.
func Filter(s []string, p Predicate) []string {
	c := make([]string, 0, len(s))
	for _, v := range s {
		if p(v) {
			c = append(c, v)
		}
	}
	return c
}
