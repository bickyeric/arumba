package comic

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func EncodeCursor(i int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", i)))
}

func DecodeCursor(after *string) (int, error) {
	if after == nil {
		return 0, nil
	}
	rawSkip, err := base64.StdEncoding.DecodeString(*after)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(rawSkip))
}
