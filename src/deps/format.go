package deps

import "fmt"

func FormatBinary32(i uint32) string {
	format := fmt.Sprintf("%%0%db", 32)
	return fmt.Sprintf(format, i)
}

func FormatBinary64(i uint64) string {
	format := fmt.Sprintf("%%0%db", 64)
	return fmt.Sprintf(format, i)
}
