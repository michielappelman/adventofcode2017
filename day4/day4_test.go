package main

import (
	"testing"
)

func TestValidPassphrase(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}
	for _, test := range tests {
		got := validPassphrase(test.input)
		if got != test.want {
			t.Errorf("for %s got %v, want %v", test.input, got, test.want)
		}
	}
}

func TestValidSecurePassphrase(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}
	for _, test := range tests {
		got := validSecurePassphrase(test.input)
		if got != test.want {
			t.Errorf("for %s got %v, want %v", test.input, got, test.want)
		}
	}
}
