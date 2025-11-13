package utils

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "a"},
		{"hello", "olleh"},
		{"Hello World", "dlroW olleH"},
		{"12345", "54321"},
		{"こんにちは", "はちにんこ"},
	}

	for _, test := range tests {
		result := Reverse(test.input)
		if result != test.expected {
			t.Errorf("Reverse(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"a", "A"},
		{"hello", "Hello"},
		{"HELLO", "HELLO"},
		{"hello world", "Hello world"},
		{"123abc", "123abc"},
		{"こんにちは", "こんにちは"},
	}

	for _, test := range tests {
		result := Capitalize(test.input)
		if result != test.expected {
			t.Errorf("Capitalize(%q) = %q, expected %q", test.input, result, test.expected)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
		{"   ", true},
		{"\t\n\r", true},
		{" \t ", true},
		{"a", false},
		{" a ", false},
		{"hello", false},
		{" hello ", false},
	}

	for _, test := range tests {
		result := IsEmpty(test.input)
		if result != test.expected {
			t.Errorf("IsEmpty(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestContainsIgnoreCase(t *testing.T) {
	tests := []struct {
		s        string
		substr   string
		expected bool
	}{
		{"Hello World", "hello", true},
		{"Hello World", "WORLD", true},
		{"Hello World", "world", true},
		{"Hello World", "xyz", false},
		{"", "a", false},
		{"a", "", true},
		{"", "", true},
		{"GoLang", "go", true},
		{"GoLang", "LANG", true},
		{"GoLang", "python", false},
	}

	for _, test := range tests {
		result := ContainsIgnoreCase(test.s, test.substr)
		if result != test.expected {
			t.Errorf("ContainsIgnoreCase(%q, %q) = %v, expected %v", test.s, test.substr, result, test.expected)
		}
	}
}

func TestTruncateWithEllipsis(t *testing.T) {
	tests := []struct {
		s         string
		maxLength int
		expected  string
	}{
		{"", 10, ""},
		{"short", 10, "short"},
		{"This is a long string", 10, "This is..."},
		{"Hello", 5, "Hello"},
		{"Hello World", 5, "He..."},
		{"abc", 3, "abc"},
		{"abcd", 3, "abc"},
		{"a", 1, "a"},
		{"ab", 1, "a"},
		{"Hello World", 8, "Hello..."},
	}

	for _, test := range tests {
		result := TruncateWithEllipsis(test.s, test.maxLength)
		if result != test.expected {
			t.Errorf("TruncateWithEllipsis(%q, %d) = %q, expected %q", test.s, test.maxLength, result, test.expected)
		}
	}
}