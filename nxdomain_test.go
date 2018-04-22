package dump

import (
	"testing"

	"github.com/mholt/caddy"
)

func TestSetup(t *testing.T) {
	tests := []struct {
		input     string
		shouldErr bool
	}{
		// positive
		{
			`nxdomain com`, false,
		},
		// negative
		{
			`nxdomain`, true,
		},
	}

	for i, test := range tests {
		c := caddy.NewTestController("dns", test.input)
		err := setup(c)

		if test.shouldErr && err == nil {
			t.Fatalf("Test %d: Expected error but found %s for input %s", i, err, test.input)
		}

		if err != nil {
			if !test.shouldErr {
				t.Fatalf("Test %d: Expected no error but found one for input %s. Error was: %v", i, test.input, err)
			}
		}
	}
}

func TestNxdomain(t *testing.T) {
	// TODO
}
