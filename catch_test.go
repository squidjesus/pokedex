package main

import (
	"fmt"
	"testing"
)

func TestCatchAttempt(t *testing.T) {

	cases := []struct {
		name   string
		baseXp int
	}{
		{
			name:   "pikachu",
			baseXp: 112,
		},
	}

	for _, c := range cases {
		tries := 1
		for {
			if catchAttempt(c.baseXp) {
				fmt.Printf("Caught %s in %d tries.\n", c.name, tries)
				break
			}
			tries += 1
		}
	}
}
