package nbt

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

func toJavaCESU8(s string) []byte {
	var buf []byte
	for _, r := range s {
		if r > 0 && r <= 0x7F {
			buf = append(buf, byte(r))
		} else if r > 0x7F && r <= 0x7FF {
			buf = append(buf, byte(0xC0|(r>>6)), byte(0x80|(r&0x3F)))
		} else if r > 0x7FF && r <= 0xFFFF {
			buf = append(buf, byte(0xE0|(r>>12)), byte(0x80|((r>>6)&0x3F)), byte(0x80|(r&0x3F)))
		} else if r > 0xFFFF {
			r -= 0x10000
			buf = append(buf, 0xED)
			buf = append(buf, byte(0xA0|((r>>10)&0x0F)))
			buf = append(buf, byte(0x80|(r&0x3F)))

			buf = append(buf, 0xED)
			buf = append(buf, byte(0xB0|((r>>10)&0x0F)))
			buf = append(buf, byte(0x80|(r&0x3F)))
		}
	}
	return buf
}

func fromJavaCESU8(b []byte) (string, error) {

	var runes []rune
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		if r != utf8.RuneError {
			runes = append(runes, r)
			i += size
			continue
		}

		if i+2 >= len(b) {
			return "", errors.New("incomplete CESU-8 sequence")
		}

		r1 := rune(b[i])
		r2 := rune(b[i+1])
		r3 := rune(b[i+2])

		if r1 >= 0xC0 && r1 <= 0xDF && r2 >= 0x80 && r2 <= 0xBF {
			runes = append(runes, ((r1&0x1F)<<6)|(r2&0x3F))
			i += 2
			continue
		} else if r1 >= 0xE0 && r1 <= 0xEF && r2 >= 0x80 && r2 <= 0xBF && r3 >= 0x80 && r3 <= 0xBF {
			runes = append(runes, ((r1&0x0F)<<12)|((r2&0x3F)<<6)|(r3&0x3F))
			i += 3
			continue
		} else if r1 == 0xED && r2 >= 0xA0 && r2 <= 0xAF && r3 >= 0x80 && r3 <= 0xBF {
			if i+5 >= len(b) {
				return "", errors.New("incomplete surrogate pair in CESU-8 sequence")
			}

			r4 := rune(b[i+3])
			r5 := rune(b[i+4])
			r6 := rune(b[i+5])

			if r4 == 0xED && r5 >= 0xB0 && r5 <= 0xBF && r6 >= 0x80 && r6 <= 0xBF {

				highSurrogate := (((r2 & 0x0F) << 10) | (r3 & 0x3F))
				lowSurrogate := (((r5 & 0x0F) << 10) | (r6 & 0x3F))

				rune := 0x10000 + (highSurrogate << 10) + lowSurrogate
				runes = append(runes, rune)
				i += 6
				continue
			}
		}

		return "", fmt.Errorf("invalid CESU-8 sequence at byte index %d", i)
	}
	return string(runes), nil
}
