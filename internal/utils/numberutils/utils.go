package numberutils

import (
	"fmt"
	"strconv"
)

func FormatInt64(value int64) string {
	return strconv.FormatInt(value, 10)
}

func FormatInt(value int) string {
	return strconv.Itoa(value)
}

func HumanBytes(bytes int64, base int64) string {
	sizeHuman := ""
	b := float64(base)
	if bytes < base {
		sizeHuman = FormatInt64(bytes) + "B"
	} else if bytes < base*base {
		sizeHuman = fmt.Sprintf("%.2fKB", float64(bytes)/b)
	} else if bytes < base*base*base {
		sizeHuman = fmt.Sprintf("%.2fMB", float64(bytes)/b/b)
	} else if bytes < base*base*base*base {
		sizeHuman = fmt.Sprintf("%.2fGB", float64(bytes)/b/b/b)
	} else if bytes < base*base*base*base*base {
		sizeHuman = fmt.Sprintf("%.2fTB", float64(bytes)/b/b/b/b)
	} else {
		sizeHuman = fmt.Sprintf("%.2fPB", float64(bytes)/b/b/b/b/b)
	}
	return sizeHuman
}

func HumanBytes1000(bytes int64) string {
	return HumanBytes(bytes, 1000)
}

func HumanBytes1024(bytes int64) string {
	return HumanBytes(bytes, 1024)
}
