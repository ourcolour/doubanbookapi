package isbn

import (
	"iamcc.cn/doubanbookapi/frameworks/constants/errs"
	"strconv"
	"strings"
)

func ConvertToIsbn10(isbn13 string) (string, error) {
	var (
		result string
		err    error
	)

	// Args

	result = strings.Replace(isbn13, "-", "", -1)
	if 13 != len(result) {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	// Remove last 1-bit checksum code
	result = result[3:12]

	// Re-calculate checksum
	cs, err := checksumIsbn10(result)
	if nil != err {
		return "", errs.ERR_INVALID_PARAMETERS
	}
	result = result + cs

	return result, err
}

func ConvertToIsbn13(isbn10 string) (string, error) {
	var (
		result string
		err    error
	)

	// Args
	result = strings.Replace(isbn10, "-", "", -1)
	if 10 != len(result) {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	// Append 978 to the header
	// Remove last 1-bit checksum code
	result = "978" + result[:9]

	// Re-calculate checksum
	cs, err := checksumIsbn13(result)
	if nil != err {
		return "", errs.ERR_INVALID_PARAMETERS
	}

	result = result + cs

	return result, err
}

func checksumIsbn10(isbn10WithNoChecksum string) (string, error) {
	var (
		result string
		err    error
	)

	var sum int = 0
	for i := 0; i < len(isbn10WithNoChecksum); i++ {
		// chr
		chr := isbn10WithNoChecksum[i : i+1]

		// char to int
		if no, err := strconv.Atoi(chr); nil != err {
			return "", errs.ERR_INVALID_PARAMETERS
		} else {
			// Plus all odd and even values
			sum += no * (10 - i)
		}
	}

	val := 11 - (sum % 11)
	if 10 == val {
		result += "X"
	} else {
		result += strconv.Itoa(val)
	}

	return result, err
}

func checksumIsbn13(isbn13WithNoChecksum string) (string, error) {
	var (
		result string
		err    error
	)

	var sum int = 0
	for i := 0; i < len(isbn13WithNoChecksum); i++ {
		// chr
		chr := isbn13WithNoChecksum[i : i+1]
		// char to int
		if no, err := strconv.Atoi(chr); nil != err {
			return "", errs.ERR_INVALID_PARAMETERS
		} else {
			// odd * 1
			// even * 3
			// Plus all odd and even values
			if 0 != i%2 {
				sum += no * 3
			} else {
				sum += no
			}
		}
	}

	val := 10 - (sum % 10)
	if 10 == val {
		val = 0
	}

	result = strconv.Itoa(val)

	return result, err
}
