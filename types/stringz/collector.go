package stringz

import "strings"

// CollectorOption configures a Collector.
type CollectorOption func(*Collector)

// SkipDuplicates skips adding duplicate strings.
func SkipDuplicates(c *Collector) {
	c.skipDuplicates = true
}

// SkipEmpties skips adding empty strings.
func SkipEmpties(c *Collector) {
	c.skipEmpties = true
}

// InitialCap sets the initial capacity of the underlying data structures.
func InitialCap(cap int) CollectorOption {
	return func(c *Collector) {
		c.initialCap = cap
	}
}

// Collector collects strings.
// It must be initialized using NewCollector, and its behavior can be regulated via options.
type Collector struct {
	skipDuplicates bool
	skipEmpties    bool
	initialCap     int

	m map[string]struct{}
	s []string
}

// NewCollector initializes a new collector.
func NewCollector(options ...CollectorOption) *Collector {
	c := &Collector{}

	for _, option := range options {
		option(c)
	}

	if c.initialCap == 0 {
		c.m = make(map[string]struct{})
		c.s = make([]string, 0)
	} else {
		c.m = make(map[string]struct{}, c.initialCap)
		c.s = make([]string, 0, c.initialCap)
	}

	return c
}

// Add a string.
func (c *Collector) Add(ss ...string) *Collector {
	for _, s := range ss {
		if c.skipEmpties && s == "" {
			return c
		}

		if c.skipDuplicates {
			if _, ok := c.m[s]; ok {
				return c
			}
		}

		c.m[s] = struct{}{}
		c.s = append(c.s, s)
	}

	return c
}

// AddPtr a string pointer, skipping it if nil.
func (c *Collector) AddPtr(ss ...*string) *Collector {
	for _, s := range ss {
		if s == nil {
			return c
		}

		if c.skipEmpties && *s == "" {
			return c
		}

		if c.skipDuplicates {
			if _, ok := c.m[*s]; ok {
				return c
			}
		}

		c.m[*s] = struct{}{}
		c.s = append(c.s, *s)
	}

	return c
}

// Slice returns the collected slice.
func (c *Collector) Slice() []string {
	return c.s
}

// Map returns the collected map.
func (c *Collector) Map() map[string]struct{} {
	return c.m
}

// Join the collected slice using the given "sep".
func (c *Collector) Join(sep string) string {
	return strings.Join(c.s, sep)
}
