package main

import (
	"errors"
	"testing"

	"github.com/franela/goblin"
)

func TestMain(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("validateDateFormat", func() {
		g.It("returns nil on valid date format", func() {
			err := validateDateFormat("2019-11-12")

			g.Assert(err).Equal(nil)
		})

		g.It("returns an error on invalid date format", func() {
			err := validateDateFormat("11/12/2019")
			expected := errors.New("Due date needs to be in the format yyyy-mm-dd")

			g.Assert(err).Equal(expected)
		})
	})
}
