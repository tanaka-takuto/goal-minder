package vo

import (
	"encoding/base64"
	"fmt"
	"goal-minder/config"
	"strconv"
	"strings"
)

// GlobalUniqueID ユニークID
type GlobalUniqueID string

const (
	// applicationName アプリケーション名
	applicationName = "goal-minder"

	// separator ユニークIDのセパレータ
	separator = ":"
)

// newGlobalUniqueID ユニークIDを作成する
func newGlobalUniqueID(key string, id int) GlobalUniqueID {
	base64ID := base64.StdEncoding.EncodeToString([]byte(strings.Join([]string{applicationName, config.Version(), key, fmt.Sprint(id)}, separator)))
	return GlobalUniqueID(base64ID)
}

// decodeByKey キーでIDを取得する
func (guID GlobalUniqueID) decodeByKey(key string) (id *int, err error) {
	decodedIDBytes, err := base64.StdEncoding.DecodeString(string(guID))
	if err != nil {
		return nil, err
	}

	separeteUniqID := strings.Split(string(decodedIDBytes), separator)
	if len(separeteUniqID) != 4 {
		return nil, fmt.Errorf("invalid GlobalUniqueID")
	}

	decodedApplicationName, decodedVersion, decodedKey, decodedIDStr := separeteUniqID[0], separeteUniqID[1], separeteUniqID[2], separeteUniqID[3]
	if decodedApplicationName != applicationName ||
		decodedVersion != config.Version() ||
		decodedKey != key {
		return nil, fmt.Errorf("invalid GlobalUniqueID")
	}

	decodedID, err := strconv.Atoi(decodedIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid GlobalUniqueID:%s", err.Error())
	}

	return &decodedID, nil
}

// String 文字列に変換する
func (guID GlobalUniqueID) String() string {
	return string(guID)
}
