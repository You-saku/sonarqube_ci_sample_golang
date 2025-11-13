package utils

import (
	"strings"
	"unicode"
)

// Reverse は文字列を逆順にして返します
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsPalindrome は文字列が回文かどうかを判定します
func IsPalindrome(s string) bool {
	// 空白文字と区切り文字を除去し、小文字に変換
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return cleaned == Reverse(cleaned)
}

// Capitalize は文字列の最初の文字を大文字にします
func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

// IsEmpty は文字列が空または空白文字のみかどうかを判定します
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Contains は文字列が指定された部分文字列を含むかを大文字小文字を区別せずに判定します
func ContainsIgnoreCase(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}

// TruncateWithEllipsis は文字列を指定された長さで切り詰め、省略記号を追加します
func TruncateWithEllipsis(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	if maxLength <= 3 {
		return s[:maxLength]
	}
	return s[:maxLength-3] + "..."
}