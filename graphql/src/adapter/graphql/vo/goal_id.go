package vo

import "goal-minder/domain/model"

type GoalID GlobalUniqueID

const goalIDKKey = "goal"

// NewGoalID 目標IDを作成する
func NewGoalID(goalID model.GoalID) GoalID {
	return GoalID(newGlobalUniqueID(goalIDKKey, int(goalID)))
}

// Decode 目標IDを取得する
func (a GoalID) Decode() (model.GoalID, error) {
	decodedID, err := GlobalUniqueID(a).decodeByKey(goalIDKKey)
	if err != nil {
		return 0, err
	}

	return model.GoalID(*decodedID), nil
}

// String 文字列に変換する
func (a GoalID) String() string {
	return string(a)
}
