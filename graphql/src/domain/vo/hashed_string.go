package vo

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"

	"goal-minder/sdk"
)

// HashedString ハッシュ化された文字列
type HashedString string

const (
	// saltLength ソルトの長さ
	saltLength = 32
)

// newHashedStringWithSalt ソルトを指定してハッシュ化された文字列を作成する
func newHashedStringWithSalt(plainStr string, salt string) HashedString {
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

// NewHashedString ハッシュ化された文字列を作成する
func NewHashedString(plainStr string) HashedString {
	salt := sdk.RandomString(saltLength)
	return newHashedStringWithSalt(plainStr, salt)
}

// ValidString 文字列が正しいかどうかを確認する
func (hs HashedString) ValidString(plainStr string) error {
	hsSplit := strings.Split(string(hs), ":")
	if len(hsSplit) != 3 {
		return fmt.Errorf("invalid hashed string")
	}
	_, salt, _ := hsSplit[0], hsSplit[1], hsSplit[2]

	hashedString := newHashedStringWithSalt(plainStr, salt)
	if hs != hashedString {
		return fmt.Errorf("invalid hashed string")
	}

	return nil

}
