package episode

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

// EncodeCursor encode episode as cursor by encoding it into base64
func EncodeCursor(no int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", no)))
}

// DecodeCursor encode episode as cursor by encoding it into base64
func DecodeCursor(cursor string) (int, error) {
	raw, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(raw))
}
