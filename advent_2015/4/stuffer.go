package advent_2015_4

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

func Stuffer(input string, compare string) int {
	if input == "" {
		return 0
	}

	counter := 1
	for {
		newString := input + strconv.Itoa(counter)
		md5Sum := md5.Sum([]byte(newString))
		md5String := hex.EncodeToString(md5Sum[:])

		if strings.Index(md5String, compare) == 0 {
			return counter
		}
		counter++
	}
}
