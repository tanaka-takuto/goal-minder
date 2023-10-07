package model

// Goal 目標
type Goal struct {
	ID        GoalID        // 目標ID
	AccountID AccountID     // アカウントID
	Name      GoalName      // 目標名
	Detail    GoalDetail    // 詳細
	Deadline  *GoalDeadline // 期限
	Scale     *GoalScale    // 目標規模
}

// NewGoal 目標を作成する
func NewGoal(accountID AccountID, name GoalName, detail GoalDetail, deadline *GoalDeadline, scale *GoalScale) Goal {
	return Goal{
		AccountID: accountID,
		Name:      name,
		Detail:    detail,
		Deadline:  deadline,
		Scale:     scale,
	}
}
