// Package suites allows callers to look up Kyber suites by name.
//
// Currently, only the "ed25519" suite is available with a constant
// time implementation and the other ones use variable time algorithms.
package suites

import (
	"errors"
	"strings"

	"github.com/dedis/kyber"
)

// Suite is the sum of all suites mix-ins in Kyber.
type Suite interface {
	kyber.Encoding
	kyber.Group
	kyber.HashFactory
	kyber.XOFFactory
	kyber.Random
}

var suites = map[string]Suite{}

// register is called by suites to make themselves known to Kyber.
//
func register(s Suite) {
	suites[strings.ToLower(s.String())] = s
}

// ErrUnknownSuite indicates that the suite was not one of the
// registered suites.
var ErrUnknownSuite = errors.New("unknown suite")

// Find looks up a suite by name.
func Find(name string) (Suite, error) {
	if s, ok := suites[strings.ToLower(name)]; ok {
		return s, nil
	}
	return nil, ErrUnknownSuite
}

// MustFind looks up a suite by name and panics if it is not found.
func MustFind(name string) Suite {
	s, err := Find(name)
	if err != nil {
		panic("Suite " + name + " not found.")
	}
	return s
}
