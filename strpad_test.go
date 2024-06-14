package iso8583

import "testing"

func TestLeftPad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{leftPad("foo", 5, "0"), "00foo"},
		{leftPad("foobar", 6, " "), "foobar"},
		{leftPad("1", 2, "0"), "01"},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s but got %s", v.expected, v.s)
		}
	}
}

func TestRightPad(t *testing.T) {
	ts := []struct {
		s, expected string
	}{
		{rightPad("foo", 5, "0"), "foo00"},
		{rightPad("foobar", 6, " "), "foobar"},
		{rightPad("1", 2, "0"), "10"},
		{rightPad("hello", 8, " "), "hello   "},
	}
	for _, v := range ts {
		if v.expected != v.s {
			t.Errorf("Expected %s but got %s", v.expected, v.s)
		}
	}
}
