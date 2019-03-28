package normalizer

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

// Functions below are copied from strconv.Unquote and strconv.Quote.
// Original functions are unable to escape/unescape values containing
// multiple characters since in Go single quotes represent a rune literal
// https://github.com/golang/go/blob/65a54aef5bedbf8035a465d12ad54783fb81e957/src/strconv/quote.go#L360

// unquoteSingle is the same as strconv.Unquote, but uses ' as a quote.
func unquoteSingle(s string) (string, error) {
	n := len(s)
	if n < 2 {
		return "", fmt.Errorf("%+q is not a quoted string", s)
	}
	quote := s[0]
	if quote != s[n-1] {
		return "", fmt.Errorf("%+q does not begin and end with a quote", s)
	}
	s = s[1 : len(s)-1]

	if contains(s, '\n') {
		return "", fmt.Errorf("%+q contains EOL", s)
	}

	// Is it trivial? Avoid allocation.
	if !contains(s, '\\') && !contains(s, quote) {
		r, size := utf8.DecodeRuneInString(s)
		if size == len(s) && (r != utf8.RuneError || size != 1) {
			return s, nil
		}
	}
	s = replaceEscapedMaybe(s, "\\0", "\x00") // treatment of special JS escape seq

	var runeTmp [utf8.UTFMax]byte
	buf := make([]byte, 0, 3*len(s)/2) // Try to avoid more allocations.
	for len(s) > 0 {
		c, multibyte, ss, err := strconv.UnquoteChar(s, '\'')
		if err != nil {
			return "", err
		}
		s = ss
		if c < utf8.RuneSelf || !multibyte {
			buf = append(buf, byte(c))
		} else {
			n := utf8.EncodeRune(runeTmp[:], c)
			buf = append(buf, runeTmp[:n]...)
		}
	}
	return string(buf), nil
}

// contains reports whether the string contains the byte c.
func contains(s string, c byte) bool {
	return strings.IndexByte(s, c) >= 0
}

// replaceEscapedMaybe returns a copy of s in which occurrences of old followed by a
// non-digit are replaced by repl.
// Is not part of the stdlib, handles the special case of JS escape sequence.
// Regexp replacement and manual expansion performance was tested against the
// current implementation and found this was fastest.
func replaceEscapedMaybe(s, old, repl string) string {
	var out strings.Builder
	for s != "" {
		pos := strings.Index(s, old)
		if pos < 0 {
			break
		}
		out.WriteString(s[:pos])
		s = s[pos+len(old):]
		r, n := utf8.DecodeRuneInString(s)
		if r >= '0' && r <= '9' {
			out.WriteString(old)
		} else {
			out.WriteString(repl)
		}
		if strings.IndexRune(old, r) == 0 { // old startsWith r
			continue
		}
		s = s[n:]
		if n != 0 {
			out.WriteRune(r)
		}
	}
	out.WriteString(s)
	return out.String()
}

const lowerhex = "0123456789abcdef"

// quoteSingle is the same as strconv.Quote, but uses ' as a quote.
// quoteSingle(unquoteSingle(s)) may not result in exact same bytes as s,
// because quoteSingle always uses the hex escape sequence format.
func quoteSingle(s string) string {
	const quote = '\''
	buf := make([]byte, 0, 3*len(s)/2)

	buf = append(buf, quote)
	for width := 0; len(s) > 0; s = s[width:] {
		r := rune(s[0])
		width = 1
		if r >= utf8.RuneSelf {
			r, width = utf8.DecodeRuneInString(s)
		}
		if width == 1 && r == utf8.RuneError {
			buf = append(buf, `\x`...)
			buf = append(buf, lowerhex[s[0]>>4])
			buf = append(buf, lowerhex[s[0]&0xF])
			continue
		}
		buf = appendEscapedRune(buf, r, quote)
	}
	buf = append(buf, quote)
	return string(buf)
}

func appendEscapedRune(buf []byte, r rune, quote byte) []byte {
	var runeTmp [utf8.UTFMax]byte
	if r == rune(quote) || r == '\\' { // always backslashed
		buf = append(buf, '\\')
		buf = append(buf, byte(r))
		return buf
	}
	if strconv.IsPrint(r) {
		n := utf8.EncodeRune(runeTmp[:], r)
		buf = append(buf, runeTmp[:n]...)
		return buf
	}
	switch r {
	case '\a':
		buf = append(buf, `\a`...)
	case '\b':
		buf = append(buf, `\b`...)
	case '\f':
		buf = append(buf, `\f`...)
	case '\n':
		buf = append(buf, `\n`...)
	case '\r':
		buf = append(buf, `\r`...)
	case '\t':
		buf = append(buf, `\t`...)
	case '\v':
		buf = append(buf, `\v`...)
	default:
		switch {
		case r < ' ':
			buf = append(buf, `\x`...)
			buf = append(buf, lowerhex[byte(r)>>4])
			buf = append(buf, lowerhex[byte(r)&0xF])
		case r > utf8.MaxRune:
			r = 0xFFFD
			fallthrough
		case r < 0x10000:
			buf = append(buf, `\u`...)
			for s := 12; s >= 0; s -= 4 {
				buf = append(buf, lowerhex[r>>uint(s)&0xF])
			}
		default:
			buf = append(buf, `\U`...)
			for s := 28; s >= 0; s -= 4 {
				buf = append(buf, lowerhex[r>>uint(s)&0xF])
			}
		}
	}
	return buf
}
