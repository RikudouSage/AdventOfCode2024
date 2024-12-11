package helper

func Pointer[TType any](input TType) *TType {
	return &input
}
