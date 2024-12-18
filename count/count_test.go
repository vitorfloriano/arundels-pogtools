package count_test

import (
	"bytes"
	"testing"
	"count"
)

func TestLinesCountsLinesInInput(t *testing.T) {
	t.Parallel()
	c, _ := count.NewCounter()
	c.input = bytes.NewBufferString("1\n2\n3")
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}
