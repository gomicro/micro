package license

import (
	"strconv"
	"testing"
	"time"
)

func TestLicenseConstructor(t *testing.T) {
	l := New("neworg")

	expected := "neworg"
	if l.Name != expected {
		t.Errorf("expected: %v\ngot: %v\n", expected, l.Name)
	}

	expected = strconv.FormatInt(int64(time.Now().Year()), 10)
	if l.Year != expected {
		t.Errorf("expected: %v\ngot: %v\n", expected, l.Year)
	}
}
