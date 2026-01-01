package utils

import "bytes"

func IntArrayToByteArray(array []int) []byte {
	var b bytes.Buffer
	for i := 0; i < len(array); i++ {
		b.WriteByte(byte(array[i]))
	}
	return b.Bytes()
}
