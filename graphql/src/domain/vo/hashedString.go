package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/tanaka-takuto/goal-minder/sdk"
)

// HashedString ハッシュ化された文字列
type HashedString string

const (
	// saltLength ソルトの長さ
	saltLength = 32
)

// NewHashedString ハッシュ化された文字列を作成する
func NewHashedString(plainStr string) HashedString {
	salt := sdk.RandomString(saltLength)
	plainStrWithSalt := plainStr + salt

	h := sha256.New()
	_, err := h.Write([]byte(plainStrWithSalt))
	if err != nil {
		panic(err)
	}

	hashedBytes := h.Sum(nil)
	hashedStr := hex.EncodeToString(hashedBytes)
	hashedStrWithSalt := fmt.Sprintf("%v:%v:%v", "sha256", salt, hashedStr)

	return HashedString(hashedStrWithSalt)
}
