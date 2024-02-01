package helper

import (
	"crypto/sha256"
	"fmt"
)

func LognToken(username, password string) string {
	sum := sha256.Sum256([]byte(username))
	return fmt.Sprintf("%x",sum)
}
