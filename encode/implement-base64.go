package encode

import (
	"errors"
)

const (
	encodeStd = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	paddingChar = '='
)

var (
	decodeMap map[byte]uint
)

func byteAt(i uint8) byte {
	return []byte(encodeStd)[i]
}

// indexOf() don't check the c is valid.
// call ValidBase64() before call the func.
func indexOf(c byte) uint {
	return decodeMap[c]
}

// indexOfSlow() don't check the c is valid.
// call ValidBase64() before call the func.
// indexOfSlow() use loop to find the index, not map.
func indexOfSLow(c byte) uint {
	var i uint
	for i = 0; i < uint(len(encodeStd)); i++ {
		if []byte(encodeStd)[i] == c {
			break
		}
	}

	return i
}

func init() {
	decodeMap = make(map[byte]uint)
	for i, c := range []byte(encodeStd) {
		decodeMap[c] = uint(i)
	}
}

func EncodeBase64(src []byte) []byte {
	n := (len(src) / 3 ) * 3
	var dst []byte
	for i := 0; i < n; i += 3 {
		var c byte
		// convert the []byte to uint, uint is 64bit.
		group := uint(src[i]) << 16 | uint(src[i+1]) << 8 | uint(src[i+2])

		c = byte(group >> 18 & 0x3f)
		dst = append(dst, byteAt(c))

		c = byte(group >> 12 & 0x3f)
		dst = append(dst, byteAt(c))

		c = byte(group >> 6 & 0x3f)
		dst = append(dst, byteAt(c))

		c = byte(group & 0x3f)
		dst = append(dst, byteAt(c))
	}

	// convert the last section, 2 chars or 3 chars
	remain := len(src) -n
	if remain == 0 {
		return dst
	}

	lastGroup := uint(src[n]) << 16
	if remain == 2 {
		lastGroup |= uint(src[n+1]) << 8
	}

	c := byte(lastGroup >> 18 & 0x3f)
	dst = append(dst, byteAt(c))

	c = byte(lastGroup >> 12 & 0x3f)
	dst = append(dst, byteAt(c))

	switch remain {
	case 1:
		dst = append(dst, paddingChar, paddingChar)
	case 2:
		c = byte(lastGroup >> 6 & 0x3f)
		dst = append(dst, byteAt(c))

		dst = append(dst, paddingChar)
	}

	return dst
}

func ValidBase64(src []byte) (bool, error)  {
	if len(src) % 4 != 0 {
		err := errors.New("The encode is not base64, the lenght is not round 4.")
		return false, err
	}

	for i, c := range src {
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' {
			continue
		}

		if c == encodeStd[len(encodeStd) - 2] || c == encodeStd[len(encodeStd) - 1] {
			continue
		}

		// The paddingChar is the last one or last two
		if c == paddingChar {
			if i == len(src) -2 && src[len(src) - 1] == paddingChar {
				continue
			}

			if i == len(src) - 1 {
				continue
			}
		}

		err := errors.New("The " + string(c) + " is not valid.")
		return false, err
	}

	return true, nil
}

func DecodeBase64(src []byte) ([]byte, error) {
	res, err := ValidBase64(src)
	if res != true {
		return nil, err
	}

	var dst []byte
	for i := 0; i < len(src); i += 4 {
		group := (indexOf(src[i]) & 0x3f << 18 | indexOf(src[i+1]) << 12 | indexOf(src[i+2]) << 6 | indexOf(src[i+3]))

		dst = append(dst, byte(group >> 16 & 0xff))
		if src[i+2] == paddingChar {
			break
		}

		dst = append(dst, byte(group >> 8 & 0xff))
		if src[i+3] == paddingChar {
			break
		}

		dst = append(dst, byte(group & 0xff))
	}

	return dst, nil
}