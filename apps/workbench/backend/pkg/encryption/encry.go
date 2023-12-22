package encryption

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptPwd(pwd string) string {
	hash := sha256.New()
	hash.Write([]byte(pwd))
	hashBytes := hash.Sum(nil)
	password := hex.EncodeToString(hashBytes)
	return password
}
