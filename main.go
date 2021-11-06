package n7m

import (
	"github.com/rivo/uniseg"
	"strconv"
	"strings"
	"unicode"
)

func isEnd(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}

func writeRunes(b *strings.Builder, runes []rune) {
	for _, v := range runes {
		b.WriteRune(v)
	}
}

func N7m(from string) string {
	gr := uniseg.NewGraphemes(from)
	var b strings.Builder

	var start []rune
	var beforeEnd []rune
	var end []rune
	cnt := -2
	for gr.Next() {
		runes := gr.Runes()
		if len(runes) == 1 && isEnd(runes[0]) {
			if len(start) != 0 {
				writeRunes(&b, start)
			}

			if cnt > 1 {
				b.WriteString(strconv.Itoa(cnt))
			} else if cnt == 1 {
				writeRunes(&b, beforeEnd)
			}

			if cnt >= 0 && len(end) != 0 {
				writeRunes(&b, end)
			}

			start = []rune{}
			end = []rune{}
			cnt = -2

			// write separator
			writeRunes(&b, runes)
		} else {
			if len(start) == 0 {
				start = runes
			}

			beforeEnd = end

			end = runes

			cnt += 1
		}
	}

	if len(start) != 0 {
		writeRunes(&b, start)
	}

	if cnt > 1 {
		b.WriteString(strconv.Itoa(cnt))
	} else if cnt == 1 {
		writeRunes(&b, beforeEnd)
	}

	if cnt >= 0 && len(end) != 0 {
		writeRunes(&b, end)
	}

	return b.String()
}
