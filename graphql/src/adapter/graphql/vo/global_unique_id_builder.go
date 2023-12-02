package vo

import "goal-minder/domain/model"

// globalUniqueIDBuilder グローバルユニークIDビルダー
type globalUniqueIDBuilder[T ~int] struct {
	key string
}

// New グローバルユニークIDを作成する
func (b *globalUniqueIDBuilder[T]) New(id T) string {
	return string(newGlobalUniqueID(b.key, int(id)))
}

// Decode グローバルユニークIDをデコードする
func (b *globalUniqueIDBuilder[T]) Decode(id string) (T, error) {
	decodedID, err := globalUniqueID(id).decodeByKey(b.key)
	if err != nil {
		return 0, err
	}

	return T(*decodedID), nil
}

// newBuilder グローバルユニークIDビルダーを作成する
func newBuilder[T ~int](key string) globalUniqueIDBuilder[T] {
	return globalUniqueIDBuilder[T]{
		key: key,
	}
}

var (
	AccountID = newBuilder[model.AccountID]("AccountID") // AccountID アカウントID
	GoalID    = newBuilder[model.GoalID]("GoalID")       // GoalID 目標ID
)
