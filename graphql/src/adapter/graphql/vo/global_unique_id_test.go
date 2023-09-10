package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalUniqueID_DecodeByKey(t *testing.T) {
	tests := []struct {
		name     string
		setKey   string
		getKey   string
		id       int
		hasError bool
	}{
		{"正常", "keyName", "keyName", 123, false},
		{"マイナス", "minus", "minus", -123, false},
		{"ゼロ", "zero", "zero", 0, false},
		{"keyなし", "", "", 123, false},
		{"keyが別", "key1", "key2", 123, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uID := NewGlobalUniqueID(tt.setKey, tt.id)
			id, err := uID.DecodeByKey(tt.getKey)

			if tt.hasError {
				assert.Error(t, err)
				assert.Nil(t, id)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, id)
				assert.Equal(t, tt.id, *id)
			}
		})
	}
}
