package decta

import (
	"os"
	"testing"
)

var decta *Decta

func TestNewDecta(t *testing.T) {
	key := os.Getenv("DECTAKEY")
	if key == "" {
		t.Fatal("environment variable DECTAKEY not found")
	}

	d, err := NewDecta(key)
	if err != nil {
		t.Fatal(err)
	}

	decta = d
}
