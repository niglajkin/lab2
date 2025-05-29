package lab2

import (
	"encoding/binary"
	"unicode/utf16"
)

// ComputeHandler should be constructed with input io.Reader and output io.Writer.
// Its Compute() method should read the expression from input and write the computed result to the output.
type ComputeHandler struct {
	// TODO: Add necessary fields.
}

func toUTF8(raw []byte) string {
	if len(raw) >= 2 {
		if raw[0] == 0xFF && raw[1] == 0xFE {
			u16 := make([]uint16, 0, (len(raw)-2)/2)
			for i := 2; i+1 < len(raw); i += 2 {
				u16 = append(u16, binary.LittleEndian.Uint16(raw[i:]))
			}
			return string(utf16.Decode(u16))
		}

		if raw[0] == 0xFE && raw[1] == 0xFF {
			u16 := make([]uint16, 0, (len(raw)-2)/2)
			for i := 2; i+1 < len(raw); i += 2 {
				u16 = append(u16, binary.BigEndian.Uint16(raw[i:]))
			}
			return string(utf16.Decode(u16))
		}
	}
	return string(raw)
}

func (ch *ComputeHandler) Compute() error {
	// TODO: Implement.
	return nil
}
