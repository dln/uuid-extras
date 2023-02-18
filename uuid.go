package uuid

import (
	"bytes"
	"encoding/base64"
	"strings"

	"github.com/gofrs/uuid/v5"
)

type UUID struct {
	uuid.UUID
}

var (
	conse = [...]byte{
		'b', 'd', 'f', 'g', 'h', 'j', 'k', 'l', 'm', 'n',
		'p', 'r', 's', 't', 'v', 'z',
	}
	vowse = [...]byte{'a', 'i', 'o', 'u'}
	consd = map[byte]uint16{
		'b': 0, 'd': 1, 'f': 2, 'g': 3,
		'h': 4, 'j': 5, 'k': 6, 'l': 7,
		'm': 8, 'n': 9, 'p': 10, 'r': 11,
		's': 12, 't': 13, 'v': 14, 'z': 15,
	}
	vowsd = map[byte]uint16{'a': 0, 'i': 1, 'o': 2, 'u': 3}
)

func NewV4() (UUID, error) {
	u, err := uuid.NewV4()
	return UUID{u}, err
}

func NewV7() (UUID, error) {
	u, err := uuid.NewV7()
	return UUID{u}, err
}

func (u *UUID) Base64Url() string {
	return base64.RawURLEncoding.EncodeToString(u.Bytes())
}

func FromBase64Url(s string) (UUID, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return UUID{}, err
	}

	u, err := uuid.FromBytes(b)
	return UUID{u}, err
}

func (u *UUID) Proquint() string {
	var out bytes.Buffer

	buf := u.Bytes()

	for i := 0; i < len(buf); i = i + 2 {
		var n uint16 = (uint16(buf[i]) * 256) + uint16(buf[i+1])

		var (
			c3 = n & 0x0f
			v2 = (n >> 4) & 0x03
			c2 = (n >> 6) & 0x0f
			v1 = (n >> 10) & 0x03
			c1 = (n >> 12) & 0x0f
		)

		out.WriteByte(conse[c1])
		out.WriteByte(vowse[v1])
		out.WriteByte(conse[c2])
		out.WriteByte(vowse[v2])
		out.WriteByte(conse[c3])

		if (i + 2) < len(buf) {
			out.WriteByte('-')
		}
	}

	return out.String()
}

func FromProquint(str string) (UUID, error) {
	var (
		out  bytes.Buffer
		bits []string = strings.Split(str, "-")
	)

	for i := 0; i < len(bits); i++ {
		var x uint16 = consd[bits[i][4]] +
			(vowsd[bits[i][3]] << 4) +
			(consd[bits[i][2]] << 6) +
			(vowsd[bits[i][1]] << 10) +
			(consd[bits[i][0]] << 12)

		out.WriteByte(byte(x >> 8))
		out.WriteByte(byte(x))
	}

	u, err := uuid.FromBytes(out.Bytes())
	if err != nil {
		return UUID{}, err
	}
	return UUID{u}, nil
}
