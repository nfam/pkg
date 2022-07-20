package crc32c

import "hash/crc32"

const hextable = "0123456789abcdef"

var crc32c_table = crc32.MakeTable(crc32.Castagnoli)

// Checksum returns the lowercase heximal CRC-32 checksum
// of data using the Castagnoli's polynomial.
func Hash(data []byte) string {
	crc := crc32.Update(0, crc32c_table, data)
	buf := make([]byte, 8)
	buf[0] = hextable[(crc>>28)&0x0f]
	buf[1] = hextable[(crc>>24)&0x0f]
	buf[2] = hextable[(crc>>20)&0x0f]
	buf[3] = hextable[(crc>>16)&0x0f]
	buf[4] = hextable[(crc>>12)&0x0f]
	buf[5] = hextable[(crc>>8)&0x0f]
	buf[6] = hextable[(crc>>4)&0x0f]
	buf[7] = hextable[(crc)&0x0f]
	return string(buf)
}
