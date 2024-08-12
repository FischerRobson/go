package variables

import "strconv"

// bool

// the number for int represent bits
// int, int8, int16, int32, int64 - int has the same amount of bits as arch (x32 || x64)
// uint, uint8, uint16, uint32, uint64 - allows only + numbers

// byte - alias for uint8

// rune - alias for uin32 -> most used for characters

// float32, float64

// complex64, complex128 - used for complex numbers

func main() {
	var x int = 65
	s := string(x)                              // Instead of be "65" this is a rune "A"
	myString := strconv.FormatInt(int64(x), 10) // "65"
	f := float64(x)

	const x = 10
	takeInt32(x) // handle x as int32
	takeInt64(x) // handle x as int64
}
