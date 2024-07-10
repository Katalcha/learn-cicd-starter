package auth

import (
	"testing"
)

func TestGetAPIKeyFailureNoHeader(test *testing.T) {
	// var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

	// expected1 := ""
	// expected2 := ErrNoAuthHeaderIncluded

}

func TestGetAPIFailureMalformedHeader(test *testing.T) {
	// var ErrMalformedAuthHeaderIncluded = errors.New("malformed authorization header")

	// expected1 := ""
	// expected2 := ErrMalformedAuthHeaderIncluded
}
