package sdk

// Ptr ポインタを生成する
func Ptr[T any](v T) *T {
	return &v
}
