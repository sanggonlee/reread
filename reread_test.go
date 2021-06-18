package reread_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/sanggonlee/reread"
)

func ExampleWith() {
	var err error
	var r io.Reader
	r, err = reread.With(r, func(r io.Reader) error {
		// Read from r
		_, err = ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		return nil
	})

	// You can read from r again
	bytes, _ := ioutil.ReadAll(r)
	fmt.Println(string(bytes))
}

func TestWith(t *testing.T) {
	s := "Variously called Grushenka, Grusha, and Grushka, Agrafena Alexandrovna, a beautiful 22-year-old, is the local Jezebel and has an uncanny charm among men."
	r := io.Reader(strings.NewReader(s))
	var err error

	r, err = reread.With(r, func(r io.Reader) error {
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		if string(bytes) != s {
			t.Fatalf("Result did not match after reading for first time")
		}

		return nil
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	r, err = reread.With(r, func(r io.Reader) error {
		bytes, err := ioutil.ReadAll(r)
		if err != nil {
			return err
		}

		if string(bytes) != s {
			t.Fatalf("Result did not match after reading for second time")
		}

		return nil
	})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	bytes, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(bytes) != s {
		t.Fatalf("Result did not match after reading for third time")
	}
}
